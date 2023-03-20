import datetime, copy

class Puzzle:
    def __init__(self):
        self.input = [[],[],[]] #입력을 받고 움직임이 이루어질 판
        self.stack = {} #히스토리를 쌓아줄 스택
        self.target = [] #목표를 지정할 변수

    def printClean(self): #예쁘게 출력
        for row in self.input:
            print(row)

    def inputClean(self): #형식에 맞게 입력 받기
        print("입력: ")
        for i in range(3):
            self.input[i] = list(map(int, input().split(" ")))

    def findBlank(self):
        for i in range(3):
            for j in range(3):
                if self.input[i][j] == 0: #//blank 좌표를 찾아서 반환
                    return (i, j)

    def change(self, inVal): #//blank를 지정한 방향과 바꿈
        a, b = self.findBlank()
        result = False
        if inVal == 0 and b-1 >= 0: #< 바꾸려는 방향이 인덱스 밖으로 튀는 경우 방지
            self.input[a][b], self.input[a][b-1] = self.input[a][b-1], self.input[a][b]
            result = True
        elif inVal == 1 and b+1 < 3:#>
            self.input[a][b], self.input[a][b+1] = self.input[a][b+1], self.input[a][b]
            result = True
        elif inVal == 2 and a-1 >= 0:#^
            self.input[a][b], self.input[a-1][b] = self.input[a-1][b], self.input[a][b]
            result = True
        elif inVal == 3 and a+1 < 3:#v
            self.input[a][b], self.input[a+1][b] = self.input[a+1][b], self.input[a][b]
            result = True
        return result

    def historyCheck(self, inVal): #true는 옛날이랑 같은 상태가 없는 것
        #temp = copy.deepcopy(self) #성능 많이 드는 곳 수정필요
        temp = Puzzle() 
        temp.input = copy.deepcopy(self.input) #현재 상태를 temp에 저장
        temp.change(inVal) #temp에서 change 수행, self.input에 영향가지 않도록
        for i, v in self.stack.items(): #change된 temp에 self.stack(과거기록) 중 같은 것이 있는지 확인
            if temp.input == v and i > 1:
                return False
        return True

    def explore(self, inVal): #실제 탐색을 실행하는 함수
        if inVal < 0:
            inVal = 9223372036854775807
        self.count = 0
        upperlevel = 0 #더이상 갈 곳이 없을 때가 반복될 때 몇개의 층을 다시 올라가야하는지 저장하는 변수
        for i in range(inVal + 1):#in번 실행
            if self.input == self.target:#타겟을 만나면 종료
                print("eureka! change 횟수는", self.count, "입니다.")
                print("만들어진 노드 수는", i, "입니다.")
                return
            arrow = 0
            while arrow <= 3:
                #print("count:",self.count ,arrow, upperlevel, self.historyCheck(arrow))
                if self.historyCheck(arrow):#true는 옛날이랑 같은 상태가 없는 것
                    if self.change(arrow):#change가 성공하면 true를 반환함
                        upperlevel = 0
                        self.count += 1
                        self.stack[self.count] = copy.deepcopy(self.input)#stack[노드번호]에 현재 판의 상태를 쌓음
                        arrow = 10#for문 탈출
                elif arrow == 3: #옛날이랑 같은 상태가 있는 경우, 그리고 모든 방향을 시도해보았을 경우
                    self.input = self.stack[self.count-upperlevel]#upperlevel만큼 올라가서 판을 갱신
                    upperlevel += 1#갱신 이후에 history에서 같은 것을 또 만날 경우 더 위에 층을 불러오기 위함.
                arrow += 1


pan = Puzzle()
pan.target = [[1,2,3],[4,5,6],[7,8,0]]  # 타겟 지정
pan.inputClean()
print("연산 시작 기다려주세용.")
startTime = datetime.datetime.now()
pan.explore(-1)  # 최대 루프 횟수 지정 -1은 int64의 최대값
pan.printClean()  # 결과 출력
endTime = datetime.datetime.now()
print("연산에 걸린 시간: ", endTime-startTime)
