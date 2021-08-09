package srv

import (
	"errors"
	"fmt"
)

type ViewNode struct {
	Name string
	Os   string
	Arch string
	Pods []ViewPod
}

type ViewPod struct {
	Name string
}

type ViewNodeData struct {
	Nodes []ViewNode
}

type View interface {
	Printout() error
}

func (vnd ViewNodeData) Printout() error {
	if vnd.Nodes == nil {
		return errors.New("list of view nodes must not be null")
	}
	l := len(vnd.Nodes)
	if l == 0 {
		fmt.Println("no nodes to display...")
		return nil
	}
	fmt.Printf("viewing %d node(s):\n", l)
	for _, n := range vnd.Nodes {
		fmt.Printf("%s (%s/%s)\n", n.Name, n.Os, n.Arch)
		for _, p := range n.Pods {
			fmt.Printf("* %s", p.Name)
		}
	}
	return nil
}
