package main

import (
	"github.com/andlabs/ui"
)

// Example showing how to update the UI using the QueueMain function
// especially if the update is coming from another goroutine
//
// see QueueMain in 'main.go' for detailed description

var countLabel *ui.Label
var count int
var inputMap map[string]*ui.Form
var inputMapLabels map[string] *ui.Entry
var inputLabels = []string{"Ширина", "Высота", "Минимум", "Максимум", "Точность"}

func newInputForm(label string) (*ui.Form, *ui.Entry) {
	inputForm := ui.NewForm()
	inputForm.SetPadded(true)
	message := ui.NewEntry()
	message.SetText("")
	inputForm.Append(label, message, false)
	return inputForm, message
}


func setupUI() {
	mainWindow := ui.NewWindow("libui Updating UI", 250, 400, true)
	mainWindow.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		mainWindow.Destroy()
		return true
	})

	vbContainer := ui.NewVerticalBox()
	vbContainer.SetPadded(true)

	inputGroup := ui.NewGroup("Вход")
	inputGroup.SetMargined(true)

	vbInput := ui.NewVerticalBox()
	vbInput.SetPadded(true)

	showMessageButton := ui.NewButton("Сгенерировать")

	for _, name := range inputLabels {
		ii, ee := newInputForm(name)
		inputMap[name] = ii
		inputMapLabels[name] = ee
		vbInput.Append(ii, false)
	}
	vbInput.Append(showMessageButton, false)

	inputGroup.SetChild(vbInput)

	vbContainer.Append(inputGroup, false)



	outputGroup := ui.NewGroup("Выход")
	outputGroup.SetMargined(true)

	vbOutput := ui.NewVerticalBox()
	vbOutput.SetPadded(true)
	outputTable := ui.NewEntry()
	outputTable.SetReadOnly(true)

	vbOutput.Append(outputTable, false)

	outputGroup.SetChild(vbOutput)
	vbContainer.Append(outputGroup, false)

	showMessageButton.OnClicked(func(*ui.Button) {
		// check input
		width, err := convertAndCheckInt(inputMapLabels["Ширина"].Text())
		if err != nil {
			ui.MsgBoxError(mainWindow,
				"Ошибка!",
				err.Error())
		}

		height, err := convertAndCheckInt(inputMapLabels["Высота"].Text())
		if err != nil {
			ui.MsgBoxError(mainWindow,
				"Ошибка!",
				err.Error())
		}

		min, err := convertAndCheckFloat(inputMapLabels["Минимум"].Text())
		if err != nil {
			ui.MsgBoxError(mainWindow,
				"Ошибка!",
				err.Error())
		}

		max, err := convertAndCheckFloat(inputMapLabels["Максимум"].Text())
		if err != nil {
			ui.MsgBoxError(mainWindow,
				"Ошибка!",
				err.Error())
		}
		precision, err := convertAndCheckInt(inputMapLabels["Точность"].Text())
		if err != nil {
			ui.MsgBoxError(mainWindow,
				"Ошибка!",
				err.Error())
		}
		//generate table based on input
		table := generateRandomTable(height, width, min, max, precision)
		tableText := sliceToString(table, precision)
		outputTable.SetText(tableText)
	})

	mainWindow.SetChild(vbContainer)
	mainWindow.Show()
}


func main() {
	inputMap = make(map[string]*ui.Form)
	inputMapLabels = make(map[string]*ui.Entry)
	ui.Main(setupUI)
}