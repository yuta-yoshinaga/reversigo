// Package test リバーシに関する構造体をまとめるパッケージ。
package model

// Copyright Yuta Yoshinaga.

import (
	"testing"
)

// TestResJson1 コンストラクタのテスト
func TestResJson1(t *testing.T) {
	r := NewResJson()
	if r == nil {
		t.Errorf("NG")
	}
}

// TestResJson2 GetAuth()のテスト
func TestResJson2(t *testing.T) {
	r := NewResJson()
	if r.GetAuth() != "" {
		t.Errorf("NG")
	}
}

// TestResJson3 SetAuth()のテスト
func TestResJson3(t *testing.T) {
	r := NewResJson()
	r.SetAuth("SetAuth")
	if r.GetAuth() != "SetAuth" {
		t.Errorf("NG")
	}
}

// TestResJson4 GetCallbacks()のテスト
func TestResJson4(t *testing.T) {
	r := NewResJson()
	if r.GetCallbacks() == nil {
		t.Errorf("NG")
	}
}

// TestResJson5 SetCallbacks()のテスト
func TestResJson5(t *testing.T) {
	r := NewResJson()
	r.SetCallbacks(NewCallbacksJson())
	if r.GetCallbacks() == nil {
		t.Errorf("NG")
	}
}
