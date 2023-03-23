BIN	=	fetch

SRC	=	src/fetch.go\
		src/print.go\
		src/str2arr.go\
		src/util.go\

CC	=	gccgo

all:
	$(CC) $(SRC) -o $(BIN)

clean:
	@rm $(BIN)

install: all
	sudo cp fetch /usr/local/bin/	

uninstall: clean
	sudo rm /usr/local/bin/fetch

re: clean all
