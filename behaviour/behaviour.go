package behaviour

import (
	"fmt"
	"time"
)
	
func Maingoroutinebehaviour(){
	go func(){
	fmt.Println("Starting the program...")
	}()
	fmt.Println("End of the program")
	time.Sleep(1 * time.Second)
}