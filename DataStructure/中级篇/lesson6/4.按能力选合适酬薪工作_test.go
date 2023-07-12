package lesson6

import (
	"fmt"
	"testing"
)

func TestAcAbilityGetWork(t *testing.T) {
	jobArr := make([]Job, 0)
	jobArr = append(jobArr, Job{Ability: 1, Salary: 2})
	jobArr = append(jobArr, Job{Ability: 2, Salary: 4})
	jobArr = append(jobArr, Job{Ability: 3, Salary: 1})
	jobArr = append(jobArr, Job{Ability: 4, Salary: 5})
	jobArr = append(jobArr, Job{Ability: 4, Salary: 2})
	jobArr = append(jobArr, Job{Ability: 4, Salary: 3})

	capArr := []int{1, 2, 3, 4, 5}

	fmt.Println(AcAbilityGetWork(jobArr, capArr))
}
