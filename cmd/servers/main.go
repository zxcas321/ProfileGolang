package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zxcas321/ProfileGolang/services/admin"
	"github.com/zxcas321/ProfileGolang/config"
	"github.com/zxcas321/ProfileGolang/routes"
	"github.com/zxcas321/ProfileGolang/utils"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	admin.AdminBoostrap()
	
	utils.InitSqids()
	r := gin.Default()

	// Register routes
	routes.ApiRoutes(r)

	r.Run()
}
