package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

//func NewFuncParser(
//	p ParserFunc, name string) *FuncParser {
//	return &FuncParser{
//		parser: p,
//		name:   name,
//	}
//}
