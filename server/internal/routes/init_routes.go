package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO
		return true
	},
}

var onlineUsers = make(map[string]chan string)

func InitRoutes(r *gin.Engine, redistHost string) {
	r.GET("/users", func(c *gin.Context) {
		users := make([]string, 0)
		for user := range onlineUsers {
			users = append(users, user)
		}
		c.JSON(200, gin.H{
			"message": users,
		})
	})

	r.POST("/login/:userID", func(c *gin.Context) {
		userID := c.Param("userID")
		if _, ok := onlineUsers[userID]; ok {
			c.JSON(400, gin.H{
				"message": "user already online",
			})
			return
		}

		onlineUsers[userID] = make(chan string)
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.POST("/logout/:userID", func(c *gin.Context) {
		userID := c.Param("userID")
		if _, ok := onlineUsers[userID]; !ok {
			c.JSON(400, gin.H{
				"message": "user already offline",
			})
			return
		}

		delete(onlineUsers, userID)
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.GET("/connect/:myUser/:otherUser", func(c *gin.Context) {
		myUser := c.Param("myUser")
		// otherUser := c.Param("otherUser")

		client := redis.NewClient(&redis.Options{
			Addr:     redistHost + ":6379", // todo: include port in env vars
			Password: "",
			DB:       0,
		})

		pubsub := client.Subscribe(context.TODO(), myUser)
		defer pubsub.Close()

		//upgrade get request to websocket protocol
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ws.Close()

		for {
			go func() {
				msg, err := pubsub.ReceiveMessage(context.TODO())
				fmt.Println(msg)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(msg.Channel, msg.Payload)
			}()

			go func() {
				_, _, err := ws.ReadMessage()
				if err != nil {
					fmt.Println(err)
					return
				}
				// fmt.Println(mt, string(b))
			}()
		}
	})
}
