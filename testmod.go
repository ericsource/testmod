package testmod

import (
    "fmt"
)

func Hi(name string) string {
    return fmt.Sprintf("Hi, %s", name)
}

func New(name string) string {
    return fmt.Sprintf("New, %s", name)
}

