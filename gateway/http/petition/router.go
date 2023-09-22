package petition

import (
	"github.com/catness812/e-petitions-project/gateway/config"
	"github.com/gin-gonic/gin"
	"log"
)

func RegisterPetitionRoutes(r *gin.Engine, c config.Config) {
	petitionService, err := NewPetitionService(c)
	if err != nil {
		log.Fatal("Failed to connect to petition service grpc: ", err)

	}

	petitionController := NewPetitionController(petitionService)

	route := r.Group("/petition")
	route.POST("/", petitionController.CreatePetition)
	route.GET("/", petitionController.GetPetition)
	route.DELETE("/", petitionController.DeletePetition)
	route.POST("/update", petitionController.UpdatePetition)
	route.GET("/all", petitionController.GetAllPetitions)

}
