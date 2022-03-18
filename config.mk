SRC	=	src/main.c

LIB =	lib/ft_strdup.c\
		lib/ft_strlen.c\
		lib/ft_strncmp.c\
		lib/ft_strstr.c\
		lib/ft_strcmp.c

CC =	gcc

OBJ =	$(SRC:.c=.o)

LOBJ =	$(LIB:.c=.o)

NAME =	libmy.a

CFLAGS +=	-Wall -Werror -Wextra -g3

BIN =	rfetch

RM =	rm -rf
