package stack

type ArrayStack struct {
	size  int
	store []interface{}
}

type IStack interface {
	isEmpty() bool
	push(item interface{})
	pop() interface{}
}

func (stack *ArrayStack) isEmpty() bool {
	return stack.size == 0
}

func (stack *ArrayStack) push(item interface{}) {
	stack.store[stack.size] = item
	stack.size++
}

func (stack *ArrayStack) pop() interface{} {
	stack.size--
	result := stack.store[stack.size]
	stack.store[stack.size] = nil
	return result
}

// ValidBrackets 检查括号是否完整
func ValidBrackets(s string) bool {
	l := len(s)
	arr := make([]interface{}, 10)
	stack := &ArrayStack{0, arr}
	brackets := map[string]string{")": "(", "]": "[", "}": "{"}
	for i := 0; i < l; i++ {
		bracket := string(s[i])
		if bracket == ")" || bracket == "}" || bracket == "]" {
			tmp, _ := stack.pop().(string)
			if tmp == brackets[bracket] {
				continue
			} else {
				return false
			}
		}
		if bracket == "(" || bracket == "{" || bracket == "[" {
			stack.push(bracket)
		}
	}
	return true
}

// CompleteBrackets 补充括号  1+2)*3-4)*5-6))) => ((1+2)*((3-4)*(5-6)))
func CompleteBrackets(s string) string {
	l := len(s)
	opsStore := make([]interface{}, l)
	valStore := make([]interface{}, l)
	ops := &ArrayStack{0, opsStore}
	val := &ArrayStack{0, valStore}
	for i := range s {
		tmp := string(s[i])
		if tmp == "+" || tmp == "-" || tmp == "*" || tmp == "/" {
			ops.push(tmp)
		} else if tmp == ")" {
			val1, _ := val.pop().(string)
			val2, _ := val.pop().(string)
			op, _ := ops.pop().(string)
			val.push("(" + val2 + op + val1 + ")")
		} else {
			val.push(tmp)
		}
	}
	result, _ := val.pop().(string)
	return result
}
