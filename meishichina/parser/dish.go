package parser

import (
	"crawler/engine"
	"regexp"
)

const dishes = `<a title="([^"]+)" href="(https://home.meishichina.com/recipe-[1-9]+.html)" target="_blank">`
const nextPage = `<a href="(https://www.meishichina.com/YuanLiao/[a-zA-Z]+/[0-9]+/)">(下一页)</a>`

func ParseDishes(contexts []byte) engine.ParseResult {
	compile := regexp.MustCompile(dishes)
	dishesMatchs := compile.FindAllSubmatch(contexts, -1)

	result := engine.ParseResult{}
	for _, m := range dishesMatchs {
		//result.Items = append(result.Items, string(m[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[2]),
			ParserFunc: ParseRecipe,
		})
	}
	parseYuanliaoListNext(contexts, &result)
	return result
}

func parseYuanliaoListNext(contexts []byte, result *engine.ParseResult) {
	compile := regexp.MustCompile(nextPage)
	nextPageMatchs := compile.FindAllSubmatch(contexts, -1)

	for _, m := range nextPageMatchs {
		//result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseDishes,
		})
	}
}
