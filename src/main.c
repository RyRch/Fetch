#include <stdio.h>
#include <stdbool.h>
#include <stdlib.h>

int ft_strlen(char *str)
{
        int i = 0;

        for ( ; str[i]; i++);
        return (i);
}

bool ft_strncpy(char *s1, char *s2, int n)
{
        for (int i = 0; i < n; i++) {
                if (s1[i] != s2[i])
                        return false;
        }
        return true;
}

char *ft_strdup(char *src)
{
        char *new = malloc(sizeof(char) * ft_strlen(src) + 1);

        for (int i = 0; src[i]; i++)
                new[i] = src[i];
        new[ft_strlen(src)] = '\0';
        return (new);
}

void gest_env(char **env)
{
        char *shell = NULL;
        char *term = NULL;

        for (int i = 0; env[i] != NULL; i++) {
                if (ft_strncpy(env[i], "SHELL", 4))
                        shell = ft_strdup(env[i]);
                if (ft_strncpy(env[i], "TERM", 4))
                        term = ft_strdup(env[i]);

        }
        printf("%s\n", shell);
        printf("%s\n", term);
}

int main(int ac, char **av, char **env)
{
        //char buf[BUFSIZ] = {0};
        gest_env(env);
        (void)ac;
        (void)av;
        return 0;
}
