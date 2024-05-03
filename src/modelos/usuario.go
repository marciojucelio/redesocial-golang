package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario representa um usuario usando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os metodos para validar e formatar o usuario recebido
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}
func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("o campo Nome é obrigatório")
	}
	if usuario.Nick == "" {
		return errors.New("o campo Nick é obrigatório")
	}
	if usuario.Email == "" {
		return errors.New("o campo E-mail é obrigatório")
	}
	if usuario.Senha == "" {
		return errors.New("o campo Senha é obrigatório")
	}
	return nil
}
func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
