package controllers

import (
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}
