package main


/*heap是常用的实现优先队列的方法。heap包对任意实现了heap接口的类型提供堆操作。堆结构继承自sort.Interface, 而sort.Interface，
需要实现三个方法：Len() int / Less(i, j int) bool / Swap(i, j int)
再加上堆接口定义的两个方法：Push(x interface{}) / Pop() interface{}。故只要实现了这五个方法，便定义了一个堆。*/

type IntHeap []int

func

