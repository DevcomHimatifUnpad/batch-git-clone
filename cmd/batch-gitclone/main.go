package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func gitClone(user, repo, dest string, done chan []byte) {
	gitURL := fmt.Sprintf("git@github.com:%s/%s", user, repo)
	cloneDir := fmt.Sprintf("%s/%s_%s", dest, user, repo)
	out, _ := exec.Command("git", "clone", gitURL, cloneDir).CombinedOutput()
	done <- out
}

func main() {
	var repoName string
	var dest string

	if len(os.Args) < 2 {
		fmt.Println("usage: ./batch-gitclone <repo name> <destinatnion folder>")
		return
	}

	repoName = os.Args[1]
	dest = os.Args[2]

	fmt.Println("clonning into " + dest)

	students := []string{
		"sheilaazhar",
		"sittiufairoh",
		"fauzanakmalh1",
		"bagasapk",
		"Libramawan",
		"Alvaniakrn",
		"alfarisg34",
		"alvin2105",
		"meiradwianaa",
		"sinamustopa1",
		"muhamadilhamh",
		"LowShort",
		"sarahnvnt",
		"ahmadf20",
		"biairmal",
		"delanikaotc",
		"fahrulazimi",
		"kefilino",
		"budiy12",
		"riosapta",
		"tykozidane",
		"okka-riswana",
		"bbyreed",
		"hanifxdp",
		"Namury",
		"daffalfa",
		"fadlanmp",
		"egyaranda02",
		"danielrama7",
		"kdknive",
		"RizkyAnggita",
		"NadhifalR",
		"ryzzzanu50",
		"rahmabatari",
		"aithrajbouty",
		"alfianfl",
		"mveestor",
		"dip23",
		"anneaudistyaf",
		"hafidhakhdan",
		"aprischa",
		"rsudanta",
		"FarizAlfairuz",
		"bipbipbop",
		"NaufalA",
	}

	done := make(chan []byte)

	for _, student := range students {
		go gitClone(student, repoName, dest, done)
		msg := <-done
		log.Printf("%s\n", msg)
	}
}
