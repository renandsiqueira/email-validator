// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"

	"strings"

	"log"

	"github.com/asaskevich/govalidator"

	"bufio"

	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	lines, err := readLines("emails_to_check.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	validEmails := []string{}
	invalidEmails := []string{}

	for i, line := range lines {

		r := strings.Split(line, ",")

		if govalidator.IsEmail(r[2]) && govalidator.IsExistingEmail(r[2]) {
			validEmails = append(validEmails, fmt.Sprintf("valid email,%s\n", line))
		} else {
			invalidEmails = append(invalidEmails, fmt.Sprintf("invalid email,%s\n", line))
		}

		fmt.Printf("line %d \n", i)
	}

	if err := writeLines(validEmails, "valid_emails.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}

	if err := writeLines(invalidEmails, "invalid_emails.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}
