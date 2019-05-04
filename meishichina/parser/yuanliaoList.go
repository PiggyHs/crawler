package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

const yunliaoList = `<li><a target="_blank" href="(https://www.meishichina.com/YuanLiao/[a-zA-Z]+/)" title="[^"]+">([^<]+)</a></li>`

func ParseYuanliaoList(contexts []byte) engine.ParseResult {
	compile := regexp.MustCompile(yunliaoList)
	log.Printf("ParseYuanliao")
	yuanliaoListMatchs := compile.FindAllSubmatch(contexts, -1)

	result := engine.ParseResult{}
	for _, m := range yuanliaoListMatchs {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseDishes,
		})
	}
	return result
}
