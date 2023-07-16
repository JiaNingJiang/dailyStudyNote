package lesson9

import (
	"fmt"
	"testing"
)

func TestPerfectActivityRoute(t *testing.T) {

	actA := NewActivity(4, 20, make([]*Activity, 0), make([]*Activity, 0))

	actB := NewActivity(3, 20, make([]*Activity, 0), make([]*Activity, 0))
	actC := NewActivity(1, 200, make([]*Activity, 0), make([]*Activity, 0))
	actD := NewActivity(4, 500, make([]*Activity, 0), make([]*Activity, 0))

	actE := NewActivity(5, 1000, make([]*Activity, 0), make([]*Activity, 0))

	// 追加依赖关系
	actA.AddNext(actB)
	actA.AddNext(actC)
	actA.AddNext(actD)

	actB.AddLast(actA)
	actB.AddNext(actE)

	actC.AddLast(actA)
	actC.AddNext(actE)

	actD.AddLast(actA)
	actD.AddNext(actE)

	actE.AddLast(actB)
	actE.AddLast(actC)
	actE.AddLast(actD)

	// 从后往前整理活动路线
	actE.OrderActRoute()
	actB.OrderActRoute()
	actC.OrderActRoute()
	actD.OrderActRoute()
	actA.OrderActRoute()

	actSet := []*Activity{actA, actB, actC, actD, actE}
	perfectRoutes := SumPerfectRoutes(actSet)

	for _, route := range perfectRoutes {
		fmt.Println(*route)
	}

}
