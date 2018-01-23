// package for random helper utilities

package wmb

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// Clear clears the terminal screen
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// JSONToString converts JSON file to string
func JSONToString(fileName string) (string, error) {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// TermNotify notifies terminal with data
func TermNotify(data string) error {
	if err := exec.Command("terminal-notifier", "-message", data).Run(); err != nil {
		return err
	}
	return nil
}
