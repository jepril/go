package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	var keyword string
	fmt.Scanln(&keyword)

	num := 1

	for i := 1; i <= 2; i++ {
		requestUrl := "http://studygolang.com/search?q=" + keyword + "&f=title&p=" + strconv.Itoa(i)
		rp, err := http.Get(requestUrl)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(rp.Body)
		if err != nil {
			panic(err)
		}
		content := string(body)
		//fmt.Println(content)
		defer rp.Body.Close()

		dom, err := goquery.NewDocumentFromReader(strings.NewReader(content))
		if err != nil {
			panic(err)
		}
		dom.Find(".row").Each(func(i int, selection *goquery.Selection) {
			//	fmt.Println(selection.Text())
			selection.Find(".website").Each(func(i int, title *goquery.Selection) {
				//		fmt.Println(title.Text())
				fmt.Printf("%3d   ", num)
				fmt.Println(title.Text())
				num++
			})
		})
	}
}
