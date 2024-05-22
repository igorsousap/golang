package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da mensagem"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco"))
		return
	}

	statment, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar  o statment"))
		return
	}
	defer statment.Close()

	insercao, erro := statment.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao inserir o usuario no banco"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obsert id inserido"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso id: %d", idInserido)))

}

// traz todos os usuarios do banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro aoprocurar usuarios"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear usuario"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuarios para JSON"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	linha, erro := db.Query("SELECT * FROM usuarios WHERE ID = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuario"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario!"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usaurio para JSON!"))
	}

}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da request"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequest, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statment, erro := db.Prepare("UPDATE usuarios SET nome = ?, email = ? WHERE id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statment"))
	}
	defer statment.Close()

	if _, erro := statment.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statment, erro := db.Prepare("DELETE FROM usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statment"))
	}
	defer statment.Close()

	if _, erro := statment.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar o usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
