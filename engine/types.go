package engine

// Parser ...
type Parser interface {
	Parse(contents []byte, url string) ParserResult
	Serialize() (name string, args interface{})
}

// ParserFunc ...
type ParserFunc func(
	contents []byte, url string) ParserResult

// Request 请求的URL，和处理该URL结果的Parser
type Request struct {
	URL    string
	Parser Parser
}

// ParserResult 提取结果
type ParserResult struct {
	Reuqests []Request
	Items    []Item
}

// Item ...
type Item struct {
	URL     string
	Type    string // The name of table
	ID      string
	Payload interface{}
}

// NilParser ...
type NilParser struct{}

// Parse ...
func (NilParser) Parse(
	_ []byte, _ string) ParserResult {
	return ParserResult{}
}

// Serialize ...
func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

// 工程函数的方法创建

// FuncParser ...
type FuncParser struct {
	parser ParserFunc
	name   string
}

// Parse ..
func (f *FuncParser) Parse(contents []byte, url string) ParserResult {
	return f.parser(contents, url)
}

// Serialize ...
func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

// NewFuncParser 工程函数的方法创建
func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
