package main

import "fmt"

type Tortilla struct{
}
func (s *Tortilla) Ready(){
	fmt.Println("Tortilla is Ready!")
}
type Taco interface {
	Ready()
}

type Burito struct{
	meat string
}

type Nachos struct{
	meal string
}

func (n *Nachos) Ready() {
	fmt.Println("Nachos is Ready!")
}

func (this Burito) Ready(){
	fmt.Println("Today's meat is "+this.meat)

}

func NewBurito() Taco{
	return Burito{
		meat: "beef",
	}
}
func NewTacofromBurito() Taco{
	burito := NewBurito()
	
	//b := Burito(burito) // need type assertion
	//b.meat = "turkey"
	
	if b, ok:=burito.(Burito); ok{
	//if b, ok:=burito.(*Burito); ok{ // will fail unless *Burito implements Taco
	
	fmt.Println("Pick Burito's Meat")
	b.meat = "pork"
	b.Ready()
	return b
	}
	//burito.Ready()
	return burito
}

func NewTacofrom() Taco{
	return &Tortilla{}
}

func NewNachos() *Nachos{

	//burito :=NewTaco()
	burito := NewTacofromBurito()
	
	/*nachos := Nachos{
	  side: NewTaco(),
	}*/
	
	if t, ok:=burito.(*Nachos); ok{
		fmt.Println("return Taco as Nachos")
		return t
	}else{
	fmt.Println("return Taco as Nachos")
		return &Nachos{}
	}

}

func main() {
	nachos := NewNachos()
	nachos.Ready()
}
