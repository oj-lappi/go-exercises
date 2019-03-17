
package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
        //Skapa en scanner
        scanner := bufio.NewScanner(os.Stdin)
        
        fmt.Print("Ge input:")

        //Läs input
        scanner.Scan()
        text := scanner.Text()

        //Gör något med inputten
        fmt.Println(text)

}
