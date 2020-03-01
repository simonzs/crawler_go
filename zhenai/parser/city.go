package parser

import (
	"crawler_go/engine"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityRe = regexp.MustCompile(
		`<a href="(http://www.zhenai.com/zhenghun/chengdu/[^"]+)"`)
)

// ParserCity ...
func ParserCity(
	contents []byte) engine.ParserResult {

	matches := profileRe.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	for _, m := range matches {
		name := string(m[2])
		result.Reuqests = append(
			result.Reuqests, engine.Request{
				URL: string(m[1]),
				ParserFunc: func(c []byte) engine.ParserResult {
					return ParserProfile(c, name)
				},
			})
	}

	matches = cityRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Reuqests = append(
			result.Reuqests, engine.Request{
				URL:        string(m[1]),
				ParserFunc: ParserCity,
			})
	}

	return result
}
