# E. Дерево поиска

Гоша понял, что такое дерево поиска, и захотел написать функцию, которая определяет, является ли заданное дерево деревом поиска. Значения в левом поддереве должны быть строго меньше, в правом —- строго больше значения в узле.

Помогите Гоше с реализацией этого алгоритма.

![IMG](image.png)

## Формат ввода

На вход функции подается корень бинарного дерева.

**Замечания про отправку решений**

По умолчанию выбран компилятор make.
Решение нужно отправлять в виде файла с расширением, которое соответствует вашему языку программирования.

Go:

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

Функция должна вернуть True, если дерево является деревом поиска, иначе - False.
