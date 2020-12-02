package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewCurso struct {
	Name     string
	IsEdit   bool
	Data     models.Curso
	Widgets  []models.Curso
	Docentes []models.Docente
}

var tmplc = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/curso/index.html", "web/curso/form.html"))

func CursoList(w http.ResponseWriter, req *http.Request) {

	lis := []models.Curso{}
	if err := cfig.DB.Preload("Docente").Find(&lis).Error; err != nil { // Preload("Alumno") carga los objetos Alumno relacionado
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := ViewCurso{
		Name:    "Curso",
		Widgets: lis,
	}
	err := tmplc.ExecuteTemplate(w, "curso/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CursoForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Curso
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	docente := models.Docente{}
	docentes, _ := docente.GetAll(cfig.DB) // para mostrar los alumnos en un combobox

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		d.Curso = r.FormValue("curso")
		d.Semestre = r.FormValue("semestre")
		//n, err := strconv.Atoi(r.FormValue("alumno_id"))
		//if err != nil {
		//	log.Printf("Invalid ID: %v - %v\n", n, err)
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		d.DocenteId = r.FormValue("docente_id") //n
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
		http.Redirect(w, r, "/curso/index", 301)
	}

	data := ViewCurso{
		Name:     "Curso",
		Data:     d,
		IsEdit:   IsEdit,
		Docentes: docentes,
	}

	err := tmplc.ExecuteTemplate(w, "curso/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CursoDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var d models.Curso
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		//log.Printf("No save  %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}

	http.Redirect(w, r, "/curso/index", 301)
}
