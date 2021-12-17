package main
import(
	"fmt"
)


type Stringer interface {
String() string
}

type Plusser interface {
Plus(string) string
}
type MyStringer string

type MyPlusser string

func (m MyStringer) String() string {
return string(m)+"stringer"
}
func (m MyPlusser) Plus(s string) string {
return string(m)+s+"plusser"
}

func concatTo[S Stringer, P Plusser](s []S, p []P) []string {
r := make([]string, len(s))
for i, v := range s {
fmt.Printf("ConcatTo1 s type %T,p type%T\n,si type %T,pp type%T\n", s, p,s[i], p[i])
r[i] = p[i].Plus(v.String())
}
return r
}
func ConcatTo2(s []Stringer, p []Plusser) []string {
r := make([]string, len(s))
for i, v := range s {
fmt.Printf("ConcatTo2 s type %T,p type%T\n,si type %T,pp type%T\n", s, p,s[i], p[i])
r[i] = p[i].Plus(v.String())
}
return r
}

func main() {

//s1 := MyStringer("hello world  name is  good")
//p1 := MyPlusser("name is good")
s2 := make([]Stringer, 1)
p2 := make([]Plusser, 1)

for i:= 0; i<1; i++ {
s2[i] = MyStringer("hello world  name is  good")
p2[i] = MyPlusser("name is good")
}
concatTo(s2, p2)
ConcatTo2(s2, p2)

}