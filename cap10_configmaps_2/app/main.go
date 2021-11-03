package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	msg := os.Getenv("MSG")
	value := os.Getenv("VALUE")
	fmt.Fprintln(w,"Mensagem = ", msg, "\nValor = ", value)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
