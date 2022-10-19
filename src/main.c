#include "../icl/proto.h"
#include "../icl/config.h"

void get_env(void)
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

char *get_gtk(const char *file, const char *info)
{
        char **arr = NULL;
        char *str = NULL;

        arr = str_to_tab((char *)file, "=\n");
        for (int i = 0; arr[i] != NULL; i++) {
                if (ft_strcmp((char *)info, arr[i]))
                        str = ft_strdup(arr[i + 1]);
        }
        return str;
}

int main(void)
{
        struct utsname uname_pointer;

        uname(&uname_pointer);
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
