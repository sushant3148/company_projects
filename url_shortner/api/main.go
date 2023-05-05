package main

import(
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/middleware/logger"
	"github.com/joho/godotenv"
	"os"
	"log"
)
func SetUpRoutes(app *fiber.App){
	app.get("/:url",routes.OriginalUrl)
	app.post("/api/v1",routes.ShortUrl)
}
func main(){
	err := godotenv.Load()
	if err != nil{
		fmt.Println(err)
	}
	app.Use(logger.New())
	app1 := fiber.New()
	SetUpRoutes(app1)
	log.Fatal(app.Listen(os.Getenv("APP")))

}