package main

/*
Assume we have a sentence like this:
I don't Think That You Have any Insight whatsoever into your capacity for good until you have some well-developed insight into your capacity for evil

And we want to count the letters which the sentence have like this:
a: 9
c: 4
d: 4
e: 10
f: 2
g: 3
h: 7
i: 12
...
*/
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){

	r:=bufio.NewReader(os.Stdin)
	fmt.Println("Give me a sentence ...")
	str,err:=r.ReadString('\n')
	if err!=nil{
		panic(err)
	}
	
	str=strings.TrimSpace(str)
	sl:=strings.Split(str, " ")
	
	output:=createOutput(sl)

	fmt.Println(output)

	for i:= range output{
		fmt.Println(string(i),":",output[i])
	}

	// fmt.Println(strings.ToLower("S"))
}

func createOutput(sl []string) map[string]int{
	output:=make(map[string]int)
	for _, word:=range sl{
		for i:=0;i<len(word);i++{
			output[strings.ToLower(string(word[i]))]+=1
		}
	}
	return output
}