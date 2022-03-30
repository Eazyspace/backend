package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Eazyspace/api"
	"github.com/Eazyspace/controller"
	"github.com/Eazyspace/db"
	"github.com/Eazyspace/repo"
	"github.com/Eazyspace/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoomController(db *gorm.DB) *controller.RoomController {
	roomRepo := repo.NewRoomRepo(db)
	floorRepo := repo.NewFloorRepo(db)
	roomService := service.NewRoomService(roomRepo, floorRepo)
	return controller.NewRoomController(roomService)
}

func InitFloorController(db *gorm.DB) *controller.FloorController {
	floorRepo := repo.NewFloorRepo(db)
	floorService := service.NewFloorService(floorRepo)
	return controller.NewFloorController(floorService)
}

func InitRootController(db *gorm.DB) *controller.RootController {
	return controller.NewRootController()
}

// func InitRootController(db *gorm.DB) *controller.RoomController {
// 	rootRepo := repo.NewRoomRepo(db)
// 	rootService := service.NewRoomService(rootRepo)
// 	return controller.NewRoomController(rootService)
// }
func main() {
	// e := echo.New()
	// e.GET("/", root)
	// e.Logger.Fatal(e.Start(":8080"))

	// init server
	server := api.InitServer()
	godotenv.Load(".env")
	PORT := ":" + os.Getenv("PORT")
	DB_URI := os.Getenv("DB_URI")
	fmt.Printf("%s %s", PORT, DB_URI)

	// init database
	AppDB := db.CreateUniversalDB(DB_URI, "eazyspace")

	// Auto Migration

	// AppDB.AutoMigrate(&model.Room{})
	// AppDB.AutoMigrate(&model.Floor{})
	// AppDB.AutoMigrate(&model.Organization{})
	// AppDB.AutoMigrate(&model.User{})
	// AppDB.AutoMigrate(&model.Token{})
	// AppDB.AutoMigrate(&model.Request{})

	// Create controllers
	rootController := InitRootController(AppDB)
	roomController := InitRoomController(AppDB)
	floorController := InitFloorController(AppDB)

	// Create groups
	server.SetGroup("/", rootController.InitRouting)
	server.SetGroup("/room", roomController.InitRouting)
	server.SetGroup("/floor", floorController.InitRouting)
	// InitRoomRouting(roomController, *server)
	server.Start(PORT)

}

func onDBConnected(c *gorm.DB) {
	// model.InitUserModel(c)
	// model.InitTestDB(c)
	// model.InitResultDB(c)
	// model.InitQuestionDB(c)
	// model.InitTopicDB(c)
	// model.InitSubjectDB(c)
}

// func InitRoomRouting(controller controller.BaseController, server api.APIServer) {
// 	server.SetGroup("/", controller.InitRouting)
// 	server.SetGroup("/room", controller.InitRouting)
// 	// server.SetGroup("/user", controller.UserControllerGroup)
// 	// server.SetGroup("/test", controller.TestControllerGroup)
// 	// server.SetGroup("/topic", controller.TopicControllerGroup)
// 	// server.SetGroup("/question", controller.QuestionControllerGroup)
// 	// server.SetGroup("/result", controller.ResultControllerGroup)
// 	// server.SetGroup("/subject", controller.SubjectControllerGroup)
// }

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Eazyspace server!")
}
