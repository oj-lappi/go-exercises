/*
En simpel webbserver
--------------------

 - Kör filen med go run 
 - hämta sidan localhost:8080 i en browser, browsern måste köra i gästoperativsystemet (Ubuntu)

 Läs mera på: 	https://golang.org/doc/articles/wiki/#tmp_3
 		https://golang.org/pkg/net/http/

*/
package main

import (
        "fmt"
        "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hej från go servern!")
}

func main(){
	//Handlefunc säger vilken endpoint som hör ihop med vilken funktion
	//Nu har vi definierat en enda endpoint som fångar alla http-anrop
	http.HandleFunc("/",handler)
	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
