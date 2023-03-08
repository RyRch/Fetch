package main

import (
    "fmt"
    "os"
)

const (
    BLACK = iota
    RED 
    GREEN
    YELLOW
    PURPLE
    BLUE
    CYAN
    WHITE
    RESET
)

var colors = []string {
    BLACK: "\033[0;37m",
    RED: "\033[1;31m",
    GREEN: "\033[1;32m",
    YELLOW: "\033[1;33m",
    PURPLE: "\033[1;34m",
    BLUE: "\033[1;35m",
    CYAN: "\033[1;36m",
    WHITE: "\033[1;37m",
    RESET: "\033[0m",
}

/*
var paths = []string {
    kernel: "/proc/sys/kernel/osrelease",
    ostype: "/proc/sys/kernel/ostype",
    host: "/proc/sys/kernel/hostname",
    cpu: "/proc/cpuinfo",
    mem: "/proc/meminfo",
}
*/

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

func print_shell() string {
    str := os.Getenv("SHELL")
    list := [6]string{"bash", "zsh", "dash", "csh", "ksh", "fish"}
         
    for i := 0; i < 6; i++ {
        if strstr(list[i], str) == true {
            return list[i]
        }
    }  
    return ""
}

func main() {
    fmt.Printf("%s%s%s\n", colors[WHITE], "┌─────────────────────────┐", colors[BLACK]);
    fmt.Printf("    Distro: %s%s\n", colors[RESET], "arch linux")
    fmt.Printf("%s    WM/DE: %s%s\n", colors[RED], colors[RESET], "dwm")
    fmt.Printf("%s    TERM: %s%s\n", colors[GREEN], colors[RESET], "st tmux")
    fmt.Printf("%s    SHELL: %s%s\n", colors[YELLOW], colors[RESET], print_shell())
    fmt.Printf("%s    CPU: %s%s\n", colors[PURPLE], colors[RESET], "i5 8250U")
    fmt.Printf("%s    GPU: %s%s\n", colors[BLUE], colors[RESET], "intel uhd 620")
    fmt.Printf("%s    MEM: %s%s\n", colors[CYAN], colors[RESET], "20gb")
    fmt.Printf("%s    DISK: %s%s\n", colors[WHITE], colors[RESET], "256gb")
    fmt.Printf("%s%s%s\n", colors[WHITE], "└─────────────────────────┘", colors[RESET]);
}
