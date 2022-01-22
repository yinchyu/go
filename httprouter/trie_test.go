package httprouter

import (
	"testing"
)

var tests = []struct {
	give string
	flag int
	want bool
}{
	{
		give: "hello",
		flag: 1,
		want: true,
	},
	{
		give: "world",
		flag: 1,
		want: true,
	},
	{
		give: "hello",
		flag: 0,
		want: true,
	},
	{
		give: "hel",
		flag: 0,
		want: true,
	},
	{
		give: "wor",
		flag: 0,
		want: true,
	},
}

func TestTrie(t *testing.T) {
	tree := newtrie()
	tree.insert("hello")
	tree.insert("world")
	for _, test := range tests {
		if test.flag == 0 {
			if s := tree.startwith(test.give); s != test.want {
				t.Errorf("startwith(%q) = %t, want %t", test.give, s, test.want)
			}
		} else {
			if s := tree.isword(test.give); s != test.want {
				t.Errorf("isword(%q) = %t, want %t", test.give, s, test.want)
			}
		}
	}
}

func TestRouter_GET(t *testing.T) {
	test()
}
