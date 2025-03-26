package main

import (
    "fmt"
)

func main() {
    result := []string{}
    str := "cat"
    strByte := []byte(str)
    permutation(strByte, []byte{}, &result)
    fmt.Println("result##", result)
}

func permutation(str, current []byte, result *[]string) {
    if len(str) == 0 {
        *result = append(*result, string(current))
        return
    }

    for i := 0; i < len(str); i++ {

        char := str[i]

        // create a new copy; which I did not do; This is important was causing the issue.
        remaining := append([]byte{}, str[:i]...)
        remaining = append(remaining, str[i+1:]...)

        // Create a new copy of current and append the character
        newCurrent := append([]byte{}, current...)
        newCurrent = append(newCurrent, char)

        permutation(remaining, newCurrent, result)
    }
}
