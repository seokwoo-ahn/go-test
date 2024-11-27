package algorithm

import (
	"fmt"
	"math"
	"math/rand"
)

// 거리 계산 함수 (맨해튼 거리)
func calcDistance(pos1, pos2 [2]int) int {
	return int(math.Abs(float64(pos1[0]-pos2[0])) + math.Abs(float64(pos1[1]-pos2[1])))
}

// 자리 배치 함수
func assignSeats(people []string, friendsMap map[string][]string) map[string][2]int {
	rand.Shuffle(len(people), func(i, j int) { people[i], people[j] = people[j], people[i] })
	// 좌석 배치 결과를 저장
	seating := make(map[string][2]int)
	positions := [][2]int{}

	// 테이블 크기 설정 (열 개수)
	tableSize := (len(people) + 1) / 2

	// 테이블 크기에 따라 가능한 좌석 좌표 생성
	for r := 1; r <= 2; r++ {
		for c := 1; c <= tableSize; c++ {
			positions = append(positions, [2]int{r, c})
		}
	}

	// 배치를 위한 초기 상태 설정
	used := make(map[[2]int]bool)

	// 각 사람을 순서대로 배치
	for _, person := range people {
		var bestPos [2]int
		maxDistance := -1

		// 가능한 자리 중 최적의 위치 탐색
		for _, pos := range positions {
			if used[pos] {
				continue
			}

			// 친구가 있다면 친구와의 거리 계산
			minFriendDistance := math.MaxInt32
			for _, friend := range friendsMap[person] {
				if friendPos, ok := seating[friend]; ok {
					distance := calcDistance(pos, friendPos)
					if distance < minFriendDistance {
						minFriendDistance = distance
					}
				}
			}

			// 친구가 없으면 모든 자리 가능
			if len(friendsMap[person]) == 0 {
				minFriendDistance = math.MaxInt32
			}

			// 최소 거리 최대화
			if minFriendDistance > maxDistance {
				maxDistance = minFriendDistance
				bestPos = pos
			}
		}

		// 최적 위치 배정
		seating[person] = bestPos
		used[bestPos] = true
	}

	return seating
}

// 배치도 출력 함수
func printSeatingArrangement(seating map[string][2]int) {
	// 테이블 크기 설정 (열 개수)
	tableSize := (len(seating) + 1) / 2

	// 테이블을 행렬 형태로 초기화
	table := [2][]string{
		make([]string, tableSize),
		make([]string, tableSize),
	}

	// 배치된 사람들 이름을 테이블에 넣기
	for person, pos := range seating {
		row := pos[0] - 1
		col := pos[1] - 1
		table[row][col] = person
	}

	// empty seat을 가장 끝으로 보내는 로직
	for r := 0; r < 2; r++ {
		for c := 0; c < tableSize; c++ {
			if table[r][c] == "" {
				table[r][c] = "empty"
				for c2 := c + 1; c2 < tableSize; c2++ {
					if table[r][c2] != "" {
						table[r][c], table[r][c2] = table[r][c2], table[r][c]
						break
					}
				}
			}
		}
	}

	// 테이블 출력
	for _, row := range table {
		for _, seat := range row {
			fmt.Printf("%s ", seat)
		}
		fmt.Println()
	}
}
