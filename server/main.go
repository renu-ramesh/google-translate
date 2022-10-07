package main

import (
	"fmt"
	"google-translate/server/services"
	"net/http"
)

func main() {

	http.HandleFunc("/getalllanguages", services.GetAllLanguagesFromGoogleTranslate)
	http.HandleFunc("/translate", services.TranslateTheText)

	err := http.ListenAndServe(":8080", nil)
	go func() {
		if err == nil {
			fmt.Println("server started on port 8080", err)
		}
	}()

}
