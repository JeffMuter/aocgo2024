package daythree

import (
	"fmt"
	"testing"
)

func TestCreateResult(t *testing.T) {
	tests := []struct {
		name      string
		textInput string
		expected  int
	}{
		{"normal good result 0", "mul(2134,8722)", 18612748},
		{"random stuff. bad result", "edwdewgytjyu765r4", -1},
		{"multiple good input", "mul(11,3)mul(12,1)", 45},
		{"nestedgood value. good result", "mul(12,mul(13mul(23,2),4),3)", 46},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateResult(tt.textInput)
			if err != nil {
				fmt.Println("test returned err...")
			}
			if got != tt.expected {
				t.Errorf("Result tested:%s\ninput: %s\nexpected: %d\ngot: %d\nerr: %v\n\n", tt.name, tt.textInput, tt.expected, got, err)
			}
		})
	}

}
