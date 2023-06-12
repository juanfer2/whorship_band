package utilities

import "fmt"

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

// Map implements the basic map function.
func Map[In, Out any](arr []In, fn func(In) Out) []Out {
	out := make([]Out, 0, len(arr))
	for _, elem := range arr {
		out = append(out, fn(elem))
	}
	return out
}

// MapError is the same as `Map` but for functions that might return an error.
func MapError[In, Out any](arr []In, fn func(In, int) (Out, error)) ([]Out, error) {
	out := make([]Out, 0, len(arr))
	for i, elem := range arr {
		mapped, err := fn(elem, i)
		if err != nil {
			return nil, fmt.Errorf("MapError got an error in index %d, value %v: %w", i, elem, err)
		}
		out = append(out, mapped)
	}
	return out, nil
}

func Includes[In comparable](value In, arr []In) bool {
	for _, elem := range arr {
		if elem == value {
			return true
		}
	}

	return false
}

// Map implements the basic map function.
func Find[In any](arr []In, fn func(In) bool) In {
	var out In
	for _, elem := range arr {
		if fn(elem) {
			return elem
		}
	}
	return out
}

/*
type Data struct {
	Name string
}

r := utilities.Includes("b", []string{"a", "b", "c"})
fmt.Println(r)

a := utilities.Find([]string{"a", "b", "c"}, func(s string) bool {
	return s == "b"
})
fmt.Println(a)

d := Data{Name: "John"}
vvv := utilities.Find([]Data{d}, func(s Data) bool {
	return s.Name == "John"
})
fmt.Println(vvv)
*/
