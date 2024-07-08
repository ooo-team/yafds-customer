package main

import (
	// "context"
	"log"
	// "github.com/ooo-team/YAFDS/internal/app"
)

func main() {

	// ctx := context.Background()
	log.Println("kekes")
	// a, err := app.NewApp(ctx)
	// if err != nil {
	// 	log.Fatalf("failed to init app: %s", err.Error())
	// }

	// err = a.Run()
	// if err != nil {
	// 	log.Fatalf("failed to run app: %s", err.Error())
	// }
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// // init is invoked before main()
// func init() {
// 	// loads values from .env into the system
// 	if err := godotenv.Load(); err != nil {
// 		log.Print("No .env file found")
// 	}
// }

// func load_dotenv_var(var_name string) string {
// 	var_, exists := os.LookupEnv(var_name)

// 	if !exists {
// 		log.Panic("Dotenv variable {} does not exist", var_name)
// 	}

// 	return var_
// }

// func main() {
// 	// Get the GITHUB_USERNAME environment variable
// 	db_user := load_dotenv_var("db_user")
// 	db_password := load_dotenv_var("db_password")

// }
