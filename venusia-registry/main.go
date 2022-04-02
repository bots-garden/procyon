package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/go-chi/chi"
)

type Api struct {
	Address string
	Port    string
	Router  *chi.Mux
}

type FileInfo struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}


type ErrResponse struct {
	HTTPStatusCode int
	Message        string
}

func (a *Api) UploadHandler(responseWriter http.ResponseWriter, request *http.Request) {
	// Maximum upload of 10 MB files
	request.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := request.FormFile("file")
	if err != nil {
		log.Println("😡 Error Retrieving the File")
		log.Println("😡", err)
		return
	}

	defer file.Close()
	// TODO: store this part somewhere (and create an API)
	// TODO: add API to get functions list
	// TODO: add API to get functions sizes
	log.Println("📦 uploaded file:", handler.Filename)
	log.Println("📦 file size:", handler.Size)
	log.Println("🌍 MIME Header:", handler.Header)

	// Create file
	dst, err := os.Create("./functions/" + handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		responseWriter.WriteHeader(404)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		responseWriter.WriteHeader(404)
		return
	}

	log.Println("🎉", handler.Filename, "successfully uploaded")
	responseWriter.WriteHeader(201)
}

func (a *Api) DeleteHandler(responseWriter http.ResponseWriter, request *http.Request) {

	wasmFile := chi.URLParam(request, "wasmFile")
	if wasmFile == "" {
		log.Println("😡 no wasm file name passed in request")
		responseWriter.WriteHeader(400)
	}

	file2Del := "./functions/" + wasmFile
	log.Println("🪣 deleting file:", file2Del)
	err := os.Remove(file2Del)
	if err != nil {
		log.Println("😡 when removing", file2Del)
		responseWriter.WriteHeader(404)
	}
	responseWriter.WriteHeader(204)

}

func (a *Api) DownloadHandler(responseWriter http.ResponseWriter, request *http.Request) {
	wasmFile := chi.URLParam(request, "wasmFile")
	if wasmFile == "" {
		log.Println("😡 no wasm file name passed in request")
		responseWriter.WriteHeader(400)
	}

	workDir := "./functions"
	fileLocation := path.Join(workDir, wasmFile)
	file, _ := os.Open(fileLocation)
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("😡 when downloading:", err)
		responseWriter.Write([]byte("something wen't wrong!!"))
    return
  }
	responseWriter.Write(bytes)

	log.Println("🎉", wasmFile, "successfully downloaded")
	responseWriter.WriteHeader(201)
}

func (a *Api) ListHandler(responseWriter http.ResponseWriter, request *http.Request) {

	dir, err := os.Open("./functions")
	if err != nil {
		log.Println("😡 when getting list:", err)
		responseWriter.WriteHeader(404)	
	}
	entries, err := dir.Readdir(0)
	if err != nil {
		log.Println("😡 when getting list:", err)
		responseWriter.WriteHeader(404)
	}

	list := []FileInfo{}

	for _, entry := range entries {
			f := FileInfo{
					Name:    entry.Name(),
					Size:    entry.Size(),
					Mode:    entry.Mode(),
					ModTime: entry.ModTime(),
					IsDir:   entry.IsDir(),
			}
			list = append(list, f)
	}

	output, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		log.Println("😡 when getting list:", err)
		responseWriter.WriteHeader(404)
	}

	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(200)
	json.NewEncoder(responseWriter).Encode(string(output)) 
}

/*
# upload wasm file
curl -v \
  -F "file=@oh.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:9999/wasm/upload

curl -v \
  -F "file=@hello.wasm" \
	-H "Content-Type: multipart/form-data" \
	-X POST https://localhost:9999/wasm/upload

# delete wasm file
curl -v \
	-X DELETE https://localhost:9999/wasm/delete/oh.wasm

# download wasm file
curl https://localhost:9999/wasm/download/oh.wasm --output oh.wasm

# get wasm files list
curl https://localhost:9999/wasm/list
*/

func (a *Api) InitRouter() {
	a.Router = chi.NewRouter()

	a.Router.Route("/wasm", func(r chi.Router) {
		r.Get("/list", a.ListHandler)
		r.Post("/upload", a.UploadHandler)
		r.Route("/download/{wasmFile}", func(r chi.Router) {
			r.Get("/", a.DownloadHandler)
		})
		r.Route("/delete/{wasmFile}", func(r chi.Router) {
			r.Delete("/", a.DeleteHandler)
		})
	})
}

func (a *Api) Start() { // Address???
	a.InitRouter()

	log.Println("🚀 starting Galago Registry...")

	crt := getEnv("REGISTRY_CRT", "certs/venusia.local.crt")
	key := getEnv("REGISTRY_KEY", "certs/venusia.local.key")

	log.Println("🌍 Listening on " + a.Port)

	log.Fatal(http.ListenAndServeTLS(":"+a.Port, crt, key, a.Router))

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {

	api := Api{
		Address: "", Port: getEnv("REGISTRY_HTTP", "9999"),
	}

	api.Start()

}
