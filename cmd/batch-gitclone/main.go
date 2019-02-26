package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func gitClone(user, repo, dest, gitBaseURL string) ([]byte, error) {
	gitURL := fmt.Sprintf("%s%s/%s", gitBaseURL, user, repo)
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
	var repoName, dest, usersFile, gitBaseURL string

	if len(os.Args) < 5 {
		fmt.Println("\nUsage: ./batch-gitclone <ssh/https> <repo name> <destinatnion folder> <listUsername.txt>")
		return
	}

	if os.Args[1] == "ssh" {
		gitBaseURL = "git@github.com:"
	} else if os.Args[1] == "https" {
		gitBaseURL = "https://github.com/"
	} else {
		log.Fatalf("Unknown protocol %s. Please choose ssh or https", os.Args[1])
		return
	}

	repoName = os.Args[2]
	dest = os.Args[3]
	usersFile = os.Args[4]

	log.Println(">>> Clonning into " + dest)

	students, err := file2slice(usersFile)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	done := make(chan []byte)

	for _, student := range students {
		go (func(student string) {
			msg, err := gitClone(student, repoName, dest, gitBaseURL)
			if err != nil {
				done <- []byte(err.Error())
				return
			}

			done <- msg
		})(student)
	}

	log.Printf("%s", <-done)
	log.Print(">>> Finished")
}
