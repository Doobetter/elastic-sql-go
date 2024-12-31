package parser

import "github.com/antlr4-go/antlr/v4"

// 对token大小写不敏感
type CaseInsensitiveStream struct {
	stream antlr.CharStream
}

func NewCaseInsensitiveStream(stream antlr.CharStream) *CaseInsensitiveStream {
	return &CaseInsensitiveStream{stream: stream}
}

func (c *CaseInsensitiveStream) Consume() {
	c.stream.Consume()
}

func (c *CaseInsensitiveStream) LA(i int) int {
	result := c.stream.LA(i)

	if result >= 'a' && result <= 'z' {
		return result - 32
	} else {
		return result
	}

}

func (c *CaseInsensitiveStream) Mark() int {
	return c.stream.Mark()
}

func (c *CaseInsensitiveStream) Release(marker int) {
	c.stream.Release(marker)
}

func (c *CaseInsensitiveStream) Index() int {
	return c.stream.Index()
}

func (c *CaseInsensitiveStream) Seek(index int) {
	c.stream.Seek(index)
}

func (c *CaseInsensitiveStream) Size() int {
	return c.stream.Size()
}

func (c *CaseInsensitiveStream) GetSourceName() string {
	return c.stream.GetSourceName()
}

func (c *CaseInsensitiveStream) GetText(i int, i2 int) string {
	return c.stream.GetText(i, i2)
}

func (c *CaseInsensitiveStream) GetTextFromTokens(start, end antlr.Token) string {
	return c.stream.GetTextFromTokens(start, end)
}

func (c *CaseInsensitiveStream) GetTextFromInterval(interval antlr.Interval) string {
	return c.stream.GetTextFromInterval(interval)
}
