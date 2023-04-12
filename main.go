package main

import (
	"github.com/muhadyan/financial-planner/database"
	"github.com/muhadyan/financial-planner/route"
)

func main() {
	database.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
