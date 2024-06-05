package string

/*
输入: message = "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

*/

// reverseMessage 维护一个单词是否记录完的done，考虑四种情况
//  @param message
//  @return string
func ReverseMessage(message string) string {
	done := false
	word := ""
	reMeg := ""

	for _, char := range message {
		//正常情况
		if !done && char != ' ' {
			word = word + string(char)
			continue
		}

		//单词记录到一半遇到空格
		if !done && char == ' ' {
			//！！！也可能是刚开始
			if word == "" {
				continue
			}
			reMeg = word + " " + reMeg
			done = true
			continue
		}

		//单词未开始记录时遇到空格
		if done && char != ' ' {
			word = string(char)
			done = false
			continue
		}
	}

	//！！！最后如果没记录完，说明有最后一个单词
	if !done {
		reMeg = word + " " + reMeg
	}

	return reMeg[:len(reMeg)-1]
}
