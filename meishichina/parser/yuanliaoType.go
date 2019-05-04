package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

const yunliaoTypeRe = `<li><a title="更多([^"]+)" href="(https://www.meishichina.com/YuanLiao/category/([^/]+)/)" target="_blank">更多</a></li>`

func ParseYuanliaoType(contexts []byte) engine.ParseResult {
	compile := regexp.MustCompile(yunliaoTypeRe)
	log.Printf("ParseYuanliao")
	yuanliaoMatchs := compile.FindAllSubmatch(contexts, -1)

	result := engine.ParseResult{}
	for _, m := range yuanliaoMatchs {
		result.Items = append(result.Items, string(m[1]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[2]),
			ParserFunc: ParseYuanliaoList,
		})
	}
	return result
}
