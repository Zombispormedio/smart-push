package main

import (
	"github.com/labstack/echo"

	"github.com/Zombispormedio/smart-push/config"
)

func main() {

	server := echo.New()

	config.Middleware(server)

	/*e.GET("/ws", standard.WrapHandler(websocket.Handler(func(ws *websocket.Conn) {
		for {
			websocket.Message.Send(ws, "Hello, Client!")
			msg := ""
			log.WithFields(log.Fields{
				"animal": "walrus",
			}).Info("A walrus appears")

			err:=websocket.Message.Receive(ws, &msg)

			if err!= nil{
				fmt.Println(err)
				break
			}

			fmt.Println(msg)


		}
	})))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})*/
	
	config.Listen(server)

}
