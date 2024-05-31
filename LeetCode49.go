package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		bs := []byte(strs[i])
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		s := string(bs)
		//_, ok := mp[s]
		//if ok {
		mp[s] = append(mp[s], strs[i])
		//}else{
		//	mp
		//}
	}

	res := make([][]string, 0)
	for _, value := range mp {
		res = append(res, value)
	}
	return res
}
