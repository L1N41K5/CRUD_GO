package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/API_GO/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "index", todosOsProdutos)

}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)

	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")

	models.DeletaProduto(idDoProduto)

	http.Redirect(w, r, "/", 301)

}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)

}

func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidoParaInt, err := strconv.Atoi(id)

		if err != nil {
			log.Println("erro na conversão do ID para int", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("erro na conversão do preco para Float64", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("erro na conversão da quantidade para int", err)
		}

		models.AtualizaProduto(idConvertidoParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)

	}

	http.Redirect(w, r, "/", 301)

}
