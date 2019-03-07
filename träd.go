
package main

import (
        "fmt"
)

type studentgrupp struct{
	name string
	undergrupper []studentgrupp
}

//Skapar datanom-studentgrupperna data18,data18st,dat17,dat17st och dat16. Sparar dem i en lista
var datanomer = []studentgrupp{
			studentgrupp{"data18",nil},
			studentgrupp{"data18st",nil},
			studentgrupp{"dat17",nil},
			studentgrupp{"dat17st",nil},
			studentgrupp{"dat16",nil}	}

//Skapar merkonom-studentgrupperna mera18,mer17 och mer16. Sparar dem i en lista
var merkonomer = []studentgrupp{
			studentgrupp{"mera18",nil},
			studentgrupp{"mer17",nil},
			studentgrupp{"mer16",nil}	}

//Skapar studentgrupperna merkonomer och datanomer med tidigare studentgrupper som undergrupper
//Skapar studentgruppen "handel och data" sparar datanomer och merkonomer som dess undergrupper
var handel_och_data = studentgrupp{"handel och data",[]studentgrupp{
			studentgrupp{"merkonomer",merkonomer},
			studentgrupp{"datanomer",datanomer}	}	}

func main() {
	fmt.Println(handel_och_data)
}
