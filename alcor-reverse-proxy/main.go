package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}


func proxy(c *gin.Context) {

	log.Println("🚀", c.Param("function_name"), c.Param("function_revision"))

	remote, err := url.Parse("http://localhost:8081")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	//Define the director func
	//This is a good place to log, for example
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {

	r := gin.Default()

	//Create catchall routes
	r.Any("/functions/:function_name", proxy)
	// 🚧 work in progress
	r.Any("/functions/:function_name/:function_revision", proxy)

	r.Run(":"+getEnv("PROXY_HTTP", "8080"))
}
