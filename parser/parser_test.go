package parser

import (
	"fmt"
	"testing"
)

func TestCreateModule(t *testing.T) {
	postMod := Module{
		name: "post",
		fields: []ModuleField{
			ModuleField{name: "title", dataType: "string"},
			ModuleField{name: "content", dataType: "string"},
			ModuleField{name: "likes", dataType: "uint"},
			ModuleField{name: "dislikes", dataType: "uint*"},
		},
	}

	var tests = []struct {
		cmd  string
		want Module
	}{
		{"post title:string content:string likes:uint dislikes:uint*", postMod},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.cmd)
		t.Run(testName, func(t *testing.T) {
			ans := CreateModule(tt.cmd)
			if ans.name != tt.want.name {
				t.Errorf("got name %s, want name %s", ans, tt.want)
			}

			if len(ans.fields) != len(tt.want.fields) {
				t.Errorf("got fields len %s, want fields len %s", ans, tt.want)
			}

			for i, _ := range ans.fields {
				if ans.fields[i].name != tt.want.fields[i].name {
					t.Errorf("got field name %s at %d, want field name %s", ans, i, tt.want)
				}
			}

		})
	}
}
