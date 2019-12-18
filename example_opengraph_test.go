package opengraph_test

import (
	"fmt"
	"github.com/masa0221/go-opengraph"
)

// ExampleFetch demonstrates fetch website information defined by OpenGraph protocol from specified url.
func Example_fetch() {
	og, err := opengraph.Fetch("http://ogp.me")
	if (err != nil) {
		fmt.Println(err)
	}

	fmt.Printf("url: %s\n", og["url"])
	fmt.Printf("title: %s\n", og["title"])
	// Output:
	// url: http://ogp.me/
	// title: Open Graph protocol
}
