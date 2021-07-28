package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//处理文件上传
func UploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(1)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	file, fileHeader, err := r.FormFile("uploadFileName")
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	defer file.Close()

	f, err := os.OpenFile("/users/johnson/desktop/" + fileHeader.Filename, os.O_CREATE | os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	defer f.Close()

	io.Copy(f, file)
}
