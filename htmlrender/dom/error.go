package dom

type Error struct{ Err error }

func (err Error) Render(w Writer) {
	w.Open("div")
	w.Attr("class", "template-compilation-error")
	w.Text("Compilation error " + err.Err.Error())
	w.Close("div")
}
