package main

import "fmt"

func imprimir(color string, msg string){
  colors := map[string]string{
    "c" : "\033[36m",
    "r" : "\033[31m",
    "m" : "\033[35m",
  }
  msg = colors[color] + msg + "\033[0m"
  fmt.Println(msg)
}


func main(){
  imprimir("m", "Bienvenido al castillo del terror\n")

  var come_años int;

  edad := 43;
  nombre := "jolines";

  fmt.Println(edad);
  fmt.Println(nombre);

  fmt.Scanf("%d", &come_años)

  fmt.Println(come_años)

}
