package main

import (
	"fmt"
	"time"
)

type puzzle struct {
	input  [3][3]int         //입력을 받고 움직임이 이루어질 판
	stack  map[int][3][3]int //히스토리를 쌓아줄 스택
	target [3][3]int         //목표를 지정할 변수
	count  int               //몇 번의 change()가 일어났는지 카운팅
}

func (r *puzzle) printClean() { //예쁘게 출력
	for _, v := range r.input {
		fmt.Println(v)
	}
}

func (r *puzzle) inputClean() { //형식에 맞게 입력 받기
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
				a, b = i, j //blank 좌표를 찾아서 반환
			}
		}
	}
	return a, b
}

func (r *puzzle) change(in uint8) bool { //blank를 지정한 방향과 바꿈
	a, b := r.findBlank()
	result := false
	switch {
	case in == 0 && b-1 >= 0: // < 바꾸려는 방향이 인덱스 밖으로 튀는 경우 방지
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
	temp := *r                  //현재 상태를 temp에 저장
	temp.change(uint8(in))      //temp에서 change 수행, r.input에 영향가지 않도록
	for i, v := range r.stack { //change된 temp에 r.stack(과거기록) 중 같은 것이 있는지 확인
		if temp.input == v && i > 1 {
			return false
		}
	}
	return true
}

func (r *puzzle) explore(in int) { // 실제 탐색을 실행하는 함수
	r.stack = make(map[int][3][3]int) // 히스토리를 저장할 스택
	if in < 0 {
		in = 9223372036854775807 //int64의 최대값
	}
	r.count = 0
	upperlevel := 0            //더이상 갈 곳이 없을 때가 반복될 때 몇개의 층을 다시 올라가야하는지 저장하는 변수
	for i := 0; i <= in; i++ { // in번 실행
		if r.input == r.target { //타겟을 만나면 종료
			fmt.Println("eureka! change 횟수는 ", r.count, "입니다.")
			fmt.Println("만들어진 노드 수는 ", i, "입니다.")
			return
		}
		for arrow := 0; arrow <= 3; arrow++ {
			//fmt.Println("count: ", r.count, "arrow:", arrow, "ulevel:", upperlevel, r.historyCheck(arrow)) //디버그용
			if r.historyCheck(arrow) { //true는 옛날이랑 같은 상태가 없는 것
				if r.change(uint8(arrow)) { //change가 성공하면 true를 반환함
					//fmt.Print(arrow) //디버그용
					upperlevel = 0
					r.count++
					r.stack[r.count] = r.input // stack[노드번호]에 현재 판의 상태를 쌓음
					arrow = 10                 //for문 탈출
				}
			} else if arrow == 3 { //옛날이랑 같은 상태가 있는 경우, 그리고 모든 방향을 시도해보았을 경우
				r.input = r.stack[r.count-upperlevel] //upperlevel만큼 올라가서 판을 갱신
				upperlevel++                          //갱신 이후에 history에서 같은 것을 또 만날 경우 더 위에 층을 불러오기 위함.
			}
		}
		//r.printClean()
	}
}

func main() {
	var pan puzzle
	pan.target = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 0}} //타겟 지정
	pan.inputClean()
	fmt.Println("연산 시작 기다려주세용.")
	startTime := time.Now()
	pan.explore(-1)  //최대 루프 횟수 지정 -1은 int64의 최대값
	pan.printClean() //결과 출력
	fmt.Println("연산에 걸린 시간(초)): ", time.Since(startTime).Seconds())
}
