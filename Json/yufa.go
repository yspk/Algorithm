package main
import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)
func main ( ) {
	var jsonBlob = [ ] byte ( ` [ 
        { "Name" : "Platypus" , "Order" : "Monotremata" } , 
        { "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" } ,
		{ "Name" : "余胜" ,     "Order" : "1234143大家"}
    ] ` )
	jsonString := hex.EncodeToString(jsonBlob)
	fmt.Println(jsonString)
	fmt.Println(jsonBlob)
	fmt.Println(string(jsonBlob))

	type Animal struct {
		Name  string
		Order string
	}
	var animals [ ] Animal

	err := json. Unmarshal ( jsonBlob , & animals )
	if err != nil {
		fmt. Println ( "error:" , err )
	}
	fmt. Printf ( "%+v" , animals )
}
