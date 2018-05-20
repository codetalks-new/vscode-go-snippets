import "fmt"

func dumpListNode(lnode *ListNode) {
	fmt.Print("ListNode[")
	for lnode != nil {
		fmt.Print(lnode.Val, ",")
		lnode = lnode.Next
	}
	fmt.Print("]\n")
}
