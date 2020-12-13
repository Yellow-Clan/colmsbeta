package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewAlumnocurso struct {
	Name    string
	IsEdit  bool
	Data    models.Alumnocurso
	Widgets []models.Alumnocurso
	Alumnos []models.Alumno
}

var tmplac = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/alumnocurso/index.html", "web/alumnocurso/form.html"))

func AlumnocursoList(w http.ResponseWriter, req *http.Request) {

	lis := []models.Alumnocurso{}
	if err := cfig.DB.Preload("Alumno").Find(&lis).Error; err != nil { // Preload("Alumno") carga los objetos Alumno relacionado
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := ViewAlumnocurso{
		Name:    "Alumnocurso",
		Widgets: lis,
	}
	err := tmplac.ExecuteTemplate(w, "alumnocurso/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AlumnocursoForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Alumnocurso
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	alumno := models.Alumno{}
	alumnos, _ := alumno.GetAll(cfig.DB) // para mostrar los alumnos en un combobox

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		d.Semestre = r.FormValue("semestre")
		//n, err := strconv.Atoi(r.FormValue("alumno_id"))
		//if err != nil {
		//	log.Printf("Invalid ID: %v - %v\n", n, err)
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		d.AlumnoId = r.FormValue("alumno_id") //n
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
		http.Redirect(w, r, "/alumnocurso/index", 301)
	}

	data := ViewAlumnocurso{
		Name:    "Alumnocurso",
		Data:    d,
		IsEdit:  IsEdit,
		Alumnos: alumnos,
	}

	err := tmplac.ExecuteTemplate(w, "alumnocurso/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AlumnocursoDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var d models.Alumnocurso
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		//log.Printf("No save  %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}

	http.Redirect(w, r, "/alumnocurso/index", 301)
}
