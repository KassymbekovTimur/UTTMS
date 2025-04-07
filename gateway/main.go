package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "gateway ok"})
	})

	r.Any("/schedules/*proxyPath", reverseProxy("http://schedule:8001"))
	r.Any("/participants/*proxyPath", reverseProxy("http://participant:8002"))

	r.Run(":8080")
}

func reverseProxy(target string) gin.HandlerFunc {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)
	return func(c *gin.Context) {
		c.Request.URL.Path = c.Param("proxyPath") // путь без префикса /schedules
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
