package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"metalink-apiserver/database"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	api := router.Group("/api")
	{
		api.Any("/*action", func(c *gin.Context) {
			paramsMap := make(map[string]string)
			for _, param := range c.Params {
				paramsMap[param.Key] = param.Value
			}
			queryParams := c.Request.URL.Query()
			for key, values := range queryParams {
				if len(values) > 0 {
					paramsMap[key] = values[0]
				}
			}
			if len(paramsMap) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": "There is no view parameter",
				})
			} else {
				if !CheckKey(paramsMap, "view") {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  http.StatusBadRequest,
						"message": "There is no view parameter",
					})
				} else {
					clv, cli := database.Check(paramsMap["view"])
					if !clv {
						c.JSON(http.StatusBadRequest, gin.H{
							"status":  http.StatusBadRequest,
							"message": "invalid value",
						})
					} else {
						c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(database.List[cli].MetaLink))
					}
				}
			}
		})
	}

	router.GET("/submit", func(c *gin.Context) {
		paramsMap := make(map[string]string)
		var images []string
		for _, param := range c.Params {
			paramsMap[param.Key] = param.Value
		}
		queryParams := c.Request.URL.Query()
		for key, values := range queryParams {
			if len(values) > 0 {
				paramsMap[key] = values[0]
			}
			if key == "image" {
				images = values
			}
		}
		if len(paramsMap) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "invalid data",
			})
		} else {
			scheme := "https"
			host := c.Request.Host
			if !CheckKey(paramsMap, "title") &&
				!CheckKey(paramsMap, "description") &&
				!CheckKey(paramsMap, "siteurl") &&
				!CheckKey(paramsMap, "color") &&
				images == nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": "You can't leave all items blank.",
				})
			} else if paramsMap["title"] == "" &&
				paramsMap["description"] == "" &&
				paramsMap["siteurl"] == "" &&
				paramsMap["color"] == "" &&
				paramsMap["image"] == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": "You can't leave all items blank.",
				})
			} else if !CheckKey(paramsMap, "sitetype") {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"message": "There is no sitetype",
				})
			} else {
				if !CheckKey(paramsMap, "title") {
					paramsMap["title"] = ""
				}
				if !CheckKey(paramsMap, "description") {
					paramsMap["description"] = ""
				}
				if !CheckKey(paramsMap, "siteurl") || paramsMap["siteurl"] == "" {
					paramsMap["siteurl"] = scheme + "://" + host
				}
				if !CheckKey(paramsMap, "color") {
					paramsMap["color"] = ""
				}
				if images == nil {
					images = []string{""}
				}
				if paramsMap["sitetype"] != "website" {
					c.JSON(http.StatusBadRequest, gin.H{
						"status":  http.StatusBadRequest,
						"message": "sitetype currently supports only website",
					})
				} else {
					id := database.New(c.ClientIP(), 60*60*24*7,
						paramsMap["title"],
						paramsMap["description"],
						paramsMap["sitename"],
						paramsMap["siteurl"],
						paramsMap["sitetype"],
						paramsMap["color"],
						images,
					)
					completeURL := scheme + "://" + host + "/api?view="
					c.JSON(http.StatusOK, gin.H{
						"status":  http.StatusOK,
						"message": completeURL + id,
					})
				}
			}
		}
	})

	return router
}

func CheckKey(data map[string]string, key string) bool {
	_, exists := data[key]
	return exists
}
