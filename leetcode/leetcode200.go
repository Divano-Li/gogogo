package main


import (
	"fmt"
	"strconv"
	"strings"
)

/*type Queue struct {
	arr []string
	head int
	tail int
}

func ConstructQueue(k int) Queue {
	a:= Queue{make([]string, k), -1, -1}
	return a
}

func (this *Queue) isEmpty() bool {
	if this.head == -1 && this.tail == -1 {
		return true
	}
	return false
}

func (this *Queue) isFull() bool {
	if this.tail == len(this.arr) - 1 {
		return true
	}
	return false
}

func (this *Queue) enQueue(value string) bool {
	if this.isFull() {
		return false
	}
	if this.isEmpty() {
		this.tail = 0
		this.head = 0
		this.arr[0] = value
		return true
	} else {
		this.tail += 1
		this.arr[this.tail] = value
		return true
	}
}

func (this *Queue) deQueue() bool {
	if this.isEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
	} else {
		this.head = this.head + 1
	}
	return true
}


func (this *Queue) front() string {
	if this.isEmpty() {
		return ""
	}
	return this.arr[this.head]
}


func (this *Queue) rear() string {
	if this.isEmpty() {
		return ""
	}
	return this.arr[this.tail]
}*/

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	a := map[string]byte{}
	ii := len(grid)
	jj := len(grid[0])
	count := 0
	key := ""
	str := ""
	queue := ConstructQueue(100000)
	var temp byte
	for i:=0;i< len(grid); i++ {
		for j:=0;j<len(grid[i]);j++ {
			key = strconv.Itoa(i) + ":" + strconv.Itoa(j)
			temp = grid[i][j]
			_, ok := a[key];  if !ok && temp == 49 {
				count++
				a[key] = temp
				queue.enQueue(key)
				for queue.isEmpty() == false {
					str = queue.front()
					queue.deQueue()
					i1 := strings.Split(str, ":")[0]
					i,_ := strconv.Atoi(i1)
					j1 := strings.Split(str, ":")[1]
					j, _ := strconv.Atoi(j1)
					if i-1 >=0 {
						key = strconv.Itoa(i-1) + ":" +  strconv.Itoa(j)
						_, ok := a[key]; if !ok && grid[i-1][j] == 49 {
							queue.enQueue(key)
							a[key] = grid[i-1][j]
						}
					}
					if i+1 <= ii -1 {
						key = strconv.Itoa(i+1) + ":" + strconv.Itoa(j)
						_, ok := a[key]; if !ok && grid[i+1][j] ==49 {
							queue.enQueue(key)
							a[key] = grid[i+1][j]
						}
					}
					if j-1 >= 0 {
						key = strconv.Itoa(i) + ":" + strconv.Itoa(j-1)
						_, ok := a[key]; if !ok && grid[i][j-1] ==49 {
							queue.enQueue(key)
							a[key] = grid[i][j-1]
						}
					}
					if j+1 <= jj-1 {
						key = strconv.Itoa(i) + ":" + strconv.Itoa(j+1)
						_, ok := a[key]; if !ok && grid[i][j+1] ==49 {
							queue.enQueue(key)
							a[key] = grid[i][j+1]
						}
					}
				}

			}
		}
	}
	return  count
}

func numIslands1(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	row := len(grid)
	col := len(grid[0])

	count := 0
	for i:=0; i<row; i++ {
		for j:=0; j<col; j++ {
			if grid[i][j] == 49{
				count++
				dfs1(i,j, row, col, grid)
			}
		}
	}
	return count
}

func dfs1(x int, y int, row int, col int, grid [][]byte)  {
	if x < 0 || x >= row || y < 0 || y >= col {

	} else if grid[x][y] == 49  {
		grid[x][y] = 48
		dfs1(x-1, y, row, col, grid)
		dfs1(x+1, y, row, col, grid)
		dfs1(x, y-1, row, col, grid)
		dfs1(x, y+1, row, col, grid)
	}
}

func main() {
	//grid :=[][]byte{{"1","1","1","1","0"},{"1","1","0","1","0"},{"1","1","0","0","0"},{"0","0","0","0","0"}}
	//grid :=[][]byte{{1,1,1,1,0},{1,1,0,1,0},{1,1,0,0,0},{0,0,0,0,0}}
	grid1 :=[][]byte{{49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,0,0,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,0,0,0,0,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49,49}}
	//grid1 := [][]byte{{"1","1","1","1","0"},["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]}
	fmt.Println(grid1)
	fmt.Println(numIslands(grid1))


}