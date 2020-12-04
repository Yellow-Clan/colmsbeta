package controllers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/models"
)

type ViewAdministrador struct {
	Name    string
	IsEdit  bool
	Data    models.Administrador
	Widgets []models.Administrador
}

var tmplad = template.Must(template.New("foo").Funcs(cfig.FuncMap).ParseFiles("web/Header.tmpl", "web/Menu.tmpl", "web/Footer.tmpl", "web/administrador/index.html", "web/administrador/form.html"))

func AdministradorList(w http.ResponseWriter, req *http.Request) {

	administrador := models.Administrador{}

	administradores, _ := administrador.FindAll(cfig.DB)

	for _, lis := range administradores {
		fmt.Println(lis.ToString())
		fmt.Println("Ingresoadmis: ", len(lis.Ingresoadmis))
		if len(lis.Ingresoadmis) > 0 {
			for _, d := range lis.Ingresoadmis {
				fmt.Println(d.ToString())
				fmt.Println("=============================")
			}
		}
		fmt.Println("--------------------")
	}

	// Create
	//cfig.DB.Create(&models.Administrador{Name: "Juan", City: "Juliaca"})
	lis := []models.Administrador{}
	if err := cfig.DB.Find(&lis).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//log.Printf("lis: %v", lis)
	data := ViewAdministrador{
		Name:    "Administrador",
		Widgets: lis,
	}

	err := tmplad.ExecuteTemplate(w, "administrador/indexPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AdministradorForm(w http.ResponseWriter, r *http.Request) {
	//log.Printf("r.Method= %v", r.Method)
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]
	log.Printf("get id=: %v", id)
	var d models.Administrador
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
		http.Redirect(w, r, "/administrador/index", 301)
	}

	data := ViewAdministrador{
		Name:   "Administrador",
		Data:   d,
		IsEdit: IsEdit,
	}

	err := tmplad.ExecuteTemplate(w, "administrador/formPage", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AdministradorDel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id") //mux.Vars(r)["id"]//log.Printf("del id=: %v", id)
	var d models.Administrador
	if err := cfig.DB.First(&d, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := cfig.DB.Unscoped().Delete(&d).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return //err
	}
	http.Redirect(w, r, "administrador/index", 301)
}
