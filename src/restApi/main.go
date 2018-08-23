package main

// import ("encoding/json"
// 		"log"
// 		"net/http"
// 		"math/rand"
// 		"strconv"
// 		"github.com/gorilla/mux")
import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux")

// Msg model(sturacture)
type Msg struct {
	ID string `json:"id"`
	Msg string `json:"msg"`
	Sender * Sender `json:"sender"`
	Receiver * Receiver `json:"receiver"`
}
// user model

type Sender struct {
	User * User `json:"user"`
}
type Receiver struct {
	User * User `json:"user"`
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
// mock data
var msgs[] Msg

// get message
func getMsgs(w http.ResponseWriter, r * http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msgs)
}
// get message by id 
func getMsg(w http.ResponseWriter, r * http.Request) {

}
// new message
func newMsg(w http.ResponseWriter, r * http.Request) {

}
// delete message
func deleteMsg(w http.ResponseWriter, r * http.Request) {

}
// edit message
func editMsg(w http.ResponseWriter, r * http.Request) {

}

// bootstrapping

func main() {
	//fmt.Printf("hello, world\n")
	msgs = append(msgs, Msg {
		ID: "2",
		Msg: "hello ganza, happy birth!",
		Sender: & Sender {
			User: & User {
				Firstname: "kiza",
				Lastname: "emmy"
			}
		},
		Receiver: & Receiver {
			User: & User {
				Firstname: "ganza",
				Lastname: "respi"
			}
		}
	})


	msgs = append(msgs, Msg {
		ID: "1",
		Msg: "Hi! kiza, am good thank u!",
		Sender: & Sender {
			User: & User {
				Firstname: "ganza",
				Lastname: "res"
			}
		},
		Receiver: & Receiver {
			User: & User {
				Firstname: "kiza",
				Lastname: "emmy"
			}
		}
	})

	//Init router
	r: = mux.NewRouter()
	r.HandleFunc("/api/msg", getMsgs).Methods("GET")
	r.HandleFunc("/api/msg/{id}", getMsg).Methods("GET")
	r.HandleFunc("/api/msg", newMsg).Methods("POST")
	r.HandleFunc("/api/msg/{id}", deleteMsg).Methods("DELETE")
	r.HandleFunc("/api/msg/{id}", editMsg).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", r))
}