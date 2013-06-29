package main

import (
	//"encoding/json"
	//"fmt"
	"strings"
)

type Element interface{}
type List []Element
type Queue struct {
	Size   int
	Values List
}

func (l *List) String() string {
	var s []string
	for _, v := range *l {
		s = append(s, v.(string))
	}
	return "[" + strings.Join(s, ",") + "]"
}

func (q *Queue) Init(n int) {
	q.Size = n
}

func (q *Queue) Len() int {
	return len(q.Values)
}

func (q *Queue) Pop() (v interface{}) {
	v, q.Values = q.Values[0], q.Values[1:]
	return
}

func (q *Queue) Push(v interface{}) {
	if q.Len() == q.Size {
		q.Pop()
	}
	q.Values = append(q.Values, v)
}

/*
func main() {
	var q Queue
	q.Init(3)

	q.Push("string")
	q.Push(1)
	q.Push("12313")

	var m interface{}
	m = map[string]interface{}{
		"name": "solu",
		"date": "2013-06-27",
		"msg":  "hello world",
	}
	q.Push(m)
	j, _ := json.Marshal(q.Values)
	fmt.Println(string(j))
	//[1,"12313",{"date":"2013-06-27","msg":"hello world","name":"solu"}]
}
*/
