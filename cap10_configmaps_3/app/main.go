package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	filePath := "/app/textfile.txt"
	dat, err := ioutil.ReadFile(filePath)

	if(err != nil){
		return
	}
	fmt.Fprintln(w,"Texto no arquivo ", filePath, "=", string(dat))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
