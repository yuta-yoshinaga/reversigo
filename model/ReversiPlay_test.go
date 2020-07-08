// Package test リバーシに関する構造体をまとめるパッケージ。
package model

// Copyright Yuta Yoshinaga.

import (
	"testing"
)

// TestReversiPlay1 コンストラクタのテスト
func TestReversiPlay1(t *testing.T) {
	r := NewReversiPlay()
	if r == nil {
		t.Errorf("NG")
	}
}
