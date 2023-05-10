package main

import(
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"os"
	"log"
	"assigment2/api/routes"
)
func SetUpRoutes(app *fiber.App){
	app.Get("/:url",routes.OriginalUrl)
	app.Post("/api/v1",routes.ShortUrl)
}
func main(){
	err := godotenv.Load()
	if err != nil{
		fmt.Println(err)
	}
    app1 := fiber.New()
	app1.Use(logger.New())
	SetUpRoutes(app1) 	
	log.Fatal(app1.Listen(os.Getenv("APP")))

}