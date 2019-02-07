package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

//HandlerAmigo get post put delete
func HandlerAmigo(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		u, err := url.Parse(req.URL.String())
		if err != nil {
			log.Fatal(err)
		}
		q := u.Query()
		identificador := q.Get("id")
		for _, item := range amigos {
			if item.ID == identificador {
				res.Header().Set("Content-Type", "application/json")
				json.NewEncoder(res).Encode(item)
				return
			}
		}
		id2 := q.Get("nombre")
		for _, item := range amigos {
			if item.Nombre == id2 {
				res.Header().Set("Content-Type", "application/json")
				json.NewEncoder(res).Encode(item)
				return
			}
		}
		if identificador == "" {
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(amigos)

		}
	case http.MethodPost:

		var amigo Amigo
		_ = json.NewDecoder(req.Body).Decode(&amigo)
		amigos = append(amigos, amigo)
		json.NewEncoder(res).Encode(amigos)

	case http.MethodDelete:
		u, err := url.Parse(req.URL.String())
		if err != nil {
			log.Fatal(err)
		}
		q := u.Query()
		identificador := q.Get("id")
		fmt.Println(identificador)
		for _, item := range amigos {
			if item.ID == identificador {
				amigos = append(amigos[:1], amigos[1+1:]...)
				return
			}
		}

	default:
	}
}

//Amigo struct
type Amigo struct {
	ID       string `json:"id"`
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Edad     int    `json:"edad"`
	Telefono string `json:"telefono"`
}

var amigos []Amigo

func main() {

	amigos = append(amigos, Amigo{ID: "1", Nombre: "andres", Correo: "andresmuquinche@gmail.com", Edad: 30, Telefono: "0979392921"})
	amigos = append(amigos, Amigo{ID: "2", Nombre: "jose", Correo: "jose@gmail.com", Edad: 31, Telefono: "0979392922"})
	amigos = append(amigos, Amigo{ID: "3", Nombre: "pedro", Correo: "pedro@gmail.com", Edad: 32, Telefono: "0979392923"})

	mux := http.NewServeMux()
	mux.Handle("/amigos", http.HandlerFunc(HandlerAmigo))
	http.ListenAndServe(":3000", mux)
}
