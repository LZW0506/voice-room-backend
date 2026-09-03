package handler

import "testing"

// TestReadText 验证请求文本的清理和字符长度限制
func TestReadText(t *testing.T) {
	if got := readText("  声屿  ", 24); got != "声屿" {
		t.Fatalf("readText() = %q, want %q", got, "声屿")
	}
	if got := readText("一二三四", 3); got != "一二三" {
		t.Fatalf("readText() = %q, want %q", got, "一二三")
	}
}
