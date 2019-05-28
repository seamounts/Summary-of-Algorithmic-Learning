package lengthOfLongestSubstring

func LengthOfLongestSubstring(s string) int {
	var (
		maxSubStr    int
		newMaxSubStr int
	)
	hashTable := make(map[byte]int)
	for i, _ := range s {
		_, ok := hashTable[s[i]]
		if ok {
			index := hashTable[s[i]]
			newMaxSubStr = LengthOfLongestSubstring(s[index+1 : len(s)])
			break
		}
		hashTable[s[i]] = i
		maxSubStr += 1
	}
	if maxSubStr > newMaxSubStr {
		return maxSubStr
	} else {
		return newMaxSubStr
	}

}

func LengthOfLongestSubstring2(s string) int {
	var (
		maxSubStr   int
		subStrCount int
		index       int
		flag        int
	)

	hashTable := make(map[byte]int)
	for {
		if index >= len(s) {
			break
		}

		i, ok := hashTable[s[index]]
		if ok && i >= flag {
			flag = i
			delete(hashTable, s[index])

			if maxSubStr < subStrCount {
				maxSubStr = subStrCount
			}
			if maxSubStr >= len(s)-i-1 {
				break
			}
			subStrCount = index - i - 1
			continue
		}

		hashTable[s[index]] = index
		subStrCount += 1
		index += 1

	}
	if maxSubStr < subStrCount {
		maxSubStr = subStrCount
	}

	return maxSubStr
}
