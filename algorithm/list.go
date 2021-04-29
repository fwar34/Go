package main
import "fmt"

type LinkNode struct {
	Data int64
	Next *LinkNode
}

func main() {
    node := new(LinkNode)
	node.Data = 2

	node1 := new(LinkNode)
	node1.Data = 3
	node.Next = node1 // 将 node1 链接到 node 节点上

	node2 := new(LinkNode)
	node2.Data = 4
	node1.Next = node2 // 将 node2 连接到 node1 节点上

	// 按顺序打印数据
	nowNode := node
	for {
		if nowNode != nil {
			fmt.Println(nowNode.Data)
			// 获取下一节点
			nowNode = nowNode.Next
			continue
		}

		// 如果下一节点为空，表示链表结束了
		break
	}
}

