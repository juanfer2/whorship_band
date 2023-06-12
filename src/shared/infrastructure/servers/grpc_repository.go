package servers

import (
	"fmt"
	"log"
	"net"

	wompi_grpc "github.com/juanfer2/whorship_band/src/wompi/infrastructure/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServerGrpcApp() {
	list, err := net.Listen("tcp", ":5060")

	if err != nil {
		log.Fatalf("Error listening: %s", err.Error())
	}

	s := grpc.NewServer()
	wompi_grpc.RegisterServices(s)

	reflection.Register(s)
	fmt.Println("Server is running on port 5060")

	if err := s.Serve(list); err != nil {
		log.Fatalf("Error serving: %s", err.Error())
	} else {
		fmt.Println("Server is running on port 5060")
	}

	// v := fmt.Sprintln("Server is running on port %s", ":5060")
	fmt.Println("Server is running on port 5060")
}
