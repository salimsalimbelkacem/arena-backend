package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"io"
)

var err error
const uploadDir = "./uploads"

func main() {
	
	err = os.MkdirAll(uploadDir, os.ModePerm)
	handleErr(err)

	
	http.HandleFunc("POST /submit", formSubHandler)

	port := "8080" 
	fmt.Println("Server is running on localhost:" + port)

	
	err = http.ListenAndServe(":"+port, nil)
	handleErr(err)
}

func formSubHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) 
	handleErr(err)

	
	teamName := strings.ToLower(r.FormValue("teamName"))
	if teamName == "" {
		http.Error(w, "Team name is required", http.StatusBadRequest)
		return
	}

	
	teamDir := uploadDir + "/" + teamName
	err = os.MkdirAll(teamDir, os.ModePerm)
	handleErr(err)

	
	for _, subdir := range []string{"code", "reverse"} {
		err = os.MkdirAll(teamDir+"/"+subdir, os.ModePerm)
		handleErr(err)
	}

    for _, which := range []string{"first", "second", "third"}{
        handleFileUpload(r, w, which+"Cod", teamDir+"/code")
        handleFileUpload(r, w, which+"Rev", teamDir+"/reverse")
    }

	
	fmt.Fprintf(w, "Thank you, team %s! Your submission has been received.", teamName)
}

func handleFileUpload(r *http.Request, w http.ResponseWriter, fieldName, subDir string) {
	
	file, _, err := r.FormFile(fieldName)
	if err != nil {
		if err.Error() != "http: no such file" {
			http.Error(w, "Error reading file: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	defer file.Close()

	
	dstFile, err := os.Create(subDir + "/" + fieldName)
	if err != nil {
		http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dstFile.Close()

	
	_, err = io.Copy(dstFile, file)
	if err != nil {
		http.Error(w, "Error saving file: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
