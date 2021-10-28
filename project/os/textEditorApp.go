/*
take template from -> "https://developer.fyne.io/started/firstapp"
*/

package main

import (
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var count int= 1
func textEditor() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(600,600))

	content:= container.NewVBox(
		container.NewHBox(
			widget.NewLabel("pep text editor"),
		),
	)

	content.Add(widget.NewButton("Add New File",func() {
		content.Add(widget.NewLabel("New File "+ strconv.Itoa(count)))
		count++
	}))


	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")
	// input.Resize(fyne.NewSize(800,400))

	saveBtn:=widget.NewButton("Save text file ",func() {
		saveFileDialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData:=[]byte(input.Text)
				uc.Write(textData)
			},w)

			saveFileDialog.SetFileName("New File"+strconv.Itoa(count-1)+".txt")
			saveFileDialog.Show()
	})

	openBtn:=widget.NewButton("Open text file",func() {
		openFileDialog:=dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				ReadData ,_ := ioutil.ReadAll(r)

				output:=fyne.NewStaticResource("New File ",ReadData)

				viewData:= widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				w:=fyne.CurrentApp().NewWindow(
					string(output.StaticName))


					sveBtn:=widget.NewButton("save ",func() {
						saveDialog:=dialog.NewFileSave(
							func(uc fyne.URIWriteCloser, _ error) {
								text:=[]byte(input.Text)
								uc.Write(text)
							},w)
							saveDialog.SetFileName("updated_textFile.txt")
							saveDialog.Show()
					})

					w.SetContent(						
						container.NewVBox(
							sveBtn,
							container.NewScroll(viewData),
						),

					)

					w.Resize(fyne.NewSize(400,400))

					w.Show()

			},w)
			openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
			openFileDialog.Show()
	})


	w.SetContent(
		container.NewVBox(
			content,
			input,

			container.NewHBox(
				saveBtn,
				openBtn,
			),
		),
	)

	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,EditorContainer),)

	w.ShowAndRun()
}