// Package model リバーシに関する構造体をまとめるパッケージ。
package model

// Copyright Yuta Yoshinaga.

import (
	"testing"
)

// TestFuncsJson1 コンストラクタのテスト
func TestFuncsJson1(t *testing.T) {
	f := NewFuncsJson()
	if f == nil {
		t.Errorf("NG")
	}
}

// TestFuncsJson2 GetFunction()のテスト
func TestFuncsJson2(t *testing.T) {
	f := NewFuncsJson()
	if f.GetFunction() != "" {
		t.Errorf("NG")
	}
}

// TestFuncsJson3 SetFunction()のテスト
func TestFuncsJson3(t *testing.T) {
	f := NewFuncsJson()
	f.SetFunction("SetFunction")
	if f.GetFunction() != "SetFunction" {
		t.Errorf("NG")
	}
}

// TestFuncsJson4 GetParam1()のテスト
func TestFuncsJson4(t *testing.T) {
	f := NewFuncsJson()
	if f.GetParam1() != "" {
		t.Errorf("NG")
	}
}

// TestFuncsJson5 SetParam1()のテスト
func TestFuncsJson5(t *testing.T) {
	f := NewFuncsJson()
	f.SetParam1("SetParam1")
	if f.GetParam1() != "SetParam1" {
		t.Errorf("NG")
	}
}

// TestFuncsJson6 GetParam2()のテスト
func TestFuncsJson6(t *testing.T) {
	f := NewFuncsJson()
	if f.GetParam2() != "" {
		t.Errorf("NG")
	}
}

// TestFuncsJson7 SetParam2()のテスト
func TestFuncsJson7(t *testing.T) {
	f := NewFuncsJson()
	f.SetParam2("SetParam2")
	if f.GetParam2() != "SetParam2" {
		t.Errorf("NG")
	}
}

// TestFuncsJson8 GetParam3()のテスト
func TestFuncsJson8(t *testing.T) {
	f := NewFuncsJson()
	if f.GetParam3() != "" {
		t.Errorf("NG")
	}
}

// TestFuncsJson9 SetParam3()のテスト
func TestFuncsJson9(t *testing.T) {
	f := NewFuncsJson()
	f.SetParam3("SetParam3")
	if f.GetParam3() != "SetParam3" {
		t.Errorf("NG")
	}
}

// TestFuncsJson10 GetParam4()のテスト
func TestFuncsJson10(t *testing.T) {
	f := NewFuncsJson()
	if f.GetParam4() != "" {
		t.Errorf("NG")
	}
}

// TestFuncsJson11 SetParam4()のテスト
func TestFuncsJson11(t *testing.T) {
	f := NewFuncsJson()
	f.SetParam4("SetParam4")
	if f.GetParam4() != "SetParam4" {
		t.Errorf("NG")
	}
}

// TestFuncsJson12 GetParam5()のテスト
func TestFuncsJson12(t *testing.T) {
	f := NewFuncsJson()
	if f.GetParam5() != "" {
		t.Errorf("NG")
	}
}

// TestFuncsJson13 SetParam5()のテスト
func TestFuncsJson13(t *testing.T) {
	f := NewFuncsJson()
	f.SetParam5("SetParam5")
	if f.GetParam5() != "SetParam5" {
		t.Errorf("NG")
	}
}
