package BackTrack

var res [][]string

func solveNQueens(n int) [][]string {
	// '.' 表示空，'Q' 表示皇后，初始化空棋盘。
	var board [][]string
	board = make([][]string,n)
	for i  := range board {
		board[i] = make([]string, n)
	}
	backtrack(board, 0)
	return res
}

// 路径：board 中小于 row 的那些行都已经成功放置了皇后
// 选择列表：第 row 行的所有列都是放置皇后的选择
// 结束条件：row 超过 board 的最后一行
func backtrack(board [][]string,row int) [][]string{
	// 触发结束条件
	if (row == len(board)) {
		res = make([][]string,row)
		for i  := range res {
			res[i] = make([]string, row)
			copy(res[i],board[i])
		}
		return res
	}

	n := len(board[row])
	for col := 0; col < n; col++ {
		// 排除不合法选择
		if !isValid(board, row, col) {
			continue
		}
		// 做选择
		board[row][col] = "Q"
		// 进入下一行决策
		backtrack(board, row + 1)
		// 撤销选择
		board[row][col] = "."
	}
	return nil
}

/* 是否可以在 board[row][col] 放置皇后？ */
func isValid(board [][]string, row, col int) bool {
	n := len(board)
	// 检查列是否有皇后互相冲突
	for i := 0; i < n; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	for i,j:= row - 1, col + 1; i >= 0 && j < n; i-- {
		if board[i][j] == "Q" {
			return false
		}
		j++
	}
	// 检查左上方是否有皇后互相冲突
	for i,j := row - 1,col - 1; i >= 0 && j >= 0; i-- {
		if board[i][j] == "Q" {
			return false
		}
		j--
	}
	return true
}