/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
	n := len(grid)
	if n == 0 {
		return nil
	}
	return buildQuadTree(grid, 0, 0, n)
}

func buildQuadTree(grid [][]int, x, y, length int) *Node {
	if length == 1 {
		return &Node{
			Val:         grid[x][y] == 1,
			IsLeaf:      true,
			TopLeft:     nil,
			TopRight:    nil,
			BottomLeft:  nil,
			BottomRight: nil,
		}
	}

	half := length / 2
	topLeft := buildQuadTree(grid, x, y, half)
	topRight := buildQuadTree(grid, x, y+half, half)
	bottomLeft := buildQuadTree(grid, x+half, y, half)
	bottomRight := buildQuadTree(grid, x+half, y+half, half)

	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf &&
		topLeft.Val == topRight.Val && topRight.Val == bottomLeft.Val && bottomLeft.Val == bottomRight.Val {
		return &Node{
			Val:         topLeft.Val,
			IsLeaf:      true,
			TopLeft:     nil,
			TopRight:    nil,
			BottomLeft:  nil,
			BottomRight: nil,
		}
	} else {
		return &Node{
			Val:         false, // Val在非葉子節點中並不重要
			IsLeaf:      false,
			TopLeft:     topLeft,
			TopRight:    topRight,
			BottomLeft:  bottomLeft,
			BottomRight: bottomRight,
		}
	}
}
