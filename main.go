package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/name", hostName).Methods("GET")
	fmt.Print("Server Started")
	http.ListenAndServe(":3000", r)

}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode("Hello! your request is processed sucessfully")
}

func hostName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
	name, _ := os.Hostname()
	uid:= os.Geteuid()
	gid:=os.Getegid()
	pid:=os.Getpid()
	fmt.Fprintf(w, " Hello from Hostname %s with UID:%v and GID:%v and PID: %v", name,uid,gid,pid)
}
