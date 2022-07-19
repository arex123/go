package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(500,500))
	var st string= "hello guys i am aditya";
	text := canvas.NewText(st, color.White)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Italic: true}
	w.SetContent(container.NewVBox(
		text,
		// widget.NewButton("Hi!", func() {
		// 	hello.SetText("Welcome :)")
		// }),
	))

	w.ShowAndRun()
}


// // SearchContent returns the results for a given query
// func SearchContent(input string) []Page {
// 	pages := []Page{}
//    ctx := context.Background()
// 	// Search for a page in the database using multi match query
// 	q := elastic.NewMultiMatchQuery(input, "title", "description", "body", "url").
// 	 Type("most_fields").
// 	 Fuzziness("2")
// 	result, err := client.Search().
// 	 Index(indexName).
// 	 Query(q).
// 	 From(0).Size(20).
// 	 Sort("_score", false).
// 	 Do(ctx)
// 	if err != nil {
// 	 log.Fatal(err)
// 	}
//    var ttyp Page
// 	for _, page := range result.Each(reflect.TypeOf(ttyp)) {
// 	 p := page.(Page)
// 	 pages = append(pages, p)
// 	}
//    return pages
//    }


// // Page struct to store in database
// type Page struct {
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// 	URL         string `json:"url"`
//    }