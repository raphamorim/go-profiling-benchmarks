package main

import (
	"fmt"

	profile "github.com/pkg/profile"
)

type Dog struct {
	Name string
	Age  int
}

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()


	bob := Dog{
		Name: "Bob",
		Age:  19,
	}

	clifford := Dog{
		Name: "Clifford",
		Age:  7,
	}

	dogs := []Dog{bob, clifford}

	fmt.Println("")

	for _, dog := range dogs {
		fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
		fmt.Printf("Addr: %p\n", &dog)
	}
}
