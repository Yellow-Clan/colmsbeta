package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewRol struct {
	Name    string
	IsEdit  bool
	Data    models.Rol
	Widgets []models.Rol
}

var tmpld = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/rol/index.html", "web/rol/form.html"))

func RolList(w http.ResponseWriter, req *http.Request) {

	rol := models.Rol{}

	roles, _ := rol.FindAll(cfig.DB)

	for _, lis := range roles {
		fmt.Println(lis.ToString())
		fmt.Println("Personas: ", len(lis.Personas))
		if len(lis.Personas) > 0 {
			for _, d := range lis.Personas {
				fmt.Println(d.ToString())
				fmt.Println("=============================")
			}
		}
		fmt.Println("--------------------")
	}

	// Create
	//cfig.DB.Create(&models.Docente{Name: "Juan", City: "Juliaca"})
	lis := []models.Rol{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Printf("lis: %v", lis)
	data := ViewRol{
		Name:    "Rol",
		Widgets: lis,
	}

	err := tmpld.ExecuteTemplate(w, "rol/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RolForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Rol
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
		d.Nombre = r.FormValue("nombre")
		d.Codigo = r.FormValue("codigo")

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
		http.Redirect(w, r, "/rol/index", 301)
	}

	data := ViewRol{
		Name:   "Rol",
		Data:   d,
		IsEdit: IsEdit,
	}

	err := tmpld.ExecuteTemplate(w, "rol/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RolDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]//log.Printf("del id=: %v", id)
	var d models.Rol
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}
	http.Redirect(w, r, "rol/index", 301)
}
