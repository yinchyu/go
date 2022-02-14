func addslice(list []int) {
	fmt.Println(list)
	fmt.Printf("%p\n", list)
	fmt.Println(len(list), cap(list))
	list = append(list, 1, 2, 3, 4)
	fmt.Println(list)
	fmt.Printf("%p\n", list)
}
func main() {
	//fmt.Println(1 | 2)
	s2 := make([]int, 0, 10)
	fmt.Printf("%p\n", s2)
	fmt.Println("=====", s2[:4])
	addslice(s2)
	fmt.Println("=====", s2, len(s2))
	s := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
	s.Len = 4
	// 就是中间经过两次转换， 一次是通过转换为unsafe 一次是断言为具体的类型
	s3 := *(*[]int)(unsafe.Pointer(s))
	fmt.Println("=====", s3, len(s3), cap(s3))
}