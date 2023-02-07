#include <stdlib.h>
#include "../icl/proto.h"

char *ft_strdup(char *src)
{
    char *new = malloc(sizeof(char) * ft_strlen(src) + 1);

    for (int i = 0; src[i]; i++)
        new[i] = src[i];
    new[ft_strlen(src)] = '\0';
    return (new);
}
