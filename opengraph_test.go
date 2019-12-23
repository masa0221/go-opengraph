package opengraph_test

import (
	"testing"

	"github.com/masa0221/go-opengraph"
)

func TestFailure (t *testing.T) {
	_, err := opengraph.Fetch("test")
	if (err == nil) {
		t.Fatal("Error was not occurred nevertheless it used incorrect url")
	}
}

func TestSuccess (t *testing.T) {
	og, err := opengraph.Fetch("http://ogp.me")
	if (err != nil) {
		t.Fatal("It was cannot fetch Opengraph data")
	}

	expectedTitle := "Open Graph protocol"
	if og["title"] != expectedTitle {
		t.Errorf("Title is not match with expected value (expected: %s, actual: %s)", og["title"], expectedTitle)
	}
	expectedURL := "http://ogp.me/"
	if og["url"] != expectedURL {
		t.Errorf("Url is not match with expected value (expected: %s, actual: %s)", og["url"], expectedURL)
	}
	expectedImage := "http://ogp.me/logo.png"
	if og["image"] != expectedImage {
		t.Errorf("Image is not match with expected value (expected: %s, actual: %s)",og["image"], expectedImage)
	}
}
