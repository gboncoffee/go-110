package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func readline(prompt string) string {
    fmt.Print(prompt)
    stdin := bufio.NewScanner(os.Stdin)
    stdin.Scan()
    return stdin.Text()
}

func readint(prompt string) int {
    for {
        i, err := strconv.Atoi(readline(prompt))
        if nil == err && 0 <= i {
            return i
        }
        fmt.Println("please enter a valid positive integer!")
    }
}

func is_binary(s string) bool {
    for _, c := range s {
        if '0' != c && '1' != c {
            return false
        }
    }
    return true
}

func readBinaryString(prompt string) string {
    for {
        s := readline(prompt)
        if is_binary(s) {
            return s
        }
        fmt.Println("please enter a string of only 1 and 0!")
    }
}

func applyRule110_chars(bef byte, act byte, aft byte) string {
    arr := [3]byte{bef, act, aft}
    switch {
    case [3]byte{'1', '1', '1'} == arr:
        return "0"
    case [3]byte{'1', '1', '0'} == arr:
        return "1"
    case [3]byte{'1', '0', '1'} == arr:
        return "1"
    case [3]byte{'1', '0', '0'} == arr:
        return "0"
    case [3]byte{'0', '1', '1'} == arr:
        return "1"
    case [3]byte{'0', '1', '0'} == arr:
        return "1"
    case [3]byte{'0', '0', '1'} == arr:
        return "1"
    case [3]byte{'0', '0', '0'} == arr:
        return "0"
    }

    return "0"
}

func applyRule110(s string) string {
    next := applyRule110_chars('0', s[0], s[1])
    for i := 1; len(s) - 1 > i; i++ {
        next += applyRule110_chars(s[i - 1], s[i], s[i + 1])
    }
    next += applyRule110_chars(s[len(s) - 2], s[len(s) - 1], '0')
    return next
}

func main() {
    fmt.Printf("===\nRule 110\n===\n")
    n := readint("number of iterations: ")
    s := readBinaryString("initial state: ")
    fmt.Println(s)
    for n := n; 0 <= n; n-- {
        s = applyRule110(s)
        for _, c := range s {
            if '0' == c {
                fmt.Print(" ")
            } else {
                fmt.Print("1")
            }
        }
        fmt.Println()
    }
    fmt.Println(s)
}
