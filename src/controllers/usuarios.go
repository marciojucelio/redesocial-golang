package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io"
	"net/http"
)

// CriarUsuario insert one user in database
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}
	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios searching all users in database
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all users"))
}

// BuscarUsuario searching just on user in database
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching just one user"))
}

// AtualizarUsuario update a user from the database
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Usuario"))
}

// DeletarUsuario delete a user from the database
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete User"))
}
