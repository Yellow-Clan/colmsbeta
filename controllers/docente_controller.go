package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewDocente struct {
	Nombre  string
	IsEdit  bool
	Data    models.Docente
	Widgets []models.Docente
}

var tmplu = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/docente/index.html", "web/docente/form.html"))

func DocenteList(w http.ResponseWriter, req *http.Request) {
	// Create
	//cfig.DB.Create(&models.Docente{Nombre: "Juan", City: "Juliaca"})
	lis := []models.Docente{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Printf("lis: %v", lis)
	data := ViewDocente{
		Nombre:  "Docente",
		Widgets: lis,
	}

	err := tmplu.ExecuteTemplate(w, "docente/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DocenteForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	Id_docente := r.URL.Query().Get("Id_docente") //mux.Vars(r)["Id_docente"]
	log.Printf("get Id_docente=: %v", Id_docente)
	var d models.Docente
	IsEdit := false
	if Id_docente != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "Id_docente = ?", Id_docente).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if r.Method == "POST" {
		log.Printf("POST Id_docente=: %v", Id_docente)
		d.Nombre = r.FormValue("nombre")
		d.Curso_acargo = r.FormValue("curso_acargo")
		if Id_docente != "" {
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
		Nombre: "Docente",
		Data:   d,
		IsEdit: IsEdit,
	}

	err := tmplu.ExecuteTemplate(w, "docente/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DocenteDel(w http.ResponseWriter, r *http.Request) {
	Id_docente := r.URL.Query().Get("Id_docente") //mux.Vars(r)["Id_docente"]//log.Printf("del Id_docente=: %v", Id_docente)
	var d models.Docente
	if err := cfig.DB.First(&d, "Id_docente = ?", Id_docente).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}
	http.Redirect(w, r, "/docente/index", 301)
}
