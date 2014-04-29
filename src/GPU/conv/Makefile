#	Xiaofan Li
#	15418	final project
# could possibly be changed for the real GPU tasks
#
#


CC = /usr/bin/gcc
CFLAGS = -Wall -g -Werror

FILES = gpu

# compile small testing program
hello: main.c gpu.c
	$(CC) $(CFLAG) -o gpu main.c gpu.c

# clean up
clean:
	rm -f $(FILES) *.o *~
