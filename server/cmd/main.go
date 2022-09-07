package main

import (
	"log"
	"server/internal/db"
)

func main() {
	app,err := InitApp()
	if err!=nil{
		log.Fatalln("fff")
	}

	db.Migration(app.DB)
}
