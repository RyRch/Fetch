SRC	=	src/main.c

CC =	gcc

OBJ =	$(SRC:.c=.o)

CFLAGS +=	-Wall -Werror -Wextra -g3

BIN =	rfetch

RM =	rm -rf
