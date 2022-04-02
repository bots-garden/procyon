package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Settings struct {
	ProcyonUrl string `json:"procyonUrl"`
	ProcyonDomain string `json:"procyonDomain"`
}

var procyonUrl = getSettings().ProcyonUrl
var procyonDomain = getSettings().ProcyonDomain

func getSettings() Settings {
	settingsFile, err := ioutil.ReadFile("./alcor.json")
	if err != nil {
		log.Fatal(err)
	}
	settings := Settings{}

	err = json.Unmarshal([]byte(settingsFile), &settings)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("📝", settings)
	return settings
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

type FunctionRecord struct {
	WasmFunctionHttpPort int
	TaskId uuid.UUID
}

var functionsMap map[string]FunctionRecord

func proxy(c *gin.Context) {

	functionUrl := procyonDomain+":"+strconv.Itoa(functionsMap[c.Param("function_name")+"-*"].WasmFunctionHttpPort)

	remote, err := url.Parse(functionUrl)

	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func proxyRevision(c *gin.Context) {

	functionUrl := procyonDomain+":"+strconv.Itoa(functionsMap[c.Param("function_name")+"-"+c.Param("function_revision")].WasmFunctionHttpPort)

	remote, err := url.Parse(functionUrl)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

// TODO: use redis to share the data
func getFunctionsList() {
	for {
		resp, err := http.Get(procyonUrl+"/functions")
		if err != nil {
			log.Println(err)
		} else {
			// read the response body
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			json.Unmarshal(body, &functionsMap)
		}

	
		time.Sleep(5 * time.Second)
	}

}



func main() {

	go getFunctionsList()

	r := gin.Default()

	//Create catchall routes
	r.Any("/functions/:function_name", proxy)
	// 🚧 work in progress
	r.Any("/functions/:function_name/:function_revision", proxyRevision)

	r.Run(":" + getEnv("PROXY_HTTP", "8080"))
}
