package main
import(
	"fmt"
)

type Animal interface{
	GetInfor() string
}

type Cat struct{
	Name string
	Age int
}

func (c Cat) GetInfor() string{
	return fmt.Sprintf("Name: %s, Age: %d",c.Name,c.Age)
}

func announce(animal Animal){
	fmt.Println(animal.GetInfor())
}


type Describer interface{
	Describe() string
}

type Person struct{
	Name string
}

func (p Person) Describe() string{
	return fmt.Sprintf("I'm %s",p.Name)
}

type Car struct{
	Model string
}

func(c Car) Describe() string{
	return fmt.Sprintf("Model: %s",c.Model)
}

func findDescriber(d Describer){
	fmt.Println(d.Describe())
}


func main(){
	ppls:=[]string{"John","Doe","Nithin","Alaska"}
	for _,ppl:=range ppls{
		fmt.Println(ppl)
	}

	mp:=map[string]string{"name":"John","age":"25","city":"New York"}
	for key,value:=range mp{
		fmt.Println(key,value)
	}

	c:=Cat{"Tom",5}
	announce(c)

	p:=Person{"John"}
	findDescriber(p)

	car:=Car{"BMW"}
	findDescriber(car)
	
}