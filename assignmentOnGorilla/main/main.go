package main
 
import (
	"assignmentOnGorilla/app"
	"assignmentOnGorilla/config"
)
 
func main() {
	config := config.GetConfig()
 
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}