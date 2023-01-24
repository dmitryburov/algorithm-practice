# H. Числовые пути

Вася и его друзья решили поиграть в игру. Дано дерево, в узлах которого записаны цифры от 0 до 9. Таким образом, каждый путь от корня до листа содержит число, получившееся конкатенацией цифр пути (склеиванием цифр пути в одно число). Нужно найти сумму всех таких чисел в дереве.

Гарантируется, что ответ не превосходит 20 000.

![IMG](image.png)

## Формат ввода

На вход подается корень дерева.


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

func Solution(root *Node) int {
    // Your code
    // “ヽ(´▽｀)ノ”
}
```

## Формат вывода

Функция должна вернуть число, равное сумме чисел всех путей в дереве.