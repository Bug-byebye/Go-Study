package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

 func mergeNodes(head *ListNode) *ListNode {

	cur, fhead:= head.Next, head
	for cur.Next != nil{

		if cur.Val !=0 {
			fhead.Val+=cur.Val
		}else{
			fhead = fhead.Next
			fhead.Val = 0
		}
		cur=cur.Next
	}
	fhead.Next=nil

	return head
 }