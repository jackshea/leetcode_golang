package src

// 二进制求和
// https://leetcode-cn.com/problems/add-binary/
func addBinary(a string, b string) string {
	var long []byte
	var short []byte
	if len(a) > len(b) {
		long = []byte(a)
		short = []byte(b)
	} else {
		long = []byte(b)
		short = []byte(a)
	}

	var add uint8
	for i := 0; i < len(long); i++ {
		lv := long[len(long)-1-i] - '0'
		var sv uint8
		if i < len(short) {
			sv = short[len(short)-1-i] - '0'
		}
		sum := lv + sv + add
		add = sum / 2
		long[len(long)-1-i] = '0' + sum%2
	}

	if add != 0 {
		long = append([]byte{'1'}, long...)
	}

	return string(long)
}
