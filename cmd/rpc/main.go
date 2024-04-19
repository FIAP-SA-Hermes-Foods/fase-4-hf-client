package main

import (
	cp "fase-4-hf-client/client_proto"
	"fase-4-hf-client/external/db/dynamo"
	l "fase-4-hf-client/external/logger"
	repositories "fase-4-hf-client/internal/adapters/driven/repositories/nosql"
	"fase-4-hf-client/internal/core/application"
	"fase-4-hf-client/internal/core/useCase"
	grpcH "fase-4-hf-client/internal/handler/rpc"
	"net"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/marcos-dev88/genv"
	"google.golang.org/grpc"
)

func init() {
	if err := genv.New(); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}
}

func main() {

	listener, err := net.Listen("tcp", os.Getenv("GRPC_CLIENT_PORT"))

	if err != nil {
		l.Errorf("error to create connection %v", " | ", err)
	}

	defer listener.Close()

	configAws := aws.NewConfig()
	configAws.Region = aws.String("us-east-1")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            *configAws,
		SharedConfigState: session.SharedConfigEnable,
	}))

	db := dynamo.NewDynamoDB(sess)

	repo := repositories.NewClientRepository(db, os.Getenv("DB_TABLE"))

	uc := useCase.NewClientUseCase()

	app := application.NewApplication(repo, uc)

	h := grpcH.NewHandler(app)

	grpcServer := grpc.NewServer()

	cp.RegisterClientServer(grpcServer, h.Handler())

	if err := grpcServer.Serve(listener); err != nil {
		l.Errorf("error in server: %v", " | ", err)
	}
}
