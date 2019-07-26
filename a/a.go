package main // import "github.com/amshinde/go-modules-by-example-submodules/a"

import (
	"fmt"
	"github.com/amshinde/go-modules-by-example-submodules/b"
)

const Name = b.Name

func main() {
	fmt.Println(Name)
}
