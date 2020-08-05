package main

import "fmt"

// Course ...
type Course struct {
	ActiveDays []int
	// NextOpenDays map[int]int
}

// WeekDayNode 是一个链表的节点
type WeekDayNode struct {
	// DayVal 是一周的第几天，默认周一是第一天，周日是第七天
	DayVal int
	// IsOpen 是否是上课日
	IsOpen bool
	// NextOpenDays 下一节课是几天后
	NextOpen int
	// Name 就是星期几的名字
	Name string
	Next *WeekDayNode
}

// NewWeekDays ...
func NewWeekDays() []*WeekDayNode {
	list := make([]*WeekDayNode, 0)
	days := [7]string{
		"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	for i := 0; i <= 6; i++ {
		nNode := &WeekDayNode{
			DayVal:   i + int(1),
			Name:     days[i],
			NextOpen: 1,
		}
		list = append(list, nNode)
		if len(list) > 1 {
			list[i-1].Next = nNode
		}
	}
	list[6].Next = list[0]
	return list
}

// ParseMap ...
func (c *Course) ParseMap() {
	// if len(c.ActiveDays) == 0 {
	// 	panic("panic ...")
	// }
	// m := make(map[int]int)
	// for i := 0; i <= 7; i++ {
	// 	// 最后一个上课日
	// 	if i == len(c.ActiveDays)-1 {
	// 		m[i] = 7 - i + c.ActiveDays[0]
	// 	} else {
	// 		// m[i] =
	// 	}
	// }
	// c.NextOpenDays = m
}

// NextCoursingDays 后续 N 节课的安排
func NextCoursingDays(n int) {

}

// const

func main() {
	w := NewWeekDays()
	fmt.Println(w)

	c := Course{ActiveDays: []int{7}}
	for _, v := range c.ActiveDays {
		// fmt.Println(v)
		w[v-1].IsOpen = true
	}
	for x := range w {
		// fmt.Println(w[x])
		currentDay := w[x]
		for !currentDay.Next.IsOpen {
			w[x].NextOpen++
			currentDay = currentDay.Next
		}
	}
	nextDict := map[string]int{}
	for j := range w {
		nextDict[w[j].Name] = w[j].NextOpen
		fmt.Println(w[j])
	}
	fmt.Println(nextDict)
	scheduleWeekly := [7]int{}
	for ck := range c.ActiveDays {
		scheduleWeekly[c.ActiveDays[ck]-1] = 1
	}
	fmt.Println(scheduleWeekly)
	offsets := [7]int{1, 1, 1, 1, 1, 1, 1}
	for sw := range scheduleWeekly {
		curr := sw
		for scheduleWeekly[NextKey(curr)] != 1 {
			offsets[sw]++
			curr = NextKey(curr)
		}
	}
	fmt.Println(offsets)
}

// getOffsetDict ...
// func getOffsetDict() [7]int {
// }

// NextKey returns the next index
func NextKey(i int) int {
	if i == 6 {
		return 0
	}
	return i + 1
}
