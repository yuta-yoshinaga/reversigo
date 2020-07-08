// Package test リバーシに関する構造体をまとめるパッケージ。
package model

// Copyright Yuta Yoshinaga.

import (
	"testing"
)

// TestReversi1 コンストラクタのテスト
func TestReversi1(t *testing.T) {
	r := NewReversi(8, 8)
	if r == nil {
		t.Errorf("NG")
	}
}
