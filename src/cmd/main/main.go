package main

import "github.com/tobbensol/elga_3_website/src/internal/App"

//go:generate templ generate -path=../../.

func main() {
	App.StartApp()
}
