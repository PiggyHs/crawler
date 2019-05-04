package parser

import (
	"crawler/engine"
	"regexp"
)

const dishes = `<a title="([^"]+)" href="(https://home.meishichina.com/recipe-[1-9]+.html)" target="_blank">`

func ParseDishes(contexts []byte) engine.ParseResult {
	compile := regexp.MustCompile(dishes)
	dishesMatchs := compile.FindAllSubmatch(contexts, -1)

	result := engine.ParseResult{}
	for _, m := range dishesMatchs {
		result.Items = append(result.Items, string(m[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[2]),
			ParserFunc: ParseRecipe,
		})
	}
	return result
}
