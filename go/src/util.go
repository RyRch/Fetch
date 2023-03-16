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

func is_charset(c byte, charset []byte) bool {
    for i := 0; i < len(charset); i++ {
        if c == charset[i] {
            return true
        }
    }
    return false
}

func count_rows(str []byte, charset []byte) int {
	rows := 0

	for i := 0; i < len(str); i++ {
		if is_charset(str[i], charset) {
			rows++
		}
	}
	return rows
}

func count_cols(str []byte, rows int, charset []byte) []int {
    cols := make([]int, rows)
    x := 0

    for y := 0; y < rows; y++ {
        for ; !is_charset(str[x], charset); x++ {
            cols[y]++;
        }
        x++
    }
    return cols
}

func str2arr(str []byte, charset []byte) [][]byte {
    rows := count_rows(str, charset)
    cols := count_cols(str, rows, charset)
    arr := make([][]byte, rows)
    x := 0

    for y := 0; y < rows; y++ {
        arr[y] = make([]byte, cols[y])
        for w := 0; w < cols[y]; x++ {
            arr[y][w] = str[x]
            w++
        }
        x++
    }
    return arr
}
