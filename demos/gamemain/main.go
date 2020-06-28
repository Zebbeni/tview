// Demo code for the Flex primitive.
package main

import (
	"strconv"

	"github.com/Zebbeni/tview"
	"github.com/gdamore/tcell"
)

var (
	topMenuItemsText = []string{
		"ESC - Quit",
		"H - Help",
		"P - Previous",
	}
	mainMenuItemsText = []string{
		"1 - Start New Game",
		"2 - Open Saved Game",
	}
	graphicText = "                             ______                              \n" +
		"                   _____/^\\/ :::: \\/^\\_____                    \n" +
		"                  /___|___|{______}|___|___\\                   \n" +
		"                    \\___[____|__| ___]___/                     \n" +
		"                           \\______/                            \n" +
		"             ◿○□○◺           )__(                              \n" +
		"                             \\__/                              \n" +
		"        ◿○□○◺         .-------)(------.                        \n" +
		"                _.---'        \\/       '----.                  \n" +
		"            _.-'              \\/             '--.              \n" +
		"         ,-'                                     '-.           \n" +
		" ________  __             __  ___                   '  ______  \n" +
		"/   ____/_/  \\__ _____ __/  \\_\\__\\   ____ _______     /  __  \\ \n" +
		"\\____   \\\\_    _\\\\_ _ \\\\__   _\\\\  \\ / __ \\\\  __  \\    \\      / \n" +
		" \\   \\   \\_ \\  \\_ /  \\ \\_ \\  \\_ \\  \\\\   \\ \\\\   \\  \\_   /  __  \\\n" +
		"  \\_______/  \\__/ \\______/ \\__/  (__\\\\____/ \\__ \\__/   \\______/\n"
)

func main() {
	setDefaultStyles()

	app := tview.NewApplication()
	topMenu := createTopMenu()
	graphic := createGraphic()
	mainMenu := createMainMenu()
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(topMenu, 3, 0, false).
		AddItem(graphic, 0, 1, false).
		AddItem(mainMenu, 10, 1, false)
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

func createGraphic() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow)
	flex.AddItem(tview.NewBox(), 0, 2, false)
	flex.AddItem(tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(graphicText), 16, 1, false)
	flex.AddItem(tview.NewBox(), 0, 1, false)
	return flex
}

func createMenuItem(text string) *tview.TextView {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}

func createMainMenu() *tview.Flex {
	menu := tview.NewFlex().SetDirection(tview.FlexRow)
	menuTitle := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Main Menu")
	menu.AddItem(menuTitle, 1, 1, false)

	mainMenuWidth := 0

	menu.AddItem(tview.NewBox(), 1, 1, false) // spacer
	for _, menuItemText := range mainMenuItemsText {
		menu.AddItem(tview.NewTextView().SetTextAlign(tview.AlignLeft).SetText(menuItemText), 1, 1, false)
		if len(menuItemText) > mainMenuWidth {
			mainMenuWidth = len(menuItemText)
		}
	}
	menu.AddItem(tview.NewBox(), 1, 1, false) // spacer

	selector := createSelector()
	menu.AddItem(selector, 3, 1, false)
	menu.AddItem(tview.NewBox(), 1, 1, false) // spacer

	flex := tview.NewFlex()
	flex.AddItem(tview.NewBox(), 0, 1, false)
	flex.AddItem(menu, mainMenuWidth+2, 1, false)
	flex.AddItem(tview.NewBox(), 0, 1, false)
	return flex
}

func createSelector() *tview.Flex {
	flex := tview.NewFlex()

	text := "\nSelect (1-" + strconv.Itoa(len(mainMenuItemsText)) + "): "
	textItem := tview.NewTextView().
		SetTextAlign(tview.AlignRight).
		SetText(text)
	flex.AddItem(textItem, 0, 1, false)

	input := tview.NewInputField().
		SetPlaceholder("1").
		SetBorder(true)
	flex.AddItem(input, 5, 1, false)

	return flex
}

func setDefaultStyles() {
	tview.Styles = tview.Theme{
		PrimitiveBackgroundColor:    tcell.ColorDefault,
		ContrastBackgroundColor:     tcell.ColorDefault,
		MoreContrastBackgroundColor: tcell.ColorDefault,
		BorderColor:                 tcell.ColorDefault,
		TitleColor:                  tcell.ColorDefault,
		GraphicsColor:               tcell.ColorDefault,
		PrimaryTextColor:            tcell.ColorDefault,
		SecondaryTextColor:          tcell.ColorDefault,
		TertiaryTextColor:           tcell.ColorDefault,
		InverseTextColor:            tcell.ColorDefault,
		ContrastSecondaryTextColor:  tcell.ColorDefault,
	}

}
