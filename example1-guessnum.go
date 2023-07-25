package main
import(
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
	"strings"
	"strconv"
)
func main(){
	maxNum:= 100
	//需要设置随机数种子否则每次生成的是相同的随机数
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ",secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	input,err :=reader.ReadString('\n')
	if err!=nil{
		fmt.Println("An err ocurred while reading input. Please try again",err)
		return
	}
	input = strings.TrimSuffix(input,"\n")

	guess,err:=strconv.Atoi(input)
	if err!=nil{
		fmt.Println("Invalid input. Please enter an integer value")
		return
	}
	fmt.Println("Your guess is",guess)
}