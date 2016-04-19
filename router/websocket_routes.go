package router



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
	})))*/