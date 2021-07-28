package sparse_matrix

import "fmt"

//稀疏矩阵
//棋盘数据记录
//仅记录有效数据, 其他数值都丢弃掉, 进而达到数据精简的效果

type ChessRecord struct {
	row int
	col int
	val string
}

func SparseMatrix() {

	var ChessArray [11][11]string
	ChessArray[1][4] = "x"
	ChessArray[0][3] = "x"
	ChessArray[5][7] = "y"

	var chessSlice [][]string
	for _, v := range ChessArray {
		//todo 新申请一个slice来装数据
		//var slice []string = make([]string, len(v))
		//copy(slice, v[:])
		//fmt.Printf("%p %p\n", slice, &v)
		//chessSlice = append(chessSlice, slice)

		//todo 重新声明变量, 并追加
		val := v
		chessSlice = append(chessSlice, val[:])
	}


	//稀疏数组 记录数据
	RecordSlice := make([]ChessRecord, 1)
	RecordSlice[0] = ChessRecord{
		row: len(chessSlice),
		col: len(chessSlice[0]),
		val: "0",
	}

	enumerator(chessSlice, func(row int, col int, val string, rowEnumertorEnd bool) {
		value := val
		if value == "" {
			value = "0"
		}else {
			RecordSlice = append(RecordSlice, ChessRecord{
				row: row,
				col: col,
				val: value,
			})
		}
		fmt.Printf("%v\t", value)
		if rowEnumertorEnd {
			fmt.Println()
		}
	})

	//输出记录得数据
	fmt.Printf("RecordSlice = %v\n", RecordSlice)

	//恢复
	var chessSlice2 [][]string = make([][]string, RecordSlice[0].row)
	for i, _ := range chessSlice2 {
		chessSlice2[i] = make([]string, RecordSlice[0].col)
	}

	for i, v := range RecordSlice {
		if i == 0 {
			continue
		}else {
			chessSlice2[v.row][v.col] = v.val
		}
	}

	enumerator(chessSlice2, func(row int, col int, val string, rowEnumertorEnd bool) {
		value := val
		if value == "" {
			value = "0"
		}
		fmt.Printf("%v\t", value)
		if rowEnumertorEnd {
			fmt.Println()
		}
	})
}

//枚举二维切片
func enumerator(chessArray [][]string, f func(row int, col int, val string, rowEnumertorEnd bool)) {
	for i, v := range chessArray {
		for j, v2 := range v {
			rowEnumertorEnd := false
			if j == len(v) - 1 {
				rowEnumertorEnd = true
			}
			f(i, j, v2, rowEnumertorEnd)
		}
	}
}