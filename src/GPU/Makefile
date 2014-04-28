#	Xiaofan Li
#	15418	final project
# could possibly be changed for the real GPU tasks
#
#


CC = /usr/bin/gcc
CFLAGS = -Wall -g -Werror

FILES = hello

# compile small testing program
hello: main.c hello_help.c
	$(CC) $(CFLAG) -o hello main.c hello_help.c

# clean up
clean:
	rm -f $(FILES) *.o *~
