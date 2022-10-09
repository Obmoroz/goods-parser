package main
package main

import "fmt"

func TransmitIt(R RadioInterface) {
	R.ReadInstruction()
	R.PlugBatary()
	R.SetFrequensy("435,500")
	R.DoTransmit("охо - хо")
	/*
		1) прочитать инструкцию
		2) вставить аккомулятор
		3) выставить частоту
		4) осуществить передачу
	*/
}

type RadioInterface interface {
	ReadInstruction()
	PlugBatary()
	SetFrequensy(Frequensy string)
	DoTransmit(message string)
}

type Radio3 struct {
	Radio1
}

type Radio1 struct {
	Frequensy string
}

func (r *Radio1) ReadInstruction() {
	fmt.Println("Radio1 ReadInstruction")
}

func (r *Radio1) PlugBatary() {
	fmt.Println("Radio1 PlugBatary")
}

func (r *Radio1) SetFrequensy(Frequensy string) {
	r.Frequensy = Frequensy
}

func (r *Radio1) DoTransmit(message string) {
	fmt.Printf("Radio1 передаю на частоте %s сообщение %s \r\n", r.Frequensy, message)

}

type Radio2 struct {
	Frequensy string
}

func (r *Radio2) ReadInstruction() {
	fmt.Println("Radio2 ReadInstruction")
}

func (r *Radio2) PlugBatary() {
	fmt.Println("Radio2 PlugBatary")
}

func (r *Radio2) SetFrequensy(Frequensy string) {
	r.Frequensy = Frequensy
}

func (r *Radio2) DoTransmit(message string) {
	fmt.Printf("Radio2 передаю на частоте %s сообщение %s \r\n", r.Frequensy, message)

}

func main() {
	R := &Radio1{}
	TransmitIt(R)
	R2 := &Radio2{}
	TransmitIt(R2)
	R3 := &Radio3{}
	TransmitIt(R3)

}

