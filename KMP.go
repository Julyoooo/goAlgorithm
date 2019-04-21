package main

import "fmt"

func main() {
	fmt.Println(Kmp("aaaa", "aaabaaabaaabaaaa"))
}

func Kmp(m, s string) int{
	lengthS := len(s)
	lengthM := len(m)
	index := -1
	table := genFailTable(m)

	i := 0
	j := 0

	for j < lengthM {
		if s[i] != m[j] {
			if j == 0 {
				i += 1
				continue
			}
			j = table[j-1] + 1 // 取j-1的最长匹配子串的下标，再加1表示下次比较的起点
			continue
		}

		i += 1
		j += 1

		if j == lengthM {
			index = i-lengthM
			break
		}

		if i == lengthS {
			break
		}
	}

	fmt.Println(table)
	return index
}

// 失效函数：根据模式串生成的数组，用于到匹配失效时候的对模式串滑动，下标是前缀长度，值是匹配前缀的前缀子串的结尾字符下标
// 0. 模式串 abcabcy
// 1. 好前缀 abcabc
// 2. 好前缀的前缀子串 a/ab/abc/abca/abcab
// 3. 好前缀的后缀子串 c/bc/abc/cabc/bcabc
// 4. 与好前缀的前缀子串匹配的好前缀的后缀子串中最长的 abc
// 5. 与好前缀的前缀子串匹配的最长后缀子串的结尾字符下标 table[5] = 2
func genFailTable(m string) []int{
	length := len(m) // 模式串的长度
	table := make([]int, length)
	for i:=0; i<length; i++ {
		table[i] = -1 // 初始化
	}

	k := -1 // 默认不匹配，下标为 -1，相当于 table[0] = k = -1

	for i:=1; i<length; i++ {
		for m[k+1] != m[i] && k != -1 {
			k = table[k]
		}

		if m[k+1] == m[i] {
			k = k + 1
		}

		table[i] = k
	}

	return table
}

// 下面是逻辑更清晰的版本
func genFailTableV2(m string) []int{
	length := len(m) // 模式串的长度
	table := make([]int, length)
	for i:=0; i<length; i++ {
		table[i] = -1 // 初始化
	}

	k := -1 // 默认不匹配，下标为 -1，相当于 table[0] = k = -1

	for i:=1; i<length; i++ {
		// 第一种情况：
		// 1. 已知 table[i-1] = k，既 m[0...k] 是 m[0...i-1] 的最长前缀子串
		// 2. 如果 m[k+1] == m[i]，则 m[0...k+1] 是 m[0...i] 的最长前缀子串
		// 3. 所以 table[i] = k+1
		if m[k+1] == m[i] {
			k = k + 1
			table[i] = k
			continue
		}

		// 第二种情况：
		// 1. 已知 table[i-1] = k，既 m[0...k] 是 m[0...i-1] 的最长前缀子串
		// 2. 如果 m[k+1] != m[i]，则 m[0...i] 的最长前缀子串在 m[0...i-1] 的次长（第二长、第三长依次类推）前缀子串中
		// 3. m[0...i-1] 的次长前缀子串在 m[0...k] 内，且为 m[0...k] 的最长前缀子串，既 m[0...table[k]]，代码 k = table[k]
		// 4. 比较 m[table[k]+1] == m[i]，相等则找到最长子串 m[0...table[k]+1]
		// 5. 不相等则继续重复步骤 3，直到 k == -1
		// 6. 因此循环结束后要判断两种情况，如果 m[table[k]+1] == m[i]，则 k = table[k] + 1，既代码中的 k = k + 1
		// 7. 如果 k == -1 则直接赋值 table[i] = -1
		for m[k+1] != m[i] && k != -1 {
			k = table[k]
		}

		if m[k+1] == m[i] {
			k = k + 1
			table[i] = k
			continue
		}

		if k == -1 {
			table[i] = k
		}
	}

	return table
}