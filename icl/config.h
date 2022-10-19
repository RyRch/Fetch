#include <stdio.h>
#include <fcntl.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/utsname.h>   /* Header for 'uname'  */

static const char *wm[] =  {
        "dwm", "i3wm", "bspwn", "xmonad", "qtile", "awesomewm", NULL
};

const char *xinitrc = "/home/rr/.xinitrc";
static const char *osrelease = "/proc/sys/kernel/osrelease";
static const char *hostname = "/proc/sys/kernel/hostname";
static const char *cpuinfo = "/proc/cpuinfo";
static const char *meminfo = "/proc/meminfo";
static const char *gtkfile = "/home/rr/.config/gtk-3.0/settings.ini";

enum {
        RED,
        YELLOW,
        PURPLE,
        BLUE,
        CYAN,
        WHITE,
        GREEN,
        RESET,
};

const char *colors[] = {
        [RED] = "\033[1;31m",
        [GREEN] = "\033[1;32m",
        [YELLOW] = "\033[1;33m",
        [PURPLE] = "\033[1;34m",
        [BLUE] = "\033[1;35m",
        [CYAN] = "\033[1;36m",
        [WHITE] = "\033[1;37m",
        [RESET] = "\033[0m",
};
