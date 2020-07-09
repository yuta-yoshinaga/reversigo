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

// TestReversiPlay2 GetmReversi()のテスト
func TestReversiPlay2(t *testing.T) {
	r := NewReversiPlay()
	if r.GetmReversi() == nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay3 SetmReversi()のテスト
func TestReversiPlay3(t *testing.T) {
	r := NewReversiPlay()
	r.SetmReversi(NewReversi(8, 8))
	if r.GetmReversi() == nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay4 GetmSetting()のテスト
func TestReversiPlay4(t *testing.T) {
	r := NewReversiPlay()
	if r.GetmSetting() == nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay5 SetmSetting()のテスト
func TestReversiPlay5(t *testing.T) {
	r := NewReversiPlay()
	r.SetmSetting(NewReversiSetting())
	if r.GetmSetting() == nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay6 GetmSetting()のテスト
func TestReversiPlay6(t *testing.T) {
	r := NewReversiPlay()
	if r.GetmCurColor() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiPlay7 SetmSetting()のテスト
func TestReversiPlay7(t *testing.T) {
	r := NewReversiPlay()
	r.SetmCurColor(1)
	if r.GetmCurColor() == 0 {
		t.Errorf("NG")
	}
}

// TestReversiPlay8 GetmCpu()のテスト
func TestReversiPlay8(t *testing.T) {
	r := NewReversiPlay()
	if r.GetmCpu() == nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay9 SetmCpu()のテスト
func TestReversiPlay9(t *testing.T) {
	r := NewReversiPlay()
	r.SetmCpu(make([]*ReversiPoint, 0))
	if r.GetmCpu() == nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay10 GetmEdge()のテスト
func TestReversiPlay10(t *testing.T) {
	r := NewReversiPlay()
	if r.GetmEdge() == nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay11 SetmEdge()のテスト
func TestReversiPlay11(t *testing.T) {
	r := NewReversiPlay()
	r.SetmEdge(make([]*ReversiPoint, 0))
	if r.GetmEdge() == nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay12 GetmPassEnaB()のテスト
func TestReversiPlay12(t *testing.T) {
	r := NewReversiPlay()
	if r.GetmPassEnaB() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiPlay13 SetmPassEnaB()のテスト
func TestReversiPlay13(t *testing.T) {
	r := NewReversiPlay()
	r.SetmPassEnaB(1)
	if r.GetmPassEnaB() == 0 {
		t.Errorf("NG")
	}
}

// TestReversiPlay14 GetmPassEnaW()のテスト
func TestReversiPlay14(t *testing.T) {
	r := NewReversiPlay()
	if r.GetmPassEnaW() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiPlay15 SetmPassEnaW()のテスト
func TestReversiPlay15(t *testing.T) {
	r := NewReversiPlay()
	r.SetmPassEnaW(1)
	if r.GetmPassEnaW() == 0 {
		t.Errorf("NG")
	}
}

// TestReversiPlay16 GetmPassEnaW()のテスト
func TestReversiPlay16(t *testing.T) {
	r := NewReversiPlay()
	if r.GetmGameEndSts() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiPlay17 SetmGameEndSts()のテスト
func TestReversiPlay17(t *testing.T) {
	r := NewReversiPlay()
	r.SetmGameEndSts(1)
	if r.GetmGameEndSts() == 0 {
		t.Errorf("NG")
	}
}

// TestReversiPlay18 GetmPlayLock()のテスト
func TestReversiPlay18(t *testing.T) {
	r := NewReversiPlay()
	if r.GetmPlayLock() != 0 {
		t.Errorf("NG")
	}
}

// TestReversiPlay19 SetmPlayLock()のテスト
func TestReversiPlay19(t *testing.T) {
	r := NewReversiPlay()
	r.SetmPlayLock(1)
	if r.GetmPlayLock() == 0 {
		t.Errorf("NG")
	}
}

// TestReversiPlay20 GetmDelegate()のテスト
func TestReversiPlay20(t *testing.T) {
	r := NewReversiPlay()
	if r.GetmDelegate() != nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay21 SetmDelegate()のテスト
func TestReversiPlay21(t *testing.T) {
	r := NewReversiPlay()
	r.SetmDelegate(NewReversiPlayDelegate(nil))
	if r.GetmDelegate() == nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay22 GetmCallbacks()のテスト
func TestReversiPlay22(t *testing.T) {
	r := NewReversiPlay()
	if r.GetmCallbacks() != nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay23 SetmCallbacks()のテスト
func TestReversiPlay23(t *testing.T) {
	r := NewReversiPlay()
	r.SetmCallbacks(NewCallbacksJson())
	if r.GetmCallbacks() == nil {
		t.Errorf("NG")
	}
}

// TestReversiPlay24 ReversiPlay()のテスト
func TestReversiPlay24(t *testing.T) {
	r := NewReversiPlay()
	r.ReversiPlay(2, 4)
}

// TestReversiPlay25 ReversiPlaySub()のテスト
func TestReversiPlay25(t *testing.T) {
	r := NewReversiPlay()
	r.ReversiPlaySub(1, REVERSI_STS_WHITE)
}

// TestReversiPlay26 ReversiPlayEnd()のテスト
func TestReversiPlay26(t *testing.T) {
	r := NewReversiPlay()
	r.ReversiPlayEnd()
}

// TestReversiPlay27 ReversiPlayPass()のテスト
func TestReversiPlay27(t *testing.T) {
	r := NewReversiPlay()
	r.ReversiPlayPass(REVERSI_STS_BLACK)
}

// TestReversiPlay28 ReversiPlayCpu()のテスト
func TestReversiPlay28(t *testing.T) {
	r := NewReversiPlay()
	r.ReversiPlayCpu(REVERSI_STS_WHITE, 1)
}

// TestReversiPlay29 DrawUpdate()のテスト
func TestReversiPlay29(t *testing.T) {
	r := NewReversiPlay()
	r.DrawUpdate(1)
}

// TestReversiPlay30 DrawUpdateForcibly()のテスト
func TestReversiPlay30(t *testing.T) {
	r := NewReversiPlay()
	r.DrawUpdateForcibly(1)
}

// TestReversiPlay31 Reset()のテスト
func TestReversiPlay31(t *testing.T) {
	r := NewReversiPlay()
	r.Reset()
}

// TestReversiPlay32 GameEndAnimExec()のテスト
func TestReversiPlay32(t *testing.T) {
	r := NewReversiPlay()
	r.GameEndAnimExec()
}

// TestReversiPlay33 SendDrawMsg()のテスト
func TestReversiPlay33(t *testing.T) {
	r := NewReversiPlay()
	r.SendDrawMsg(0, 0)
}

// TestReversiPlay34 SendDrawInfoMsg()のテスト
func TestReversiPlay34(t *testing.T) {
	r := NewReversiPlay()
	r.SendDrawInfoMsg(0, 0)
}

// TestReversiPlay35 ExecMessage()のテスト
func TestReversiPlay35(t *testing.T) {
	r := NewReversiPlay()
	r.ExecMessage(LC_MSG_CUR_COL, nil)
}

// TestReversiPlay36 ViewMsgDlgLocal()のテスト
func TestReversiPlay36(t *testing.T) {
	r := NewReversiPlay()
	r.ViewMsgDlgLocal("", "")
}

// TestReversiPlay37 DrawSingleLocal()のテスト
func TestReversiPlay37(t *testing.T) {
	r := NewReversiPlay()
	r.DrawSingleLocal(0, 0, 0, 0, "")
}

// TestReversiPlay38 CurColMsgLocal()のテスト
func TestReversiPlay38(t *testing.T) {
	r := NewReversiPlay()
	r.CurColMsgLocal("")
}

// TestReversiPlay39 CurStsMsgLocal()のテスト
func TestReversiPlay39(t *testing.T) {
	r := NewReversiPlay()
	r.CurStsMsgLocal("")
}

// TestReversiPlay40 WaitLocal()のテスト
func TestReversiPlay40(t *testing.T) {
	r := NewReversiPlay()
	r.WaitLocal(0)
}
