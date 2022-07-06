package main

import (
	"go.mod/handlers"
	"go.mod/model"
)

func main(){
	model.InitDB()
	handlers.InitHandlers()

}

