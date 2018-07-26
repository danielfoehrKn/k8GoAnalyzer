package tutorial

import (
	"io/ioutil"
	"net/http"
	"log"
	"html/template"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	// warum brauche ich hier den Pointer mit & ? -> funktioniert auch ohne
	//p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	//p1.save()
	//p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}


func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Webserver handler
// with Fprintf write the string to the response writer to return to client
//func viewHandler(w http.ResponseWriter, r *http.Request) {
//	title := r.URL.Path[len("/view/"):]
//	p, _ := loadPage(title)
//	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
//}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	//parse file with template
	renderTemplate(w, "edit", p)
}


func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}


func renderTemplate(w http.ResponseWriter, tmplName string,  p *Page) {
	t, _ := template.ParseFiles(tmplName + ".html")
	// execute template (reading in page pointer p , filling in vars in template and then writing html to the response writer
	t.Execute(w, p)
}