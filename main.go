package main

import "fmt"

type puzzle struct {
	input  [3][3]int
	stack  map[int][3][3]int
	target [3][3]int
	count  int
}

func (r *puzzle) printClean() {
	for _, v := range r.input {
		fmt.Println(v)
	}
}

func (r *puzzle) inputClean() {
	fmt.Println("입력: ")
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d %d %d", &r.input[i][0], &r.input[i][1], &r.input[i][2])
	}
}

func (r *puzzle) findBlank() (int, int) {
	var a, b int
	for i, _ := range r.input {
		for j, v := range r.input[i] {
			if v == 0 {
				a, b = i, j //blank 좌표를 반환
			}
		}
	}
	return a, b
}

func (r *puzzle) change(in uint8) bool { //blank를 지정한 방향과 바꿈
	a, b := r.findBlank()
	result := false
	switch {
	case in == 0 && b-1 >= 0: // <
		r.input[a][b], r.input[a][b-1] = r.input[a][b-1], r.input[a][b]
		//fmt.Print("<-")
		result = true
	case in == 1 && b+1 < 3: // >
		r.input[a][b], r.input[a][b+1] = r.input[a][b+1], r.input[a][b]
		//fmt.Print("<-")
		result = true
	case in == 2 && a-1 >= 0: // ^
		r.input[a][b], r.input[a-1][b] = r.input[a-1][b], r.input[a][b]
		//fmt.Print("^-")
		result = true
	case in == 3 && a+1 < 3: // v
		r.input[a][b], r.input[a+1][b] = r.input[a+1][b], r.input[a][b]
		//fmt.Print("v-")
		result = true
	}
	return result
}

func (r *puzzle) historyCheck(in int) bool { //true는 옛날이랑 같은 상태가 없는 것
	temp := *r
	temp.change(uint8(in))
	for i, v := range r.stack {
		if temp.input == v && i > 1 {
			return false
		}
	}
	return true
}

func (r *puzzle) explore(in int) {
	r.stack = make(map[int][3][3]int)
	if in < 0 {
		in = 9223372036854775807 //int64의 최대값
	}
	r.count = 0
	for i := 0; i <= in; i++ { // 일단 n번 실행, 결과 찾으면 종료 조건 추가예정
		if r.input == r.target { //타겟을 만나면 종료
			fmt.Println("eureka! 턴 횟수는 ", r.count, "입니다.")
			return
		}
		r.stack[i] = r.input
		upperlevel := 0
		for arrow := 0; arrow <= 3; arrow++ {
			//fmt.Println(arrow, tempstate, r.historyCheck(arrow)) //debug용
			if r.historyCheck(arrow) { //true는 옛날이랑 같은 상태가 없는 것
				if r.change(uint8(arrow)) {
					//fmt.Print(arrow) //디버그용
					r.count++
					arrow = 10
					upperlevel = 0
				}
			} else { //옛날이랑 같은 상태가 있는 경우
				r.input = r.stack[i-upperlevel]
				upperlevel++
			}
		}
	}
}

func main() {
	var pan puzzle
	pan.target = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}} //타겟 지정
	pan.inputClean()
	fmt.Println("연산 시작 기다려주세용.")
	pan.explore(-1)  //최대 루프 횟수 지정 -1은 int64의 최대값
	pan.printClean() //결과 출력
}
