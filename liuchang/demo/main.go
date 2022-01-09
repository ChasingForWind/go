package main

import "fmt"

//定义结构体
type Student struct {
	ID    int
	Name  string
	Score int
}

//主函数
func main() {
	s1 := Student{12346, "张三", 90}
	s2 := Student{12347, "李四", 85}
	s := s1.compareScore(s2)
	fmt.Println(s)
	s = compareSore(s1, s2)
	fmt.Println(s)

}

//方法
func (s1 Student) compareScore(s2 Student) Student {
	if s1.Score > s2.Score {
		return s1
	} else {
		return s2
	}
}

//函数
func compareSore(s1 Student, s2 Student) Student {
	if s1.Score > s2.Score {
		return s1
	} else {
		return s2
	}
}
