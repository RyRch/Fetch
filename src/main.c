#include "../icl/proto.h"

static const char *osrelease = "/proc/sys/kernel/osrelease";
static const char *hostname = "/proc/sys/kernel/hostname";
static const char *cpuinfo = "/proc/cpuinfo";
static const char *meminfo = "/proc/meminfo";
static const char *gtkfile = "/home/rr/.config/gtk-3.0/settings.ini";
static const char *xinitrc = "/home/rr/.xinitrc";

#include <unistd.h>
#include <fcntl.h>
#include <stdio.h>

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

#include "../icl/config.h"

int main(void)
{
    struct utsname uname_pointer;

    uname(&uname_pointer);
    (void)wm;
    get_env();
    printf("%sWinm   %s\n", colors[GREEN], 
            (char *)get_wm(get_file(xinitrc)));
    printf("%sKrnl   %s", colors[YELLOW], get_file(osrelease));
    printf("%sArch   %s\n", colors[PURPLE], uname_pointer.machine);
    printf("%sHost   %s", colors[BLUE], get_file(hostname));
    printf("%sMmry   %d GB\n", colors[CYAN], get_mem(get_file(meminfo)));
    printf("%sProc  %s\n", colors[RESET], get_cpu(get_file(cpuinfo)));
    printf("%sFont   %s\n", colors[WHITE], 
            get_gtk(get_file(gtkfile), "gtk-font-name"));
    printf("%sThem   %s\n", colors[RED], 
            get_gtk(get_file(gtkfile), "gtk-theme-name"));
    printf("%sIcon   %s\n%s", colors[GREEN], get_gtk(get_file(gtkfile), 
                "gtk-icon-theme-name"), colors[RESET]);
    return 0;
}
