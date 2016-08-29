package bengine

import "github.com/aki237/salt"

var URLS salt.URLS = salt.URLS{
	{Routename: "favicon", Pattern: "^/favicon.ico$", Handler: fav},
	{Routename: "payload", Pattern: "^/payload$", Handler: payload},
	{Routename: "post", Pattern: "^/posts/<any:post>$", Handler: showpost},
	{Routename: "home", Pattern: "^/$", Handler: home},
}

func fav(w salt.ResponseBuffer, r *salt.RequestBuffer) {
	salt.Redirect(w, r, "/static/favicon.png", 302)
}
