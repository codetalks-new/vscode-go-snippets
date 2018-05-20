import "fmt"

func dumpTree(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}
	branches := []*TreeNode{root}
	fmt.Print("[")
	for len(branches) > 0 {
		branch := branches[0]
		branches = branches[1:]
		if branch.Left != nil {
			branches = append(branches, branch.Left)
		}
		if branch.Right != nil {
			branches = append(branches, branch.Right)
		}
		if len(branches) > 0 {
			fmt.Print(branch.Val, ",")
		} else {
			fmt.Print(branch.Val, "")
		}
	}
	fmt.Println("]")
}