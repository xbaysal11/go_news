package main

import (
	"go_news/news"
	"go_news/routes"
)

func main() {
	a := news.NewArchive()
	r := routes.New()

	go a.CollectNews()
	r.Run()
}
