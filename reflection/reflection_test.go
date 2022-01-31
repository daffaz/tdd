package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"Wawan"},
			[]string{"Wawan"},
		},
		{
			"Struct with two fields",
			struct {
				Name string
				City string
			}{"Wawan", "Bekasi"},
			[]string{"Wawan", "Bekasi"},
		},
		{
			"Struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Wawan", 2},
			[]string{"Wawan"},
		},
		{
			"Nested fields",
			Person{
				"Wawan",
				Profile{2, "Bekasi"},
			},
			[]string{"Wawan", "Bekasi"},
		},
		{
			"Pointer to things",
			&Person{
				"Wawan",
				Profile{2, "Bekasi"},
			},
			[]string{"Wawan", "Bekasi"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
}
