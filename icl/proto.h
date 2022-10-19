#ifndef PROTO_H
# define PROTO_H

#include <stdbool.h>

int         ft_strlen(char *str);
bool        ft_strncmp(char *s1, char *s2, int n);
char        *ft_strdup(char *src);
bool        ft_strstr(char *str, char *to_find);
bool        ft_strcmp(char *s1, char *s2);
char        **str_to_tab(char *str, char *chars);
void        get_env(void);
const char *get_file(const char *file);
char        *get_wm(const char *file); 
int         get_mem(const char *file);
char        *get_cpu(const char *file);
char        *get_gtk(const char *file, const char *info);

#endif
