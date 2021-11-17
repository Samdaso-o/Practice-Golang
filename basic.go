package basic

import "fmt"

//struct 사용해보기
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"g", "h"}
	han := person{name: "han", age: 12, favFood: favFood}
	fmt.Println(han)
}

//map 사용, 파이썬의 dict을 생각하면 되지만, 속도를 위해서 타입을 지정해줘야 한다.
// func main() {
// 	han := map[string]string{"name": "han", "age": "1"}
// 	for key, value := range han {
// 		fmt.Println(key, value)
// 	}
// }

//리스트에 관한여, 변수명:= 리스트크기,타입,{값}
// func main() {
// 	name := []string{"g", "h", "f"}
// 	name = append(name, "w", "d", "e", "qw")
// 	fmt.Println(name)
// }

// func main() {
// 	name := [5]string{"g", "h", "f"}
// 	name[3] = "e"
// 	name[4] = "w"
// 	fmt.Println(name)
// }

//메모리 할당에 관하여...
// func main() {
// 	a := 1
// 	b := a
// 	fmt.Println(a, b)
// }

// func address() {
// 	a := 1
// 	b := a
// 	fmt.Print(&a, &b)
// }

// func memory() {
// 	a := 1
// 	b := &a
// 	fmt.Println(a, b)
// }

// func memoryValue() {
// 	a := 1
// 	b := &a
// 	fmt.Println(*b)
// }

// func usememory() {
// 	a := 1
// 	b := &a
// 	*b = 20
// 	fmt.Println(a)
// 	//a의 메모리 값을 할당한 b로 a 업데이트 하기
// }

//switch 사용하기
// func check(age int) bool {
// 	switch {
// 	case age < 10:
// 		return false
// 	case age == 20:
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	result := check(20)
// 	println(result)
// }

//if문 사용하기
// func check(age int) bool {
// 	if age < 18 {
// 		return false
// 	} else {
// 		return true
// 	}
// }

// func main() {
// 	result := check(20)
// 	println(result)
// }

//index는 날리고, number 더해서 반환하기
// func alladd(number ...int) int {
// 	total := 0
// 	for _, number := range number {
// 		total = total + number
// 	}
// 	return total
// }

// func main() {
// 	result := alladd(1, 2, 3, 4, 5, 6)
// 	println(result)
// }

// 나중에 참고할 것 for문
// func alladd(number ...int) int {
// 	for i := 0; i < len(number); i++ {
// 		fmt.Println(number[i])
// 	}
// 	return 1
// }

// func main() {
// 	alladd(1, 2, 3, 4, 5, 6)
// }

//range는 index와 값을 같이 준다 index, number 순으로...
// func alladd(number ...int) int {
// 	for number := range number {
// 		fmt.Println(number)
// 	}
// 	return 1
// }

// func main() {
// 	total := alladd(1, 2, 3, 4, 5, 6)
// 	fmt.Println(total)
// }

//반환할 return값 미리 선언해서 반환하기, return에 적어도 됨^^, defer을 사용해 func이 끝나면 실행할 작업 선언
// func lenAndUpper(name string) (length int, uppercase string) {
// 	defer fmt.Println("done")
// 	length = len(name)
// 	uppercase = strings.ToUpper(name)
// 	return
// }

// func main() {
// 	length, uppername := lenAndUpper("hannnnnn")
// 	fmt.Println(length, uppername)
// }

//여러개의 값 받기, 리스트 형식으로 출력됨
// func repeat(word ...string) {
// 	fmt.Println(word)
// }

// func main() {
// 	repeat("h", "t", "f", "r", "d")
// }

//기본 문법 핸들링
// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// func main() {
// 	length, uppername := lenAndUpper("han")
// 	h, _ := lenAndUpper("ggggg")
// 	fmt.Println(length, uppername)
// 	fmt.Println(h)
// }

//기본 곱셉 연산
// func multiply(a, b int) int {
// 	return a * b
// }

// func main() {
// 	fmt.Println(multiply(6, 3))
// }

// 기본 프린트
// func main() {
// 	const lastname string = "han"
// 	fmt.Println(lastname)
// 	hi := "hihihihi"
// 	fmt.Println(hi)
// 	var firstname string = "???"
// 	fmt.Println(firstname)
// 	firstname = "seunghun"
// 	fmt.Println((firstname))
// }
