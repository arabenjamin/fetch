package main


import (

  "os"
  "fmt"
  "log"

  "github.com/arabenjamin/fetch/server"
)


func main(){


  fmt.Printf("Running Fetch App\n")

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Starting receipt processor service ...")

  
  err := server.StartServer()
  if err != nil {
    logger.Println(err)
    logger.Print("Holy critical malfunction Batman!")
  }

}
