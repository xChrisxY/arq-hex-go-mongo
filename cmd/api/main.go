package main

import (
	"log"
	"os"
	"github.com/arq-hex-go/cmd/api/handlers/player"
	"github.com/arq-hex-go/internal/repositories/mongo"
	playerMongo "github.com/arq-hex-go/internal/repositories/mongo/player"
	playerService "github.com/arq-hex-go/internal/services/player"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
 
func main() {

	err := godotenv.Load()
	
	if err != nil {

		log.Fatal("Error loading .env file")

	}

	ginEngine := gin.Default()

	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))

	if err != nil {

		log.Fatal(err.Error())
	}

	playerRepo := playerMongo.Repository {

		Client: client,

	}

	playerSrv := playerService.Service{

		Repo: playerRepo,

	}

	playerHandler := player.Handler{

		PlayerService: playerSrv,

	}

	ginEngine.POST("/players", playerHandler.CreatePlayer)
	ginEngine.GET("/players/:id", playerHandler.GetPlayer)
	ginEngine.DELETE("/players/:id", playerHandler.DeletePlayer)

	log.Fatalln(ginEngine.Run(":8001"))

}


