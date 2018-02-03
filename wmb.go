// package for random helper utilities

package wmb

import (
	"bufio"
	"io/ioutil"
	"log"
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
		log.Fatal("[!] Error encountered when reading JSON file\n", err)
	}
	return string(raw), nil
}

// TermNotify notifies terminal with data
func TermNotify(data string) error {
	if err := exec.Command("terminal-notifier", "-message", data).Run(); err != nil {
		log.Fatal("[!] Error encountered when attemping to notify user via TermNotify\n", err)
	}
	return nil
}

// ReadFileByLine reads a file line-by-line
func ReadFileByLine(path string, data []string) ([]string, int, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal("[!] Error encountered when opening file\n", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, len(data), nil
}

// WriteFile writes strings to a file, delimited in some fashion prior to being passed in as an argument
func WriteFile(path string, data string) error {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("[!] Error encountered when creating file\n", err)
	}
	defer file.Close()
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("[!] Error encountered when opening file\n", err)
	}
	defer f.Close()
	if _, err = f.WriteString(data); err != nil {
		log.Fatal("[!] Error encountered when writing to file\n", err)
	}
	return nil
}
