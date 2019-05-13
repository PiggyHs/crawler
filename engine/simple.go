package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (*SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Go Item %s", item)
		}
	}

}

//进行请求网页和解析取值网页内容
func Worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fercher(r.Url)
	if err != nil {
		log.Printf("Fetcher error"+
			"fetcher Url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
