package entity

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gotd/td/tg"
)

func TestItalicUnderline(t *testing.T) {
	a := assert.New(t)
	b := Builder{}
	a.NoError(Markdown(strings.NewReader("___italic underline___"), &b, nil))

	msg, entities := b.Complete()
	a.Equal("italic underline", msg)
	a.Equal(getEntities(Underline(), Italic())(msg), entities)
}

func TestMarkdown(t *testing.T) {
	tests := []struct {
		md       string
		msg      string
		entities func(msg string) []tg.MessageEntityClass
	}{
		{`*bold \*text*`, "bold *text", getEntities(Bold())},
		{`_italic \*text_`, "italic *text", getEntities(Italic())},
		{"__underline__", "underline", getEntities(Underline())},
		{"___italic underline_\r__", "italic underline", getEntities(Italic(), Underline())},
		{"~strikethrough~", "strikethrough", getEntities(Strike())},
		{"`code`", "code", getEntities(Code())},
		{`[inline URL](http://www.example.com/)`, "inline URL",
			getEntities(TextURL("http://www.example.com/"))},
		{`[inline mention of a user](tg://user?id=123456789)`, "inline mention of a user",
			getEntities(MentionName(&tg.InputUser{
				UserID: 123456789,
			}))},
		{"`code`", "code",
			getEntities(Code())},
		{"```\ncode\n```", "code",
			getEntities(Code())},
		{"```python\npython code\n```", "python code",
			getEntities(Pre("python"))},
	}

	for _, test := range tests {
		t.Run(strings.Title(test.msg), func(t *testing.T) {
			a := assert.New(t)
			b := Builder{}
			a.NoError(Markdown(strings.NewReader(test.md), &b, nil))

			msg, entities := b.Complete()
			a.Equal(test.msg, msg)
			a.Equal(test.entities(test.msg), entities)
		})
	}
}
