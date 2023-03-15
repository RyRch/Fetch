package main

import (
    "fmt"
    "os"
    "ioutil"
    "strconv"
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
	Used uint64 `json:"Used"`
	Free uint64 `json:"Free"`
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
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
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

func print_distro() string {
    // we still don't know other outputs
    list := [6]string{"arch", "fedora", "gentoo", "kali", "debian", "bsd"}
    release, _ := ioutil.ReadFile("/proc/sys/kernel/osrelease")

    for i := 0; i < 6; i++ {
        if strstr(list[i], string(release)) {
            return list[i] + " linux"
        }
    }
    return ""
}

func print_memory() int {
    file, _ := ioutil.ReadFile("/proc/meminfo")
    arr := str2arr(file, []byte("\n:"))
    part := string(arr[1])
    mem, _ := strconv.Atoi(part[7:len(part)-3])
    mem /= 1000000
    arr = nil

    return mem
}

func print_cpu() string {
    file, _ := ioutil.ReadFile("/proc/cpuinfo");
    arr := str2arr(file, []byte("\n:"))
    part := string(arr[9])
    arr = nil

    return part
}

func main() {
	  disk := DiskUsage("/")
	  fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
    fmt.Printf("%s%s%s\n", colors[WHITE], "┌─────────────────────────┐", colors[BLACK]);
    fmt.Printf("    Distro: %s%s\n", colors[RESET], print_distro())
    fmt.Printf("%s    WM/DE: %s%s\n", colors[RED], colors[RESET], "dwm")
    fmt.Printf("%s    TERM: %s%s\n", colors[GREEN], colors[RESET], os.Getenv("TERM_PROGRAM"))
    fmt.Printf("%s    SHELL: %s%s\n", colors[YELLOW], colors[RESET], print_shell())
    fmt.Printf("%s    CPU: %s%s\n", colors[PURPLE], colors[RESET], print_cpu())
    fmt.Printf("%s    GPU: %s%s\n", colors[BLUE], colors[RESET], "intel uhd 620")
    fmt.Printf("%s    MEM: %s%d gb\n", colors[CYAN], colors[RESET], print_memory())
    fmt.Printf("%s    DISK: %s%.2f gb%s\n", colors[WHITE], colors[RESET], float64(disk.All)/float64(GB))
    fmt.Printf("%s%s%s\n", colors[WHITE], "└─────────────────────────┘", colors[RESET]);
}
