package main

import (
	"goapi/controller"
	"goapi/database"
	"goapi/middleware"
	"goapi/repository"
	"goapi/service"
	"log"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	// DB接続
	db, err := database.NewDB("user:p@ssw0rd@tcp(localhost:3306)/goapi?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	// マイグレーション
	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	// リポジトリ、サービス、コントローラーの生成
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	authController := controller.NewAuthController(userService)
	userController := controller.NewUserController(userService)

	// ルーティング
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//　認証関連
			auth := v1.Group("/auth")
			{
				auth.POST("/sign_in", authController.SignIn)
				auth.POST("/sign_up", authController.SignUp)
			}

			// ユーザー関連
			user := v1.Group("/user").Use(middleware.Auth())
			{
				user.GET("/", userController.GetUser)
				user.PUT("/", userController.UpdateUser)
				user.DELETE("/", userController.DeleteUser)
			}
		}
	}
	return r
}
