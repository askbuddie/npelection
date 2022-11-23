package main

import (
	_ "embed"

	"npelection/cmd"
)

//go:embed data/data.json
var districtsLink string

func main() {
	cmd.Execute(districtsLink)
}
