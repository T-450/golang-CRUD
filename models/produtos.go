package models

import (
	"fmt"
	"strconv"

	"github.com/edward-teixeira/src/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

// BuscaTodosOsProdutos retorna []Produto
func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	queryResult, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for queryResult.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := queryResult.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	query, err := db.Prepare("insert into produtos (nome, descricao , preco , quantidade) values ($1, $2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}
	query.Exec(nome, descricao, preco, quantidade)
	fmt.Println(query)
	defer db.Close()
}

func BuscarProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	convertedId, err := strconv.Atoi(id)
	if err != nil {
		panic(err.Error())
	}
	query, err := db.Query("select * from produtos where id=$1", convertedId)

	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err := query.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}
	defer db.Close()
	return p
}

func DeleteProduto(id string) {
	db := db.ConectaComBancoDeDados()
	query, err := db.Prepare("DELETE FROM produtos WHERE id=$1;")

	if err != nil {
		panic(err.Error())
	}

	query.Exec(id)
	fmt.Println("Deleteando produto => ", query)
	defer db.Close()

}

func EditarProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	query, err := db.Prepare("UPDATE produtos SET nome = $1, descricao = $2, preco = $3, quantidade = $4 WHERE id=$5;")

	if err != nil {
		panic(err.Error())
	}

	query2, err := query.Exec(nome, descricao, preco, quantidade, id)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(query2)
	defer db.Close()
}
