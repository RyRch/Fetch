#include <stdbool.h>
#include "../icl/proto.h"

bool ft_strstr(char *str, char *to_find)
{
    for (int i = 0; str[i] != '\0'; i++) {
        if (ft_strncmp(&str[i], to_find, ft_strlen(to_find)))
            return true;
    }
    return false;
}
