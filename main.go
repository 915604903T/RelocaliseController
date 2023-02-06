package main

import (
	"github.com/915604903T/RelocaliseController/handlers"

	"github.com/gorilla/mux"
)

const NameExpression = "-a-zA-Z_0-9."

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/info", handlers.MakeReceiveRelocInfoHandler()).
		Queries("name1", "{name1:["+NameExpression+"]+}", "name2", "{name2:["+NameExpression+"]+}")
	router.HandleFunc("/scene/{name}", handlers.MakeSendSceneModelHandler())
}
