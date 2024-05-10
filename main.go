package main

import (
	"log"
	"os"

	"github.com/Samueelx/issues/github"
	"github.com/Samueelx/issues/report"
)

/*Print a table of Github issues matching the search params*/
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.IssueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
