// ------------------------------------
// Everything is about  pass by value
// ------------------------------------

// Pointer server only 1 purpose: sharing
// Pointer shares values across the program boundary.
// There are several types of program boundary. The most common one  is between function calls.
// We can also have a boundary between Goroutines when we will discuss it later.

package main

type user struct {
	name  string
	email string
}

func main() {
	count := 10
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	increment1(count)
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	increment2(&count)
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	stayOnStack()
	escapeToHeap()
}

func increment1(inc int) {
	inc++
	println("inc1:\tValue Of[", inc, "]\tAddr Of[", &inc, "]")
}

func increment2(inc *int) {
	*inc++
	println("inc2:\tValue Of[", inc, "]\tAddr Of[", &inc, "]\tValue Points To[", *inc, "]")
}

func stayOnStack() user {
	u := user{
		name:  "ycliu912",
		email: "ycliu912@qq.com",
	}

	return u
}

func escapeToHeap() *user {
	u := user{
		name:  "ycliu912",
		email: "ycliu912@qq.com",
	}
	return &u
}
