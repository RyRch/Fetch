package main

import (
    "os"
    "fmt"
    "syscall"
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

type DiskStatus struct {
	All  uint64 `json:"All"`
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func DiskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = fs.Blocks * uint64(fs.Bsize)
	return
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

func print_term() string {
    term := os.Getenv("TERM_PROGRAM")
    if term == "" {
        term = os.Getenv("TERM")
    }
    return term
}

func main() {
    disk := DiskUsage("/home/")
    fmt.Printf("\n%s%s%s%s\n", colors[WHITE], "      .___.      ", "┌───────────────────────────────────┐", colors[BLACK]);
    fmt.Printf("     /     \\           DISTRO: %s%s\n", colors[RESET], print_distro())
    fmt.Printf("%s    | O _ O |          WM/DE: %s%s\n", colors[RED], colors[RESET], print_wm())
    fmt.Printf("%s    /  \\_/  \\          TERMINAL: %s%s\n", colors[GREEN], colors[RESET], print_term())
    fmt.Printf("%s  .' /     \\ '.        SHELL: %s%s\n", colors[YELLOW], colors[RESET], print_shell())
    fmt.Printf("%s / _|       |_ \\       CPU: %s%s\n", colors[PURPLE], colors[RESET], print_cpu())
    fmt.Printf("%s(_/ |       | \\_)      GPU: %s%s\n", colors[BLUE], colors[RESET], print_gpu())
    fmt.Printf("%s    \\       /          MEMORY: %s%d gb\n", colors[CYAN], colors[RESET], print_memory())
    fmt.Printf("%s   __\\_>-<_/__         DISK: %s%d gb\n", colors[WHITE], colors[RESET], uint64(disk.All)/uint64(GB))
    fmt.Printf("%s%s%s%s\n\n", colors[WHITE], "   ~;/     \\;~   ", "└───────────────────────────────────┘", colors[RESET]);
}
