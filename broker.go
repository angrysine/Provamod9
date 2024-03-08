package main

import (
	"os/exec"
)


func main() {
	cmd := exec.Command("mosquitto", "-c", "mosquitto.conf")
	cmd.Output()
}




  



