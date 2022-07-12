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
			Solution(v.input, &buf)
			if buf.String() != v.output {
				t.Errorf("Неверный ответ решения!\nОтвет: \n%v \nВерно: \n%v", buf.String(), v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{strings.NewReader(`3
2[a]2[ab]
3[a]2[r2[t]]
a2[aa3[b]]
`), "aaa"},
		{strings.NewReader(`3
abacabaca
2[abac]a
3[aba]
`), "aba"},
		{strings.NewReader(`1
2[aba]
`), "abaaba"},
		{strings.NewReader(`5
2[jwcn]2[1[emgu]]1[bzgy]
jwcnjwcnemguemgu2[ctr]2[1[cnma]3[y]2[b]3[v]3[gas]]3[oc]
jwcnjwcnemguemgu2[2[mou]1[y]2[dd]]1[1[ynt]]1[nnq]2[qfw]2[e]1[1[k]1[utz]1[fal]2[g]]
jwcnjwcnemguemgu1[j]3[2[h]3[iqd]3[xen]1[ia]]2[3[i]]2[li]
2[jwcn]2[1[emgu]]2[whub]3[qdrz]3[3[potn]2[l]1[wc]1[snzs]]1[3[u]3[myi]2[fdk]2[ot]]2[qp]
`), "jwcnjwcnemguemgu"},
		{strings.NewReader(`5
vkskkkcuqaqaqauuisissxjesxjewqwqwqbbgz2[fm]1[3[tama]2[xyk]3[j]2[rp]1[af]]3[3[atp]]1[3[jzs]3[cah]2[o]3[sfmr]2[lp]3[v]]2[2[u]3[ah]]3[ydrr]
vkskkkcuqaqaqauuisissxjesxjewqwqwqbbgz2[rgr]1[x]3[1[ded]]1[1[qx]2[ang]]1[wx]
vkskkkcuqaqaqauuisissxjesxjewqwqwqbbgz1[3[qif]3[we]3[ljt]2[rlps]3[i]]
1[vks]3[k]1[cu]1[3[qa]2[u]2[is]2[sxje]3[wq]1[bbgz]]3[1[xfob]]2[s]3[h]2[2[fz]3[qs]]1[jxrf]2[ze]
1[vks]3[k]1[cu]1[3[qa]2[u]2[is]2[sxje]3[wq]1[bbgz]]3[fe]3[2[hbll]3[jvt]]2[eca]
`), "vkskkkcuqaqaqauuisissxjesxjewqwqwqbbgz"},
	}
}
