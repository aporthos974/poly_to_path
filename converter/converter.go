package converter

import (
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
	"log"
)

func Convert(svg string) string {
	document, err := gokogiri.ParseXml([]byte(svg))
	if err != nil {
		log.Fatalf("Unable to parse SVG : %s", err.Error())
	}
	defer document.Free()

	expression := xpath.Compile("//*[local-name() = 'polygon']")
	polygons, err := document.Search(expression)
	if err != nil {
		log.Fatalf("Unable to find path : %s", err.Error())
	}
	for _, polygon := range polygons {
		polygon.SetName("path")
	}
	return document.InnerHtml()
}
