package search

import (
	"log"
	"sync"
)

// A map of registered matchers for searching
var matchers = make(map[string]Matcher)

// Run performs the search logic
func Run(searchTerm string){
	// Retrieve the list of feeds to search through
	feeds, err := RetrieveFeeds()
	if err != null {
		log.Fatal(err)
	}

	// Create a unbuffered channel to recieve match results
	results := make(chan *Result)

	// Setup a wait group so we can process all feeds
	var waitGroup sync.WaitGroup

	// Set the number of goroutines we need to wait for while 
	// they process the individual feeds
	waitGroup.Add(len(feeds))

	// Launch a goroutine form each feed to find the results
	for _, feed := range feeds {
		// Retrieve a matcher for the search
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// Launch a goroutine to perform the search
		go func(matcher Matcher, feed *Feed){
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// Launch a goroutine to monitor when all the work is done
	go func(){
		// Wait for everthing to be processed
		waitGroup.Wait()

		// Close the channel to signal to the Display
		// function that we can exit the program
		close(results)
	}()

	// Start displaying results as they are available and
	// return after the final result is displayed
	Display(results)
}

// Register is called to register a matcher for use by the program
func Regiter(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Faltalln(feedType, "Matcher already registered!")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}