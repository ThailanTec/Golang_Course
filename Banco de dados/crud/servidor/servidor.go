package servidor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ThailanTec/banco-de-dados/crud/banco"
)

/* Criando User no banco:

Pacote ioutil serve para que ele consiga lê todos os dados que vão está no corpo da nossa requisição.

Para inserir os dados no banco utilizamos o PREPARE STATEMENT, para evitar SQL INJECTION e cria um comando de inserção ao banco.

*/

type user struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CriarUser(w http.ResponseWriter, r *http.Request) {

	corpoReq, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Não foi possivel vê corpo da requisição."))
		return
	}

	var usuario user

	if erro = json.Unmarshal(corpoReq, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter user em struct"))
	}

	db, erro := banco.Conecta()
	if erro != nil {
		w.Write([]byte("Não foi possivel conectar ao banco de dados."))
		return
	}
	defer db.Close()
	// Montando o PREPARE STATEMENT
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?,?)")
	if erro != nil {
		w.Write([]byte("Error ao criar o STATEMENT."))
		return
	}
	defer statement.Close()

	/* Passamos aqui os dados que não preenchemos no statement, para que com segurança adicionemos ao banco de dados. Logo em seguida, podemos adicionar os campos, na ordem que passamos acima, no caso, seria primeiro o nome e depois o email. */
	insert, erro := statement.Exec(usuario.Name, usuario.Email)
	if erro != nil {
		w.Write([]byte("Não foi possivel inserir os dados no banco de dados."))
		return
	}
	// Usamos essa função para fazer o auto preenchimento do ID do usuario no banco de dados.
	idIserido, erro := insert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Não foi possivel inserir o ID inserido"))
		return
	}
	// Status code, informamos qual foi o status recebido, e passa passar isso no golang, podemos passar da seguinte forma:
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso, id: %d", idIserido)))
}

// Buscar users do banco de dados
func buscarUsers(w http.ResponseWriter, r *http.Request) {

}

// Buscar user especifico do banco de dados
func buscarUser(w http.ResponseWriter, r *http.Request) {

}
