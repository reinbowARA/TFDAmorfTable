package main

import (
	"TFDAmorfTable/handler"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	srv := handler.Server{}

	file, err := os.ReadFile("configs/secrets.json")
	if err != nil {
		log.Fatalf("Error reading gonfig file: %s", err.Error())
		return
	}
	srv.GoogleTable = &handler.DataSecret{}
	json.Unmarshal(file, &srv.GoogleTable)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	srv.SetRouter(router)

	router.LoadHTMLGlob("web/page.html")

	router.GET("/:object", srv.GetTable)
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/descendants")
	})
	log.Println("server start")
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
