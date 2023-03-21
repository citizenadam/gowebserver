package controllers

import (
	"html/template"
	"net/http"

	"views"
)

type Static struct {
	Template views.Template
}

func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "can it be free?",
			Answer:   "Yes there is a free version.",
		},
		{
			Question: "What are the hours?",
			Answer:   "super support 24 7",
		},
		{
			Question: "how do i contact support?",
			Answer:   "call or email us.",
		},
		{
			Question: "Where is the office?",
			Answer:   "OnMoonbase12",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
