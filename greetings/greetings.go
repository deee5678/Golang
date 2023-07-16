package greetings

import "fmt"

func Hello(name string) string{
// return a greeting that embeds the name in a message

message := fmt.Sprintf("Hi,%v. Welcome!",name)
//the above line is same as below 2 lines
//var message string
//message = fmt.Sprintf("Hi, %v. Welcome!", name)

return message 
}
