package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	asciiArt "ascii-art-web/ascii-art"
)

type Ascii struct {
	Style string `json:"style"`
	Text  string `json:"text"`
	Out   string `json:"out"`
}

var data Ascii

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println(r.URL, r.Method)
	if r.URL.Path != "/" || r.Method != "GET" {
		app.infoLog.Println("Not found", r.URL)
		app.notFound(w)
		return
	}

	tmpl, err := template.ParseFiles("ui/html/index.html")
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(w, err)
	}
	tmpl.Execute(w, data)
}

func (app *application) asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.badRequest(w)
	}

	app.infoLog.Println(r.URL, r.Method, r.Form)
	if r.URL.Path != "/ascii-art" {
		app.infoLog.Println("Not Found", r.URL)
		app.notFound(w)
		return
	}
	if r.Method != "POST" {
		app.infoLog.Println(" ", r.URL, r.Method)
		app.methodNotAllowed(w)
		return
	}
	text, textOk := r.PostForm["text"]
	style, styleOk := r.PostForm["style"]

	if !textOk || !styleOk {
		app.infoLog.Println("Bad request", r.URL, r.Method, r.Form)
		app.badRequest(w)
		return
	}

	out, err := asciiArt.GetStyled(strings.ReplaceAll(text[0], "\r\n", "\n"), style[0])
	fmt.Println(out)
	if err != nil {
		app.errorLog.Println(err)
		switch err.Error() {
		case "invalid banner type, avilable : standard, shadow, thinkertoy":
			app.badRequest(w)
			app.infoLog.Println("Bad request", r.URL, r.Method, r.Form)
			return
		case "invalid text":
			app.badRequest(w)
			app.infoLog.Println("invalid text", r.URL, r.Method, r.Form)
			return
		default:
			app.errorLog.Println(err)
			app.serverError(w, err)
			return
		}
	}
	data = Ascii{
		Text:  text[0],
		Style: style[0],
		Out:   out,
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
