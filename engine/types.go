package engine

// Request 请求的URL，和处理该URL结果的Parser
type Request struct {
	URL        string
	ParserFunc func([]byte) ParserResult
}

// ParserResult 提取结果
type ParserResult struct {
	Reuqests []Request
	Items    []interface{}
}

// NilParser ...
func NilParser([]byte) ParserResult {
	return ParserResult{}
}
