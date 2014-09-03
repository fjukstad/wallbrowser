package main

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/fjukstad/gowebsocket"
)

type Key struct {
	Code int
}

func main() {
	ip := "10.1.1.60"
	port := ":9192"

	c, err := gowebsocket.NewClient(ip, port)

	if err != nil {
		log.Panic(err)
	}

	for {
		recv := c.Receive()
		log.Println(recv)
		key := new(Key)

		err = json.Unmarshal([]byte(recv), key)
		if err != nil {
			log.Panic(err)
		}

		var keyCode string
		switch key.Code {
		case 37:
			log.Println("Left")
			keyCode = "0xff51"
		case 38:
			log.Println("Up")
			keyCode = "0xff52"
		case 39:
			log.Println("Right")
			keyCode = "0xff53"
		case 40:
			log.Println("Down")
			keyCode = "0xff54"
		default:
			log.Println(key.Code, " - don't know that key")
		}

		log.Println("Translates to", keyCode)
		if keyCode != "" {
			cmd := exec.Command("xdotool", "key", keyCode)
			out, err := cmd.Output()
			if err != nil {
				log.Println("Could not execute command", err, string(out))
				log.Println(cmd)
			}
		}

	}
}
