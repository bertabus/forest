package rat

import (
	"strings"
	"testing"
)

func TestJSONPathRoot(t *testing.T) {
	src := map[string]interface{}{
		"key": "value",
		"nested": map[string]interface{}{
			"sub": "super",
		},
		"kids": []interface{}{
			map[string]interface{}{"name": "dennis"},
			map[string]interface{}{"name": "lisa"},
		},
	}
	found := pathFindIn(0, strings.Split(".", ".")[1:], src)
	if found == nil {
		t.Errorf("expected root document, got:%v", found)
	}
	found = pathFindIn(0, strings.Split(".key", ".")[1:], src)
	if found != "value" {
		t.Errorf("expected value, got:%v", found)
	}
	found = pathFindIn(0, strings.Split(".nested.sub", ".")[1:], src)
	if found != "super" {
		t.Errorf("expected super, got:%v", found)
	}
	found = pathFindIn(0, strings.Split(".kids.1.name", ".")[1:], src)
	if found != "lisa" {
		t.Errorf("expected lisa, got:%v", found)
	}

}