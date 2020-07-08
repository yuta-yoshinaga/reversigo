// Package test リバーシに関する構造体をまとめるパッケージ。
package model

// Copyright Yuta Yoshinaga.

import (
	"testing"
)

// TestReversiHistory1 コンストラクタのテスト
func TestReversiHistory1(t *testing.T) {
	r := NewReversiHistory()
	if r == nil {
		t.Errorf("NG")
	}
}

// TestReversiHistory2 GetPoint()のテスト
func TestReversiHistory2(t *testing.T) {
	r := NewReversiHistory()
	if r.GetPoint() == nil {
		t.Errorf("NG")
	}
}

// TestReversiHistory3 SetPoint()のテスト
func NewReversiHistory3(t *testing.T) {
	r := NewReversiHistory()
	r.SetPoint(NewReversiPoint())
	if r.GetPoint() == nil {
		t.Errorf("NG")
	}
}

// TestReversiHistory4 GetColor()のテスト
func TestReversiHistory4(t *testing.T) {
	r := NewReversiHistory()
	if r.GetColor() != -1 {
		t.Errorf("NG")
	}
}

// NewReversiHistory5 SetColor()のテスト
func NewReversiHistory5(t *testing.T) {
	r := NewReversiHistory()
	r.SetColor(1)
	if r.GetColor() != 1 {
		t.Errorf("NG")
	}
}
