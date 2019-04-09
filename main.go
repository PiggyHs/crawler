package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.11yinyuan.com/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",
			resp.StatusCode)
		return
	}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(
		resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", all)
	printAllCityList(all)
}

func determineEncoding(r io.Reader) encoding.Encoding  {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printAllCityList(contexts []byte)  {
	compile := regexp.MustCompile(`<a href="(http://www.11yinyuan.com/citylist-[0-9]+-0.html)">([^<]+)</a>`)

	match := compile.FindAllSubmatch(contexts, -1)


	for _,m := range match{
		fmt.Printf("City: %s, URL: %s \n", m[2],m[1])
	}

	fmt.Println("Match length:", len(match))
}