package main

import(
	"github.com/zxcas321/ProfileGolang/config"
)

func main(){
	config.LoadEnv()
	config.ConnectDB()
}