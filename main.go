package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/c-bata/go-prompt"
)

var (
	version  string
	revision string
)

// Executor handle input.
func Executor(h, c string) func(string) {
	return func(s string) {
		s = strings.TrimSpace(s)
		if s == "" {
			return
		} else if s == "quit" || s == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
			return
		}

		c = strings.TrimSpace(strings.Join([]string{c, s}, " "))
		c = fmt.Sprintf("ssh %s '%s'", h, c)

		cmd := exec.Command("/bin/sh", "-c", c)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
		}
	}
}

// Completer returns prompt suggestions.
func Completer(d prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No ssh host provided")
	}

	host := os.Args[1]
	cmd := ""
	if len(cmd) == 0 && !strings.Contains(host, "dokku") {
		cmd = "dokku"
	}

	fmt.Printf("dokku-prompt %s (rev-%s)\n", version, revision)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")

	defer fmt.Println("Bye!")

	p := prompt.New(
		Executor(host, cmd),
		Completer,
		prompt.OptionTitle("dokku: interactive dokku client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Yellow),
	)
	p.Run()
}
