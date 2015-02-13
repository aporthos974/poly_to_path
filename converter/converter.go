package converter

import (
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
	"log"
	"strings"
)

func Convert(svg string) string {
	document, err := gokogiri.ParseXml([]byte(svg))
	if err != nil {
		log.Fatalf("Unable to parse SVG : %s", err.Error())
	}
	defer document.Free()

	expression := xpath.Compile("//*[local-name() = 'polygon' or local-name() = 'polyline']")
	tags, err := document.Search(expression)
	if err != nil {
		log.Fatalf("Unable to find path : %s", err.Error())
	}
	for _, tag := range tags {
		dAttribut := tag.Attribute("points")
		dAttribut.SetName("d")

		points := "M"
		coordonates := strings.SplitN(dAttribut.Value(), " ", 2)
		points = points + coordonates[0] + " L" + coordonates[1]

		if tag.Name() == "polygon" {
			points = points + "Z"
		}

		dAttribut.SetValue(points)
		tag.SetName("path")
	}
	return document.InnerHtml()
}
