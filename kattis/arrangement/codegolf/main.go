package main
import ("fmt"
"strings")
func main(){var r,t int
fmt.Scanf("%d\n%d",&r,&t)
for r>0{
a:=t/r
r-=1
t-=a
fmt.Println(strings.Repeat("*",a))}}
