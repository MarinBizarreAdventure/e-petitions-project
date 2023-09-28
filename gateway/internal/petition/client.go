package petition

import (
	"github.com/catness812/e-petitions-project/gateway/internal/config"
	"github.com/catness812/e-petitions-project/gateway/internal/petition/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type PetitionClient struct {
	Client pb.PetitionServiceClient
}

func InitPetitonServiceClient(c *config.Config) pb.PetitionServiceClient {
	cc, err := grpc.Dial(c.PetitionPort, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Could not connect:", err)
	}

	return pb.NewPetitionServiceClient(cc)
}
