// Demo code for the Table primitive.
package main

import (
	"github.com/Zebbeni/tview"
	"github.com/gdamore/tcell"
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
	app := tview.NewApplication()
	table := tview.NewTable().SetBorders(true)

	cols := 2
	rows := len(attributes)

	rowValues := make([][]string, 0, 4)
	for k, v := range attributes {
		rowValues = append(rowValues, []string{k, v})
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite
			if c == 0 {
				color = tcell.ColorYellow
			}

			cellValue := rowValues[r][c]
			table.SetCell(r, c,
				tview.NewTableCell(cellValue).
					SetTextColor(color).
					SetAlign(tview.AlignLeft).
					SetMaxWidth(20))
		}
	}
	table.Select(0, 0).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			table.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		table.SetSelectable(false, false)
	})
	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
