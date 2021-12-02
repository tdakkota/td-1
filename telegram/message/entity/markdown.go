package entity

import (
	"bytes"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/go-faster/errors"

	"github.com/gotd/td/internal/ascii"
	"github.com/gotd/td/tg"
)

type markdown struct {
	builder      *Builder
	userResolver UserResolver
	offset       int
	stack        []stackElem
}

func (m *markdown) addEntity(offset, utf8offset int, f Formatter) {
	length := ComputeLength(m.builder.message.String())
	u8 := utf8entity{
		offset: utf8offset,
		length: m.builder.message.Len() - utf8offset,
	}
	m.builder.appendEntities(offset, length-offset, u8, f)
}

func (m *markdown) pop(c byte) error {
	if len(m.stack) < 1 {
		return io.ErrUnexpectedEOF
	}

	last := m.stack[len(m.stack)-1]
	if last.tag[0] != c {
		return errors.Errorf("unexpected %c", c)
	}
	m.stack = m.stack[:len(m.stack)-1]
	length := ComputeLength(m.builder.message.String())
	u8 := utf8entity{
		offset: last.utf8offset,
		length: m.builder.message.Len() - last.utf8offset,
	}
	m.builder.appendEntities(last.offset, length-last.offset, u8, last.format)
	return nil
}

func (m *markdown) push(c string, f Formatter) {
	m.stack = append(m.stack, stackElem{
		offset:     m.offset,
		utf8offset: m.builder.message.Len(),
		tag:        c,
		format:     f,
	})
}

func (m *markdown) writeByte(c byte) {
	m.builder.message.WriteByte(c)
	m.offset++
}

func (m *markdown) writeRune(c rune) {
	m.builder.message.WriteRune(c)
	m.offset++
	if c >= 0xF0 {
		m.offset++
	}
}

func (m *markdown) write(s []byte) {
	m.builder.message.Write(s)
	m.offset += ComputeLength(string(s))
}

func (m *markdown) open(data []byte, c, next byte, i int) (int, error) {
	switch c {
	case '_':
		if next == '_' {
			i++
			m.push("__", Underline())
			break
		}
		m.push(string(c), Italic())
	case '*':
		m.push(string(c), Bold())
	case '~':
		m.push(string(c), Strike())
	case '[':
		start := m.offset
		startUTF8 := m.builder.message.Len()

		bracketIdx := bytes.IndexByte(data[i:], ']')
		if bracketIdx < 0 {
			return 0, errors.New("] expected")
		}
		m.write(data[i+1 : bracketIdx])

		i += bracketIdx
		if bracketIdx+2 >= len(data) || data[bracketIdx+1] != '(' {
			break
		}

		i++
		parenthesesIdx := bytes.IndexByte(data[i:], ')')
		if parenthesesIdx < 0 {
			return 0, errors.New(") expected")
		}
		rawURL := string(data[i+1 : i+parenthesesIdx])
		i += parenthesesIdx

		f, err := getURLFormatter(rawURL, m.userResolver)
		if err != nil {
			return 0, errors.Errorf("invalid URL %q: %w", rawURL, err)
		}
		m.addEntity(start, startUTF8, f)

	case '`':
		if next != '`' {
			m.push(string(c), Code())
			break
		}
		if i+2 < len(data) && data[i+2] != '`' {
			return 0, errors.New("` expected")
		}
		i += 3

		var langEnd int
		for langEnd = i; langEnd < len(data); langEnd++ {
			r := rune(data[langEnd])
			if ascii.IsSpace(r) || r == '`' {
				break
			}
		}

		f := Code()
		if i != langEnd && data[langEnd] != '`' {
			f = Pre(string(data[i:langEnd]))
			i = langEnd
		}

		if i+1 < len(data) && (data[i] == '\n' || data[i] == '\r') {
			if (data[i+1] == '\n' || data[i+1] == '\r') && data[i] != data[i+1] {
				i += 2
			} else {
				i++
			}
		}

		i--
		m.push("```", f)
	}

	return i, nil
}

// parseV2 parses Telegram MarkdownV2.
func (m *markdown) parseV2(data []byte) error {
	const (
		reservedChars = "_*[]()~`>#+-=|{}.!"
		reservedQuote = "`"
	)

	var (
		length   = len(data)
		reserved = reservedChars
	)
	for i := 0; i < len(data); i++ {
		r, size := utf8.DecodeRune(data[i:])
		i += size - 1

		var next byte
		if i+1 < length {
			next = data[i+1]
		}

		if r == '\\' && next > 0 && next <= 126 {
			m.writeByte(next)
			i++
			continue
		}

		if !strings.ContainsRune(reserved, r) {
			m.writeRune(r)
			continue
		}

		c := byte(r)
		var (
			closing bool
			last    stackElem
		)
		if len(m.stack) > 0 {
			last = m.stack[len(m.stack)-1]
			switch last.tag {
			case "__":
				closing = last.tag[0] == c && next == c
			default:
				closing = last.tag[0] == c
			}
		}
		if !closing {
			idx, err := m.open(data, c, next, i)
			if err != nil {
				return errors.Wrapf(err, "at %d", i)
			}
			i = idx
		} else {
			last := m.stack[len(m.stack)-1]
			switch last.tag {
			case "```":
				i += 2
			case "__":
				i++
			}
			if err := m.pop(c); err != nil {
				return errors.Wrapf(err, "at %d", i)
			}
		}
	}

	return nil
}

// Markdown parses given input from reader and adds parsed entities to given builder.
//
// Parameter userResolver is used to resolve user by ID during formatting. May be nil.
// If userResolver is nil, formatter will create tg.InputUser using only ID.
// Notice that it's okay for bots, but not for users.
//
// See https://core.telegram.org/bots/api#markdownv2-style.
func Markdown(r io.Reader, b *Builder, userResolver UserResolver) error {
	if userResolver == nil {
		userResolver = func(id int64) (tg.InputUserClass, error) {
			return &tg.InputUser{
				UserID: id,
			}, nil
		}
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return errors.Wrap(err, "read data")
	}
	m := markdown{
		builder:      b,
		userResolver: userResolver,
	}

	return m.parseV2(data)
}
