package testuplaod

import (
	"fmt"
	"io"
	genKey "iqdev/ss/libs/key"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     nil}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Expect") == "100-continue" {
		w.WriteHeader(http.StatusContinue)
	}
	r.ParseMultipartForm(100000000)

	fmt.Println("reading from upload")
	file, header, err := r.FormFile("uploadfile")
	if err != nil {
		http.Error(w, "faild to get upload file ", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fmt.Println("making new file")
	cfile, cfileErr := os.OpenFile(
		genKey.RandomKey(20)+header.Filename,
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	if cfileErr != nil {
		http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
		return
	}
	fmt.Println("copying the file")
	size, cpyErr := io.Copy(cfile, file)
	fmt.Println("done copy")
	if cpyErr != nil {
		http.Error(w, "faild to read upload file ", http.StatusInternalServerError)
		return
	}
	fmt.Println(size)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("file uploaded yay"))
}
