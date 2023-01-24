# G. Стек - MaxEffective

Реализуйте класс StackMaxEffective, поддерживающий операцию определения максимума среди элементов в стеке. Сложность
операции должна быть O(1). Для пустого стека операция должна возвращать None. При этом push(x) и pop() также должны
выполняться за константное время.

## Формат ввода

В первой строке записано одно число — количество команд, оно не превосходит 100000. Далее идут команды по одной в
строке. Команды могут быть следующих видов:

- push(x) — добавить число x в стек;
- pop() — удалить число с вершины стека;
- get_max() — напечатать максимальное число в стеке;

Если стек пуст, при вызове команды get_max нужно напечатать «None», для команды pop — «error».

## Формат вывода

Для каждой команды get_max() напечатайте результат её выполнения. Если стек пустой, для команды get_max() напечатайте
«None». Если происходит удаление из пустого стека — напечатайте «error».

### Пример 1

<table>
  <thead>
     <tr>
        <th>Ввод</th>
        <th>Вывод</th>
     </tr>
  </thead>
  <tbody>
     <tr>
        <td>
            10<br>
            pop<br>
            pop<br>
            push 4<br>
            push -5<br>
            push 7<br>
            pop<br>
            pop<br>
            get_max<br>
            pop<br>
            get_max<br>
        </td>
        <td>
            error<br>
            error<br>
            4<br>
            None<br>
            <br>
            <br>
            <br>
            <br>
            <br>
            <br>
            <br>
        </td>
     </tr>
  </tbody>
</table>

<br> 

### Пример 2

<table>
  <thead>
     <tr>
        <th>Ввод</th>
        <th>Вывод</th>
     </tr>
  </thead>
  <tbody>
     <tr>
        <td>
            10<br>
            get_max<br>
            push -6<br>
            pop<br>
            pop<br>
            get_max<br>
            push 2<br>
            get_max<br>
            pop<br>
            push -2<br>
            push -6<br>
        </td>
        <td>
            None<br>
            error<br>
            None<br>
            2<br>
            <br>
            <br>
            <br>
            <br>
            <br>
            <br>
            <br>
        </td>
     </tr>
  </tbody>
</table>