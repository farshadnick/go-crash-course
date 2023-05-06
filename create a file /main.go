package main
import (
	"fmt"
	"os"

)


func main (){
	fmt.Println("input Number to call:")
	var number string;
    fmt.Scanln(&number)
    err := writeToFile(number)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("File created successfully")
    }
}


func writeToFile (number string)  error{
	file, err := os.Create("d:/mycall.call")
	if err != nil {
  	return  err
	}
	fmt.Fprintf(file, "number: %s", number)

	defer file.Close()
	return nil
}

 

