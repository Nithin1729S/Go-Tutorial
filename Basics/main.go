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
	
}