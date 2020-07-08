// Package test リバーシに関する構造体をまとめるパッケージ。
package model

// Copyright Yuta Yoshinaga.

import (
	"testing"
)

// TestCallbacksJson1 コンストラクタのテスト
func TestCallbacksJson1(t *testing.T) {
	c := NewCallbacksJson()
	if c == nil {
		t.Errorf("NG")
	}
}

// TestCallbacksJson2 GetFuncs()のテスト
func TestCallbacksJson2(t *testing.T) {
	c := NewCallbacksJson()
	if c.GetFuncs() == nil {
		t.Errorf("NG")
	}
}

// TestCallbacksJson3 SetFuncs()のテスト
func TestCallbacksJson3(t *testing.T) {
	c := NewCallbacksJson()
	c.SetFuncs(make([]*FuncsJson, 0))
	if c.GetFuncs() == nil {
		t.Errorf("NG")
	}
}
