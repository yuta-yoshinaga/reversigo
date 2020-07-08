// Package test リバーシに関する構造体をまとめるパッケージ。
package model

// Copyright Yuta Yoshinaga.

import (
	"testing"
)

// TestReversiAnz1 コンストラクタのテスト
func TestReversiAnz1(t *testing.T) {
	r := NewReversiAnz()
	if r == nil {
		t.Errorf("NG")
	}
}

// TestReversiAnz2 GetMin()のテスト
func TestReversiAnz2(t *testing.T) {
	r := NewReversiAnz()
	if r.GetMin() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz3 SetMin()のテスト
func TestReversiAnz3(t *testing.T) {
	r := NewReversiAnz()
	r.SetMin(1)
	if r.GetMin() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz4 GetMax()のテスト
func TestReversiAnz4(t *testing.T) {
	r := NewReversiAnz()
	if r.GetMax() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz5 SetMax()のテスト
func TestReversiAnz5(t *testing.T) {
	r := NewReversiAnz()
	r.SetMax(1)
	if r.GetMax() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz6 GetAvg()のテスト
func TestReversiAnz6(t *testing.T) {
	r := NewReversiAnz()
	if r.GetAvg() != 0.0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz7 SetAvg()のテスト
func TestReversiAnz7(t *testing.T) {
	r := NewReversiAnz()
	r.SetAvg(1.1)
	if r.GetAvg() != 1.1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz8 GetPointCnt()のテスト
func TestReversiAnz8(t *testing.T) {
	r := NewReversiAnz()
	if r.GetPointCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz9 SetPointCnt()のテスト
func TestReversiAnz9(t *testing.T) {
	r := NewReversiAnz()
	r.SetPointCnt(1)
	if r.GetPointCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz10 GetEdgeCnt()のテスト
func TestReversiAnz10(t *testing.T) {
	r := NewReversiAnz()
	if r.GetEdgeCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz11 SetEdgeCnt()のテスト
func TestReversiAnz11(t *testing.T) {
	r := NewReversiAnz()
	r.SetEdgeCnt(1)
	if r.GetEdgeCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz12 GetEdgeSideOneCnt()のテスト
func TestReversiAnz12(t *testing.T) {
	r := NewReversiAnz()
	if r.GetEdgeSideOneCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz13 SetEdgeSideOneCnt()のテスト
func TestReversiAnz13(t *testing.T) {
	r := NewReversiAnz()
	r.SetEdgeSideOneCnt(1)
	if r.GetEdgeSideOneCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz14 GetEdgeSideTwoCnt()のテスト
func TestReversiAnz14(t *testing.T) {
	r := NewReversiAnz()
	if r.GetEdgeSideTwoCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz15 SetEdgeSideTwoCnt()のテスト
func TestReversiAnz15(t *testing.T) {
	r := NewReversiAnz()
	r.SetEdgeSideTwoCnt(1)
	if r.GetEdgeSideTwoCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz16 GetEdgeSideThreeCnt()のテスト
func TestReversiAnz16(t *testing.T) {
	r := NewReversiAnz()
	if r.GetEdgeSideThreeCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz17 SetEdgeSideTwoCnt()のテスト
func TestReversiAnz17(t *testing.T) {
	r := NewReversiAnz()
	r.SetEdgeSideThreeCnt(1)
	if r.GetEdgeSideThreeCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz18 GetEdgeSideOtherCnt()のテスト
func TestReversiAnz18(t *testing.T) {
	r := NewReversiAnz()
	if r.GetEdgeSideOtherCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz19 SetEdgeSideTwoCnt()のテスト
func TestReversiAnz19(t *testing.T) {
	r := NewReversiAnz()
	r.SetEdgeSideOtherCnt(1)
	if r.GetEdgeSideOtherCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz20 GetOwnMin()のテスト
func TestReversiAnz20(t *testing.T) {
	r := NewReversiAnz()
	if r.GetOwnMin() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz21 SetOwnMin()のテスト
func TestReversiAnz21(t *testing.T) {
	r := NewReversiAnz()
	r.SetOwnMin(1)
	if r.GetOwnMin() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz22 GetOwnMax()のテスト
func TestReversiAnz22(t *testing.T) {
	r := NewReversiAnz()
	if r.GetOwnMax() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz23 SetOwnMax()のテスト
func TestReversiAnz23(t *testing.T) {
	r := NewReversiAnz()
	r.SetOwnMax(1)
	if r.GetOwnMax() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz24 GetOwnAvg()のテスト
func TestReversiAnz24(t *testing.T) {
	r := NewReversiAnz()
	if r.GetOwnAvg() != 0.0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz25 SetOwnAvg()のテスト
func TestReversiAnz25(t *testing.T) {
	r := NewReversiAnz()
	r.SetOwnAvg(1.1)
	if r.GetOwnAvg() != 1.1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz26 GetOwnPointCnt()のテスト
func TestReversiAnz26(t *testing.T) {
	r := NewReversiAnz()
	if r.GetOwnPointCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz27 SetOwnAvg()のテスト
func TestReversiAnz27(t *testing.T) {
	r := NewReversiAnz()
	r.SetOwnPointCnt(1)
	if r.GetOwnPointCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz28 GetOwnPointCnt()のテスト
func TestReversiAnz28(t *testing.T) {
	r := NewReversiAnz()
	if r.GetOwnEdgeCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz29 SetOwnEdgeCnt()のテスト
func TestReversiAnz29(t *testing.T) {
	r := NewReversiAnz()
	r.SetOwnEdgeCnt(1)
	if r.GetOwnEdgeCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz30 GetOwnEdgeSideOneCnt()のテスト
func TestReversiAnz30(t *testing.T) {
	r := NewReversiAnz()
	if r.GetOwnEdgeSideOneCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz31 SetOwnEdgeSideOneCnt()のテスト
func TestReversiAnz31(t *testing.T) {
	r := NewReversiAnz()
	r.SetOwnEdgeSideOneCnt(1)
	if r.GetOwnEdgeSideOneCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz32 GetOwnEdgeSideTwoCnt()のテスト
func TestReversiAnz32(t *testing.T) {
	r := NewReversiAnz()
	if r.GetOwnEdgeSideTwoCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz33 SetOwnEdgeSideTwoCnt()のテスト
func TestReversiAnz33(t *testing.T) {
	r := NewReversiAnz()
	r.SetOwnEdgeSideTwoCnt(1)
	if r.GetOwnEdgeSideTwoCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz34 GetOwnEdgeSideThreeCnt()のテスト
func TestReversiAnz34(t *testing.T) {
	r := NewReversiAnz()
	if r.GetOwnEdgeSideThreeCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz35 SetOwnEdgeSideThreeCnt()のテスト
func TestReversiAnz35(t *testing.T) {
	r := NewReversiAnz()
	r.SetOwnEdgeSideThreeCnt(1)
	if r.GetOwnEdgeSideThreeCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz36 GetOwnEdgeSideOtherCnt()のテスト
func TestReversiAnz36(t *testing.T) {
	r := NewReversiAnz()
	if r.GetOwnEdgeSideOtherCnt() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz37 SetOwnEdgeSideOtherCnt()のテスト
func TestReversiAnz37(t *testing.T) {
	r := NewReversiAnz()
	r.SetOwnEdgeSideOtherCnt(1)
	if r.GetOwnEdgeSideOtherCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz38 GetBadPoint()のテスト
func TestReversiAnz38(t *testing.T) {
	r := NewReversiAnz()
	if r.GetBadPoint() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz39 SetBadPoint()のテスト
func TestReversiAnz39(t *testing.T) {
	r := NewReversiAnz()
	r.SetBadPoint(1)
	if r.GetBadPoint() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz40 GetGoodPoint()のテスト
func TestReversiAnz40(t *testing.T) {
	r := NewReversiAnz()
	if r.GetGoodPoint() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiAnz41 SetBadPoint()のテスト
func TestReversiAnz41(t *testing.T) {
	r := NewReversiAnz()
	r.SetGoodPoint(1)
	if r.GetGoodPoint() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiAnz42 Reset()のテスト
func TestReversiAnz42(t *testing.T) {
	r := NewReversiAnz()
	r.Reset()
	if CheckFieldsReversiAnz(r) == false {
		t.Errorf("NG")
	}
}

// CheckFields Reset()の結果確認
func CheckFieldsReversiAnz(r *ReversiAnz) bool {
	ret := true
	if r.GetMin() != 0 {
		ret = false
	}
	if r.GetMax() != 0 {
		ret = false
	}
	if r.GetAvg() != 0.0 {
		ret = false
	}
	if r.GetPointCnt() != 0 {
		ret = false
	}
	if r.GetEdgeCnt() != 0 {
		ret = false
	}
	if r.GetEdgeSideOneCnt() != 0 {
		ret = false
	}
	if r.GetEdgeSideTwoCnt() != 0 {
		ret = false
	}
	if r.GetEdgeSideThreeCnt() != 0 {
		ret = false
	}
	if r.GetEdgeSideOtherCnt() != 0 {
		ret = false
	}
	if r.GetOwnMin() != 0 {
		ret = false
	}
	if r.GetOwnMax() != 0 {
		ret = false
	}
	if r.GetOwnAvg() != 0.0 {
		ret = false
	}
	if r.GetOwnPointCnt() != 0 {
		ret = false
	}
	if r.GetOwnEdgeCnt() != 0 {
		ret = false
	}
	if r.GetOwnEdgeSideOneCnt() != 0 {
		ret = false
	}
	if r.GetOwnEdgeSideTwoCnt() != 0 {
		ret = false
	}
	if r.GetOwnEdgeSideThreeCnt() != 0 {
		ret = false
	}
	if r.GetOwnEdgeSideOtherCnt() != 0 {
		ret = false
	}
	if r.GetBadPoint() != 0 {
		ret = false
	}
	if r.GetGoodPoint() != 0 {
		ret = false
	}
	return ret
}
