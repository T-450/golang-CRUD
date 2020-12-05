package routes

import (
	"net/http"

	"github.com/edward-teixeira/src/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.DeleteProduto)
	http.HandleFunc("/edit", controllers.EditarProduto)
	http.HandleFunc("/update", controllers.UpdateProduto)
}
