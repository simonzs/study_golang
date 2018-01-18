package engine

type ParserFunc func([]byte) ParseResult
type Request struct {
	Url        string
	ParserFunc ParserFunc
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
