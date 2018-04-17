package main

import (
	"github.com/projects/newsgrabber/news"
	"github.com/projects/newsgrabber/router"
)

func main() {
	r := router.New()
	a := news.New()

	go a.Serve()
	r.Run()

}
