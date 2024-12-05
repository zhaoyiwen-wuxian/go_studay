package main

import (
	"go_dome/cache"
	"go_dome/controllers"
	"log"

	"fmt"
	"go_dome/config"
	"go_dome/threadpool"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	cache.InitRedis()
	cache.InitMySQL()

	pool := threadpool.New(5)
	for i := 0; i < 10; i++ {
		task := func(i int) func() {
			return func() {
				fmt.Printf("Task %d completed\n", i)
			}
		}(i)
		pool.Add(task)
	}
	pool.Wait()
	r := gin.Default()
	port := config.AppConfig.Server.Port
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Go!")
	})
	fmt.Printf("Server running at http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

	r.POST("/user", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)

	// 文章路由
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts/user/:user_id", controllers.GetPost)

}
