package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/edward-teixeira/src/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {

	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")

		precoConvertidoFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		quantidadeParaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertidoFloat, quantidadeParaInt)
	}

	http.Redirect(w, r, "/", 301)
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeleteProduto(idDoProduto)

	http.Redirect(w, r, "/", 301)
}

func EditarProduto(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.BuscarProduto(id)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")

		precoConvertidoFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		quantidadeParaInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		IdConvertido, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		models.EditarProduto(IdConvertido, nome, descricao, precoConvertidoFloat, quantidadeParaInt)
		http.Redirect(w, r, "/", 301)
	}
}
