package main


import (
  "log"
  "net/http"

  "github.com/gin-gonic/gin"
  _ "github.com/Akbarkhon-Avazkhonov/GO-REACT/goapp/docs/goapp"
  swaggerFiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
)

// @title GO REACT API
// @version 0.1.0



func main() {
	// Gin instance
	r := gin.New()

	// Routes
	r.GET("/", HealthCheck)
	r.GET("/ping", Pong )

	url := ginSwagger.URL("http://localhost:8000/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Start server
	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Root
// @Accept */ping*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

// Pong godoc
// @Summary Say pong.
// @Description Say pong.
// @Tags Root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /ping [get]
func Pong (c *gin.Context) {
	res := map[string]interface{}{
		"data": "pong",
	}

	c.JSON(http.StatusOK, res)
}