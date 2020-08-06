package exercise

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (n Node) String() string {
	//return fmt.Sprintf("Val: %d", n.Val)
	p := [][]string{}
	p = append(p, []string{fmt.Sprintf("Val: %d ", n.Val)})
	StringRec(n.Left, &p, 1)
	StringRec(n.Right, &p, 1)

	return fmt.Sprintf("%s", p)
}

func StringRec(n *Node, p *[][]string, lvl int) {
	if n == nil {
		return
	}

	if len(*p) > lvl {
		(*p)[lvl] = append((*p)[lvl], fmt.Sprintf("Val: %d ", n.Val))
	} else {
		*p = append(*p, []string{fmt.Sprintf("Val: %d ", n.Val)})
	}

	StringRec(n.Left, p, lvl+1)
	StringRec(n.Right, p, lvl+1)
}

func getMaxWidth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	fmt.Println("height: ", height(root))

	return getMaxWidthRec(root)
}
func getMaxWidthRec(root *Node) int {
	h := height(root)
	var max int
	for i := 1; i < h; i++ {
		w := getWidth(root, i)
		if max < w {
			max = w
		}
	}
	return max
}

func getWidth(root *Node, lvl int) int {

	if root == nil {
		return 0
	}
	if lvl == 1 {
		return 1
	}

	lvl--
	return getWidth(root.Left, lvl) + getWidth(root.Right, lvl)

}

func height(root *Node) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
