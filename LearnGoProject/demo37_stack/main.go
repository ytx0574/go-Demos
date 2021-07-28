package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
栈的应用场景  FILO
1) 子程序的调用:在跳往子程序前，会先将下个指令的地址存到堆栈中，直到子程序执行完后再 将地址取出，以回到原来的程序中。
2) 处理递归调用:和子程序的调用类似，只是除了储存下一个指令的地址外，也将参数、区域变 量等数据存入堆栈中。
3) 表达式的转换与求值。
4) 二叉树的遍历。
5) 图形的深度优先(depth 一 first)搜索法。
*/

type Stack struct {
	top int
	info []string
}

func (this *Stack)isEmpty() bool {
	return this.top == -1
}

func (this *Stack)Push(val string) {
	this.top++
	this.info[this.top] = val
}

func (this *Stack)Pop() (val string, err error) {
	if this.isEmpty() {
		return "", errors.New("stack is empty")
	}

	val = this.info[this.top]
	this.top--
	return
}

func isNumber(val string) bool {
	_, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return false
	}else {
		return true
	}
}

func operatorLevel(val string) int {
	switch val {
	case "*", "/", "%":
		return 1
	default:
		return 0
	}
}

func cal(num1, num2, operator string) string {

	//x := float64([]byte(num1)[0])
	//y := float64([]byte(num2)[0])
	x, _ := strconv.ParseFloat(num1, 64)
	y, _ := strconv.ParseFloat(num2, 64)
	val := 0.0

	switch operator {
		case "*":
			val = x * y
	case "/":
		val = x / y
	case "%":
		val = float64(int(x) % int(y))
	case "+":
		val = x + y
	case "-":
		val = x - y
	default:

	}



	str_val := strconv.FormatFloat(val, 'f', -1,64)
	//var ss number.Formatter = number.Decimal(val, number.NoSeparator())
	//精度丢失 四舍五入了
	//str_val = fmt.Sprintf("%f %v", val, ss)

	return str_val
}

/*
解析数学表达式
*/
func ParseMathFormula(formula string) (result string) {
	numStack := &Stack{
		top: -1,
		info: make([]string, 10),
	}
	symbolStack := &Stack{
		top: -1,
		info: make([]string, 10),
	}


	index := 0
	num1 := ""
	num2 := ""
	operator := ""
	word := ""
	for {
		//提取每次遍历的char
		ch := string(formula[index])
		fmt.Printf("%T, %v\n", ch, ch)

		if isNumber(ch) || strings.Contains(ch, ".") {
			word += ch
			ch_next := ""
			//提取下一个char
			if index + 1 < len(formula) {
				ch_next = string(formula[index + 1])
			}
			split_word := strings.Split(word, ".")
			if len(split_word) > 2 {
				fmt.Printf("表达式中的数值有问题\n")
				break
			}

			if !isNumber(ch_next) && !strings.Contains(ch_next, ".") {
				numStack.Push(word)
			}else {
				index++
				continue
			}
		}else {
			word = ""
			if !symbolStack.isEmpty() {
				//拿到运算符时, 与上一个运算符比较优先级, 低于则, 把之前的提前进行计算
				if operatorLevel(ch) < operatorLevel(symbolStack.info[symbolStack.top]) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					operator, _ = symbolStack.Pop()

					//此处注意, 参与运算的数值的前后顺序. 栈顶的作为第二参数
					result = cal(num2, num1, operator)
					numStack.Push(result)
					symbolStack.Push(ch)
				}else {
					symbolStack.Push(ch)
				}
			}else {
				symbolStack.Push(ch)
			}
		}

		if index == len(formula) - 1 {
			break
		}
		index++
	}

	for !symbolStack.isEmpty() {
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		operator, _ = symbolStack.Pop()

		result = cal(num2, num1, operator)
		numStack.Push(result)
	}

	result, _ = numStack.Pop()
	return
}

/*
迷宫回溯问题
0. 表示可走
1. 表示墙
2. 表示成功
3. 表示不通

//数组 前行后列
*/
func MazeBack(mazeMap [][]int, startx, starty int) bool {
	//终点要选对.  如果终点就是个墙, 肯定走不通
	if mazeMap[6][5] == 2 {
		return true
	}else {
		val := mazeMap[startx][starty]
		//可探测
		if val == 0 {
			//假设探对
			mazeMap[startx][starty] = 2
			if MazeBack(mazeMap, startx + 1, starty) { //下
				return true
			} else if MazeBack(mazeMap, startx, starty + 1) { //右
				return true
			} else if MazeBack(mazeMap, startx - 1, starty) { //上
				return true
			} else if MazeBack(mazeMap, startx, starty - 1) { //左
				return true
			} else {
				//都探不了, 标记不通
				mazeMap[startx][starty] = 3
				return false
			}
		}
	}
	return false
}

func main()  {
	exp := "100+30*80.33333333-40"
	result := ParseMathFormula(exp)
	fmt.Printf("%v = %v\n", exp, result)

	row := 8
	colume := 7

	mazeMap := make([][]int, row)
	for i, _ := range mazeMap {
		mazeMap[i] = make([]int, colume)

		//最前最后列为墙
		mazeMap[i][0] = 1
		mazeMap[i][6] = 1
	}

	for i := 0; i < colume; i++ {
		//最前最后行为墙
		mazeMap[0][i] = 1
		mazeMap[7][i] = 1
	}

	mazeMap[2][1] = 1
	mazeMap[2][2] = 1


	for i, v := range mazeMap {
		for j, _ := range v {
			fmt.Printf("%v ", mazeMap[i][j])
		}
		fmt.Println()
	}

	fmt.Println("--------------")

	MazeBack(mazeMap, 1, 1)

	for i, v := range mazeMap {
		for j, _ := range v {
			fmt.Printf("%v ", mazeMap[i][j])
		}
		fmt.Println()
	}

}

