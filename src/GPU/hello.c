/* 15418 project
 * GPU on rpi testing code
 * Tests capability with cgo
 */

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <strings.h>
#include "sys/types.h"
#include "sys/wait.h"

int hello_test (){
  //multi-process
  pid_t pid;
  int status;

  //for execve
  char file[73] = "/opt/vc/src/hello_pi/hello_triangle/hello_triangle.bin";
  char *argv[] = {NULL};
  extern char **environ;
  
  //fork and execute
  if((pid=fork())!=0){
    printf("executing! from %d\n",(int)pid);
    if (execve(file,argv,environ) < 0){ 
      printf("execution failure, file: %s\n",file);
      exit(1);
    }   
  }

  //wait for the child to finish
  wait(&status);

  //checks return status
  if (WIFEXITED(status)){
    printf("child exited correctly!");
  }
  else {
    printf("problem with child process!");
  }

  return 0;
}
