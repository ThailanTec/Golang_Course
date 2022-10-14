package crud

import (
	"database/sql"
	"fmt"
	"log"

	/* Precisamos fazer esse import de forma anonima, mesmo que não seja aqui, que consumir os dados, precisamos dele aqui, importado de forma anonima, quando passamos o _, o golang entende que ele é dessa forma. */
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// Conexão com o banco de dados
	/* Passamos uma configuração dessa forma para acessar o banco, mas vamos complementar mais as informações, para ele entender os caracteres, entender questões de hora e que o banco é local*/
	stringConexao := "golang:Golang123#@/devbook?charset=utf8&parseTime=true&loc=Local"
	// Chamando o banco de dados: 1-Passamos o nome/tipo do banco de dados que estamos nos referindo, segundo a string de conexão com as informações do banco de dados.

	/* Para verificar se estamos realmente conectados dentro do banco, precisamos fazer uma validação chamada 'ping',  */
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println("Conexão aberta!")

	// Criando uma Query de busca no banco:

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
