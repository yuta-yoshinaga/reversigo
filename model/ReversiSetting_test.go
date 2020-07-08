// Package test リバーシに関する構造体をまとめるパッケージ。
package model

// Copyright Yuta Yoshinaga.

import (
	"testing"
)

// TestReversiSetting1 コンストラクタのテスト
func TestReversiSetting1(t *testing.T) {
	r := NewReversiSetting()
	if r == nil {
		t.Errorf("NG")
	}
}

// TestReversiSetting2 GetmMode()のテスト
func TestReversiSetting2(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmMode() != DEF_MODE_ONE {
		t.Errorf("NG")
	}
}

// TestReversiSetting3 SetmMode()のテスト
func TestReversiSetting3(t *testing.T) {
	r := NewReversiSetting()
	r.SetmMode(1)
	if r.GetmMode() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting4 GetmType()のテスト
func TestReversiSetting4(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmType() != DEF_TYPE_HARD {
		t.Errorf("NG")
	}
}

// TestReversiSetting5 SetmType()のテスト
func TestReversiSetting5(t *testing.T) {
	r := NewReversiSetting()
	r.SetmType(1)
	if r.GetmType() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting6 GetmPlayer()のテスト
func TestReversiSetting6(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmPlayer() != REVERSI_STS_BLACK {
		t.Errorf("NG")
	}
}

// TestReversiSetting7 SetmPlayer()のテスト
func TestReversiSetting7(t *testing.T) {
	r := NewReversiSetting()
	r.SetmPlayer(1)
	if r.GetmPlayer() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting8 GetmAssist()のテスト
func TestReversiSetting8(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmAssist() != DEF_ASSIST_ON {
		t.Errorf("NG")
	}
}

// TestReversiSetting9 SetmAssist()のテスト
func TestReversiSetting9(t *testing.T) {
	r := NewReversiSetting()
	r.SetmAssist(1)
	if r.GetmAssist() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting10 GetmGameSpd()のテスト
func TestReversiSetting10(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmGameSpd() != DEF_GAME_SPD_MID {
		t.Errorf("NG")
	}
}

// TestReversiSetting11 SetmGameSpd()のテスト
func TestReversiSetting11(t *testing.T) {
	r := NewReversiSetting()
	r.SetmGameSpd(1)
	if r.GetmGameSpd() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting12 GetmEndAnim()のテスト
func TestReversiSetting12(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmEndAnim() != DEF_END_ANIM_ON {
		t.Errorf("NG")
	}
}

// TestReversiSetting13 SetmEndAnim()のテスト
func TestReversiSetting13(t *testing.T) {
	r := NewReversiSetting()
	r.SetmEndAnim(1)
	if r.GetmEndAnim() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting14 GetmMasuCntMenu()のテスト
func TestReversiSetting14(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmMasuCntMenu() != DEF_MASU_CNT_8 {
		t.Errorf("NG")
	}
}

// TestReversiSetting15 SetmMasuCntMenu()のテスト
func TestReversiSetting15(t *testing.T) {
	r := NewReversiSetting()
	r.SetmMasuCntMenu(1)
	if r.GetmMasuCntMenu() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting16 GetmMasuCnt()のテスト
func TestReversiSetting16(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmMasuCnt() != DEF_MASU_CNT_8_VAL {
		t.Errorf("NG")
	}
}

// TestReversiSetting17 SetmMasuCnt()のテスト
func TestReversiSetting17(t *testing.T) {
	r := NewReversiSetting()
	r.SetmMasuCnt(1)
	if r.GetmMasuCnt() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting18 GetmPlayCpuInterVal()のテスト
func TestReversiSetting18(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmPlayCpuInterVal() != DEF_GAME_SPD_MID_VAL2 {
		t.Errorf("NG")
	}
}

// TestReversiSetting19 SetmPlayCpuInterVal()のテスト
func TestReversiSetting19(t *testing.T) {
	r := NewReversiSetting()
	r.SetmPlayCpuInterVal(1)
	if r.GetmPlayCpuInterVal() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting20 GetmPlayDrawInterVal()のテスト
func TestReversiSetting20(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmPlayDrawInterVal() != DEF_GAME_SPD_MID_VAL {
		t.Errorf("NG")
	}
}

// TestReversiSetting21 SetmPlayDrawInterVal()のテスト
func TestReversiSetting21(t *testing.T) {
	r := NewReversiSetting()
	r.SetmPlayDrawInterVal(1)
	if r.GetmPlayDrawInterVal() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting22 GetmEndDrawInterVal()のテスト
func TestReversiSetting22(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmEndDrawInterVal() != 100 {
		t.Errorf("NG")
	}
}

// TestReversiSetting23 SetmEndDrawInterVal()のテスト
func TestReversiSetting23(t *testing.T) {
	r := NewReversiSetting()
	r.SetmEndDrawInterVal(1)
	if r.GetmEndDrawInterVal() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting24 GetmEndInterVal()のテスト
func TestReversiSetting24(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmEndInterVal() != 500 {
		t.Errorf("NG")
	}
}

// TestReversiSetting25 SetmEndInterVal()のテスト
func TestReversiSetting25(t *testing.T) {
	r := NewReversiSetting()
	r.SetmEndInterVal(1)
	if r.GetmEndInterVal() != 1 {
		t.Errorf("NG")
	}
}

// TestReversiSetting26 GetmPlayerColor1()のテスト
func TestReversiSetting26(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmPlayerColor1() != "#000000" {
		t.Errorf("NG")
	}
}

// TestReversiSetting27 SetmPlayerColor1()のテスト
func TestReversiSetting27(t *testing.T) {
	r := NewReversiSetting()
	r.SetmPlayerColor1("SetmPlayerColor1")
	if r.GetmPlayerColor1() != "SetmPlayerColor1" {
		t.Errorf("NG")
	}
}

// TestReversiSetting28 GetmPlayerColor2()のテスト
func TestReversiSetting28(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmPlayerColor2() != "#FFFFFF" {
		t.Errorf("NG")
	}
}

// TestReversiSetting29 SetmPlayerColor2()のテスト
func TestReversiSetting29(t *testing.T) {
	r := NewReversiSetting()
	r.SetmPlayerColor2("SetmPlayerColor2")
	if r.GetmPlayerColor2() != "SetmPlayerColor2" {
		t.Errorf("NG")
	}
}

// TestReversiSetting30 GetmBackGroundColor()のテスト
func TestReversiSetting30(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmBackGroundColor() != "#00FF00" {
		t.Errorf("NG")
	}
}

// TestReversiSetting31 SetmBackGroundColor()のテスト
func TestReversiSetting31(t *testing.T) {
	r := NewReversiSetting()
	r.SetmBackGroundColor("SetmBackGroundColor")
	if r.GetmBackGroundColor() != "SetmBackGroundColor" {
		t.Errorf("NG")
	}
}

// TestReversiSetting32 GetmBorderColor()のテスト
func TestReversiSetting32(t *testing.T) {
	r := NewReversiSetting()
	if r.GetmBorderColor() != "#000000" {
		t.Errorf("NG")
	}
}

// TestReversiSetting33 SetmBorderColor()のテスト
func TestReversiSetting33(t *testing.T) {
	r := NewReversiSetting()
	r.SetmBorderColor("SetmBorderColor")
	if r.GetmBorderColor() != "SetmBorderColor" {
		t.Errorf("NG")
	}
}

// TestReversiSetting33 Reset()のテスト
func TestReversiSetting34(t *testing.T) {
	r := NewReversiSetting()
	r.Reset()
	if CheckFieldsReversiSetting(r) == false {
		t.Errorf("NG")
	}
}

// CheckFields Reset()の結果確認
func CheckFieldsReversiSetting(r *ReversiSetting) bool {
	ret := true
	if r.GetmMode() != DEF_MODE_ONE {
		ret = false
	}
	if r.GetmType() != DEF_TYPE_HARD {
		ret = false
	}
	if r.GetmPlayer() != REVERSI_STS_BLACK {
		ret = false
	}
	if r.GetmAssist() != DEF_ASSIST_ON {
		ret = false
	}
	if r.GetmGameSpd() != DEF_GAME_SPD_MID {
		ret = false
	}
	if r.GetmEndAnim() != DEF_END_ANIM_ON {
		ret = false
	}
	if r.GetmMasuCntMenu() != DEF_MASU_CNT_8 {
		ret = false
	}
	if r.GetmMasuCnt() != DEF_MASU_CNT_8_VAL {
		ret = false
	}
	if r.GetmPlayCpuInterVal() != DEF_GAME_SPD_MID_VAL2 {
		ret = false
	}
	if r.GetmPlayDrawInterVal() != DEF_GAME_SPD_MID_VAL {
		ret = false
	}
	if r.GetmEndDrawInterVal() != 100 {
		ret = false
	}
	if r.GetmEndInterVal() != 500 {
		ret = false
	}
	if r.GetmPlayerColor1() != "#000000" {
		ret = false
	}
	if r.GetmPlayerColor2() != "#FFFFFF" {
		ret = false
	}
	if r.GetmBackGroundColor() != "#00FF00" {
		ret = false
	}
	if r.GetmBorderColor() != "#000000" {
		ret = false
	}
	return ret
}
