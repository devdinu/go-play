package main

import "fmt"

type quote struct {
	Quote  string
	Author string
	Cat    string
}

func (q *quote) SetQuote(qt string) {
	fmt.Println("quote is being overridden...")
	q.Quote = qt
}

func (q quote) String() string {
	return fmt.Sprintf("Quote: %s\n  - %s\n", q.Quote, q.Author)
}

func checkReceivers() {
	var q1 quote
	q1 = quote{
		Quote:  "Always code as if the guy, who ends up maintaining your code will be a violent psychopath, who knows where you live.",
		Author: "Rick Osborne",
	}

	fmt.Println(q1)

	q1.SetQuote("Debuggers don't remove bugs. They only show them in slow motion.")
	fmt.Println(q1)

	q1.Quote = "modified"
	fmt.Println(q1)

}
