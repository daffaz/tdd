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
		{
			"Slices",
			[]Profile{
				{33, "Bekasi"},
				{2, "Jakarta"},
			},
			[]string{"Bekasi", "Jakarta"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "Bekasi"},
				{2, "Jakarta"},
			},
			[]string{"Bekasi", "Jakarta"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Bekasi"}
			aChannel <- Profile{2, "Jakarta"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Bekasi", "Jakarta"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, val := range haystack {
		if val == needle {
			contains = true
			break
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
