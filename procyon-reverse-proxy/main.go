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
	resty "github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

type Settings struct {
	ProcyonUrl    string `json:"procyonUrl"`
	ProcyonDomain string `json:"procyonDomain"`
}

var procyonUrl = getSettings().ProcyonUrl
var procyonDomain = getSettings().ProcyonDomain

var procyonAdminToken = getEnv("PROCYON_ADMIN_TOKEN", "")

func getSettings() Settings {
	settingsFile, err := ioutil.ReadFile("./procyon-reverse.json")
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
	TaskId               uuid.UUID
	DefaultRevision      bool
}

var functionsMap map[string]FunctionRecord
var defaultRevisionsMap map[string]FunctionRecord

func proxy(c *gin.Context) {

	functionUrl := procyonDomain + ":" + strconv.Itoa(defaultRevisionsMap[c.Param("function_name")].WasmFunctionHttpPort)

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

	functionUrl := procyonDomain + ":" + strconv.Itoa(functionsMap[c.Param("function_name")+"-"+c.Param("function_revision")].WasmFunctionHttpPort)

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
// TODO: empty map?
func getFunctionsList() {
	client := resty.New()
	for {

		//log.Println("🔵 functions", procyonAdminToken)

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("PROCYON_ADMIN_TOKEN", procyonAdminToken).
			Get(procyonUrl + "/functions")

		if err != nil {
			log.Println("😡", err)
		} else {
			// eg 401 Unauthorized
			if resp.IsError() {
				log.Println("😡", resp.Status())
			} else {
				json.Unmarshal(resp.Body(), &functionsMap)
			}
		}

		log.Println("🌍", functionsMap)

		time.Sleep(5 * time.Second)
	}

}

// TODO: use redis to share the data
// TODO: empty map?
func getDefaultRevisionsList() {
	client := resty.New()
	for {

		//log.Println("🔴 revisions", procyonAdminToken)

		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("PROCYON_ADMIN_TOKEN", procyonAdminToken).
			Get(procyonUrl + "/revisions/default")

		if err != nil {
			log.Println("😡", err)
		} else {
			// eg 401 Unauthorized
			if resp.IsError() {
				log.Println("😡", resp.Status())
			} else {
				json.Unmarshal(resp.Body(), &defaultRevisionsMap)
			}
		}

		log.Println("🌕", defaultRevisionsMap)

		time.Sleep(5 * time.Second)
	}

}

func main() {

	go getFunctionsList()
	go getDefaultRevisionsList()

	r := gin.Default()

	//Create catchall routes
	r.Any("/functions/:function_name", proxy)
	// 🚧 work in progress
	r.Any("/functions/:function_name/:function_revision", proxyRevision)

	if getEnv("PROXY_CRT", "") != "" {
		r.RunTLS(":"+getEnv("PROXY_HTTPS", "4443"), getEnv("PROXY_CRT", "certs/procyon-registry.local.crt"), getEnv("PROXY_KEY", "certs/procyon-registry.local.key"))
	} else {
		r.Run(":" + getEnv("PROXY_HTTP", "8080"))
	}

}
