package main

import (
//    "fmt"
)

func strncmp(size int, str string, src string) bool {
    for i := 0; i < size; i++ {
        if src[i] != str[i] {
            return false
        }
    }
    return true
}

func strstr(str string, src string) bool {
    for i := 0; i < len(src); i++ {
        test := src[i:]
        if strncmp(len(str), str, test) == true {
            return true
        }
    }
    return false
}

func is_num(c byte) bool {
    if c >= '0' && c <= '9' {
        return true
    }
    return false
}

func clean4atoi(arr []byte) []byte {
    size := 0
    for i := 0; i < len(arr); i++ {
        if is_num(arr[i]) {
            size++
        }
    }
    cln := make([]byte, size)
    w := 0
    for i := 0; w < size; i++ {
        if is_num(arr[i]) {
            cln[w] = arr[i]
            w++
        }
    }
    return cln
}
