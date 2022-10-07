package services

import (
	"encoding/json"
	"fmt"
	"google-translate/server/handlers"
	"io/ioutil"
	"net/http"
)

func GetAllLanguagesFromGoogleTranslate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	language, err := handlers.GetLanguages()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	langBytes, err := json.Marshal(language)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.Write(langBytes)
}

func TranslateTheText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error on reading body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	var rqBody handlers.ReqBody

	err = json.Unmarshal(reqbody, &rqBody)
	if err != nil {
		fmt.Printf("error in unmarshaling body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	translatedText, err := handlers.ReqTranslate(&rqBody)
	if err != nil {
		fmt.Println("Error in translating", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	w.Write(translatedText)
}
