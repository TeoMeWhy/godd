package utils

import (
	"reflect"
	"testing"
)

func TestFindElement(t *testing.T) {
	tests := []struct {
		name     string
		elements []string
		e        string
		want     int
	}{
		{
			name:     "nil slice",
			elements: nil,
			e:        "some element",
			want:     -1,
		},
		{
			name:     "empty slice",
			elements: []string{},
			e:        "some element",
			want:     -1,
		},
		{
			name:     "element not in slice",
			elements: []string{"foo", "bar"},
			e:        "some element",
			want:     -1,
		},
		{
			name:     "element in slice",
			elements: []string{"foo", "target", "bar"},
			e:        "target",
			want:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindElement(tt.elements, tt.e); got != tt.want {
				t.Errorf("FindElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		elements []string
		e        string
		want     []string
	}{
		{
			name:     "nil slice",
			elements: nil,
			e:        "some element",
			want:     nil,
		},
		{
			name:     "empty slice",
			elements: []string{},
			e:        "some element",
			want:     []string{},
		},
		{
			name:     "element not in slice",
			elements: []string{"foo", "bar"},
			e:        "some element",
			want:     []string{"foo", "bar"},
		},
		{
			name:     "element in slice",
			elements: []string{"foo", "target", "bar"},
			e:        "target",
			want:     []string{"foo", "bar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveElement(tt.elements, tt.e)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveElement() = %v, want %v", got, tt.want)
			}

			// Ensures that the result uses the same underlying array as the input elements
			if len(got) > 0 && &got[0:cap(got)][cap(got)-1] != &tt.elements[0:cap(tt.elements)][cap(tt.elements)-1] {
				t.Errorf("RemoveElement() should not allocate a new array")
			}
		})
	}
}
