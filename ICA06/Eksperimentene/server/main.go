package main

import (


	"net/http"
	"html/template"
	"path"

)

func main() {


	http.HandleFunc("/", foo)
	http.HandleFunc("/speech", redirect)
	http.ListenAndServe(":8080", nil)

}

func foo(w http.ResponseWriter, r *http.Request) {


	fp3 := path.Join("templates","index3.html")
	tmpl3, err := template.ParseFiles(fp3)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl3.Execute(w, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func redirect(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	text := values.Get("text")
	sprok := values.Get("sprok")
	if sprok == "Norsk" {
		http.Redirect(w, r, "http://158.37.63.148:8080/speech?text="+text+"&voice=no", 301)

	} else if sprok == "Engelsk" {
		http.Redirect(w, r, "http://158.37.63.148:8080/speech?text=" + text + "&voice=en", 301)
}	  else if sprok == "Russisk" {
		http.Redirect(w, r, "http://158.37.63.148:8080/speech?text=" + text + "&voice=sr", 301)
	}



}
