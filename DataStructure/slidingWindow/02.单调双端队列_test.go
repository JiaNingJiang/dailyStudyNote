package slidingWindow

import (
	"fmt"
	"testing"
)

func TestMonotonicDeque(t *testing.T) {

	mdq := NewMonotonicDeque(true)
	mdq.PushNew(3)
	mdq.PushNew(2)
	mdq.PushNew(4)
	mdq.PushNew(6)
	fmt.Println("第一轮: ", mdq.BackPeak()) // 3 2 4 6

	mdq.PushNew(3)
	mdq.RemoveOldRecord()
	fmt.Println("第二轮: ", mdq.BackPeak()) // 2 4 6 3

	mdq.PushNew(5)
	mdq.RemoveOldRecord()
	fmt.Println("第三轮: ", mdq.BackPeak()) // 4 6 3 5

	mdq.PushNew(5)
	mdq.RemoveOldRecord()
	fmt.Println("第四轮: ", mdq.BackPeak()) // 6 3 5 5

	mdq.PushNew(3)
	mdq.RemoveOldRecord()
	fmt.Println("第五轮: ", mdq.BackPeak()) // 3 5 5 3

	mdq.PushNew(2)
	mdq.RemoveOldRecord()
	fmt.Println("第六轮: ", mdq.BackPeak()) // 5 5 3 2

	mdq.RemoveOldRecord()
	fmt.Println("第七轮: ", mdq.BackPeak()) // 5 3 2

	mdq.RemoveOldRecord()
	fmt.Println("第八轮: ", mdq.BackPeak()) // 3 2

	mdq.RemoveOldRecord()
	fmt.Println("第九轮: ", mdq.BackPeak()) // 2

	mdq.RemoveOldRecord()
	fmt.Println("第十轮: ", mdq.BackPeak()) //

}
