package search

import "log"

type Result struct {
	Field   string
	Content string
}
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, Results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}
	for _, result := range searchResults {
		Results <- result
	}
}
func Display(results chan *Result) {
	for result := range results {
		log.Println("%s:\n%s\n\n", result.Field, result.Content)
	}
}
