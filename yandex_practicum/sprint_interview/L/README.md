# K. Пересечение отрезков

Вам даны две последовательности отрезков. Каждый отрезок задаётся координатой начала и конца — [start<sub>i</sub>, end<sub>i</sub>]. 

Отрезки каждой последовательности отсортированы слева направо и не имеют общих точек.
Найдите пересечение двух последовательностей отрезков. То есть третью последовательность отрезков, такую, что:

1. Каждый отрезок содержится в некотором отрезке и первой, и второй последовательности;
2. Никакой отрезок нельзя увеличить; 
3. Отрезки этой последовательности не имеют общих точек; 
4. Отрезки в последовательности также отсортированы в порядке возрастания.


## Формат ввода

В первой строке дано число отрезков в первой последовательности n (0 ≤ n ≤ 100000) В следующих n строках даны отрезки первой последовательности по одному на строке. 

Каждый отрезок записан в формате start<sub>i</sub> end<sub>i</sub>, координаты начала и конца целые неотрицательные числа, не превосходящие по модулю 10<sup>9</sup>.

## Формат вывода
Выведите по одному в строке отсортированные отрезки из пересечения последовательностей в том же формате, что и во входных данных. Заметьте, что длина отрезков в пересечении может быть нулевой, в этом случае start<sub>i</sub> = end<sub>i</sub>.

### Пример 1

<table><tr>
<td>
3<br>
1 4<br>
5 10<br>
15 16<br>
2<br>
0 2<br>
4 5
</td>
<td>
1 2<br>
4 4<br>
5 5<br>
<br>
<br>
<br>
<br>
</td>
</tr></table>

### Пример 2

<table><tr>
<td>
1<br>
1 4<br>
1<br>
1 4
</td>
<td>
1 4<br>
<br>
<br>
<br>
</td>
</tr></table>



