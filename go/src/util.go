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

func is_num(arr []byte) bool {
  for i := 0; i < len(arr); i++ {
    if arr[i] >= '0' && arr[i] <= '9' {
      return true
    }
  }
  return false
}
