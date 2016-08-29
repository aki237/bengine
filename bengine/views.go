package bengine

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/aki237/salt"
)

//YOUR VIEWS GO HERE
func payload(w salt.ResponseBuffer, r *salt.RequestBuffer) {
	b, _ := ioutil.ReadAll(r.Body)
	var gwh GitPushWH
	err := json.Unmarshal(b, &gwh)
	if err != nil {
		fmt.Fprint(w, "Json Error : "+err.Error())
		return
	}
	_, err = os.Stat("./posts/")
	_, errdg := os.Stat("./posts/.git")
	if os.IsNotExist(err) {
		cmd := exec.Command("git", "clone", "https://github.com/aki237/bengine-posts", "posts")
		cmd.Output()
	} else {
		if os.IsNotExist(errdg) {
			exec.Command("rm", "-rf", "posts/*").Output()
			cmd := exec.Command("git", "clone", "https://github.com/aki237/bengine-posts", "posts")
			cmd.Output()
		}
		cmd := exec.Command("git", "-C", "posts", "pull")
		cmd.Output()
	}
}

func showpost(w salt.ResponseBuffer, r *salt.RequestBuffer) {
	a, ok := r.URLParameters["post"].(string)
	if !ok {
		return
	}
	_, err := os.Stat("./posts/" + a)
	if err != nil {
		fmt.Fprintln(w, "Post Not found :", err.Error())
		return
	}
	b, err := GetPost(a)
	if err != nil {
		fmt.Fprintln(w, "Post Not found", err.Error())
		return
	}
	fmt.Fprint(w, b.Body)
}

func home(w salt.ResponseBuffer, r *salt.RequestBuffer) {
	ls, err := ioutil.ReadDir("./posts/")
	if os.IsNotExist(err) {
		fmt.Fprint(w, "Sorry No posts available right now.")
		return
	}

	var barray []PostHeaders
	for _, val := range ls {
		if !val.IsDir() {
			b, err := GetHeaders(val.Name())
			if err == nil {
				barray = append(barray, b)
			}
		}
	}
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		fmt.Fprint(w, "Some error occured : ", err.Error())
		return
	}
	t.Execute(w, barray)
}
