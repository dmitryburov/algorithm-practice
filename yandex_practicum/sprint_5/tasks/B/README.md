# B. Сбалансированное дерево

Гоше очень понравилось слушать рассказ Тимофея про деревья. Особенно часть про сбалансированные деревья. Он решил написать функцию, которая определяет, сбалансировано ли дерево.

Дерево считается сбалансированным, если левое и правое поддеревья каждой вершины отличаются по высоте не больше, чем на единицу.

![IMG](image.png)

## Формат ввода

На вход функции подаётся корень бинарного дерева.

**Замечания про отправку решений**

По умолчанию выбран компилятор make. Шаблон для Go:

```go
package main

/**
Comment it before submitting
type Node struct {
	value  int
	left   *Node
	right  *Node
}
**/

func Solution(root *Node) bool {
	// Your code
	// “ヽ(´▽｀)ノ”
}
```

## Формат вывода

Функция должна вернуть True, если дерево сбалансировано в соответствии с критерием из условия, иначе - False.