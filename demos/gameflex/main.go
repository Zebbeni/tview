// Demo code for the Flex primitive.
package main

import (
	"github.com/Zebbeni/tview"
)

var (
	attributes = map[string]string{
		"Name": "Casey Jones",
		"Age":  "24",
		"Sex":  "M",
		"Biog": "New York City, where Jones is from, was overrun by criminals. Enticed by examples from TV, Casey decided to do something about the crime in the streets. After equipping himself with a hockey mask and various sports clubs, he began his vigilante campaign.",
	}
)

func main() {
	maxAttributeLength := 0
	for k := range attributes {
		if len(k) > maxAttributeLength {
			maxAttributeLength = len(k)
		}
	}

	app := tview.NewApplication()
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	for k := range attributes {
		// attributeRow :=
		attributeRow := tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(tview.NewBox().SetBorder(true).SetTitle(k), maxAttributeLength+2+1, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Value"), 0, 1, false)
		flex.AddItem(attributeRow, 0, 1, false)
	}

	flex.SetBorder(true).SetTitle("Info")

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
