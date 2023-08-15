package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kc0ffee/server/database"
	"github.com/kc0ffee/server/server"
)

func EnvLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(-1)
	}
	fmt.Println(".env was loaded.")
	fmt.Printf("SERVER_PORT : %s\n", os.Getenv("SERVER_PORT"))
}

func main() {
	EnvLoad()

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		fmt.Printf("Error : %+v\n", err)
		os.Exit(-1)
	}

	db := database.NewDBConnection(os.Getenv("DATABASE_ADDRESS"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_NAME"))
	defer db.Close()
	database.Insert(db, 123, "test")

	e := server.NewAPIServer()
	server.StartServer(e, port)
}
