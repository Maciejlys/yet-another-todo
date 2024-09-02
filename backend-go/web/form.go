package web

import "encoding/gob"

func init() {
	gob.Register(CreateTodoForm{})
	gob.Register(FormErrors{})
}

type FormErrors map[string]string

type CreateTodoForm struct {
	Task string
	Done string

	Errors FormErrors
}

func (e FormErrors) String() (result string) {
	for k := range e {
		result += e[k] + "\n"
	}
	return result
}

func (f *CreateTodoForm) Validate() bool {
	f.Errors = FormErrors{}

	if f.Task == "" {
		f.Errors["Task"] = "Please enter a task."
	}
	if f.Done == "" {
		f.Errors["Done"] = "Please enter a done."
	}

	return len(f.Errors) == 0
}
