package stack

import "testing"

func TestValidBrackets(t *testing.T) {
	tests := map[string]bool{"[()]{}{[()()]()}": true, "[(])": false, "[]": true, "[({})}": false}
	for key, val := range tests {
		if ans := ValidBrackets(key); ans != val {
			t.Errorf("param:%s expected:%v got:%v", key, val, ans)
		}
	}
}

func TestCompleteBrackets(t *testing.T) {
	if ans := CompleteBrackets("1+2)*3-4)*5-6)))"); ans != "((1+2)*((3-4)*(5-6)))" {
		t.Errorf("param:1+2)*3-4)*5-6))) expected:((1+2)*((3-4)*(5-6))) got:%v", ans)
	}
}
