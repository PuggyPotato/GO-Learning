package main


import(
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main(){
	var userInput string
	reader := bufio.NewReader(os.Stdin)
	userInput,_ = reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	words := strings.Fields(userInput)

	wordCount := make(map[string]int)

	for _,word:= range words{
		wordCount[word]++
	}

	fmt.Println("Word Frequency:")
	for word,count := range wordCount{
		fmt.Println(word + ":",count)
	}
	

	
}