package main

import (
	"github.com/gin-contrib/cors"
	"log"
	"time"

	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition"
	"github.com/catness812/e-petitions-project/gateway/internal/security"
	"github.com/catness812/e-petitions-project/gateway/internal/user"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	rbacCfg := config.LoadConfigRBAC()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:1337"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	user.RegisterUserRoutes(r, &config.Cfg, rbacCfg)
	petition.RegisterPetitionRoutes(r, &config.Cfg)
	security.RegisterSecurityRoutes(r, &config.Cfg)

	err := r.Run(":1337")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
