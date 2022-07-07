package gee

import (
	"fmt"
	"reflect"
	"testing"
)

var nr = NewRouter()

func TestParsePattern(t *testing.T) {
	want := []string{"p", ":name", "doc"}
	got := parsePattern("/p/:name/doc")
	fmt.Println(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v but want %v\n", got, want)
	}

	want = []string{"p", "doc"}
	got = parsePattern("/p/*path/doc")
	fmt.Println(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v but want %v\n", got, want)
	}
}

func TestAddRoute(t *testing.T) {
	nr.addRoute("GET", "/", nil)
	nr.addRoute("GET", "/hello", nil)
	nr.addRoute("GET", "/world/:name/doc", nil)
	nr.addRoute("GET", "/world/*path", nil)
	nr.addRoute("POST", "/article", nil)
	fmt.Printf("%#v\n", nr)
}

func TestGetRoutes(t *testing.T){
	nr.addRoute("GET", "/hello/:name", nil)
	fmt.Println(nr, "------------------")
	node, params := nr.getRoutes("GET", "/hello/fanzone")
	if node == nil{
		t.Fatal("nil err")
	}
	if node.pattern != "/hello/:name"{
		t.Fatal("should match /hello/:name")
	}
	if params["name"] != "fanzone"{
		t.Fatal("name should be equal to fanzone")
	}
	fmt.Println(node, params)
}
