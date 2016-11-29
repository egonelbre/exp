package main

import (
	"fmt"

	"github.com/egonelbre/exp/dos/ui"
)

var (
	Global *ui.Screen
)

func main() {
	Global = &ui.Screen{}
	Global.Start()
	// defer Global.Stop()

	ShowRecord(&MainMenu{})
}

func ShowRecord(record ui.Record) {
	form := Global.NewForm()
	form.Record = record

	// TODO: choose sizes more intelligently
	form.BoundsRect = ui.Rect{0, 0, 30, 30}
	if len(Global.Forms) > 0 {
		last := Global.Forms[len(Global.Forms)-1]
		form.BoundsRect.Left = last.BoundsRect.Left + last.BoundsRect.Width + 1
	}

	form.Show()
}

type MainMenu struct {
	Patient   ui.Button
	PatientID ui.Input

	Exit ui.Button
}

func (menu *MainMenu) Caption() string {
	return "Main Menu"
}

func (menu *MainMenu) Load() error {
	menu.Patient = ui.Button{"Open Patient", menu.HandlePatient}
	menu.PatientID = ui.Input{"> %v", "0000"}
	menu.Exit = ui.Button{"Exit", menu.HandleExit}

	return nil
}
func (menu *MainMenu) Save() error { return nil }

func (menu *MainMenu) HandlePatient(screen *ui.Form) {
	info := &PatientInfo{}
	info.ID = menu.PatientID.Text

	ShowRecord(info)
}

func (menu *MainMenu) HandleExit(screen *ui.Form) {
	screen.Close = true
}

type PatientInfo struct {
	ID string

	Name ui.Input
	DOB  ui.Input
}

func (patient *PatientInfo) Widgets() []ui.Widget {
	return []ui.Widget{
		&patient.Name,
		&patient.DOB,
	}
}

func (patient *PatientInfo) Caption() string {
	return fmt.Sprintf("%s : %s", patient.ID, patient.Name.Text)
}

func (patient *PatientInfo) Load() error {
	// TODO: load from DB, possibly with dynamic mapping
	patient.Name = ui.Input{"Name: %v", "John Smith"}
	patient.DOB = ui.Input{" DOB: %v", "1962 02 21"}

	return nil
}

func (patient *PatientInfo) Save() error {
	// TODO: save to DB
	return nil
}
