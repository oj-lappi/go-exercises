
package main

import (
    "fmt"
    "os"
    "bufio"
)

//Tillåtna färger
var tillåtna_färger = []string{"svart","vit","röd","grön","blå","gul"}
var rätt_rad = []string{"röd","svart","blå","svart"}

var antal_gissningar = 12

func main() {
        //Skapa en scanner
        scanner := bufio.NewScanner(os.Stdin)
        
        fmt.Printf("Tillåtna färger: %v",tillåtna_färger)
        
        for ;antal_gissningar > 0;antal_gissningar--{
          fmt.Print("Gissa:")
          //Läs input
          scanner.Scan()
          text := scanner.Text()

          //Gör något med inputten
          fmt.Println("Du gissade:",text)
          
          //Testa om inputten är tillåten
          //4 färger, tillåtna färger
          
          //Kolla hur många som är rätt
          
          //Printa ut hur många som var rätt
          
          //Minska på antalet gissningar
        }
        fmt.Printf("Rätt rad var %v",rätt_rad)
}
