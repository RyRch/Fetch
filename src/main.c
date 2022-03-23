#include <stdio.h>
#include "../icl/proto.h"
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
/*
static const char *border[] = {
        " _____________________",
        "|                     |",
};
*/

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

void gest_env(void)
{
        char *str = NULL;

        printf("%sShll   %s\n", colors[WHITE], getenv("SHELL"));
        str = ft_strdup(getenv("TERM"));
        if (str == NULL || ft_strcmp(str, "screen")) {
                str = NULL;
                str = ft_strdup(getenv("TERM_PROGRAM"));
        }
        printf("%sTerm   %s\n", colors[RED], str);
}

const char *get_file(const char *file)
{
        char buf[BUFSIZ] = {0};
        char *bufp = NULL;
        int fd = 0;
        int r = 0;

        fd = open(file, O_RDONLY);
        if (fd == -1)
                return NULL;
        r = read(fd, buf, BUFSIZ);
        if (r <= 0)
                return NULL;
        bufp = ft_strdup(buf);
        return bufp;
}

char *get_wm(const char *file)
{
        char *str = NULL;

        for (int i = 0; wm[i] != NULL; i++) {
                if (ft_strstr((char *)file, (char *)wm[i])) {
                        str = ft_strdup((char *)wm[i]);
                        break;
                }
        }
        return str;
}

int get_mem(const char *file)
{
        char **arr = NULL;
        int mem = 0;

        arr = str_to_tab((char *)file, ":\n");
        mem = atoi((char *)arr[1]);
        mem /= 1000000;
        for (int i = 0; arr[i] != NULL; i++)
                free(arr[i]);
        return mem;
}

char *get_cpu(const char *file)
{
        char **arr = NULL;
        char *str = NULL;

        arr = str_to_tab((char *)file, ":\n");
        str = ft_strdup((char *)arr[9]);
        for (int i = 0; arr[i] != NULL; i++)
                free(arr[i]);
        return str;
}

int main(void)
{
        struct utsname uname_pointer;

        uname(&uname_pointer);
        printf("%s", colors[RED]);
        gest_env();
        printf("%s", colors[GREEN]);
        printf("Winm   %s\n", (char *)get_wm(get_file(xinitrc)));
        printf("%s", colors[YELLOW]);
        printf("Krnl   %s", get_file(osrelease));
        printf("%s", colors[PURPLE]);
        printf("Arch   %s\n", uname_pointer.machine);
        printf("%s", colors[BLUE]);
        printf("Host   %s", get_file(hostname));
        printf("%s", colors[CYAN]);
        printf("Mmry   %d GB\n", get_mem(get_file(meminfo)));
        printf("%s", colors[RESET]);
        printf("Proc  %s\n", get_cpu(get_file(cpuinfo)));
        return 0;
}
