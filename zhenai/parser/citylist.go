package parser

import (
	"regexp"

	"github.com/simonzs/crawler_go/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// CSS 选择器
// 使用XPath
// 使用正则表达式

// ParserCityList 根据城市列表， 提取城市链接
// 城市列表解析器 Parser 输入:utf-8编码的文本 输出:Request{URL, 对应Parser}列表, Items列表
func ParserCityList(
	contents []byte, _ string) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}

	for _, m := range matchs {
		result.Reuqests = append(
			result.Reuqests, engine.Request{
				URL: string(m[1]),
				Parser: engine.NewFuncParser(
					ParserCity, "ParserCity"),
			})
	}
	return result
}
