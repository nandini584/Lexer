package main 
import ("fmt"
		"os"
		"os/user"
		
	)
func main(){
	user,err:=user.Current()
	if err!=nil{
		panic(err)
	}
	fmt.Println("Hello "+user.Name+"!")
	fmt.Println("This is the Shakespeare programming language!")
	fmt.Println("Feel free to type in commands")
	repl.Start(os.Stdin, os.Stdout)

}