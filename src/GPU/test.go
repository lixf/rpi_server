
package main

//#include "gpu.c"
import "C"
import "fmt"

func main() {

  c := C.gpu_prog();
  fmt.Println("done",c)

}

