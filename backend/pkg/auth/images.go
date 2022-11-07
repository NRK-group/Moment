package auth

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func WriteImage(dir string, r *http.Request) (bool,string) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Error getting file from form data: ", err)
		return false, "500 Internal Server Error"
	}
	defer file.Close()
	getFilePrefix := strings.Split(handler.Filename, ".")
	var imgType string
	imageTypes := "img png gif svg jpg jpeg JPG JPEG"
	if !strings.Contains(imageTypes, getFilePrefix[len(getFilePrefix)-1]) {
		log.Println("Error getting file from form data: ", err)
		return false, "Invalid file format"
	}
	if handler.Size > int64(20000000) {
		log.Println("File size exceeded")
		return false, "File size limit exceeded"
	}

	imgType = getFilePrefix[len(getFilePrefix)-1]
	tempFile, err := os.CreateTemp(dir, uuid.NewV4().String()+"*."+imgType)
	if err != nil {
		log.Println("Error creating temp file: ", err)
	}
	defer tempFile.Close()
	imgUrl := tempFile.Name()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	return true, imgUrl
}
