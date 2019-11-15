package main


import (
    "fmt"
)




type People interface {
    ReturnName() string
}


type Student struct {
    Name string
}

func (s *Student) ReturnName() string {
    return s.Name
}

func dumpName( p People){
    fmt.Println(p.ReturnName())
}


func main(){
    cbs := Student{Name: "Einsn Liu"}

    var a People

    a = &cbs

    name := a.ReturnName()

    fmt.Println(name)

    dumpName(&cbs)

    var b [2]interface {}

    b[0] = 10
    b[1] = "string"

    fmt.Printf("%+v", b)

    _, isInt := b[0].(int)

    fmt.Println(isInt)

}