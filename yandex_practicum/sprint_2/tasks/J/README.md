# J. Списочная очередь

Любимый вариант очереди Тимофея — очередь, написанная с использованием связного списка. Помогите ему с реализацией. Очередь должна поддерживать выполнение трёх команд:

- get() — вывести элемент, находящийся в голове очереди, и удалить его. Если очередь пуста, то вывести «error».
- put(x) — добавить число x в очередь
- size() — вывести текущий размер очереди


## Формат ввода

В первой строке записано количество команд n — целое число, не превосходящее 1000. В каждой из следующих n строк записаны команды по одной строке.

## Формат вывода

Выведите ответ на каждый запрос по одному в строке.

## Примечания

Все операции должны выполняться за O(1).

### Пример 1

<table class="sample-tests">
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
            put -34<br>
            put -23<br>
            get<br>
            size<br>
            get<br>
            size<br>
            get<br>
            get<br>
            put 80<br>
            size<br>
        </td>
        <td>
            -34<br>
            1<br>
            -23<br>
            0<br>
            error<br>
            error<br>
            1<br>
            <br>
            <br>
            <br>
            <br>
        </td>
     </tr>
  </tbody>
</table>


### Пример 2

<table class="sample-tests">
  <thead>
     <tr>
        <th>Ввод</th>
        <th>Вывод</th>
     </tr>
  </thead>
  <tbody>
     <tr>
        <td>
            6<br>
            put -66<br>
            put 98<br>
            size<br>
            size<br>
            get<br>
            get<br>
        </td>
        <td>
            2<br>
            2<br>
            -66<br>
            98<br>
            <br>
            <br>
            <br>
        </td>
     </tr>
  </tbody>
</table>

