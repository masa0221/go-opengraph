/*
opengraph packages provide feature that to fetch website information defined by Open Graph protcol from specified url.
The Open Graph protocol
https://ogp.me/
*/
package opengraph

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Fetch the website information specified by OpenGraph protcol
func Fetch(url string) (map[string]interface{}, error) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	ogTypePrefix := "og:"
	ogTypeAttributeName := "property"
	ogValueAttributeName := "content"
	selection := doc.Find(fmt.Sprintf("meta[property^='%s']", ogTypePrefix))

	og := make(map[string]interface{})
	selection.Each(func(index int, s *goquery.Selection) {
		if attr, exists := s.Attr(ogTypeAttributeName); exists == true {
			ogTypeName := strings.Replace(attr, ogTypePrefix, "", 1)
			if ogValue, exists := s.Attr(ogValueAttributeName); exists == true {
				og[ogTypeName] = ogValue
			}
		}
	})

	return og, nil
}
