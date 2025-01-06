package server

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
	"github.com/mako8231/disctock/server/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitDatabase(fileName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(fileName+".db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return db
}

func Serve() {
	//Intialize the template engine
	engine := html.New("./server/view", ".html")

	//initialize the database
	Database = InitDatabase("test")

	//Run migrations
	Migrate(Database)

	//Enable this to make changes without reseting the server
	engine.ShouldReload = true

	//Initialize the App
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//Static assets (Javascript code + css files)
	app.Use("/assets", static.New("./server/view/assets"))

	//call the route module
	routes.WebRoutes(app)

	//Listen the server
	log.Fatal(app.Listen(":9090"))
}
