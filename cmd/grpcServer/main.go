package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/Cavalcantefilipe/grpc-full-cycle/internal/database"
	"github.com/Cavalcantefilipe/grpc-full-cycle/internal/pb"
	"github.com/Cavalcantefilipe/grpc-full-cycle/internal/service"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/mysql")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Success!")

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
