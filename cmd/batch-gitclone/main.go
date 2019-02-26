package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func gitClone(user, repo, dest string) ([]byte, error) {
	gitURL := fmt.Sprintf("git@github.com:%s/%s", user, repo)
	cloneDir := fmt.Sprintf("%s/%s_%s", dest, user, repo)
	out, err := exec.Command("git", "clone", gitURL, cloneDir).Output()
	if err != nil {
		return nil, err
	}

	fmt.Print(string(out))

	return out, nil
}

func file2slice(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var names []string
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	return names, nil
}

func main() {
	var repoName string
	var dest string
	var usersFile string

	if len(os.Args) < 4 {
		fmt.Println("\nUsage: ./batch-gitclone <repo name> <destinatnion folder> <list of username>")
		return
	}

	repoName = os.Args[1]
	dest = os.Args[2]
	usersFile = os.Args[3]

	fmt.Println("clonning into " + dest)

	students, err := file2slice(usersFile)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	done := make(chan []byte)

	for _, student := range students {
		go (func(student string) {
			msg, err := gitClone(student, repoName, dest)
			if err != nil {
				done <- []byte(err.Error())
				return
			}

			done <- msg
		})(student)
	}

	log.Printf("%s", <-done)
}
