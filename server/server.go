package server

import (
	"log"
)

func Init() {
	r := NewRouter()
	err := r.Run(":8001") //
	if err != nil {
		log.Fatalln(err)
	}
}
