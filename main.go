package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile(".env")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("no .env file found")
			return
		}
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(file), "\n")

	if len(lines) == 0 {
		fmt.Println("no values found in .env file")
		return
	}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		pair := strings.Split(line, "=")
		key := strings.TrimSpace(pair[0])
		value := strings.TrimSpace(pair[1])

		fmt.Printf("export %s=%s\n", key, value)
	}
}
