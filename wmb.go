// package for random helper utilities

package wmb

import (
	"os"
	"os/exec"
)

func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
