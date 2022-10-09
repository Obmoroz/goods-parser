package main

import "fmt"

func TransmitIt(R RadioInterface, a Antenna) {
	R.ReadInstruction()
	R.PlugBatary()
	a.SetAntenna("antennas")
	R.SetChannal("435,500")
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
	SetChannal(Frequensy string)
	DoTransmit(message string)
}
type Antenna interface {
	SetAntenna(Frequensy string)
}

type Radio3 struct {
	Radio1
}

func (r *Radio3) DoTransmit(message string) {
	fmt.Printf("Radio3 передаю на частоте %s сообщение %s \r\n", r.Frequensy, message)

}

type Radio1 struct {
	Frequensy string
}

func (r *Radio1) SetAntenna(Frequensy string) {
	fmt.Println("antennna")
}

func (r *Radio1) ReadInstruction() {
	fmt.Println("Radio1 ReadInstruction")
}

func (r *Radio1) PlugBatary() {
	fmt.Println("Radio1 PlugBatary")
}

func (r *Radio1) SetChannal(Frequensy string) {
	r.Frequensy = Frequensy
	fmt.Println("Radio1 Frequensy")

}

func (r *Radio1) DoTransmit(message string) {
	fmt.Printf("Radio1 передаю на частоте %s сообщение %s \r\n", r.Frequensy, message)

}

type Radio2 struct {
	Frequensy string
}

func (r *Radio2) SetAntenna(Frequensy string) {
	fmt.Println("antennna")
}

func (r *Radio2) ReadInstruction() {
	fmt.Println("Radio2 ReadInstruction")
}

func (r *Radio2) PlugBatary() {
	fmt.Println("Radio2 PlugBatary")
}

func (r *Radio2) SetChannal(Frequensy string) {
	r.Frequensy = Frequensy
	fmt.Println("Radio1 Programm Channal")

}

func (r *Radio2) DoTransmit(message string) {
	fmt.Printf("Radio2 передаю на частоте %s сообщение %s \r\n", r.Frequensy, message)

}

func main() {
	R := &Radio1{}
	TransmitIt(R, R)
	R2 := &Radio2{}
	TransmitIt(R2, R2)
	R3 := &Radio3{}
	TransmitIt(R3, R3)

}
