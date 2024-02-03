package main

import (
	"fmt"

	_ "github.com/joho/godotenv/autoload"
	c "github.com/rootxrishabh/chocoGram/config"
	s "github.com/rootxrishabh/chocoGram/server"
)

func main() {
	c.Migrate()
	fmt.Println("Server running on port 8080...")
	s.Server()
}
