package entity

import (
	"io"
	"strings"

	"github.com/go-faster/errors"
	"golang.org/x/net/html"

	"github.com/gotd/td/tg"
)

type stackElem struct {
	offset     int
	utf8offset int
	tag        string
	format     Formatter
}

type htmlParser struct {
	tokenizer    *html.Tokenizer
	builder      *Builder
	offset       int
	stack        []stackElem
	attr         map[string]string
	userResolver func(id int64) (tg.InputUserClass, error)
}

func (p *htmlParser) fillAttrs() {
	// Clear old attrs.
	for k := range p.attr {
		delete(p.attr, k)
	}

	// Fill with new attributes.
	for {
		key, value, ok := p.tokenizer.TagAttr()
		p.attr[string(key)] = string(value)
		if !ok {
			break
		}
	}
}

func (p *htmlParser) startTag() error {
	const pre = "pre"

	var e stackElem
	tn, hasAttr := p.tokenizer.TagName()
	e.tag = string(tn)
	if hasAttr {
		p.fillAttrs()
	}

	e.offset = p.offset
	e.utf8offset = p.builder.message.Len()
	// See https://core.telegram.org/bots/api#html-style.
	switch e.tag {
	case "b", "strong":
		e.format = Bold()
	case "i", "em":
		e.format = Italic()
	case "u", "ins":
		e.format = Underline()
	case "s", "strike", "del":
		e.format = Strike()
	case "a":
		href, ok := p.attr["href"]
		if !ok {
			return errors.Errorf("tag %q must have attribute href", e.tag)
		}

		f, err := getURLFormatter(href, p.userResolver)
		if err != nil {
			return errors.Errorf("href must be a valid URL, got %q", href)
		}
		e.format = f
	case "code":
		e.format = Code()

		// BotAPI docs says:
		// > Use nested <pre> and <code> tags, to define programming language for <pre> entity.
		if len(p.stack) > 0 && p.stack[len(p.stack)-1].tag == pre {
			lang, ok := p.attr["class"]
			if ok {
				e.format = Pre(strings.TrimPrefix(lang, "language-"))
			}
		}
	case pre:
		e.format = Code()
	}

	p.stack = append(p.stack, e)
	return nil
}

func (p *htmlParser) endTag() error {
	tn, _ := p.tokenizer.TagName()
	if len(p.stack) == 0 {
		return errors.Errorf("unexpected end tag %q", string(tn))
	}

	var s stackElem
	// Pop from SliceTricks.
	s, p.stack = p.stack[len(p.stack)-1], p.stack[:len(p.stack)-1]
	if s.tag != string(tn) {
		return errors.Errorf("expected tag %q, got %q", s.tag, string(tn))
	}

	length := ComputeLength(p.builder.message.String())
	if s.format != nil {
		u8 := utf8entity{
			offset: s.utf8offset,
			length: p.builder.message.Len() - s.utf8offset,
		}
		p.builder.appendEntities(s.offset, length-s.offset, u8, s.format)
	}
	return nil
}

func (p *htmlParser) parse() error {
	for {
		tt := p.tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			if err := p.tokenizer.Err(); !errors.Is(err, io.EOF) {
				return err
			}
			return nil
		case html.TextToken:
			text := p.tokenizer.Text()
			p.builder.message.Write(text)
			p.offset += ComputeLength(string(text))
		case html.StartTagToken:
			if err := p.startTag(); err != nil {
				return err
			}
		case html.EndTagToken:
			if err := p.endTag(); err != nil {
				return err
			}
		}
	}
}

// HTML parses given input from reader and adds parsed entities to given builder.
// Notice that this parser ignores unsupported tags.
//
// Parameter userResolver is used to resolve user by ID during formatting. May be nil.
// If userResolver is nil, formatter will create tg.InputUser using only ID.
// Notice that it's okay for bots, but not for users.
//
// See https://core.telegram.org/bots/api#html-style.
func HTML(r io.Reader, b *Builder, userResolver UserResolver) error {
	if userResolver == nil {
		userResolver = func(id int64) (tg.InputUserClass, error) {
			return &tg.InputUser{
				UserID: id,
			}, nil
		}
	}

	p := htmlParser{
		tokenizer:    html.NewTokenizer(r),
		builder:      b,
		attr:         map[string]string{},
		userResolver: userResolver,
	}

	return p.parse()
}
