//文件属于main包的一部分，main包是程序的入口包
package main

import(
	"fmt" //导入标准库的format包，用于输入输出字符串，格式化字符串
	"math"
	"time"
)

func main(){
	//变量
		fmt.Println("hello world")
		var a = "initial" //自动推导变量类型
		var e float64
		f := float32(e)
		g := a+"foo"
		const s1 string = "constant" //常量，go中常量无确定类型,会根据使用的上下文自动确定类型
		const h =500000000
		fmt.Println(a,g,f,s1,math.Sin(h))
	//if else
		if 7%2 == 0{ //if后面必须直接接大括号
			fmt.Println("7 is seven")
		}else{
			fmt.Println("7 is odd")
		}
	//循环——只有for循环
		for{ //死循环
			fmt.Println("loop")
			break;
		}
		for j:=7; j<9;j++ {
			fmt.Println(j)
		}
	//switch
		//c++中case中不加break会默认跑完所有分支，go不会
		//switch后可以使用任意的变量类型，甚至可以不写
		w:= 2
		switch w{
			case 1:
				fmt.Println("one")
			default:
				fmt.Println("other")
		}
		//可以取代if else语句
		t := time.Now()
		switch{
			case t.Hour()< 12 :
				fmt.Println("It's before noon")
			default :
				fmt.Println("It's after noon")
		}
	//数组
		var array[5] int
		array[4] = 100
		fmt.Println(array[4],len(a))

		b :=[4]int{1,2,3,4}
		fmt.Println(b)
		var twoD [2][3]int
		for i:=0;i<2;i++{
			for j:=0;j<3;j++{
				twoD[i][j]=i+j
			}
		}
	//切片:可变长度
		//创建切片make
		s := make([]string,3) 
		fmt.Println("get:",s[2])
		//追加元素append
		s = append(s,"d")
		fmt.Println(s)	

}