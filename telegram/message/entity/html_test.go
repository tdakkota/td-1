package entity

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"

	"github.com/gotd/td/tg"
)

func getEntities(formats ...Formatter) func(msg string) []tg.MessageEntityClass {
	return func(msg string) []tg.MessageEntityClass {
		length := ComputeLength(msg)
		r := make([]tg.MessageEntityClass, len(formats))
		for i := range formats {
			r[i] = formats[i](0, length)
		}
		return r
	}
}

func TestHTML(t *testing.T) {
	tests := []struct {
		html     string
		msg      string
		entities func(msg string) []tg.MessageEntityClass
	}{
		{"<b>bold</b>", "bold", getEntities(Bold())},
		{"<strong>bold</strong>", "bold", getEntities(Bold())},
		{"<i>italic</i>", "italic", getEntities(Italic())},
		{"<em>italic</em>", "italic", getEntities(Italic())},
		{"<u>underline</u>", "underline", getEntities(Underline())},
		{"<ins>underline</ins>", "underline", getEntities(Underline())},
		{"<s>strikethrough</s>", "strikethrough", getEntities(Strike())},
		{"<strike>strikethrough</strike>", "strikethrough", getEntities(Strike())},
		{"<del>strikethrough</del>", "strikethrough", getEntities(Strike())},
		{"<code>code</code>", "code", getEntities(Code())},
		{"<pre>abc</pre>", "abc", getEntities(Code())},
		{`<a href="http://www.example.com/">inline URL</a>`, "inline URL",
			getEntities(TextURL("http://www.example.com/"))},
		{`<a href="tg://user?id=123456789">inline mention of a user</a>`, "inline mention of a user",
			getEntities(MentionName(&tg.InputUser{
				UserID: 123456789,
			}))},
		{`<pre><code class="language-python">python code</code></pre>`, "python code",
			getEntities(Pre("python"), Code())},
		{"<b>&lt;</b>", "<", getEntities(Bold())},
	}

	for _, test := range tests {
		t.Run(strings.Title(test.msg), func(t *testing.T) {
			a := require.New(t)
			b := Builder{}
			a.NoError(HTML(strings.NewReader(test.html), &b, nil))

			msg, entities := b.Complete()
			a.Equal(test.msg, msg)
			a.Equal(test.entities(test.msg), entities)
		})
	}
}

func TestIssue525(t *testing.T) {
	test := func(text string, expected []tg.MessageEntityClass) func(t *testing.T) {
		return func(t *testing.T) {
			a := require.New(t)

			b := Builder{}
			p := htmlParser{
				tokenizer:    html.NewTokenizer(strings.NewReader(text)),
				builder:      &b,
				attr:         map[string]string{},
				userResolver: nil,
			}

			a.NoError(p.parse())
			_, entities := b.Complete()
			a.Equal(expected, entities)
		}
	}

	t.Run("Ru", test(`Строка
<i>Строка текста курсивом</i>

Обычный текст с <a href="https://google.com">Ссылкой</a> внутри, и
ещё одна ссылка - <a href="https://go.dev">Здесь</a>.

Ещё одна строка.
`,
		[]tg.MessageEntityClass{
			&tg.MessageEntityItalic{
				Offset: 7,
				Length: 22,
			},
			&tg.MessageEntityTextURL{
				Offset: 47,
				Length: 7,
				URL:    "https://google.com",
			},
			&tg.MessageEntityTextURL{
				Offset: 83,
				Length: 5,
				URL:    "https://go.dev",
			},
		}),
	)
	t.Run("En", test(`Line
<i>Italic line of text</i>

Normal line of text with <a href="https://google.com">Link</a> inside, and
another link now - <a href="https://go.dev">Here</a>.

One more line.
`,
		[]tg.MessageEntityClass{
			&tg.MessageEntityItalic{
				Offset: 5,
				Length: 19,
			},
			&tg.MessageEntityTextURL{
				Offset: 51,
				Length: 4,
				URL:    "https://google.com",
			},
			&tg.MessageEntityTextURL{
				Offset: 87,
				Length: 4,
				URL:    "https://go.dev",
			},
		}),
	)
}
