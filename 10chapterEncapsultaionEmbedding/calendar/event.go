package calendar

type Event struct {
	title string
	Date
}

func (e *Event) SetTitle(title string) {
	e.title = title
}

func (e *Event) Title() string {
	return e.title
}
