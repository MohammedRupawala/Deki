package trie

import "math"


func checkPrefixTillIndex(idx int, matches *[]string) bool{
	prev := (*matches)[0][:idx]
	for _,v := range *matches{
		if prev != v[:idx]{
			return false
		}
	}

	return true
}
func (t *Trie) FindLongestMatch(matches []string) string{
	if(len(matches) == 0){
		return ""
	}
	var left = 0
	var right = math.MaxInt
	right --;
	var ans  = ""
	for _, v := range matches {
		right = min(right, len(v))
	}

	for (left <= right){
		var mid = left + ((right - left) >> 1)

		if(checkPrefixTillIndex(mid,&matches)){
			left = mid + 1;
			ans = matches[0][:mid]
		}else{
			right = mid - 1;
		}
	}

	return ans;
}