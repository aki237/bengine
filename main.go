package main

import (
	"fmt"
	"os/exec"

	"os"

	"github.com/aki237/bengine/bengine"
	"github.com/aki237/salt"
)

func main() {
	if _, err := os.Stat("./posts/"); os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", "https://github.com/aki237/bengine-posts", "posts")
		cmd.Output()
	}
	fmt.Println(salt.Configure("salt.json"))
	salt.Add404(NotFound)
	fmt.Println(salt.AddRootApp(bengine.App))
	fmt.Println(salt.RunAt(":" + os.Getenv("PORT")))
}

//Not found function
func NotFound(w salt.ResponseBuffer, r *salt.RequestBuffer) {
	w.Write([]byte("The page you are looking for doesn't exist"))
}
