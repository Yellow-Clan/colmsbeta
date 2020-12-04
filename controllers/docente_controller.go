package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewDocente struct {
	Name    string
	IsEdit  bool
	Data    models.Docente
	Widgets []models.Docente
}

var tmpld = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/docente/index.html", "web/docente/form.html"))

func DocenteList(w http.ResponseWriter, req *http.Request) {

	docente := models.Docente{}

	docentes, _ := docente.FindAll(cfig.DB)

	for _, lis := range docentes {
		fmt.Println(lis.ToString())
		fmt.Println("Cursos: ", len(lis.Cursos))
		if len(lis.Cursos) > 0 {
			for _, d := range lis.Cursos {
				fmt.Println(d.ToString())
				fmt.Println("=============================")
			}
		}
		fmt.Println("--------------------")
	}

	// Create
	//cfig.DB.Create(&models.Docente{Name: "Juan", City: "Juliaca"})
	lis := []models.Docente{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Printf("lis: %v", lis)
	data := ViewDocente{
		Name:    "Docente",
		Widgets: lis,
	}

	err := tmpld.ExecuteTemplate(w, "docente/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DocenteForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Docente
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		d.Nombres = r.FormValue("nombres")
		d.Codigo = r.FormValue("codigo")
		d.Email = r.FormValue("email")
		if id != "" {
			if err := cfig.DB.Save(&d).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}

		} else {
			if err := cfig.DB.Create(&d).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return //err
			}
		}
		http.Redirect(w, r, "/docente/index", 301)
	}

	data := ViewDocente{
		Name:   "Docente",
		Data:   d,
		IsEdit: IsEdit,
	}

	err := tmpld.ExecuteTemplate(w, "docente/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DocenteDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]//log.Printf("del id=: %v", id)
	var d models.Docente
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}
	http.Redirect(w, r, "docente/index", 301)
}
