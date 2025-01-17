package main

type Bigstruct struct{
	buff [1<<16]byte
}

var obj Bigstruct = Bigstruct{}

func fooPBV(obj Bigstruct) {
}

func fooPBP(obj *Bigstruct) {
}

func main() {
	fooPBV(obj)
	fooPBP(&obj)
}