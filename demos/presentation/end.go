package main

import (
	"fmt"

	"github.com/Zebbeni/tview"
	"github.com/gdamore/tcell"
)

// End shows the final slide.
func End(nextSlide func()) (title string, content tview.Primitive) {
	textView := tview.NewTextView().SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/Zebbeni/tview"
	fmt.Fprint(textView, url)
	return "End", Center(len(url), 1, textView)
}
