package main

import (
    "fmt"
    "os/exec"
    "os"
    "io/ioutil"
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

func print_distro() string {
    list := [7]string{"arch", "fedora", "gentoo", "kali", "debian", "bsd", "ubuntu"}
    release, _ := ioutil.ReadFile("/proc/version")

    for i := 0; i < 7; i++ {
        if strstr(list[i], string(release)) {
            return list[i] + " linux"
        }
    }
    return ""
}

func print_memory() int {
    file, _ := ioutil.ReadFile("/proc/meminfo")
    arr := str2arr(file, []byte("\n:"))
    cln := clean4atoi(arr[1])
    mem, _ := strconv.Atoi(string(cln))
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

func print_term() string {
    term := os.Getenv("TERM_PROGRAM")
    if term == "" {
        term = os.Getenv("TERM")
    }
    return term
}

func print_gpu() string {
    out, err := exec.Command("lspci").Output()
    part := ""

    if err != nil {
        fmt.Printf("ERROR: no output for gpu")
        return "unknow"
    }
    arr := str2arr(out, []byte("\n:"))
    for i := 0; i < len(arr); i++ {
        if strstr("VGA", string(arr[i])) {
            part = string(arr[i+1]) 
            break
        }
    }
    arr = nil
    return part
}

func get_wm_id() string {
    cmd := exec.Command("xprop", "-root", "-notype")
    out, err := cmd.Output()
    part := ""
    index := 0
    cnt := 0

    if err != nil {
        fmt.Printf("ERROR: no output for xprop")
        return ""
    }
    arr := str2arr(out, []byte("\n:"))
    for i := 0; i < len(arr); i++ {
        if strstr("_NET_SUPPORTING_WM_CHECK", string(arr[i])) {
            index = i+2
            break
        }
    }
    for i := 0; i < len(arr[index]); i++ {
        if is_num(arr[index][i]) {
            cnt = i
            break
        }
    }
    part = string(arr[index])
    arr = nil
    part = part[cnt:len(part)]
    return part
}

func print_wm() string {
    id := get_wm_id()

    if id == "" {
        return "unknow"
    }
    cmd := exec.Command("xprop", "-id", id, "-notype", "-f", "_NET_WM_NAME", "8t")
    list := [7]string{"dwm", "i3", "bspwm", "sway", "awesome", "xfwm", "xmonad"}
    res := ""

    out, err := cmd.Output()
    if err != nil {
        fmt.Printf("ERROR: no output for wm")
    }
    for i := 0; i < 7; i++ {
        if strstr(list[i], string(out)) {
            res = list[i]
            break
        }
    }
    return res
}

func main() {
    disk := DiskUsage("/home/")
    fmt.Printf("%s%s%s\n", colors[WHITE], "┌─────────────────────────┐", colors[BLACK]);
    fmt.Printf("    Distro: %s%s\n", colors[RESET], print_distro())
    fmt.Printf("%s    WM/DE: %s%s\n", colors[RED], colors[RESET], print_wm())
    fmt.Printf("%s    TERM: %s%s\n", colors[GREEN], colors[RESET], print_term())
    fmt.Printf("%s    SHELL: %s%s\n", colors[YELLOW], colors[RESET], print_shell())
    fmt.Printf("%s    CPU: %s%s\n", colors[PURPLE], colors[RESET], print_cpu())
    fmt.Printf("%s    GPU: %s%s\n", colors[BLUE], colors[RESET], print_gpu())
    fmt.Printf("%s    MEM: %s%d gb\n", colors[CYAN], colors[RESET], print_memory())
    fmt.Printf("%s    DISK: %s%d gb\n", colors[WHITE], colors[RESET], uint64(disk.All)/uint64(GB))
    fmt.Printf("%s%s%s\n", colors[WHITE], "└─────────────────────────┘", colors[RESET]);
}
