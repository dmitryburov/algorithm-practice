package main

import (
	"fmt"
	"testing"
)

type task struct {
	Str1, Str2 string
	Result     string
}

func TestTask(t *testing.T) {
	var taskItems = generateTask()
	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := solution(taskItems[i].Str1, taskItems[i].Str2)
			if res != taskItems[i].Result {
				t.Errorf("Неверный ответ решения!\nОтвет: %s \nВерно: %s", res, taskItems[i].Result)
			}
		})
	}

}

func generateTask() []task {
	return []task{
		{
			Str1:   "abcp",
			Str2:   "ahpc",
			Result: "False",
		},
		{
			Str1:   "abc",
			Str2:   "ahbgdcu",
			Result: "True",
		},
		{
			Str1:   "islx",
			Str2:   "yoytgtshldmogkdburkbcfvoapepjpcuwemusfkfztrzxstytrnarlizjhuoscuzlraezlaweipuuqdgvhwkhhoufexojaps",
			Result: "True",
		},
		{
			Str1:   "ijha",
			Str2:   "hmrqvftefyixinahlzgbkidroxiptbbkjmtwpsujevkulgrjiwiwzyhngulrodiwyg",
			Result: "False",
		},
		{
			Str1:   "m",
			Str2:   "zzkqxfdxbbjqhatygmtmpgbhumicrhtjkrfblwwnjlebsfdawznznxwyzehpubvdukmgwrivygosfkdquwxvkvtfvwjfwtvjvtplpckktmnhfnxprjetpxnddoiiqotzrjjfdwnzdjvtclcqwsvsegnuajookwppzfsf",
			Result: "True",
		},
		{
			Str1:   " ",
			Str2:   "ggzjxvsyleybxdjtmrfdvvuiwhamerqnavroixmxkyirnktbfsyijmtqmapjwuttpwefgfdtisrmzpfalwgefvbeecbdcllvfkylplpjmqquqhmjsplpmancofucoymnbtoitltfhuolrgzlwikinojtjvhgeszkmitqgslrvmfyuxnlkzhguvfqxfbknvmxflyxuxcrgwnhbonuqjvmzmzfdapsukwqqdzpicbnhrntgdqxnajwrrgjqlncuezesuzpscjtihrfadapwslrflbtsnvismrhpjfjqvgc",
			Result: "True",
		},
	}
}
