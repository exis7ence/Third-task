package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandle(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		http.Error(w, "Unable to determine working directory", http.StatusInternalServerError)
		return
	}

	indexPath := filepath.Join(wd, "index.html")

	if _, err := os.Stat(indexPath); os.IsNotExist(err) {
		http.Error(w, "index.html not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, indexPath)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	input := string(data)
	result, err := service.AutomaticCodeDetection(input)
	if err != nil {
		http.Error(w, "Error converting", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("%s%s", time.Now().UTC().String(), ext)

	outputFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = outputFile.WriteString(result)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
