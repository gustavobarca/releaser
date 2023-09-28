package main

import (
	"fmt"
	"os"
	"os/exec"
)

func createBranch(branchName string) {
	cmd := exec.Command("git", "checkout", "-b", branchName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func push(branchName string) {
	cmd := exec.Command("git", "push", "origin", branchName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func open(path string) {
	err := os.Chdir(path)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: releaser <repo_path> <version>")
		os.Exit(1)
	}

	repo_path := os.Args[1]
	version := os.Args[2]

	branchName := fmt.Sprintf("release/%s", version)

	open(repo_path)
	createBranch(branchName)

	fmt.Printf("Created and checked out branch '%s'\n", branchName)

	push(branchName)

	fmt.Printf("Pushed branch '%s' to remote\n", branchName)
}
