// package for random helper utilities

package wmb

import (
	"log"
	"os"
	"os/exec"

	"github.com/waymobetta/goslack"
)

func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func Slack(payload string) {
	_, err := goslack.Send(payload)
	if err != nil {
		log.Fatal(err)
	}
}
