package main

import (
	_ "fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/spider"
	"strings"
)

type MyPageProcesser struct {
}

func NewMyPageProcesser() *MyPageProcesser {
	return &MyPageProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (this *MyPageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	query := p.GetHtmlParser()

	title := query.Find("title").Text()
	//title = strings.Trim(title, " \t\n")

	keywords := ""
	description := ""
	query.Find("meta").Each(func(i int, s *goquery.Selection) {
		name, _ := s.Attr("name")
		name = strings.ToLower(name)
		if name == "keywords" {
			keywords, _ = s.Attr("content")
		}
		if name == "description" {
			description, _ = s.Attr("content")
		}
	})

	// the entity we want to save by Pipeline
	p.AddField("title", title)
	p.AddField("keywords", keywords)
	p.AddField("description", description)
}

func main() {
	// spider input:
	//  PageProcesser ;
	//  task name used in Pipeline for record;
	sp := spider.NewSpider(NewMyPageProcesser(), "TaskName")
	pageItems := sp.Get("http://www.qq.com", "html") // url, html is the responce type ("html" or "json" or "jsonp" or "text")

	url := pageItems.GetRequest().GetUrl()
	println("-----------------------------------spider.Get---------------------------------")
	println("url\t:\t" + url)
	for name, value := range pageItems.GetAll() {
		println(name + "\t:\t" + value)
	}

	/*
		println("\n--------------------------------spider.GetAll---------------------------------")
		urls := []string{
			"http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn",
			"http://baike.baidu.com/view/383720.htm?fromtitle=html&fromid=97049&type=syn",
		}
		pageItemsArr := sp.SetThreadnum(2).GetAll(urls, "html")
		for _, item := range pageItemsArr {
			url = item.GetRequest().GetUrl()
			println("url\t:\t" + url)
			fmt.Printf("item\t:\t%s\n", item.GetAll())
		}*/
}
