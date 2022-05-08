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
	requestRepo := repo.NewRequestRepo(db)
	roomService := service.NewRoomService(roomRepo, floorRepo, requestRepo)
	return controller.NewRoomController(roomService)
}

func InitFloorController(db *gorm.DB) *controller.FloorController {
	floorRepo := repo.NewFloorRepo(db)
	floorService := service.NewFloorService(floorRepo)
	return controller.NewFloorController(floorService)
}

func InitRequestController(db *gorm.DB) *controller.RequestController {
	roomRepo := repo.NewRoomRepo(db)
	requestRepo := repo.NewRequestRepo(db)
	requestService := service.NewRequestService(roomRepo, requestRepo)
	return controller.NewRequestController(requestService)
}

func InitUserController(db *gorm.DB) *controller.UserController {
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	return controller.NewUserController(userService)
}

func InitAdminController(db *gorm.DB) *controller.AdminController {
	requestRepo := repo.NewRequestRepo(db)
	userRepo := repo.NewUserRepo(db)
	adminService := service.NewAdminService(requestRepo, userRepo)
	return controller.NewAdminController(adminService)
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
	PORT := ":" + os.Getenv("PORT") //localhost
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
	requestController := InitRequestController(AppDB)
	userController := InitUserController(AppDB)
	adminController := InitAdminController(AppDB)

	// Create groups
	server.SetGroup("/", rootController.InitRouting)
	server.SetGroup("/room", roomController.InitRouting)
	server.SetGroup("/floor", floorController.InitRouting)
	server.SetGroup("/request", requestController.InitRouting)
	server.SetGroup("/admin", adminController.InitRouting)
	server.SetGroup("/user", userController.InitRouting)

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
