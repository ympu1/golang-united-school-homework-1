package solution

import (
	"github.com/kyokomi/emoji"
	"testing"
)

func TestGetMessage(t *testing.T) {
	got := GetMessage()
	want := emoji.Sprintf("Hello :world_map:!")
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
