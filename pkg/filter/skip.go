package filter

import "sort"

// the char in the sortedSkipList will be skipped
const sortedSkipList = "\n\r!\"#$%&'()*+-:;=@[]^_{|}~¤§¨°±·×÷ˉΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩαβγδεζηθικλμνξοπρστυφχψω—―‖‘’“”…‰※€℃№ⅠⅡⅢⅣⅤⅥⅦⅧⅨⅩⅪⅫ←↑→↓∈∏∑√∝∞∠∥∧∨∩∪∫∮∴∵∶∷∽≈≌≠≡≤≥≮≯⊙⊥⌒①②③④⑤⑥⑦⑧⑨⑩⑴⑵⑶⑷⑸⑹⑺⑻⑼⑽⑾⑿⒀⒁⒂⒃⒄⒅⒆⒇⒈⒉⒊⒋⒌⒍⒎⒏⒐⒑⒒⒓⒔⒕⒖⒗⒘⒙⒚⒛─━│┃┄┅┆┇┈┉┊┋┌┍┎┐┑┒┓└┕┖┗┘┙┚┛├┝┞┟┠┡┢┣┤┥┦┧┨┩┪┫┬┭┮┯┰┱┲┳┴┵┶┷┸┹┺┻┼┽┾┿╀╁╂╃╄╅╆╇╈╉╊╋■□▲△◆◇○◎●★☆♀♂、、、。。〃々〈〉《《》》「」『』【【】】〓〔〕〖〗㈠㈡㈢㈣㈤㈥㈦㈧㈨㈩︿！＂＃＆＇（）＋，，－．／：；＜＝＞？？＠［＼］＿｀｛｜｝～￣"

func SortedSkipList() string {
	return sortedSkipList
}

type Skip struct {
	list []rune
}

// Set sorted skip list
func (s *Skip) Set(word string) {
	list := []rune(word)
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	s.list = list
}

func (s *Skip) ShouldSkip(r rune) bool {
	left, right := 0, len(s.list)
	if right == 0 {
		return false
	}
	if r < s.list[0] || r > s.list[right-1] {
		return false
	}
	for left < right {

		mid := left + (right-left)>>1
		if s.list[mid] == r {
			return true
		} else if s.list[mid] > r {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return false
}
