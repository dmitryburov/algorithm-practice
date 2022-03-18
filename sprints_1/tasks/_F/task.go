package main

// F. Палиндром
// Помогите Васе понять, будет ли фраза палиндромом‎. Учитываются только буквы и цифры, заглавные и строчные буквы считаются одинаковыми.
// Решение должно работать за O(N), где N — длина строки на входе.
//
// Формат ввода:
// В единственной строке записана фраза или слово. Буквы могут быть только латинские. Длина текста не превосходит 20000 символов.
// Фраза может состоять из строчных и прописных латинских букв, цифр, знаков препинания.
//
// Формат вывода:
// Выведите «True», если фраза является палиндромом, и «False», если не является.
//
// Примеры:
// inp: 								>>> 	out:
//    A man, a plan, a canal: Panama				True
//
// inp: 	>>> 	out:
//    zo				False
//func main() {
//	inp := getInputFromFile()
//	defer inp.Close()
//
//
//	reader := bufio.NewReader(inp)
//	line, _ := reader.ReadString('\n')
//	line = strings.ToLower(line)
//	line = strings.TrimSpace(line)
//	line = strings.
//
//	fmt.Println(line)
//
//	words := strings.Split()
//	for i := 0; i < len(strItems); i++ {
//		trimStr := strings.TrimSpace(strItems[i])
//		if len(trimStr) > len(find) {
//			find = trimStr
//		}
//	}
//
//	fmt.Println(checkPalindrom())
//}
//
//func checkPalindrom(phrase string) (result string) {
//	result = "False"
//
//	for i := 0; i < len(phrase)/2; i++ {
//		fmt.Println(phrase[i], phrase[len(phrase)-i+1:])
//	}
//
//	return "True"
//}
//
//func getInputFromFile() *os.File {
//	file, err := os.Open("./input.txt")
//	if err != nil {
//		panic(err)
//	}
//
//	return file
//}
