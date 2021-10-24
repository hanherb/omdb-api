package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hanherb/omdb-api/answers"
	content "github.com/hanherb/omdb-api/lib/controllerGrpc"
	controllerRest "github.com/hanherb/omdb-api/lib/controllerRest"
	"github.com/hanherb/omdb-api/lib/services"
	"github.com/hanherb/omdb-api/lib/variables"
	contentsGrpc "github.com/hanherb/omdb-api/omdb-generate/contents"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const (
	PortRest = 3000
	PortGrpc = 8180
)

func main() {
	godotenv.Load("devel.env")
	variables.Initialization()

	// Initialization services
	services.MysqlInitialization()

	errChan := make(chan error)
	defer close(errChan)

	// Start REST & gRPC server asynchronously
	go startRESTServer(errChan)
	go startGRPCServer(errChan)

	<-errChan
}

func startRESTServer(errChan chan error) {
	//Define endpoints
	router := gin.Default()
	api := router.Group("/api/omdb")
	{
		api.GET("/", controllerRest.HealthCheck)
		api.GET("/search", controllerRest.SearchMovies)
		api.GET("/unit-test", controllerRest.UnitTest)
	}

	answer := router.Group("/answer")
	{
		answer.GET("/3", answers.NumberThree)
		answer.GET("/4", answers.NumberFour)
	}

	//Running REST Server
	if err := router.Run(":" + strconv.Itoa(PortRest)); err != nil {
		errChan <- err
		log.Fatal(err)
	}
}

func startGRPCServer(errChan chan error) {
	port := flag.Int("port", PortGrpc, "port for gRPC server listen")
	flag.Parse()

	// Listen on port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen gRPC server: %s", err)
	}
	log.Printf("Listen gRPC server on %d port \n", *port)

	// Define gRPC Server
	server := grpc.NewServer()

	contentsGrpc.RegisterContentServiceServer(server, &content.Controller{})

	// Running gRPC Server
	if err := server.Serve(lis); err != nil {
		errChan <- err
		log.Fatalf("failed to serve: %s", err)
	}
}
