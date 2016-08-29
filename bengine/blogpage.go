package bengine

import (
	"html/template"
	"io/ioutil"
	"strings"

	md "github.com/russross/blackfriday"
)

type PostHeaders struct {
	URL        string
	Title      string
	Date       string
	PostLength int
	PostImage  string
}

type Post struct {
	PostHeaders
	Body template.HTML
}

func GetHeaders(file string) (PostHeaders, error) {
	b, err := GetPost(file)
	return b.PostHeaders, err
}

func getKeyValuePair(what string) (string, string) {
	what = strings.TrimSpace(what)
	i := strings.Index(what, "=")
	if i == -1 {
		return "", ""
	}
	return strings.TrimSpace(what[:i]), strings.TrimSpace(what[i+1:])
}

func GetPost(file string) (Post, error) {
	b, err := ioutil.ReadFile("./posts/" + file)
	if err != nil {
		return Post{}, err
	}
	resp := string(b)
	splitted := strings.Split(resp, "\n")
	var headerstarted bool = false
	var headerended bool = false
	var bh Post
	bh.PostImage = "/static/no-image.png"
	var body string = ""
	for _, val := range splitted {
		if headerended && val != "" {
			body += val + "\n"
		}
		if strings.TrimSpace(val) == "+++" {
			if !headerstarted {
				headerstarted = true
			} else {
				headerended = true
			}
		}
		if headerstarted {
			key, value := getKeyValuePair(val)
			if strings.ToUpper(key) == "TITLE" {
				bh.Title = value
			}
			if strings.ToUpper(key) == "DATE" {
				bh.Date = value
			}
			if strings.ToUpper(key) == "IMAGE" {
				bh.PostImage = value
			}
		}
	}
	bh.URL = "/posts/" + file
	bh.Body = template.HTML(string(md.MarkdownCommon([]byte(body))))
	bh.PostLength = len(body)
	return bh, nil
}
