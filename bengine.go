package main

import (
	"fmt"

	"github.com/aki237/bengine/bengine"
	"github.com/aki237/salt"
)

func main() {
	fmt.Println(salt.Configure("salt.json"))
	salt.Add404(NotFound)
	fmt.Println(salt.AddRootApp(bengine.App))
	fmt.Println(salt.Run())
}

//Not found function
func NotFound(w salt.ResponseBuffer, r *salt.RequestBuffer) {
	w.Write([]byte("The page you are looking for doesn't exist"))
}
