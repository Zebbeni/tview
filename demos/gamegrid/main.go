// Demo code for the Grid primitive.
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
	grid := createGameGrid()
	if err := tview.NewApplication().SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func createGameGrid() *tview.Grid {
	maxAttributeWidth := 0
	for attribute := range attributes {
		if len(attribute) > maxAttributeWidth {
			maxAttributeWidth = len(attribute)
		}
	}

	grid := tview.NewGrid()
	grid.SetTitle("Info")
	grid.SetColumns(maxAttributeWidth+2, 50)

	row := 0
	for attribute, value := range attributes {
		attributeItem := tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(attribute + ":")
		attributeItem.SetRect(0, 0,10, 5)
		valueItem := tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(value).
			SetTextAlign(tview.AlignLeft).
			SetWordWrap(true)
		valueItem.SetRect(0, 0,14, 5)
		grid.AddItem(attributeItem, row, 0, 1, 1, 0, 0, false).
			AddItem(valueItem, row, 1, 1, 1, 0, 0, false)
		x, y, wid, hei := valueItem.GetRect()
		println(">>> x:", x, "y:", y, "wid:", wid, "hei:", hei)
		row++
	}
	return grid
}
