package gopher_test

import (
	"testing"

	gopher "github.com/daviidbartlett/go-mod"
)

func TestFavFood(t *testing.T){

	g := gopher.New("gary")

	got := g.FavFood()
	want := "crisps"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

}