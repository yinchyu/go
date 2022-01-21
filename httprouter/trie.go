package httprouter

type trie struct {
	isend    bool
	childern [26]*trie
}

func newtrie() *trie {
	root := trie{childern: [26]*trie{}}
	return &root
}

func (t *trie) insert(a string) {
	for i := 0; i < len(a); i++ {
		//只考虑小写字母
		index := a[i] - 'a'
		if t.childern[index] == nil {
			t.childern[index] = newtrie()
		}
		t = t.childern[index]
	}
	t.isend = true
}

// 字典树和布隆过滤器一样只能添加不能删除
//func delete(t *trie)delete(){
//
//}

func (t trie) startwith(s string) bool {
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if t.childern[index] == nil {
			return false
		} else {
			t = *t.childern[index]
		}

	}
	return true
}

func (t trie) isword(s string) bool {
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if t.childern[index] == nil {
			return false
		} else {
			t = *t.childern[index]
		}

	}
	return t.isend
}
