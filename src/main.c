#include <stdio.h>
#include "../icl/proto.h"
#include <fcntl.h>
#include <unistd.h>
#include <stdlib.h>

static const char *wm[] =  {
        "dwm", "i3wm", "bspwn", "xmonad", "qtile", "awesomewm", NULL
};

const char *xinitrc = "/home/rr/.xinitrc";
static const char *osrelease = "/proc/sys/kernel/osrelease";
/*
static const char *cpuinfo = "/proc/cpuinfo";
static const char *meminfo = "/proc/meminfo";
*/
void red () {
  printf("\033[1;31m");
}
void yellow() {
  printf("\033[1;33m");
}
void reset () {
  printf("\033[0m");
}
void gest_env(char **env)
{
        (void)env;
//        char *shell = NULL;
 //       char *term = NULL;
/*
        for (int i = 0; env[i] != NULL; i++) {
                if (ft_strlen(env[i]) > 4 && ft_strncmp(env[i], "SHELL", 4))
                        shell = ft_strdup(env[i]);
                else if (ft_strlen(env[i]) > 5 
                                && ft_strncmp(env[i], "TERM=", 5)) {
                        term = ft_strdup(env[i]);
                        term = ft_strdup(&term[5]);
                } else if (ft_strlen(env[i]) > 13 
                                && ft_strncmp(env[i], "TERM_PROGRAM=", 13)) {
                        term = ft_strdup(env[i]);
                        term = ft_strdup(&term[13]);
                }

        }
        */
        yellow();
        printf("Shell : %s\n", getenv("SHELL"));
        //printf("Shell: %s\n", &shell[15]);
        red();
        printf("Term  : %s\n", getenv("TERM"));
        //printf("Term : %s\n", term);
        reset();
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

int main(int ac, char **av, char **env)
{
        (void)ac;
        (void)av;
        gest_env(env);
        printf("Wm   : %s\n", (char *)get_wm(get_file(xinitrc)));
        printf("Kernel: %s\n", get_file(osrelease));
        return 0;
}
