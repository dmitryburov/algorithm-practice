# L. Фибоначчи по модулю

У Тимофея было очень много стажёров, целых N (0 ≤ N ≤ 10<sup>6</sup>) человек. Каждый стажёр хотел быть лучше своих предшественников, поэтому i-й стажёр делал столько коммитов, сколько делали два предыдущих стажёра в сумме.
Два первых стажёра были менее инициативными — они сделали по одному коммиту.

Пусть F<sub>i</sub> —– число коммитов, сделанных i-м стажёром (стажёры нумеруются с нуля). Первые два стажёра сделали по одному коммиту: F<sub>0</sub>=F<sub>1</sub>=1.
Для всех i ≥ 2 выполнено F<sub>i</sub>=F<sub>i-1</sub>+F<sub>i-2</sub>.

Определите, сколько кода напишет следующий стажёр – найдите последние k цифр числа F<sub>n</sub>.

**Как найти <i>k</i> последних цифр**

Чтобы вычислить k последних цифр некоторого числа x, достаточно взять остаток от его деления на число 10<sup>k</sup>. 
Эта операция обозначается как <i>x mod 10<sup>k</sup></i>. Узнайте, как записывается операция взятия остатка по модулю в вашем языке программирования.


## Формат ввода

В первой строке записаны через пробел два целых числа n (0 ≤ n ≤ 10<sup>6</sup>) и k (1 ≤ k ≤ 8).


## Формат вывода

Выведите единственное число – последние k цифр числа F<sub>n</sub>.
Если в искомом числе меньше k цифр, то выведите само число без ведущих нулей.


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
        <td>3 1</td>
        <td>3</td>
     </tr>
  </tbody>
</table>

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
        <td>10 1</td>
        <td>9</td>
     </tr>
  </tbody>
</table>