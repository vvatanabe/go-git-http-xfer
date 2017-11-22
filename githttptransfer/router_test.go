package githttptransfer

import (
	"net/http"
	"regexp"
	"testing"
)

func Test_Router_Append_should_append_route(t *testing.T) {
	router := &router{}
	router.Add(&Route{
		http.MethodPost,
		regexp.MustCompile("(.*?)/foo"),
		func(ctx Context) {},
	})
	router.Add(&Route{
		http.MethodPost,
		regexp.MustCompile("(.*?)/bar"),
		func(ctx Context) {},
	})
	length := len(router.routes)
	expected := 2
	if expected != length {
		t.Errorf("router length is not %d . result: %d", expected, length)
	}
}

func Test_Router_Match_should_match_route(t *testing.T) {
	router := &router{}
	router.Add(&Route{
		http.MethodPost,
		regexp.MustCompile("(.*?)/foo"),
		func(ctx Context) {},
	})
	match, route, err := router.Match(http.MethodPost, "/base/foo")
	if err != nil {
		t.Errorf("error is %s", err.Error())
	}
	if http.MethodPost != route.Method {
		t.Errorf("http method is not %s . result: %s", http.MethodPost, route.Method)
	}
	if "/base/foo" != match[0] {
		t.Errorf("match index 0 is not %s . result: %s", "/base/foo", match[0])
	}
	if "/base" != match[1] {
		t.Errorf("match index 1 is not %s . result: %s", "/base", match[1])
	}
}

func Test_Router_Match_should_return_UrlNotFound_error(t *testing.T) {
	router := &router{}
	router.Add(&Route{
		http.MethodPost,
		regexp.MustCompile("(.*?)/foo"),
		func(ctx Context) {},
	})
	match, route, err := router.Match(http.MethodPost, "/base/hoge")
	if err == nil {
		t.Error("error is nil.")
	}
	if match != nil {
		t.Error("match is not nil.")
	}
	if route != nil {
		t.Error("route is not nil.")
	}
	switch err.(type) {
	case *URLNotFoundError:
		return
	}
	t.Errorf("error is not UrlNotFound. %s", err.Error())
}

func Test_Router_Match_should_return_MethodNotAllowed_error(t *testing.T) {
	router := &router{}
	router.Add(&Route{
		http.MethodPost,
		regexp.MustCompile("(.*?)/foo"),
		func(ctx Context) {},
	})
	match, route, err := router.Match(http.MethodGet, "/base/foo")
	if err == nil {
		t.Error("error is nil.")
	}
	if match != nil {
		t.Error("match is not nil.")
	}
	if route != nil {
		t.Error("route is not nil.")
	}
	if _, is := err.(*MethodNotAllowedError); !is {
		t.Errorf("error is not MethodNotAllowed. %s", err.Error())
		return
	}
}
