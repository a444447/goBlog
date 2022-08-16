package main

import (
	"goBlog/model"
	"goBlog/routers"
)

func main() {
	model.InitDB()
	routers.InitRoute()
}
