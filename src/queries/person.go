package queries

import(
	"fmt"
	"bufio"
	"strings"
	"os"
)



func PersonSpec() string{
	var input string 
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you want to watch a  specific Actor/Director/Producer(If any type any)?")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}
