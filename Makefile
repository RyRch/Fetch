BIN	=	fetch
SRC	=	src/*.go
CC	=	gccgo

all:
	$(CC) $(SRC) -o $(BIN)

clean:
	@rm $(BIN)

re: clean all
