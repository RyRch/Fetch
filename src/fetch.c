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
