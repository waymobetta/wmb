package wmb

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime/debug"
	"time"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

// Clear clears the terminal screen
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// Elapsed used to time function execution
func Elapsed(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// TermNotify notifies terminal with data
// @return error
func TermNotify(data string) error {
	if err := exec.Command("terminal-notifier", "-message", data).Run(); err != nil {
		log.Fatal("[!] Error encountered when attemping to notify user via TermNotify\n", err)
	}
	return nil
}

// ReadFileByLine reads a file line-by-line
// @return slice
// @return error
func ReadFileByLine(path string) ([]string, error) {
	var data []string

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Println("[!] Error encountered when opening file\n", err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data, nil
}

// Require is a mirror of Require in Solidity
// rather than reverting a transaction, will panic loudly and print stacktrace
func Require(a interface{}, b interface{}) {
	if a != b {
		fmt.Println("")
		color.Red("[!] Execution halted: requirement not fulfilled!\nDetails: %v != %v\n\n", a, b)
		debug.PrintStack()
		os.Exit(0)
	}
}
