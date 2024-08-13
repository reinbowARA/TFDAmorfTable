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
		log.Println(err.Error())
		return
	}
	srv.GoogleTable = &handler.DataSecret{}
	json.Unmarshal(file, &srv.GoogleTable)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	srv.SetRouter(router)

	router.LoadHTMLGlob("web/page.html")

	//router.POST("/switch", srv.CheckSwitch)
	router.GET("/:object", srv.GetTable)
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/descdendants")
	})
	log.Println("server start")
	err = router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
