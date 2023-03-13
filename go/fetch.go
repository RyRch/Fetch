package main

import (
    "fmt"
    "os"
    "strconv"
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

func print_distro() string {
    // we still don't know other outputs
    list := [6]string{"arch", "fedora", "gentoo", "kali", "debian", "bsd"}
    release, _ := os.ReadFile("/proc/sys/kernel/osrelease")

    for i := 0; i < 6; i++ {
        if strstr(list[i], string(release)) {
            return list[i] + " linux" 
        }
    }
    return ""
}

func print_memory() int {
    file, _ := os.ReadFile("/proc/meminfo")
    arr := str2arr(file, []byte("\n:")) 
    part := string(arr[1])
    mem, _ := strconv.Atoi(part[7:len(part)-3])
    mem /= 1000000
    arr = nil

    return mem
}

func print_cpu() string {
    file, _ := os.ReadFile("/proc/cpuinfo");
    arr := str2arr(file, []byte("\n:"))
    part := string(arr[9])
    arr = nil

    return part
}


func main() {
    fmt.Printf("%s%s%s\n", colors[WHITE], "┌─────────────────────────┐", colors[BLACK]);
    fmt.Printf("    Distro: %s%s\n", colors[RESET], print_distro())
    fmt.Printf("%s    WM/DE: %s%s\n", colors[RED], colors[RESET], "dwm")
    fmt.Printf("%s    TERM: %s%s\n", colors[GREEN], colors[RESET], os.Getenv("TERM_PROGRAM"))
    fmt.Printf("%s    SHELL: %s%s\n", colors[YELLOW], colors[RESET], print_shell())
    fmt.Printf("%s    CPU: %s%s\n", colors[PURPLE], colors[RESET], print_cpu())
    fmt.Printf("%s    GPU: %s%s\n", colors[BLUE], colors[RESET], "intel uhd 620")
    fmt.Printf("%s    MEM: %s%d%s\n", colors[CYAN], colors[RESET], print_memory(), " gb")
    fmt.Printf("%s    DISK: %s%s\n", colors[WHITE], colors[RESET], "256 gb")
    fmt.Printf("%s%s%s\n", colors[WHITE], "└─────────────────────────┘", colors[RESET]);
}
