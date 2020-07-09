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

// TestReversi42 Reset()のテスト
func TestReversi42(t *testing.T) {
	r := NewReversi(8, 8)
	r.Reset()
}

// TestReversi43 makeMasuSts()のテスト
func TestReversi43(t *testing.T) {
	r := NewReversi(8, 8)
	r.makeMasuSts(REVERSI_STS_BLACK)
}

// TestReversi44 revMasuSts()のテスト
func TestReversi44(t *testing.T) {
	r := NewReversi(8, 8)
	r.revMasuSts(REVERSI_STS_BLACK, 0, 0)
}

// TestReversi45 checkPara()のテスト
func TestReversi45(t *testing.T) {
	r := NewReversi(8, 8)
	if r.checkPara(0, 0, 8) != 0 || r.checkPara(-1, 0, 8) != -1 || r.checkPara(1, 0, 8) != 0 || r.checkPara(8, 0, 8) != 0 || r.checkPara(9, 0, 8) != -1 {
		t.Errorf("NG")
	}
}

// TestReversi46 analysisReversiBlack()のテスト
func TestReversi46(t *testing.T) {
	r := NewReversi(8, 8)
	r.analysisReversiBlack()
}

// TestReversi47 analysisReversiWhite()のテスト
func TestReversi47(t *testing.T) {
	r := NewReversi(8, 8)
	r.analysisReversiWhite()
}

// TestReversi48 analysisReversiWhite()のテスト
func TestReversi48(t *testing.T) {
	r := NewReversi(8, 8)
	r.AnalysisReversi(0, 0)
}

// TestReversi49 GetMasuSts()のテスト
func TestReversi49(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetMasuSts(0, 0)
	if sts != REVERSI_STS_NONE {
		t.Errorf("NG")
	}
}

// TestReversi50 GetMasuStsOld()のテスト
func TestReversi50(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetMasuStsOld(0, 0)
	if sts != REVERSI_STS_NONE {
		t.Errorf("NG")
	}
}

// TestReversi51 GetMasuStsEna()のテスト
func TestReversi51(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetMasuStsEna(REVERSI_STS_BLACK, 0, 0)
	if sts == 1 {
		t.Errorf("NG")
	}
}

// TestReversi52 GetMasuStsCnt()のテスト
func TestReversi52(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetMasuStsCnt(REVERSI_STS_BLACK, 0, 0)
	if sts != 0 {
		t.Errorf("NG")
	}
}

// TestReversi53 GetColorEna()のテスト
func TestReversi53(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetColorEna(REVERSI_STS_BLACK)
	if sts != 0 {
		t.Errorf("NG")
	}
}

// TestReversi54 GetGameEndSts()のテスト
func TestReversi54(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetGameEndSts()
	if sts != 0 {
		t.Errorf("NG")
	}
}

// TestReversi55 SetMasuSts()のテスト
func TestReversi55(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.SetMasuSts(REVERSI_STS_BLACK, 2, 4)
	if sts != 0 {
		t.Errorf("NG")
	}
}

// TestReversi56 SetMasuStsForcibly()のテスト
func TestReversi56(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.SetMasuStsForcibly(REVERSI_STS_BLACK, 2, 4)
	if sts != 0 {
		t.Errorf("NG")
	}
}

// TestReversi57 SetMasuCnt()のテスト
func TestReversi57(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.SetMasuCnt(6)
	if sts != 0 {
		t.Errorf("NG")
	}
}

// TestReversi58 GetPoint()のテスト
func TestReversi58(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetPoint(REVERSI_STS_BLACK, 0)
	if sts == nil {
		t.Errorf("NG")
	}
}

// TestReversi59 GetPointCnt()のテスト
func TestReversi59(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetPointCnt(REVERSI_STS_BLACK)
	if sts == 0 {
		t.Errorf("NG")
	}
}

// TestReversi60 GetBetCnt()のテスト
func TestReversi60(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetBetCnt(REVERSI_STS_BLACK)
	if sts == 0 {
		t.Errorf("NG")
	}
}

// TestReversi61 GetPassEna()のテスト
func TestReversi61(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetPassEna(REVERSI_STS_BLACK, 0, 0)
	if sts != 0 {
		t.Errorf("NG")
	}
}

// TestReversi62 GetHistory()のテスト
func TestReversi62(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetHistory(0)
	if sts == nil {
		t.Errorf("NG")
	}
}

// TestReversi63 GetHistoryCnt()のテスト
func TestReversi63(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetHistoryCnt()
	if sts != 0 {
		t.Errorf("NG")
	}
}

// TestReversi64 GetPointAnz()のテスト
func TestReversi64(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetPointAnz(REVERSI_STS_BLACK, 0, 0)
	if sts == nil {
		t.Errorf("NG")
	}
}

// TestReversi65 CheckEdge()のテスト
func TestReversi65(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.CheckEdge(REVERSI_STS_BLACK, 0, 0)
	if sts != 0 {
		t.Errorf("NG")
	}
}

// TestReversi66 GetEdgeSideZero()のテスト
func TestReversi66(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetEdgeSideZero(0, 0)
	if sts != 0 {
		t.Errorf("NG")
	}
}

// TestReversi67 GetEdgeSideOne()のテスト
func TestReversi67(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetEdgeSideOne(0, 0)
	if sts == 0 {
		t.Errorf("NG")
	}
}

// TestReversi68 GetEdgeSideTwo()のテスト
func TestReversi68(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetEdgeSideTwo(0, 0)
	if sts == 0 {
		t.Errorf("NG")
	}
}

// TestReversi69 GetEdgeSideThree()のテスト
func TestReversi69(t *testing.T) {
	r := NewReversi(8, 8)
	sts := r.GetEdgeSideThree(0, 0)
	if sts == 0 {
		t.Errorf("NG")
	}
}
