// 418 final proj 
// GPU code on RPi
// Xiaofan Li
// Apr. 28th 2014



#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <assert.h>
#include <unistd.h>

#include "bcm_host.h"

#include "GLES/gl.h"
#include "EGL/egl.h"
#include "EGL/eglext.h"


#define PATH "./"
#define IMAGE_SIZE 128

// loads the file to convolve
static void load_file(char *iname, char *img_buf) {
  FILE *infile = NULL, *outfile = NULL;

  int bytes_read, image_sz = IMAGE_SIZE*IMAGE_SIZE*3;
  img_buf = malloc(image_sz);
   
  infile = fopen(PATH inname, "rb");
  
  if (infile && state->tex_buf1)
  {
     bytes_read=fread(img_buf, 1, image_sz, infile);
     assert(bytes_read == image_sz);  // some problem with file?
     fclose(infile);
  }

  return;
}

//now ready to convolve
static void process(char *img_buf){
  


}



int main(char* argv[], int argc){
  //initialize and pass in the file name
  //TODO parsing
  char *file; 

  //allocate for the file read
  char *img_buf = malloc(IMAGE_SIZE*IMAGE_SIZE*3);

  //load the file 
  load_file(file,img_buf);
  
  //process the buffer and output at "out.ppm"
  process(img_buf);
  return 0;
}



