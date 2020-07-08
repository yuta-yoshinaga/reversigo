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

// TestReversiAnz2 GetmMasuSts()のテスト
func TestReversi2(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuSts() == nil {
		t.Errorf("NG")
	}
}

// TestReversiAnz3 SetmMasuSts()のテスト
func TestReversi3(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuSts(make([][]int, 8))
	if r.GetmMasuSts() == nil {
		t.Errorf("NG")
	}
}

// TestReversi4 GetmMasuStsOld()のテスト
func TestReversi4(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuStsOld() == nil {
		t.Errorf("NG")
	}
}

// TestReversi5 SetmMasuStsOld()のテスト
func TestReversi5(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuStsOld(make([][]int, 8))
	if r.GetmMasuStsOld() == nil {
		t.Errorf("NG")
	}
}

// TestReversi6 GetmMasuStsEnaB()のテスト
func TestReversi6(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuStsEnaB() == nil {
		t.Errorf("NG")
	}
}

// TestReversi7 SetmMasuStsEnaB()のテスト
func TestReversi7(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuStsEnaB(make([][]int, 8))
	if r.GetmMasuStsEnaB() == nil {
		t.Errorf("NG")
	}
}

// TestReversi8 GetmMasuStsCntB()のテスト
func TestReversi8(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuStsCntB() == nil {
		t.Errorf("NG")
	}
}

// TestReversi9 SetmMasuStsCntB()のテスト
func TestReversi9(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuStsCntB(make([][]int, 8))
	if r.GetmMasuStsCntB() == nil {
		t.Errorf("NG")
	}
}

// TestReversi10 GetmMasuStsPassB()のテスト
func TestReversi10(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuStsPassB() == nil {
		t.Errorf("NG")
	}
}

// TestReversi11 SetmMasuStsPassB()のテスト
func TestReversi11(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuStsPassB(make([][]int, 8))
	if r.GetmMasuStsPassB() == nil {
		t.Errorf("NG")
	}
}

// TestReversi12 GetmMasuStsAnzB()のテスト
func TestReversi12(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuStsAnzB() == nil {
		t.Errorf("NG")
	}
}

// TestReversi13 SetmMasuStsAnzB()のテスト
func TestReversi13(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuStsAnzB(make([][]*ReversiAnz, 8))
	if r.GetmMasuStsAnzB() == nil {
		t.Errorf("NG")
	}
}

// TestReversi14 GetmMasuPointB()のテスト
func TestReversi14(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuPointB() == nil {
		t.Errorf("NG")
	}
}

// TestReversi15 SetmMasuPointB()のテスト
func TestReversi15(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuPointB(make([]*ReversiPoint, 8))
	if r.GetmMasuPointB() == nil {
		t.Errorf("NG")
	}
}

// TestReversi16 GetmMasuPointCntB()のテスト
func TestReversi16(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuPointCntB() != 4 {
		t.Errorf("NG")
	}
}

// TestReversi17 SetmMasuPointB()のテスト
func TestReversi17(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuPointCntB(1)
	if r.GetmMasuPointCntB() != 1 {
		t.Errorf("NG")
	}
}

// TestReversi18 GetmMasuBetCntB()のテスト
func TestReversi18(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuBetCntB() != 2 {
		t.Errorf("NG")
	}
}

// TestReversi19 SetmMasuBetCntB()のテスト
func TestReversi19(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuBetCntB(1)
	if r.GetmMasuBetCntB() != 1 {
		t.Errorf("NG")
	}
}

// TestReversi20 GetmMasuStsEnaW()のテスト
func TestReversi20(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuStsEnaW() == nil {
		t.Errorf("NG")
	}
}

// TestReversi21 SetmMasuStsEnaW()のテスト
func TestReversi21(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuStsEnaW(make([][]int, 8))
	if r.GetmMasuStsEnaW() == nil {
		t.Errorf("NG")
	}
}

// TestReversi22 GetmMasuStsCntW()のテスト
func TestReversi22(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuStsCntW() == nil {
		t.Errorf("NG")
	}
}

// TestReversi23 SetmMasuStsCntW()のテスト
func TestReversi23(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuStsCntW(make([][]int, 8))
	if r.GetmMasuStsCntW() == nil {
		t.Errorf("NG")
	}
}

// TestReversi24 GetmMasuStsPassW()のテスト
func TestReversi24(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuStsPassW() == nil {
		t.Errorf("NG")
	}
}

// TestReversi25 SetmMasuStsPassW()のテスト
func TestReversi25(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuStsPassW(make([][]int, 8))
	if r.GetmMasuStsPassW() == nil {
		t.Errorf("NG")
	}
}

// TestReversi26 GetmMasuStsAnzW()のテスト
func TestReversi26(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuStsAnzW() == nil {
		t.Errorf("NG")
	}
}

// TestReversi27 SetmMasuStsAnzW()のテスト
func TestReversi27(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuStsAnzW(make([][]*ReversiAnz, 8))
	if r.GetmMasuStsAnzW() == nil {
		t.Errorf("NG")
	}
}

// TestReversi28 GetmMasuPointW()のテスト
func TestReversi28(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuPointW() == nil {
		t.Errorf("NG")
	}
}

// TestReversi29 SetmMasuStsAnzW()のテスト
func TestReversi29(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuPointW(make([]*ReversiPoint, 8))
	if r.GetmMasuPointW() == nil {
		t.Errorf("NG")
	}
}

// TestReversi30 GetmMasuPointCntW()のテスト
func TestReversi30(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuPointCntW() != 4 {
		t.Errorf("NG")
	}
}

// TestReversi31 SetmMasuStsAnzW()のテスト
func TestReversi31(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuPointCntW(1)
	if r.GetmMasuPointCntW() != 1 {
		t.Errorf("NG")
	}
}

// TestReversi32 GetmMasuBetCntW()のテスト
func TestReversi32(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuBetCntW() != 2 {
		t.Errorf("NG")
	}
}

// TestReversi33 SetmMasuBetCntW()のテスト
func TestReversi33(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuBetCntW(1)
	if r.GetmMasuBetCntW() != 1 {
		t.Errorf("NG")
	}
}

// TestReversi34 GetmMasuCnt()のテスト
func TestReversi34(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuCnt() != 8 {
		t.Errorf("NG")
	}
}

// TestReversi35 SetmMasuCnt()のテスト
func TestReversi35(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuCnt(1)
	if r.GetmMasuCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversi36 GetmMasuCntMax()のテスト
func TestReversi36(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuCntMax() != 8 {
		t.Errorf("NG")
	}
}

// TestReversi37 SetmMasuCntMax()のテスト
func TestReversi37(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuCntMax(1)
	if r.GetmMasuCntMax() != 1 {
		t.Errorf("NG")
	}
}

// TestReversi38 GetmMasuCntMax()のテスト
func TestReversi38(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuHist() == nil {
		t.Errorf("NG")
	}
}

// TestReversi39 SetmMasuHist()のテスト
func TestReversi39(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuHist(make([]*ReversiHistory, 0))
	if r.GetmMasuHist() == nil {
		t.Errorf("NG")
	}
}

// TestReversi40 GetmMasuHistCur()のテスト
func TestReversi40(t *testing.T) {
	r := NewReversi(8, 8)
	if r.GetmMasuHistCur() != 0 {
		t.Errorf("NG")
	}
}

// TestReversi41 SetmMasuHistCur()のテスト
func TestReversi41(t *testing.T) {
	r := NewReversi(8, 8)
	r.SetmMasuHistCur(1)
	if r.GetmMasuHistCur() != 1 {
		t.Errorf("NG")
	}
}
