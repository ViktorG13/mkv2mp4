package main

// All Imports.
import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
) // End imports.

// Instance All Components.
var (
	// App Instance.
	application = app.New()
	// Main Window Instance With The Title "Mkv To Mp4" in Brazilian Portuguese.
	mainWindow = application.NewWindow("Mkv Para Mp4")
	// Label to Show The Path of File, With The Text "Select A File" in Brazilian Portuguese.
	filePathUri = widget.NewLabel("Selecione um Arquivo.")
	// Button With The Text "Select" in Brazilian Portuguese And a Callback Func to "selectFile" on Clicked.
	fileSelectButton = widget.NewButton("Selecionar", func() { selectFile() })
	// Button With The Text "Convert" in Brazilian Portuguese And a Callback Func to "convertFile" on Clicked.
	converterButton = widget.NewButton("Converter", func() { convertFile() })
	// Label to Show The Status of Conversion, Started With Nil Text.
	statusMessage = widget.NewLabel("")
) // End "var".

// Function To Get a File.
func selectFile() {
	// Create a Dialog Window to Select a File.
	dialog.ShowFileOpen(func(file fyne.URIReadCloser, err error) {
		// Verify if Have Any Errors on Select The File.
		if err != nil {
			// If Have Any Error, Create a New Dialog Window Error And Show it With Error as Message.
			dialog.NewError(err, mainWindow).Show()
			return
		} // End "if".

		// Verify if The Path of Selected File is Empty.
		if file != nil {
			// If The Path isn't Empty, Setup the "filePathUri" And Close The File.
			filePathUri.SetText(file.URI().String())
			file.Close()
		} // End "if".
	},
		// Instance of The Parent of This Dialog.
		mainWindow)
} // End Func "selectFile".

// Function To Convert The Selected File.
func convertFile() {
	// Verify if The File Has be Selected.
	if filePathUri.Text == "Selecione um Arquivo." {
		// Created a Erro Dialog Window If file hasn't be Selected.
		dialog.NewError(errors.New("por favor selecione um arquivo"), mainWindow).Show()
		return
	} // End "if".

	// Get The Path of The File in The Text of "filePathUri" Component, And Formate it.
	path := strings.ReplaceAll(filePathUri.Text[7:], " ", "\\ ")

	// Setup The "statusMessage" To Show "Converting..." in Brazilian Portuguese.
	statusMessage.SetText("Convertendo...")

	// Create A Variable to Contains The Command To Convert The File.
	command := fmt.Sprintf("ffmpeg -i %s -y -s hd1080 -c copy ~/Downloads/output.mp4", path)
	// Call The Command To Convert The File And Verify Is has any error.
	if _, err := exec.Command("sh", "-c", command).Output(); err != nil {
		// If Have Any Error, Create a New Error Dialog Window And Show it With Message "Error While Converting" in Brazilian Portuguese.
		dialog.NewError(fmt.Errorf("erro durante a conversão: %s", err), mainWindow).Show()
		// Setup The "statusMessage" To Show "Error... in Brazilian Portuguese"
		statusMessage.SetText("Erro...")
		return
	} // End "if".

	// If Hasn't Any Error, Show a Dialog Window With The Title "Success" And Message "File Successfully Converted." in Brazilian Portuguese.
	dialog.ShowInformation("Sucesso", "Arquivo Convertido Com Sucesso.", mainWindow)
	// Setup The "statusMessage" With Message "Done." in Brazilian Portuguese.
	statusMessage.SetText("Pronto.")
} // End Func "convertFile".

// Main Function. All Instance are Configured here.
func main() {
	// Set Main Window Default Size.
	mainWindow.Resize(fyne.NewSize(860, 640))
	// Set Main Window Content.
	mainWindow.SetContent(
		// Inserte in Main Window a Vertical Box.
		container.NewVBox(
			// Inserte in Vertical Box a Horizontal Box.
			container.NewHBox(
				widget.NewFileIcon(nil),
				filePathUri,
				fileSelectButton,
			), // End "NewVBox".
			// Insert in Vertical Box The "statusMessage".
			statusMessage,
			// Insert in Vertical Box The "converterButton."
			converterButton,
		), // End "NewHBox".
	) // End "SetContent".

	// Execute The Main Window And Show it.
	mainWindow.ShowAndRun()
} // End "main".
