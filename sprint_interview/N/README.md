# M. Атака клонов

**В этой задаче по умолчанию выбран компилятор Make. Решение отправляется в виде файла с расширением, соответствующим используемому языку.**

В этой задаче требуется создать копию связного графа. Оригинальный граф задается одной вершиной. Вершина содержит свое уникальное значение – value, и список соседей neighbours. Граф будет считаться копией, если в графе-копии есть связь между вершинами со значениями v1 и v2 тогда и только тогда, когда она есть в оригинальном графе.

Все вершины графа-копии должны быть созданы заново, то есть нельзя переиспользовать вершины из оригинального графа. Создавайте новые вершины с помощью публичных конструкторов и фабричных методов, указанных в шаблонах.

## Формат ввода
Функция cloneGraph принимает стартовую вершину, которая принадлежит оригинальному графу.

Используйте шаблон:

```go
package main

/** Comment it before submitting
type Node struct {
	val        int
	neighbours []*Node
}

func (n *Node) setVal(val int) {
	n.val = val
}

func (n *Node) getVal() int {
	return n.val
}

func (n *Node) setNeighbours(neighbours []*Node) {
	n.neighbours = neighbours
}

func (n *Node) getNeighbours() []*Node {
	return n.neighbours
}

func (n *Node) addNeighbour(neighbour *Node) {
	n.neighbours = append(n.neighbours, neighbour)
}

func newNode(val int) *Node {
	n := &Node{}
	n.setNeighbours(make([]*Node, 0))
	n.setVal(val)
	return n
}
*/

func cloneGraph(node *Node) *Node {
	// Your code
	// “ヽ(´▽｀)ノ”
}

```

Гарантируется, что число вершин и рёбер графа не превосходит 10<sup>5</sup>.


## Формат вывода

Функция cloneGraph должна возвращать вершину, являющуюся копией стартовой вершины в оригинальном графе.

