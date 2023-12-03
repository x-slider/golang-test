package main

import (
	"fmt"
	"os"
)

func main() {
	secret := os.Getenv("TEST_SECRET")
	fmt.Println("The secret is:", secret)
}
