package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão
)

func Conecta() (*sql.DB, error) {

	stringConexao := "golang:Golang123#@/devbook?charset=utf8&parseTime=true&loc=Local"
	//Abertura do banco

	/* Obs: Quando estamos trabalhando com a configuração do banco dessa forma, não podemos passar um log.Fatal, nem o defer, prq precisamos que ele chegue aberto até na chamada da aplicação
	e como nossa aplicação tem um retorno, precisamos validar a sua saida de erro e sucesso, dessa forma, podemos fazer como nós exemplos abaixo para consegui um resultado correto.
	*/
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil

}
