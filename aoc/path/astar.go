package path

import (
	"iter"

	"github.com/AntonKosov/advent-of-code-2019/aoc/math"
	"github.com/AntonKosov/advent-of-code-2019/aoc/slice"
)

func AStar(
	start, finish math.Vector2[int], heuristicFn func(math.Vector2[int], math.Vector2[int]) bool,
	nextPositions func(math.Vector2[int]) iter.Seq[math.Vector2[int]],
) []math.Vector2[int] {
	type Node struct {
		pos  math.Vector2[int]
		prev *Node
	}
	queue := slice.NewPriorityQueue(func(n1, n2 Node) bool { return heuristicFn(n1.pos, n2.pos) })
	queue.Push(Node{pos: start, prev: nil})
	visited := map[math.Vector2[int]]bool{}
	for queue.Len() > 0 {
		currentNode := queue.Pop()
		pos := currentNode.pos
		if visited[pos] {
			continue
		}
		visited[pos] = true

		if currentNode.pos == finish {
			var path []math.Vector2[int]
			for node := &currentNode; node != nil; node = node.prev {
				path = append(path, node.pos)
			}
			slice.Reverse(path)
			return path
		}

		for nextPos := range nextPositions(pos) {
			if visited[nextPos] {
				continue
			}
			nextNode := Node{pos: nextPos, prev: &currentNode}
			queue.Push(nextNode)
		}
	}

	return nil
}
