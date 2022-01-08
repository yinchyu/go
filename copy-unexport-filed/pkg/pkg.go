package pkg




type ommit   struct{

name string

age int

oldder  bool

}

func New() *ommit{
o:=ommit{
	"12",1,true,
}
	return  &o
}


func (o *ommit)Getname() string{
	return o.name
}

func (o *ommit)Getage() int{
	return o.age
}
func (o *ommit)Isolder() bool{
	return o.oldder
}