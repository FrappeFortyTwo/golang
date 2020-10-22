package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"io/ioutil"
)

import (
	"log"
	s "strings"
)

var editorData *walk.TextEdit

var saveMode = walk.NewMutableCondition()
var editorWindow *walk.MainWindow

var title string = ""

func main() {
	MustRegisterCondition("saveMode", saveMode)

	if err := (MainWindow{
		AssignTo: &editorWindow,

		Title:  "Untitled - Gopad",
		Layout: VBox{},
		MenuItems: []MenuItem{
			Menu{
				Text: "File",
				Items: []MenuItem{
					Action{
						Text: "New",
						OnTriggered: func() {
							editorData.SetTextSelection(4, 6)
						},
					},
					Action{
						Text: "Open",
						OnTriggered: func() {
							files, err := ioutil.ReadDir(".")
							if err != nil {
								log.Fatal(err)
							}

							for _, file := range files {
								fmt.Println(file.Name())
							}

						},
					},
					Separator{},
					Action{
						Text:    "Save",
						Enabled: Bind("saveMode"),

						OnTriggered: func() {
							ioutil.WriteFile(title, []byte(editorData.Text()), 0666)
						},
					},

					Action{
						Text: "Save as",
						OnTriggered: func() {

							saveAs(editorWindow)

						},
					},
					Separator{},
					Action{
						Text: "Exit",
						OnTriggered: func() {
							editorWindow.Close()
						},
					},
				},
			},
			Menu{
				Text: "Edit",
				Items: []MenuItem{
					Action{
						Text: "Find",
						OnTriggered: func() {
							//editorData.SetTextSelection(4, 6)

							findReplace(editorWindow)
						},
					},
				},
			},
		},
		ToolBar: ToolBar{
			Accessibility:      Accessibility{},
			Persistent:         false,
			RightToLeftReading: false,
			ButtonStyle:        ToolBarButtonImageOnly,

			Items: []MenuItem{
				Action{
					Text:  "Special",
					Image: "open.png",
					OnTriggered: func() {
						fmt.Println("yoo hoo")
					},
				},
			},
			MaxTextRows: 0,
			Orientation: 0,
		},
		Children: []Widget{
			VSpacer{
				GreedyLocallyOnly: false,
				Size:              30,
			},
			HSplitter{
				Accessibility: Accessibility{},
				Children: []Widget{
					TreeView{
						Accessibility: Accessibility{},
						Background:    SolidColorBrush{Color: walk.RGB(46, 46, 46)},
					},
					TextEdit{
						AssignTo: &editorData,
						Font: Font{
							Family:    "",
							PointSize: 30,
						},
						VScroll: true,
					},
				},
				DataBinder:  DataBinder{},
				AssignTo:    nil,
				HandleWidth: 0,
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	editorWindow.Run()

}

func saveAs(owner walk.Form) (int, error) {

	var dlg *walk.Dialog
	var okPB *walk.PushButton
	var fileName *walk.LineEdit

	return Dialog{
		AssignTo:      &dlg,
		Title:         "Save As",
		DefaultButton: &okPB,
		MinSize:       Size{Width: 300, Height: 100},
		Layout:        VBox{},
		Children: []Widget{
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "File Name :-",
					},
					LineEdit{
						AssignTo: &fileName,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &okPB,
						Text:     "OK",
						OnClicked: func() {

							// Store Value of Filename Field as 'title'
							title = fileName.Text()

							// Write file to disk named 'title'
							ioutil.WriteFile(title, []byte(editorData.Text()), 0666)

							// Enable Save
							saveMode.SetSatisfied(true)

							// Set 'editorWindow'
							owner.SetTitle(title + " - Gopad")

							dlg.Accept()

						},
					},
				},
			},
		},
	}.Run(owner)
}

func findReplace(owner walk.Form) (int, error) {

	var dlg *walk.Dialog
	var okPB *walk.PushButton
	var findStr *walk.LineEdit

	return Dialog{
		AssignTo: &dlg,
		Title:    "Find & Replace",
		//DefaultButton: &okPB,
		MinSize: Size{Width: 400, Height: 300},
		Layout:  VBox{},
		Children: []Widget{
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "Find :-",
					},
					LineEdit{
						AssignTo: &findStr,
					},
				},
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					HSpacer{},
					PushButton{
						AssignTo: &okPB,
						Text:     "Find",
						OnClicked: func() {

							// Check if String Exists :
							if s.Contains(editorData.Text(), findStr.Text()) == false {
								walk.MsgBox(owner, "FIND & REPLACE", "NO MATCHES FOUND!\nIf you dont know what youre looking for you will never find it !", walk.MsgBoxIconInformation)
							}

							editorData.SetTextSelection(s.Index(editorData.Text(), findStr.Text()), s.Index(editorData.Text(), findStr.Text())+len(findStr.Text()))
							dlg.Accept()
						},
					},
				},
			},
		},
	}.Run(owner)
}
