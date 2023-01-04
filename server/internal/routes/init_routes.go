package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO
		return true
	},
}

var onlineUsers = make(map[string]chan string)

func InitRoutes(r *gin.Engine) {
	r.GET("/users", func(c *gin.Context) {
		// Todo: use a real database with gin gonic
		c.JSON(200, gin.H{
			"message": []string{"user1", "user2", "user3", "user4"},
		})
	})

	r.POST("/online/:userID", func(c *gin.Context) {
		userID := c.Param("userID")
		onlineUsers[userID] = make(chan string)
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.GET("/connect/:myUser/:otherUser", func(c *gin.Context) {
		myUser := c.Param("myUser")
		otherUser := c.Param("otherUser")

		myUserChan, ok := onlineUsers[myUser]
		if !ok {
			c.JSON(400, gin.H{
				"message": "user not online",
			})
			return
		}

		otherUserChan, ok := onlineUsers[otherUser]
		if !ok {
			c.JSON(400, gin.H{
				"message": "user not online",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "success",
		})

		//upgrade get request to websocket protocol
		// 	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 		return
		// 	}
		// 	defer ws.Close()
		// 	for {
		// 		//Read Message from client
		// 		mt, message, err := ws.ReadMessage()
		// 		if err != nil {
		// 			fmt.Println(err)
		// 			break
		// 		}
		// 		//If client message is ping will return pong
		// 		if string(message) == "ping" {
		// 			message = []byte("pong")
		// 		}
		// 		//Response message to client
		// 		err = ws.WriteMessage(mt, message)
		// 		if err != nil {
		// 			fmt.Println(err)
		// 			break
		// 		}
		// 	}
	})
}
