package models

import (
		"testing"
)

func TestGetArticle(t *testing.T)  {
	a := GetArticle(1)
	if a.TagID == 0 {
		t.Fatal("错误")
	}
	t.Logf("成功")

}