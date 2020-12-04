package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewIngresoadmi struct {
	Name            string
	IsEdit          bool
	Data            models.Ingresoadmi
	Widgets         []models.Ingresoadmi
	Administradores []models.Administrador
}

var tmplia = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/ingresoadmi/index.html", "web/ingresoadmi/form.html"))

func IngresoadmiList(w http.ResponseWriter, req *http.Request) {

	lis := []models.Ingresoadmi{}
	if err := cfig.DB.Preload("Administrador").Find(&lis).Error; err != nil { // Preload("Alumno") carga los objetos Alumno relacionado
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := ViewIngresoadmi{
		Name:    "Ingresoadmi",
		Widgets: lis,
	}
	err := tmplia.ExecuteTemplate(w, "ingresoadmi/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IngresoadmiForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Ingresoadmi
	IsEdit := false
	if id != "" {
		IsEdit = true
		if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	administrador := models.Administrador{}
	administradores, _ := administrador.GetAll(cfig.DB) // para mostrar los alumnos en un combobox

	if r.Method == "POST" {
		log.Printf("POST id=: %v", id)
		d.Semestre = r.FormValue("semestre")
		//n, err := strconv.Atoi(r.FormValue("alumno_id"))
		//if err != nil {
		//	log.Printf("Invalid ID: %v - %v\n", n, err)
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		d.AdministradorId = r.FormValue("administrador_id") //n
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
		http.Redirect(w, r, "/ingresoadmi/index", 301)
	}

	data := ViewIngresoadmi{
		Name:            "Ingresoadmi",
		Data:            d,
		IsEdit:          IsEdit,
		Administradores: administradores,
	}

	err := tmplia.ExecuteTemplate(w, "ingresoadmi/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IngresoadmiDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var d models.Ingresoadmi
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		//log.Printf("No save  %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}

	http.Redirect(w, r, "/ingresoadmi/index", 301)
}
