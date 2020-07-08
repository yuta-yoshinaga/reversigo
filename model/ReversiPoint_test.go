// Package test リバーシに関する構造体をまとめるパッケージ。
package model

// Copyright Yuta Yoshinaga.

import (
	"testing"
)

// TestReversiPoint1 コンストラクタのテスト
func TestReversiPoint1(t *testing.T) {
	r := NewReversiPoint()
	if r == nil {
		t.Errorf("NG")
	}
}

// TestReversiPoint2 GetX()のテスト
func TestReversiPoint2(t *testing.T) {
	r := NewReversiPoint()
	if r.GetX() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiPoint3 SetX()のテスト
func TestReversiPoint3(t *testing.T) {
	r := NewReversiPoint()
	r.SetX(1)
	if r.GetX() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiPoint4 GetY()のテスト
func TestReversiPoint4(t *testing.T) {
	r := NewReversiPoint()
	if r.GetY() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiPoint5 SetY()のテスト
func TestReversiPoint5(t *testing.T) {
	r := NewReversiPoint()
	r.SetY(1)
	if r.GetY() != 1 {
		t.Errorf("NG")
	}
}
