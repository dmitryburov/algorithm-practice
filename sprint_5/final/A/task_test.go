package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

type TestItem struct {
	input  io.Reader
	output string
}

func TestTask(t *testing.T) {

	tasks := generateTasks()

	for i, v := range tasks {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			buf := strings.Builder{}
			PyramidSort(v.input, &buf)
			if buf.String() != v.output {
				t.Errorf("Неверный ответ решения!\nОтвет: \n%v \nВерно: \n%v", v.output, buf.String())
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{{strings.NewReader(`5
alla 4 100
gena 6 1000
gosha 2 90
rita 2 90
timofey 4 80`), `gena
timofey
alla
gosha
rita
`}, {strings.NewReader(`5
alla 0 0
gena 0 0
gosha 0 0
rita 0 0
timofey 0 0`), `alla
gena
gosha
rita
timofey
`}, {strings.NewReader(`5
alla 0 10
gena 0 9
gosha 0 8
rita 0 7
timofey 0 0
debilniy_comment`), `timofey
rita
gosha
gena
alla
`}, {strings.NewReader(`5
a 0 0
ab 0 0
abc 0 0
abcd 0 0
abcde 0 0
Проверка лексикографической сортировки`), `a
ab
abc
abcd
abcde
`}, {strings.NewReader(`13
tufhdbi 76 58
rqyoazgbmv 59 78
qvgtrlkmyrm 35 27
tgcytmfpj 70 27
xvf 84 19
jzpnpgpcqbsmczrgvsu 30 3
evjphqnevjqakze 92 15
wwzwv 87 8
tfpiqpwmkkduhcupp 1 82
tzamkyqadmybky 5 81
amotrxgba 0 6
easfsifbzkfezn 100 28
kivdiy 70 47
Просто тест`), `easfsifbzkfezn
evjphqnevjqakze
wwzwv
xvf
tufhdbi
tgcytmfpj
kivdiy
rqyoazgbmv
qvgtrlkmyrm
jzpnpgpcqbsmczrgvsu
tzamkyqadmybky
tfpiqpwmkkduhcupp
amotrxgba
`}}
}
