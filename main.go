/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/raihankhan/jarvis/cmd"
	"github.com/raihankhan/jarvis/database"
	"log"
)

func main() {
	err := database.OpenDatabase()
	if err != nil {
		log.Println(err)
		return
	}
	cmd.Execute()
}
