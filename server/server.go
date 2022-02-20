package server

import (
	"userbalance/config"
	"userbalance/controllers"
	"userbalance/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.ConfigDatabase()
	authController controllers.AuthController = controllers.NewAuthController(util.ProvideUserAuthService(), util.ProvideUserJwtService())
)

func RegisterRoute(r *gin.Engine) {
	defer config.CloseDBConnection(db)

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	r.Run()

}
