package httprouter

func CleanPath(path string) string {
	if path == "" {
		return "/"
	}
	//init pointer
	n := len(path)
	r := 0
	w := 1
	buf := make([]byte, n+1)
	buf[0] = '/'
	if path[0] == '/' {
		r++
	}

	for r < n {
		switch path[r] {
		case '.':
			// 返回上级目录
			if r+1 < n && path[r+1] == '.' {
				w = goback(buf, w)
				r++
			}
		case '/':
			// 过滤多余的
			if buf[w-1] != '/' {
				buf[w] = '/'
				w++
			}
		default:
			buf[w] = path[r]
			w++
		}
		r++
	}
	return string(buf[:w])
}

func goback(buf []byte, w int) int {
	var i int
	if buf[w-1] == '/' {
		i = w - 2
	} else {
		i = w - 1
	}
	for ; i >= 1; i-- {
		if buf[i] == '/' {
			return i
		}
	}
	return 1
}
