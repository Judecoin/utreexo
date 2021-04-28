package main

import "fmt"

func LookAheadResetSlice(
	allCBlocks []cBlock, maxResets []int, maxHold int) ([]int, []int) {
	currRemembers := make([][]int, len(maxResets))
	for i := 0; i < len(maxResets); i++ {
		currRemembers[i] = make([]int, maxHold)
	}
	totalRemembers := make([]int, len(maxResets))
	maxRemembers := make([]int, len(maxResets))
	prevSum := make([]int, len(maxResets))
	currSum := make([]int, len(maxResets))
	for i := 0; i < len(allCBlocks); i++ {
		if i%100 == 0 {
			fmt.Println("On block: ", i)
		}
		for j := 0; j < len(maxResets); j++ {
			if i%maxResets[j] == 0 {
				prevSum[j] = 0
				currSum[j] = 0
				currRemembers[j] = make([]int, maxHold)
			}
		}
		numRemember := make([]int, len(maxResets))
		for _, ttl := range allCBlocks[i].ttls {
			for j := 0; j < len(maxResets); j++ {
				if ttl <= int32(maxHold) &&
					int32(maxResets[j]-(i%maxResets[j])) >= ttl {
					numRemember[j] += 1
				}
			}
		}
		for j := 0; j < len(maxResets); j++ {
			if i < maxHold {
				currRemembers[j][i] = numRemember[j]
				currSum[j] = prevSum[j] + numRemember[j]
				prevSum[j] = currSum[j]
			} else {
				currRemembers[j] = append(currRemembers[j], numRemember[j])
				currSum[j] = prevSum[j] + numRemember[j] - currRemembers[j][0]
				currRemembers[j] = currRemembers[j][1:]
				prevSum[j] = currSum[j]
			}
			if currSum[j] > maxRemembers[j] {
				maxRemembers[j] = currSum[j]
			}
			totalRemembers[j] += numRemember[j]
		}
	}
	return totalRemembers, maxRemembers
}

func LookAheadSlice(allCBlocks []cBlock, maxHolds []int) ([]int, []int) {
	currRemembers := make([][]int, len(maxHolds))
	for i := 0; i < len(maxHolds); i++ {
		currRemembers[i] = make([]int, maxHolds[i])
	}
	totalRemembers := make([]int, len(maxHolds))
	maxRemembers := make([]int, len(maxHolds))
	prevSum := make([]int, len(maxHolds))
	currSum := make([]int, len(maxHolds))
	for i := 0; i < len(allCBlocks); i++ {
		if i%100 == 0 {
			fmt.Println("On block: ", i)
		}
		numRemember := make([]int, len(maxHolds))
		for _, ttl := range allCBlocks[i].ttls {
			for k := 0; k < len(maxHolds); k++ {
				if ttl <= int32(maxHolds[k]) {
					numRemember[k] += 1
				}
			}
		}
		for j := 0; j < len(maxHolds); j++ {
			if i < maxHolds[j] {
				currRemembers[j][i] = numRemember[j]
				currSum[j] = prevSum[j] + numRemember[j]
				prevSum[j] = currSum[j]
			} else {
				currRemembers[j] = append(currRemembers[j], numRemember[j])
				currSum[j] = prevSum[j] + numRemember[j] - currRemembers[j][0]
				currRemembers[j] = currRemembers[j][1:]
				prevSum[j] = currSum[j]
			}
			if currSum[j] > maxRemembers[j] {
				maxRemembers[j] = currSum[j]
			}
			totalRemembers[j] += numRemember[j]
		}
	}
	return totalRemembers, maxRemembers
}

func LookAhead(allCBlocks []cBlock, maxHold int) (int, int, [][]string) {
	currRemembers := make([]int, maxHold)
	totalRemembers := 0
	maxRemembers := 0
	prevSum := 0
	currSumStores := make([][]string, len(allCBlocks))
	for i := 0; i < len(allCBlocks); i++ {
		currSumStores[i] = make([]string, 2)
		currSumStores[i][0] = fmt.Sprint(i)
		if i%100 == 0 {
			fmt.Println("On block: ", i)
		}
		numRemember := 0
		for _, ttl := range allCBlocks[i].ttls {
			if ttl <= int32(maxHold) {
				numRemember += 1
			}
		}
		var currSum int
		if i < maxHold {
			currRemembers[i] = numRemember
			currSum = prevSum + numRemember
			prevSum = currSum
		} else {
			currRemembers = append(currRemembers, numRemember)
			currSum = prevSum + numRemember - currRemembers[0]
			currRemembers = currRemembers[1:]
			prevSum = currSum
		}
		currSumStores[i][1] = fmt.Sprint(currSum)
		if currSum > maxRemembers {
			maxRemembers = currSum
		}
		totalRemembers += numRemember
	}
	fmt.Println("total number of remembers for gen10: ", totalRemembers)
	fmt.Println("max number of remembers for gen10: ", maxRemembers)
	return totalRemembers, maxRemembers, currSumStores
}
