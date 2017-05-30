package main

import(
	"log"
	"os"
	_ "data-feeds/sample/matchers"
	"data-feeds/sample/search"
)

//init is called prior to main
func init() {
	//Change the device for logging to stdout
	log.setOutput(os.Stdout)
}

//main is the entry point form the program
func main() {
	//Perform the search for the specified term.
	search.Run("president")
}
