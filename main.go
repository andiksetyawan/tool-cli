package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
)

var (
	app        = kingpin.New("tool-cli", "tool-cli : A command-line for getting log file")
	source     = app.Arg("source", "Source File").Required().String()
	output     = app.Flag("output", "Output / Destination File").Short('o').String()
	typeOutput = app.Flag("type", "File Type Extension (default: 'text')").Short('t').Default("text").Enum("json", "text")
)

func main() {
	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))

	if *output == "" {
		if *typeOutput == "json" {
			*output = *source + ".json"
		} else if *typeOutput == "text" {
			*output = *source + ".txt"
		}
	}

	err := convert(*source, *output)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	log.Println("log file is converted to " + *output)
}
