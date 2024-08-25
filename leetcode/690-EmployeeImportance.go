package leetcode

import "fmt"

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getsubImportance(employees []*Employee, subordinatesArray []int, impSum int)int{
	i := -1
	for _, val := range subordinatesArray {
		for index, _ := range employees {
			if employees[index].Id == val {
				impSum += employees[index].Importance
				fmt.Println(impSum,index)
				i = index
				break
			}
		}
		fmt.Println(impSum,i)
	
		if len(employees[i].Subordinates)>0 {
			impSum += getsubImportance(employees, employees[i].Subordinates, impSum)
		}
		
		fmt.Println(employees[i].Subordinates,impSum)
	}
	return impSum
}
func GetImportance(employees []*Employee, id int) int {
	impSum := 0
	i := -1
	for index, _ := range employees {
		if employees[index].Id == id {
			impSum += employees[index].Importance
			i = index
			break
		}
	}
	fmt.Println(employees[i].Subordinates,impSum)
	impSum = getsubImportance(employees, employees[i].Subordinates, impSum)
	return impSum
}