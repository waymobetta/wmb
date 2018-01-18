// package for random helper utilities

package wmb

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func JsonToString(fileName string) (string, error) {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}
