// go run main.go -> to run
/* steps to make a go app
1. make a file with .go extension
2. if needed , go mod init name , it is needed when it is
   not done in the same root folder
3. after init write go mod tidy

*/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func gallaryApp(w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Hello")
	// w.Resize(fyne.NewSize(300,300))
	src_path:= "C:\\Users\\dell\\Pictures\\Saved Pictures";
	files, err := ioutil.ReadDir(src_path)
    if err != nil {
        log.Fatal(err)
    }
	
	tabs := container.NewAppTabs(
		// container.NewTabItem("image", canvas.NewImageFromFile(picsArr[0])),
	)
    for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir(){
			extension:= strings.Split(file.Name(),".")[1];
			if extension == "png" || extension == "jpeg" || extension=="jpg" {
				
				image:= canvas.NewImageFromFile(src_path+"\\"+file.Name());
				image.FillMode = canvas.ImageFillContain
				tabs.Append(container.NewTabItem(file.Name(),image))


			}
		}

    }

		
	// w.SetContent(tabs)
	// image := canvas.NewImageFromFile(picsArr[0])
	// w.SetContent(container.NewVBox(image,tabs))
	// w.SetContent(image)
	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,tabs),)
	
	w.Show()
}