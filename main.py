import datetime, copy

class Puzzle:
    def __init__(self):
        self.input = [[],[],[]]
        self.stack = {}
        self.target = []

    def printClean(self):
        for row in self.input:
            print(row)

    def inputClean(self):
        print("입력: ")
        for i in range(3):
            self.input[i] = list(map(int, input().split(" ")))

    def findBlank(self):
        for i in range(3):
            for j in range(3):
                if self.input[i][j] == 0:
                    return (i, j)

    def change(self, inVal):
        a, b = self.findBlank()
        result = False
        if inVal == 0 and b-1 >= 0:
            self.input[a][b], self.input[a][b-1] = self.input[a][b-1], self.input[a][b]
            result = True
        elif inVal == 1 and b+1 < 3:
            self.input[a][b], self.input[a][b+1] = self.input[a][b+1], self.input[a][b]
            result = True
        elif inVal == 2 and a-1 >= 0:
            self.input[a][b], self.input[a-1][b] = self.input[a-1][b], self.input[a][b]
            result = True
        elif inVal == 3 and a+1 < 3:
            self.input[a][b], self.input[a+1][b] = self.input[a+1][b], self.input[a][b]
            result = True
        return result

    def historyCheck(self, inVal):
        #temp = copy.deepcopy(self.input) #이곳이 문제 추정
        temp = Puzzle()
        temp.input = copy.deepcopy(self.input)
        temp.change(inVal)
        for i, v in self.stack.items():
            if temp.input == v and i > 1:
                return False
        return True

    def explore(self, inVal):
        if inVal < 0:
            inVal = 9223372036854775807
        self.count = 0
        upperlevel = 0
        for i in range(inVal + 1):
            if self.input == self.target:
                print("eureka! change 횟수는", self.count, "입니다.")
                print("만들어진 노드 수는", i, "입니다.")
                return
            arrow = 0
            while arrow <= 3:
                #print("count:",self.count ,arrow, upperlevel, self.historyCheck(arrow))
                if self.historyCheck(arrow):
                    if self.change(arrow):
                        upperlevel = 0
                        self.count += 1
                        self.stack[self.count] = copy.deepcopy(self.input)
                        arrow = 10
                elif arrow == 3:
                    self.input = self.stack[self.count-upperlevel]
                    upperlevel += 1
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
