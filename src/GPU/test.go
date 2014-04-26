
package main

//#include "hello.c"
import "C"
import "fmt"

func main() {

  c := C.hello_test();
  fmt.Println("done",c)

}

