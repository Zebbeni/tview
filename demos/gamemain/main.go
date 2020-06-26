// Demo code for the Flex primitive.
package main

import (
	"github.com/Zebbeni/tview"
)

var (
	topMenuItemsText = []string{
		"ESC - Quit",
		"H - Help",
		"P - Previous",
	}
)

func main() {
	app := tview.NewApplication()

	topMenu := createTopMenu()
	mainMenu := createMainMenu()
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topMenu, 3, 0, false).
		AddItem(mainMenu, 0, 1, false)
	flex.SetBorder(true)
	flex.SetBorderPadding(0, 0, 1, 1)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func createTopMenu() *tview.Flex {
	topMenu := tview.NewFlex()
	topMenu.SetBorder(true)
	topMenu.AddItem(tview.NewBox(), 0, 1, false)

	for _, itemText := range topMenuItemsText {
		itemWidth := len(itemText) + 6 // add 4 for padding
		topMenu.AddItem(createMenuItem(itemText), itemWidth, 1, false)
	}

	topMenu.AddItem(tview.NewBox(), 0, 1, false)

	return topMenu
}

func createMenuItem(text string) *tview.TextView {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}

func createMainMenu() *tview.Box {
	return tview.NewBox().SetTitle("Bottom (5 rows)")
}
