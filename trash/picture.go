package main

import "fmt"
import "os"
import "image"
import "image/png"
import "image/color"

func main(){

    const xsize = 400
    const ysize = 600

    pic := image.NewGray(image.Rect(0, 0, xsize, ysize))

    for x:= 0; x < xsize; x ++{

        for y := 0; y < ysize; y ++{
            pic.SetGray(x, y, color.Gray{200})
        }
    }

    file, err := os.Create("my.png")

    if err != nil{
        fmt.Println(err)
        return 
    }

    png.Encode(file, pic)

    file.Close()
}