package parser

import (
	"crawler_go/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParserCity(
	contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	for _, m := range matchs {
		name := string(m[2])
		result.Items = append(
			result.Items, "User "+name)
		result.Reuqests = append(
			result.Reuqests, engine.Request{
				URL: string(m[1]),
				ParserFunc: func(c []byte) engine.ParserResult {
					return ParserProfile(c, name)
				},
			})
	}
	return result
}
