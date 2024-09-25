package main

import (
	//"fmt"
	//"net/http"
	// "strings"
	// "log"
	"GoStudy/leetcode"
	"fmt"
	//	"fmt"
)

// User struct
/*type User struct {
    Username string
}
func sayhelloName(w http.ResponseWriter,r *http.Request)  {
    r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}*/

func main() {
	/*fmt.Println("6666666")
	   http.HandleFunc("/", sayhelloName) //设置访问的路由
		err := http.ListenAndServe(":9090", nil) //设置监听的端口
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}*/
	/*nums := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	leetcode.MoveZeroes(nums)*/
	//nums := []int{18,4,13,-1,2,7,19,14,11,0,-9,9,4,2,-2,-7,18,18,-7,-5,9,6,-8,3,13,0,15,9,10,-1,17,13,13,-8,3,8,4,19,10,-8,18,15,11,13,11,1,14,2,10,1,11,5,18,5,-8,13,-10,5,-8,-9,-5,9,10,-10,-3,-3,-4,-4,-8,-10}
	//leetcode.MaxRotateFunction(nums)
	//leetcode.MaxRotateFunction(nums)
	//nums1 := []int{490693, 900498, 448195, 24359, 126032, 584252, 26132, 124479, 586672, 855404, 24359, 418495, 243450, 106232, 690685, 410981, 871863, 419180, 242524, 23549, 284759, 26132, 271146, 966337, 781863, 418495, 242524, 126032, 411980, 621032, 271641, 25349, 900894, 411980, 997268, 671059, 649498, 781836, 312273, 15727, 671095}
	//nums := []int{11,13,11,14,11}
	//nums :=[]int{5,12,8,5,5,1,20,3,10,10,5,5,5,5,1}
	/*var numst = []*leetcode.Employee{
		&leetcode.Employee{
			Id:           1,
			Importance:   5,
			Subordinates: []int{2, 3},
		},
		&leetcode.Employee{
			Id:           2,
			Importance:   3,
			Subordinates: []int{},
		},
		&leetcode.Employee{
			Id:           3,
			Importance:   3,
			Subordinates: []int{},
		},
	}
	
	co := leetcode.GetImportance(numst,1)
	fmt.Println(co)*/

	nums := []int{13,23,12}

	n := leetcode.SumDigitDifferencesPro(nums)
	fmt.Print(n)
}
/*	numst := []int{1,1203,20304}
	co := leetcode.ApproximatePairs(numst)
	fmt.Println(co)
}*/
