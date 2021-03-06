package parser

import (
	"regexp"

	"github.com/simonzs/crawler_go/engine"
)

var (
	profileRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityRe = regexp.MustCompile(
		`<a href="(http://www.zhenai.com/zhenghun/chengdu/[^"]+)"`)
)

// ParserCity ...
func ParserCity(
	contents []byte, _ string) engine.ParserResult {

	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	for _, m := range matches {
		result.Reuqests = append(
			result.Reuqests, engine.Request{
				URL:    string(m[1]),
				Parser: NewProfileParser(string(m[2])),
			})
	}

	matches = cityRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Reuqests = append(
			result.Reuqests, engine.Request{
				URL: string(m[1]),
				Parser: engine.NewFuncParser(
					ParserCity, "ParserCity"),
			})
	}

	return result
}
