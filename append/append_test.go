package append

import (
	"reflect"
	"testing"
)

func TestAppendByte(t *testing.T) {
	s := []byte{1, 2, 3}
	bytesToAppend := []byte{4, 5}
	got := AppendByte(s, bytesToAppend...)
	want := []byte{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
