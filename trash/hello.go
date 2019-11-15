package main

import "fmt"
//import "encoding/json"
import "os"

/*
type Event struct {
	Device   string        `json:"device"`
	Origin   float         `json:"origin"`
}
*/

type point struct {
    x, y int
}

func printTest() {
    //Go 为常规 Go 值的格式化设计提供了多种打印方式。例如，这里打印了 point 结构体的一个实例。
    p := point{1, 2}
    fmt.Printf("%v\n", p) // {1 2}
    //如果值是一个结构体，%+v 的格式化输出内容将包括结构体的字段名。
    fmt.Printf("%+v\n", p) // {x:1 y:2}
    //%#v 形式则输出这个值的 Go 语法表示。例如，值的运行源代码片段。
    fmt.Printf("%#v\n", p) // main.point{x:1, y:2}
    //需要打印值的类型，使用 %T。
    fmt.Printf("%T\n", p) // main.point
    //格式化布尔值是简单的。
    fmt.Printf("%t\n", true)
    //格式化整形数有多种方式，使用 %d进行标准的十进制格式化。
    fmt.Printf("%d\n", 123)
    //这个输出二进制表示形式。
    fmt.Printf("%b\n", 14)
    //这个输出给定整数的对应字符。
    fmt.Printf("%c\n", 33)
    //%x 提供十六进制编码。
    fmt.Printf("%x\n", 456)
    //对于浮点型同样有很多的格式化选项。使用 %f 进行最基本的十进制格式化。
    fmt.Printf("%f\n", 78.9)
    //%e 和 %E 将浮点型格式化为（稍微有一点不同的）科学技科学记数法表示形式。
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)
    //使用 %s 进行基本的字符串输出。
    fmt.Printf("%s\n", "\"string\"")
    //像 Go 源代码中那样带有双引号的输出，使用 %q。
    fmt.Printf("%q\n", "\"string\"")
    //和上面的整形数一样，%x 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示。
    fmt.Printf("%x\n", "hex this")
    //要输出一个指针的值，使用 %p。
    fmt.Printf("%p\n", &p)
    //当输出数字的时候，你将经常想要控制输出结果的宽度和精度，可以使用在 % 后面使用数字来控制输出宽度。默认结果使用右对齐并且通过空格来填充空白部分。
    fmt.Printf("|%6d|%6d|\n", 12, 345)
    //你也可以指定浮点型的输出宽度，同时也可以通过 宽度.精度 的语法来指定输出的精度。
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
    //要最对齐，使用 - 标志。
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
    //你也许也想控制字符串输出时的宽度，特别是要确保他们在类表格输出时的对齐。这是基本的右对齐宽度表示。
    fmt.Printf("|%6s|%6s|\n", "foo", "b")
    //要左对齐，和数字一样，使用 - 标志。
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
    //到目前为止，我们已经看过 Printf了，它通过 os.Stdout输出格式化的字符串。Sprintf 则格式化并返回一个字符串而不带任何输出。
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)
    //你可以使用 Fprintf 来格式化并输出到 io.Writers而不是 os.Stdout。
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}


type Name struct{
    series string
    product string
}

type Books struct {
   title string
   author string
   subject string
   book_id int
   name Name
}


func getStruct(b Books){


}


func structTest() {

    // 创建一个新的结构体
    //fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

    // 也可以使用 key => value 格式
    b := Books{
        title: "Go 语言", 
        author: "www.runoob.com", 
        subject: "Go 语言教程", 
        book_id: 6495407, 
        name: Name { series: "SNAME", product: "Product"},
    }

   // b.name.series = "serialName"
    //b.name.product = "productName"


    c := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407, Name {"aaa", "bbbb"}}

    // 忽略的字段为 0 或 空
    // 
    //b := Books{title: "Go 语言", author: "www.runoob.com"}

   fmt.Printf("%+v\n", b)
   fmt.Printf("%+v\n", c)
}

func arrayTest() {
  var a = [3][5]int {{1, 2, 3, 4, 5}, 
  {0, 9, 8, 7, 6}, 
  {3, 4, 5, 6, 7},
}
  var i, j int
  for i = 0; i < 3; i++ {
    for j = 0; j < 5; j++ {
    fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j])
    }
  }
}


func baseTest(){

    var a, b int = 3, 4
    var ptr *int
    ptr = &a

    if (a == 3){
        fmt.Println("I am fine")
    }else{
        fmt.Println("hello")
    }

    if b == 5 {
        fmt.Println("b == 5")
    }

    fmt.Println(a,b)
    fmt.Println("ptr=", ptr)
    fmt.Println("ptr=", *ptr)


    for a < 20{

        fmt.Printf("a=%v\n", a)

        a ++
    }

    //var array [10] int
    //var array = [10]int {0, 0, 34, 67, 89}
   var array = [...]int {0, 0, 34, 67, 89} 

    array[0] = 99

    fmt.Println(len(array))

    i := 0
    for i < len(array){
        fmt.Printf("a[%d]=%v\n", i, array[i])
        i ++
    }


}

func main(){


    structTest()

    //arrayTest()

    //printTest()
}

/*
func main(){
   fmt.Printf("hello world\n")

    // 现在我们看看解码JSON数据为Go数值
    //byt := []byte(`{"device":"RandomDevice1","origin":1.5610105130094372e+18,"readings":[{"name":"SensorOne","value":"65","origin":1.5610105130094372e+18}]}`)
 	byt := []byte(`{"device":"RandomDevice1","origin":1.5610105130094372e+10}`)


    // 我们需要提供一个变量来存储解码后的JSON数据，这里
    // 的`map[string]interface{}`将以Key-Value的方式
    // 保存解码后的数据，Value可以为任意数据类型
    var dat map[string]interface{}
 
    // 解码过程，并检测相关可能存在的错误
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)


    res := &Event{}
    err := json.Unmarshal(byt, &res)
    if err != nil {
    	panic(err)
    } else {
    	fmt.Println(res)
    }
}
*/
