package main

import (
	profile "github.com/pkg/profile"
)

func main() {
    defer profile.Start(profile.CPUProfile).Stop()
}