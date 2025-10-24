package main

import (
	// "fmt"
	// "time"
	"example.com/m/v2/config"
	"example.com/m/v2/models"
	"example.com/m/v2/route"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	config.InitDB()
	models.RunMigrations()

	server := gin.Default()
	route.RegisterRoutes(server)
	server.Run(":8080")

	// inputChan := make(chan int)
	// processChan := make(chan int)
	// outputChan := make(chan int)

	// go func() {
	// 	for i := 1; i <= 5; i++ {
	// 		fmt.Println("Send to inputChan: ", i)
	// 		inputChan <- i
	// 		time.Sleep(300 * time.Millisecond)
	// 	}
	// 	close(inputChan)
	// }()

	// go func() {
	// 	for num := range inputChan {
	// 		fmt.Println("Processing: ", num)
	// 		processChan <- num * 2
	// 	}
	// 	close(processChan)
	// }()

	// go func() {
	// 	for num := range processChan {
	// 		outputChan <- num + 1
	// 	}
	// 	close(outputChan)
	// }()

	// for result := range outputChan {
	// 	fmt.Println("Final Result: ", result)
	// }
}
