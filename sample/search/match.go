package search

import(
	"log"
)

//Result contains the result of a search
type Result struct {
	Field string
	Content string
}

// Matcher defines the behavior required by types that want
// to implement a new search type
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is launched as a goroutine for each individual feed to run
// searches concurrently
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan <- *Result) {
	// Perform the search against the specified matcher
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	//Write the results to the chanel
	for _,result := range searchResults {
		results <- result
	}
}

// Display writes results to the terminal window as the
// are received by the individual goroutines
func Display(results chan *Result) {
	// The channel blocks until a request is written to the channel
	// Once the channel is closed the for loop terminates
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}