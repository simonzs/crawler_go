package parser

import (
	"crawler_go/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// CSS 选择器
// 使用XPath
// 使用正则表达式

// ParserCityList 根据城市列表， 提取城市链接
// 城市列表解析器 Parser 输入:utf-8编码的文本 输出:Request{URL, 对应Parser}列表, Items列表
func ParserCityList(
	contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	limit := 5
	for _, m := range matchs {
		result.Items = append(
			result.Items, "City "+string(m[2]))
		result.Reuqests = append(
			result.Reuqests, engine.Request{
				URL:        string(m[1]),
				ParserFunc: ParserCity,
			})

		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
