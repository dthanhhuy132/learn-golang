package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// RenderTemplate renders template using html
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing templates______:", err)
		return
	}
}

var tc = make(map[string]*template.Template) //template cache

func RenderTemplate(w http.ResponseWriter, t string) { // không return về cái gì hết
	var tmpl *template.Template
	var err error

	// check to see if we already hae the template in our cache
	_, inMap := tc[t]

	fmt.Println("inMap", inMap)

	if !inMap {
		// need to create the template
		log.Println("create template and adding to cache")
		err = createTemplateCache(t)

		if err != nil {
			log.Println("err", err)
		}
	} else {
		log.Println(("using cached template"))
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("err", err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t), "./templates/base.html.tmpl",
	}

	// fmt.Println("in cho nay ra ne_____________", fmt.Sprintf("./templates/%s", t))

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache (map)
	tc[t] = tmpl

	return nil
}
