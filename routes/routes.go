package routes

import (
	"github.com/202lp1/colms/controllers"
	"github.com/202lp1/colms/mid"
	"github.com/gorilla/mux" //gin or chi u other
)

func RoutesMain(r *mux.Router) {

	r.HandleFunc("/", controllers.Home).Methods("GET")

	r.HandleFunc("/item/index", controllers.ItemList).Methods("GET")

	r.HandleFunc("/employee/index", controllers.EmployeeList).Methods("GET")
	r.HandleFunc("/employee/form", controllers.EmployeeForm).Methods("GET", "POST")
	r.HandleFunc("/employee/delete", controllers.EmployeeDel).Methods("GET")

	//Multi-Middleware per-route usage example
	r.HandleFunc("/alumno/index", mid.Chain(controllers.AlumnoList,
		mid.AuthRequired(), mid.Logging())).Methods("GET")
	r.HandleFunc("/alumno/form", mid.Chain(controllers.AlumnoForm,
		mid.AuthRequired(), mid.Logging())).Methods("GET", "POST")
	r.HandleFunc("/alumno/delete", mid.Chain(controllers.AlumnoDel,
		mid.AuthRequired(), mid.Logging())).Methods("GET")

	r.HandleFunc("/matricula/index", controllers.MatriculaList).Methods("GET")
	r.HandleFunc("/matricula/form", controllers.MatriculaForm).Methods("GET", "POST")
	r.HandleFunc("/matricula/delete", controllers.MatriculaDel).Methods("GET")

	r.HandleFunc("/alumnocurso/index", controllers.AlumnocursoList).Methods("GET")
	r.HandleFunc("/alumnocurso/form", controllers.AlumnocursoForm).Methods("GET", "POST")
	r.HandleFunc("/alumnocurso/delete", controllers.AlumnocursoDel).Methods("GET")

	r.HandleFunc("/user/index", controllers.UserList).Methods("GET")
	r.HandleFunc("/user/form", controllers.UserForm).Methods("GET", "POST")
	r.HandleFunc("/user/delete", controllers.UserDel).Methods("GET")

	r.HandleFunc("/rol/index", controllers.RolList).Methods("GET")
	r.HandleFunc("/rol/form", controllers.RolForm).Methods("GET", "POST")
	r.HandleFunc("/rol/delete", controllers.RolDel).Methods("GET")

	r.HandleFunc("/persona/index", controllers.PersonaList).Methods("GET")
	r.HandleFunc("/persona/form", controllers.PersonaForm).Methods("GET", "POST")
	r.HandleFunc("/persona/delete", controllers.PersonaDel).Methods("GET")

}
