package queries

import(
	"fmt"
	"bufio"
	"strings"
	"os"
)

func MovieDate() string{
	var input string 
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you want to watch any specific time period? (Type the year you want the oldest movie to be)?")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}