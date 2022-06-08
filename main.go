package main

import (
    "fmt"
    "strings"
)

func filter(data []string, f func(string) bool) []string {
    fltd := make([]string, 0)

    for _, e := range data {
        if f(e) {
            fltd = append(fltd, e)
        }
    }

    return fltd
}

func main() {
    words := []string{"war", "water", "cup", "tree", "storm"}

    p := "w"
    res := filter(words, func(s string) bool {
        return strings.HasPrefix(s, p)
    })

    fmt.Println(res)

}