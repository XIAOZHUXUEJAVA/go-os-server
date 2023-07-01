package main

import (
	"fmt"
	"oschina/model"
	"oschina/routes"
)

func main() {
	fmt.Println("hello, oschina")
	model.InitDb()
	routes.InitRouter()
	// rss, err := model.GetAllRss()
	// if err != nil {
	// 	log.Fatalf("%s", "find all rss error")
	// }
	// for _, r := range rss {
	// 	fmt.Printf("Title: %s\n", r.Title)
	// 	fmt.Printf("Description: %s\n", r.Description)
	// 	fmt.Println()
	// }
}
