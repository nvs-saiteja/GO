package main
import "fmt"
import "github.com/nvs-saiteja/GO"

func main() {
  var a = 10 
  var p *int
  p = &a
  fmt.Println(p)
  fmt.Println(*p)
  fmt.Println(&a)
  fmt.Println(&p)
  abc(10)
  cal.add(10,12)
}

func abc(a int){
   fmt.Println(a*a)
}
