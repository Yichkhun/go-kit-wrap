package work_pool

import (
	"fmt"
	"testing"
	"time"
)

type WorkDerive struct {
}

func NewWorkDerive() WorkBase {
	return &WorkDerive{}
}

func(WorkDerive) Task(params interface{}) {
	threadId, _ := params.(int)
	fmt.Printf("id = %d task is working \n", threadId)
	time.Sleep(2 * time.Second)
}


func TestWorkPool(t *testing.T) {
	pool := NewWorkPool(10)

	for i:=0; i <10; i++ {
		task := NewWorkDerive()
		pool.Run(&task)
		fmt.Println("set pool index=", i)
	}

	pool.Close()
	time.Sleep(20 * time.Second)
}