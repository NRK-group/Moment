package auth

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func WriteImage(dir string, r *http.Request) (bool, string) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("Error getting file from form data: ", err)
		return false, ""
	}
	defer file.Close()
	getFilePrefix := strings.Split(handler.Filename, ".")
	var imgType string
	imageTypes := "img png gif svg jpg jpeg JPG JPEG"
	if !strings.Contains(imageTypes, getFilePrefix[len(getFilePrefix)-1]) {
		log.Println("Error getting file from form data: ", err)
		return false, "Invalid File Format"
	}
	if handler.Size > int64(20000000) {
		log.Println("File size exceeded")
		return false, "File Size Limit Exceeded"
	}
	imgType = getFilePrefix[len(getFilePrefix)-1]
	tempFile, err := os.CreateTemp(dir, "*."+imgType)
	if err != nil {
		log.Println("Error creating temp file: ", err)
		return false, ""
	}
	defer tempFile.Close()
	imgUrl := tempFile.Name()
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return false, ""
	}
	tempFile.Write(fileBytes)
	return true, imgUrl
}
