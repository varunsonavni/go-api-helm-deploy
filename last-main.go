package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main1() {

	fmt.Println("fjdg")
	log.Println("ehrer")
	cmd := exec.Command("touch", "pp.go")
	cmd.Run()
	fmt.Println("After command sleep 10 seconds")
	fmt.Println(time.Now())
	fmt.Println(time.DateOnly)
	fmt.Println("After command")

}
