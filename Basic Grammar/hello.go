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
		//存储长度、容量、指向数组的指针
		//创建切片make
		s := make([]string,3) 
		//用copy在两个slice之间拷贝数据
		
		fmt.Println("get:",s[2])
		//追加元素append
		s = append(s,"d")
		fmt.Println(s)	
		c := make([]string,len(s))
		copy(c,s)
		fmt.Println(c)	
		//取出2-5个元素，但不包括第五个元素
		fmt.Println(s[2:5])
	//map:完全无序
		m := make(map[string]int) //string为key的类型，int为value的类型
		m["one"]=1
		m["two"]=2
		fmt.Println(m)
		fmt.Println(m["one"])

		r,ok:= m["unkown"] //用于获取map中是否有此key的存在
		fmt.Println(r,ok)
		delete(m,"one")
	//range:快速遍历slice和map,会返回索引和对应位置的值
		//遍历slice
		nums:=[]int{2,3,4}
		sum:=0
		for i,num:=range nums{
			sum += num
			if num == 2{
				fmt.Println("index",i,"num:",num)
			}
		}
		fmt.Println(sum)
		//遍历map

	//函数
		//变量类型后置
		func add(a int, b int) int{
			return a+b
		}
		res:=add(1,2)
		fmt.Println(res)


}