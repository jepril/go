package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	var keyword string
	fmt.Scanln(&keyword)

	num := 1

	for i := 1; i <= 10; i++ {
		requestUrl := "http://search.zongheng.com/s?keyword=" + keyword + "&pageNo=" + strconv.Itoa(i) + "&sort="
		rp, err := http.Get(requestUrl)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(rp.Body)
		if err != nil {
			panic(err)
		}
		content := string(body)
		defer rp.Body.Close()

		dom, err := goquery.NewDocumentFromReader(strings.NewReader(content))
		if err != nil {
			panic(err)
		}
		dom.Find(".search-tab").Each(func(i int, selection *goquery.Selection){
		//	fmt.Println(selection.Text())
			selection.Find(".tit").Each(func(i int, title *goquery.Selection) {
		//		fmt.Println(title.Text())
				fmt.Printf("%3d   ", num)
				fmt.Println(title.Text())
				num++
			})
		})
	}
}
