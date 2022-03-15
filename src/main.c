#include <stdio.h>
#include "../icl/proto.h"

/*
static const char *wm[] =  {
        "dwm", "i3wm", "bspwn", "xmonad", "qtile", NULL
};
*/

void gest_env(char **env)
{
        char *shell = NULL;
        char *term = NULL;

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
        printf("Shell: %s\n", &shell[15]);
        printf("Term : %s\n", term);
}

int main(int ac, char **av, char **env)
{
        (void)ac;
        (void)av;
        gest_env(env);
        return 0;
}
