package main

import (
    "os/exec"
    "io/ioutil"
    "strconv"
    "fmt"
)

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
        //fmt.Printf("ERROR: no output for xprop")
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
