////////////////////////////////////////////////////////////////////////////////
///	@file			ReversiPlay.go
///	@brief			リバーシプレイクラス実装ファイル
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///	$Version:		$
///	$Revision:		$
///
/// Copyright (c) 2020 Yuta Yoshinaga. All Rights reserved.
///
/// - 本ソフトウェアの一部又は全てを無断で複写複製（コピー）することは、
///   著作権侵害にあたりますので、これを禁止します。
/// - 本製品の使用に起因する侵害または特許権その他権利の侵害に関しては
///   当社は一切その責任を負いません。
///
////////////////////////////////////////////////////////////////////////////////
// Package model リバーシに関する構造体をまとめるパッケージ。
package model

import (
	"math/rand"
	"strconv"
	"time"
)

////////////////////////////////////////////////////////////////////////////////
///	@class		ReversiPlay
///	@brief		リバーシプレイクラス
///
////////////////////////////////////////////////////////////////////////////////
type ReversiPlay struct {
	Reversi    *Reversi             //!< リバーシクラス
	Setting    *ReversiSetting      //!< リバーシ設定クラス
	CurColor   int                  //!< 現在の色
	Cpu        []*ReversiPoint      //!< CPU用ワーク
	Edge       []*ReversiPoint      //!< CPU用角マスワーク
	PassEnaB   int                  //!< 黒のパス有効フラグ
	PassEnaW   int                  //!< 白のパス有効フラグ
	GameEndSts int                  //!< ゲーム終了ステータス
	PlayLock   int                  //!< プレイロック
	Delegate   *ReversiPlayDelegate //!< デリゲート
	Callbacks  *CallbacksJson       //!< コールバック
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			コンストラクタ
///	@fn				NewReversiPlay() *ReversiPlay
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func NewReversiPlay() *ReversiPlay {
	r := new(ReversiPlay)
	r.Reversi = NewReversi(DEF_MASU_CNT_MAX_VAL, DEF_MASU_CNT_MAX_VAL)
	r.Setting = NewReversiSetting()
	r.Cpu = make([]*ReversiPoint, DEF_MASU_CNT_MAX_VAL*DEF_MASU_CNT_MAX_VAL)
	r.Edge = make([]*ReversiPoint, DEF_MASU_CNT_MAX_VAL*DEF_MASU_CNT_MAX_VAL)
	for i := 0; i < (DEF_MASU_CNT_MAX_VAL * DEF_MASU_CNT_MAX_VAL); i++ {
		r.Cpu[i] = NewReversiPoint()
		r.Edge[i] = NewReversiPoint()
	}
	r.CurColor = 0
	r.PassEnaB = 0
	r.PassEnaW = 0
	r.GameEndSts = 0
	r.PlayLock = 0
	r.Delegate = nil
	r.Callbacks = nil
	rand.Seed(time.Now().UnixNano())
	r.Reset()
	return r
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmReversi() *Reversi
///	@return			mReversi *Reversi
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlay) GetmReversi() *Reversi {
	return r.Reversi
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmReversi(mReversi *Reversi)
///	@param[in]		mReversi *Reversi
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SetmReversi(mReversi *Reversi) {
	r.Reversi = mReversi
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmSetting() *ReversiSetting
///	@return			mSetting *ReversiSetting
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlay) GetmSetting() *ReversiSetting {
	return r.Setting
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmSetting(mSetting *ReversiSetting)
///	@param[in]		mSetting *ReversiSetting
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SetmSetting(mSetting *ReversiSetting) {
	r.Setting = mSetting
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmCurColor() int
///	@return			mCurColor int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlay) GetmCurColor() int {
	return r.CurColor
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmCurColor(mCurColor int)
///	@param[in]		mCurColor int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SetmCurColor(mCurColor int) {
	r.CurColor = mCurColor
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmCpu() []*ReversiPoint
///	@return			mCpu []*ReversiPoint
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlay) GetmCpu() []*ReversiPoint {
	return r.Cpu
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmCpu(mCpu []*ReversiPoint)
///	@param[in]		mCpu []*ReversiPoint
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SetmCpu(mCpu []*ReversiPoint) {
	r.Cpu = mCpu
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmEdge() []*ReversiPoint
///	@return			mEdge []*ReversiPoint
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlay) GetmEdge() []*ReversiPoint {
	return r.Edge
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmEdge(mEdge []*ReversiPoint)
///	@param[in]		mEdge []*ReversiPoint
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SetmEdge(mEdge []*ReversiPoint) {
	r.Edge = mEdge
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmPassEnaB() int
///	@return			mPassEnaB int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlay) GetmPassEnaB() int {
	return r.PassEnaB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmPassEnaB(mPassEnaB int)
///	@param[in]		mPassEnaB int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SetmPassEnaB(mPassEnaB int) {
	r.PassEnaB = mPassEnaB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmPassEnaW() int
///	@return			mPassEnaW int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlay) GetmPassEnaW() int {
	return r.PassEnaW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmPassEnaW(mPassEnaW int)
///	@param[in]		mPassEnaW int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SetmPassEnaW(mPassEnaW int) {
	r.PassEnaW = mPassEnaW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmGameEndSts() int
///	@return			mGameEndSts int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlay) GetmGameEndSts() int {
	return r.GameEndSts
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmGameEndSts(mGameEndSts int)
///	@param[in]		mGameEndSts int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SetmGameEndSts(mGameEndSts int) {
	r.GameEndSts = mGameEndSts
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmPlayLock() int
///	@return			mPlayLock int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlay) GetmPlayLock() int {
	return r.PlayLock
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmPlayLock(mPlayLock int)
///	@param[in]		mPlayLock int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SetmPlayLock(mPlayLock int) {
	r.PlayLock = mPlayLock
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmDelegate() *ReversiPlayDelegate
///	@return			mDelegate *ReversiPlayDelegate
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlay) GetmDelegate() *ReversiPlayDelegate {
	return r.Delegate
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmDelegate(mDelegate *ReversiPlayDelegate)
///	@param[in]		mDelegate *ReversiPlayDelegate
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SetmDelegate(mDelegate *ReversiPlayDelegate) {
	r.Delegate = mDelegate
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmCallbacks() *CallbacksJson
///	@return			mCallbacks *CallbacksJson
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlay) GetmCallbacks() *CallbacksJson {
	return r.Callbacks
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmCallbacks(mCallbacks *CallbacksJson)
///	@param[in]		mCallbacks *CallbacksJson
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SetmCallbacks(mCallbacks *CallbacksJson) {
	r.Callbacks = mCallbacks
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			リバーシプレイ
///	@fn				ReversiPlay(y int, x int)
///	@param[in]		y int			Y座標
///	@param[in]		x int			X座標
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) ReversiPlay(y int, x int) {
	update := 0
	cpuEna := 0
	tmpCol := r.CurColor
	pass := 0
	if r.PlayLock == 1 {
		return
	}
	r.PlayLock = 1
	if r.Reversi.GetColorEna(r.CurColor) == 0 {
		if r.Reversi.SetMasuSts(r.CurColor, y, x) == 0 {
			if r.Setting.GetmType() == DEF_TYPE_HARD {
				r.Reversi.AnalysisReversi(r.PassEnaB, r.PassEnaW)
			}
			if r.Setting.GetmAssist() == DEF_ASSIST_ON {
				// *** メッセージ送信 *** //
				r.ExecMessage(LC_MSG_ERASE_INFO_ALL, nil)
			}
			// 描画
			r.SendDrawMsg(y, x)
			// その他コマ描画
			r.DrawUpdate(DEF_ASSIST_OFF)
			if r.Reversi.GetGameEndSts() == 0 {
				if tmpCol == REVERSI_STS_BLACK {
					tmpCol = REVERSI_STS_WHITE
				} else {
					tmpCol = REVERSI_STS_BLACK
				}
				if r.Reversi.GetColorEna(tmpCol) == 0 {
					if r.Setting.GetmMode() == DEF_MODE_ONE {
						// CPU対戦
						cpuEna = 1
					} else {
						// 二人対戦
						r.CurColor = tmpCol
						// 次のプレイヤーコマ描画
						r.DrawUpdate(r.Setting.GetmAssist())
					}
				} else {
					// *** パスメッセージ *** //
					r.ReversiPlayPass(tmpCol)
					pass = 1
				}
			} else {
				// *** ゲーム終了メッセージ *** //
				r.ReversiPlayEnd()
			}
			update = 1
		} else {
			// *** エラーメッセージ *** //
			r.ViewMsgDlgLocal("エラー", "そのマスには置けません。")
		}
	} else {
		if r.Reversi.GetGameEndSts() == 0 {
			if tmpCol == REVERSI_STS_BLACK {
				tmpCol = REVERSI_STS_WHITE
			} else {
				tmpCol = REVERSI_STS_BLACK
			}
			if r.Reversi.GetColorEna(tmpCol) == 0 {
				if r.Setting.GetmMode() == DEF_MODE_ONE {
					// CPU対戦
					update = 1
					cpuEna = 1
				} else {
					// 二人対戦
					r.CurColor = tmpCol
				}
			} else {
				// *** パスメッセージ *** //
				r.ReversiPlayPass(tmpCol)
				pass = 1
			}
		} else {
			// *** ゲーム終了メッセージ *** //
			r.ReversiPlayEnd()
		}
	}
	if pass == 1 {
		if r.Setting.GetmMode() == DEF_MODE_ONE {
			// CPU対戦
			if r.Setting.GetmAssist() == DEF_ASSIST_ON {
				// *** メッセージ送信 *** //
				r.ExecMessage(LC_MSG_DRAW_INFO_ALL, nil)
			}
		}
	}
	if update == 1 {
		waitTime := 0
		if cpuEna == 1 {
			waitTime = r.Setting.GetmPlayCpuInterVal()
		}
		r.WaitLocal(waitTime)
		r.ReversiPlaySub(cpuEna, tmpCol)
		r.PlayLock = 0
	} else {
		r.PlayLock = 0
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			リバーシプレイサブ
///	@fn				ReversiPlaySub(cpuEna int, tmpCol int)
///	@param[in]		cpuEna int
///	@param[in]		tmpCol int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) ReversiPlaySub(cpuEna int, tmpCol int) {
	for {
		ret := r.ReversiPlayCpu(tmpCol, cpuEna)
		cpuEna = 0
		if ret == 1 {
			if r.Reversi.GetGameEndSts() == 0 {
				if r.Reversi.GetColorEna(r.CurColor) != 0 {
					// *** パスメッセージ *** //
					r.ReversiPlayPass(r.CurColor)
					cpuEna = 1
				}
			} else {
				// *** ゲーム終了メッセージ *** //
				r.ReversiPlayEnd()
			}
		}
		if cpuEna == 0 {
			break
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			リバーシプレイ終了
///	@fn				ReversiPlayEnd()
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) ReversiPlayEnd() {
	if r.GameEndSts == 0 {
		r.GameEndSts = 1
		// 終了アニメ実行
		waitTime := r.GameEndAnimExec()
		r.PlayLock = 1
		r.WaitLocal(waitTime)
		// *** ゲーム終了メッセージ *** //
		var tmpMsg1 string
		var tmpMsg2 string
		var msgStr string
		var blk int
		var whi int
		blk = r.Reversi.GetBetCnt(REVERSI_STS_BLACK)
		whi = r.Reversi.GetBetCnt(REVERSI_STS_WHITE)
		tmpMsg1 = "プレイヤー1 = " + strconv.Itoa(blk) + " プレイヤー2 = " + strconv.Itoa(whi)
		if r.Setting.GetmMode() == DEF_MODE_ONE {
			if whi == blk {
				tmpMsg2 = "引き分けです。"
			} else if whi < blk {
				if r.CurColor == REVERSI_STS_BLACK {
					tmpMsg2 = "あなたの勝ちです。"
				} else {
					tmpMsg2 = "あなたの負けです。"
				}
			} else {
				if r.CurColor == REVERSI_STS_WHITE {
					tmpMsg2 = "あなたの勝ちです。"
				} else {
					tmpMsg2 = "あなたの負けです。"
				}
			}
		} else {
			if whi == blk {
				tmpMsg2 = "引き分けです。"
			} else if whi < blk {
				tmpMsg2 = "プレイヤー1の勝ちです。"
			} else {
				tmpMsg2 = "プレイヤー2の勝ちです。"
			}
		}
		msgStr = tmpMsg1 + tmpMsg2
		r.ViewMsgDlgLocal("ゲーム終了", msgStr)
		if r.Setting.GetmEndAnim() == DEF_END_ANIM_ON {
			// *** メッセージ送信 *** //
			r.ExecMessage(LC_MSG_CUR_COL, nil)
			// *** メッセージ送信 *** //
			r.ExecMessage(LC_MSG_CUR_STS, nil)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			リバーシプレイパス
///	@fn				ReversiPlayPass(color int)
///	@param[in]		color int		パス色
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) ReversiPlayPass(color int) {
	// *** パスメッセージ *** //
	if r.Setting.GetmMode() == DEF_MODE_ONE {
		if color == r.CurColor {
			r.ViewMsgDlgLocal("", "あなたはパスです。")
		} else {
			r.ViewMsgDlgLocal("", "CPUはパスです。")
		}
	} else {
		if color == REVERSI_STS_BLACK {
			r.ViewMsgDlgLocal("", "プレイヤー1はパスです。")
		} else {
			r.ViewMsgDlgLocal("", "プレイヤー2はパスです。")
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			リバーシプレイコンピューター
///	@fn				ReversiPlayCpu(color int, cpuEna int) int
///	@param[in]		color int		CPU色
///	@param[in]		cpuEna int		CPU有効フラグ
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) ReversiPlayCpu(color int, cpuEna int) int {
	update := 0
	var setY int
	var setX int
	for {
		if cpuEna == 1 {
			cpuEna = 0
			// *** CPU対戦 *** //
			pCnt := r.Reversi.GetPointCnt(color)
			pInfo := r.Reversi.GetPoint(color, rand.Intn(pCnt))
			if pInfo != nil {
				setY = pInfo.GetY()
				setX = pInfo.GetX()
				if r.Setting.GetmType() != DEF_TYPE_EASY {
					// 強いコンピューター
					cpuflg0 := 0
					cpuflg1 := 0
					cpuflg2 := 0
					cpuflg3 := 0
					mem := -1
					mem2 := -1
					mem3 := -1
					mem4 := -1
					rcnt1 := 0
					rcnt2 := 0
					kadocnt := 0
					loop := r.Setting.GetmMasuCnt() * r.Setting.GetmMasuCnt()
					pcnt := 0
					passCnt := 0
					othColor := 0
					if color == REVERSI_STS_BLACK {
						othColor = REVERSI_STS_WHITE
					} else {
						othColor = REVERSI_STS_BLACK
					}
					// 対戦相手のコマ数
					othBet := r.Reversi.GetBetCnt(othColor)
					// 自分のコマ数
					ownBet := r.Reversi.GetBetCnt(color)
					endZone := 0
					if (loop - (othBet + ownBet)) <= 16 {
						// ゲーム終盤フラグON
						endZone = 1
					}
					for i := 0; i < loop; i++ {
						r.Cpu[i].SetX(0)
						r.Cpu[i].SetY(0)
						r.Edge[i].SetX(0)
						r.Edge[i].SetY(0)
					}

					for i := 0; i < r.Setting.GetmMasuCnt(); i++ {
						for j := 0; j < r.Setting.GetmMasuCnt(); j++ {
							if r.Reversi.GetMasuStsEna(color, i, j) != 0 {
								// *** 角の一つ手前なら別のとこに格納 *** //
								if r.Reversi.GetEdgeSideOne(i, j) == 0 {
									r.Edge[kadocnt].SetX(j)
									r.Edge[kadocnt].SetY(i)
									kadocnt++
								} else {
									r.Cpu[rcnt1].SetX(j)
									r.Cpu[rcnt1].SetY(i)
									rcnt1++
								}
								if r.Setting.GetmType() == DEF_TYPE_NOR {
									// *** 角に置けるなら優先的にとらせるため場所を記憶させる *** //
									if r.Reversi.GetEdgeSideZero(i, j) == 0 {
										cpuflg1 = 1
										rcnt2 = (rcnt1 - 1)
									}
									// *** 角の二つ手前も優先的にとらせるため場所を記憶させる *** //
									if cpuflg1 == 0 {
										if r.Reversi.GetEdgeSideTwo(i, j) == 0 {
											cpuflg2 = 1
											rcnt2 = (rcnt1 - 1)
										}
									}
									// *** 角の三つ手前も優先的にとらせるため場所を記憶させる *** //
									if cpuflg1 == 0 && cpuflg2 == 0 {
										if r.Reversi.GetEdgeSideThree(i, j) == 0 {
											cpuflg0 = 1
											rcnt2 = (rcnt1 - 1)
										}
									}
								}
								// *** パーフェクトゲームなら *** //
								if r.Reversi.GetMasuStsCnt(color, i, j) == othBet {
									setY = i
									setX = j
									pcnt = 1
								}
								// *** 相手をパスさせるなら *** //
								if pcnt == 0 {
									if r.Reversi.GetPassEna(color, i, j) != 0 {
										setY = i
										setX = j
										passCnt = 1
									}
								}
							}
						}
					}

					if pcnt == 0 && passCnt == 0 {
						badPoint := -1
						goodPoint := -1
						pointCnt := -1
						ownPointCnt := -1
						var tmpAnz *ReversiAnz
						if rcnt1 != 0 {
							for i := 0; i < rcnt1; i++ {
								if r.Setting.GetmType() == DEF_TYPE_HARD {
									tmpAnz = r.Reversi.GetPointAnz(color, r.Cpu[i].GetY(), r.Cpu[i].GetX())
									if tmpAnz != nil {
										if badPoint == -1 {
											badPoint = tmpAnz.GetBadPoint()
											goodPoint = tmpAnz.GetGoodPoint()
											pointCnt = tmpAnz.GetPointCnt()
											ownPointCnt = tmpAnz.GetOwnPointCnt()
											mem2 = i
											mem3 = i
											mem4 = i
										} else {
											if tmpAnz.GetBadPoint() < badPoint {
												badPoint = tmpAnz.GetBadPoint()
												mem2 = i
											}
											if goodPoint < tmpAnz.GetGoodPoint() {
												goodPoint = tmpAnz.GetGoodPoint()
												mem3 = i
											}
											if tmpAnz.GetPointCnt() < pointCnt {
												pointCnt = tmpAnz.GetPointCnt()
												ownPointCnt = tmpAnz.GetOwnPointCnt()
												mem4 = i
											} else if tmpAnz.GetPointCnt() == pointCnt {
												if ownPointCnt < tmpAnz.GetOwnPointCnt() {
													ownPointCnt = tmpAnz.GetOwnPointCnt()
													mem4 = i
												}
											}
										}
									}
								}
								if r.Reversi.GetMasuStsEna(color, r.Cpu[i].GetY(), r.Cpu[i].GetX()) == 2 {
									mem = i
								}
							}
							if mem2 != -1 {
								if endZone != 0 {
									// 終盤なら枚数重視
									if mem3 != -1 {
										mem2 = mem3
									}
								} else {
									if mem4 != -1 {
										mem2 = mem4
									}
								}
								mem = mem2
							}
							if mem == -1 {
								mem = rand.Intn(rcnt1)
							}
						} else if kadocnt != 0 {
							for i := 0; i < kadocnt; i++ {
								if r.Setting.GetmType() == DEF_TYPE_HARD {
									tmpAnz = r.Reversi.GetPointAnz(color, r.Edge[i].GetY(), r.Edge[i].GetX())
									if tmpAnz != nil {
										if badPoint == -1 {
											badPoint = tmpAnz.GetBadPoint()
											goodPoint = tmpAnz.GetGoodPoint()
											pointCnt = tmpAnz.GetPointCnt()
											ownPointCnt = tmpAnz.GetOwnPointCnt()
											mem2 = i
											mem3 = i
											mem4 = i
										} else {
											if tmpAnz.GetBadPoint() < badPoint {
												badPoint = tmpAnz.GetBadPoint()
												mem2 = i
											}
											if goodPoint < tmpAnz.GetGoodPoint() {
												goodPoint = tmpAnz.GetGoodPoint()
												mem3 = i
											}
											if tmpAnz.GetPointCnt() < pointCnt {
												pointCnt = tmpAnz.GetPointCnt()
												ownPointCnt = tmpAnz.GetOwnPointCnt()
												mem4 = i
											} else if tmpAnz.GetPointCnt() == pointCnt {
												if ownPointCnt < tmpAnz.GetOwnPointCnt() {
													ownPointCnt = tmpAnz.GetOwnPointCnt()
													mem4 = i
												}
											}
										}
									}
								}
								if r.Reversi.GetMasuStsEna(color, r.Edge[i].GetY(), r.Edge[i].GetX()) == 2 {
									mem = i
								}
							}
							if mem2 != -1 {
								if endZone != 0 {
									// 終盤なら枚数重視
									if mem3 != -1 {
										mem2 = mem3
									}
								} else {
									if mem4 != -1 {
										mem2 = mem4
									}
								}
								mem = mem2
							}
							if mem == -1 {
								mem = rand.Intn(kadocnt)
							}
							// *** 置いても平気な角があればそこに置く*** //
							for i := 0; i < kadocnt; i++ {
								if r.Reversi.CheckEdge(color, r.Edge[i].GetY(), r.Edge[i].GetX()) != 0 {
									if cpuflg0 == 0 && cpuflg1 == 0 && cpuflg2 == 0 {
										cpuflg3 = 1
										rcnt2 = i
									}
								}
							}
						}
						if cpuflg1 == 0 && cpuflg2 == 0 && cpuflg0 == 0 && cpuflg3 == 0 {
							rcnt2 = mem
						}
						if rcnt1 != 0 {
							setY = r.Cpu[rcnt2].GetY()
							setX = r.Cpu[rcnt2].GetX()
						} else if kadocnt != 0 {
							setY = r.Edge[rcnt2].GetY()
							setX = r.Edge[rcnt2].GetX()
						}
					}
				}
				if r.Reversi.SetMasuSts(color, setY, setX) == 0 {
					if r.Setting.GetmType() == DEF_TYPE_HARD {
						r.Reversi.AnalysisReversi(r.PassEnaB, r.PassEnaW)
					}
					// 描画
					r.SendDrawMsg(setY, setX)
					update = 1
				}
			}
		} else {
			break
		}
	}
	if update == 1 {
		r.DrawUpdate(DEF_ASSIST_OFF)
		if r.Setting.GetmAssist() == DEF_ASSIST_ON {
			// *** メッセージ送信 *** //
			r.ExecMessage(LC_MSG_DRAW_INFO_ALL, nil)
		}
	}
	return update
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			マス描画更新
///	@fn				DrawUpdate(assist int)
///	@param[in]		assist int	アシスト設定
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) DrawUpdate(assist int) {
	if assist == DEF_ASSIST_ON {
		for i := 0; i < r.Setting.GetmMasuCnt(); i++ {
			for j := 0; j < r.Setting.GetmMasuCnt(); j++ {
				r.SendDrawInfoMsg(i, j)
			}
		}
	}
	waitTime := r.Setting.GetmPlayDrawInterVal()
	for i := 0; i < r.Setting.GetmMasuCnt(); i++ {
		for j := 0; j < r.Setting.GetmMasuCnt(); j++ {
			if r.Reversi.GetMasuSts(i, j) != r.Reversi.GetMasuStsOld(i, j) {
				r.WaitLocal(waitTime)
				r.SendDrawMsg(i, j)
			}
		}
	}
	// *** メッセージ送信 *** //
	r.ExecMessage(LC_MSG_CUR_COL, nil)
	// *** メッセージ送信 *** //
	r.ExecMessage(LC_MSG_CUR_STS, nil)
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			マス描画強制更新
///	@fn				DrawUpdateForcibly(assist int)
///	@param[in]		assist int	アシスト設定
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) DrawUpdateForcibly(assist int) {
	// *** メッセージ送信 *** //
	r.ExecMessage(LC_MSG_DRAW_ALL, nil)
	if assist == DEF_ASSIST_ON {
		// *** メッセージ送信 *** //
		r.ExecMessage(LC_MSG_DRAW_INFO_ALL, nil)
	} else {
		// *** メッセージ送信 *** //
		r.ExecMessage(LC_MSG_ERASE_INFO_ALL, nil)
	}
	// *** メッセージ送信 *** //
	r.ExecMessage(LC_MSG_CUR_COL, nil)
	// *** メッセージ送信 *** //
	r.ExecMessage(LC_MSG_CUR_STS, nil)
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			リセット処理
///	@fn				Reset()
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) Reset() {
	r.PassEnaB = 0
	r.PassEnaW = 0
	if r.Setting.GetmGameSpd() == DEF_GAME_SPD_FAST {
		r.Setting.SetmPlayDrawInterVal(DEF_GAME_SPD_FAST_VAL) // 描画のインターバル(msec)
		r.Setting.SetmPlayCpuInterVal(DEF_GAME_SPD_FAST_VAL2) // CPU対戦時のインターバル(msec)
	} else if r.Setting.GetmGameSpd() == DEF_GAME_SPD_MID {
		r.Setting.SetmPlayDrawInterVal(DEF_GAME_SPD_MID_VAL) // 描画のインターバル(msec)
		r.Setting.SetmPlayCpuInterVal(DEF_GAME_SPD_MID_VAL2) // CPU対戦時のインターバル(msec)
	} else {
		r.Setting.SetmPlayDrawInterVal(DEF_GAME_SPD_SLOW_VAL) // 描画のインターバル(msec)
		r.Setting.SetmPlayCpuInterVal(DEF_GAME_SPD_SLOW_VAL2) // CPU対戦時のインターバル(msec)
	}
	r.CurColor = r.Setting.GetmPlayer()
	if r.Setting.GetmMode() == DEF_MODE_TWO {
		r.CurColor = REVERSI_STS_BLACK
	}
	r.Reversi.SetMasuCnt(r.Setting.GetmMasuCnt()) // マスの数設定
	r.Reversi.Reset()
	if r.Setting.GetmMode() == DEF_MODE_ONE {
		if r.CurColor == REVERSI_STS_WHITE {
			pCnt := r.Reversi.GetPointCnt(REVERSI_STS_BLACK)
			pInfo := r.Reversi.GetPoint(REVERSI_STS_BLACK, rand.Intn(pCnt))
			if pInfo != nil {
				r.Reversi.SetMasuSts(REVERSI_STS_BLACK, pInfo.GetY(), pInfo.GetX())
				if r.Setting.GetmType() == DEF_TYPE_HARD {
					r.Reversi.AnalysisReversi(r.PassEnaB, r.PassEnaW)
				}
			}
		}
	}
	r.PlayLock = 1
	r.GameEndSts = 0
	r.DrawUpdateForcibly(r.Setting.GetmAssist())
	// *** 終了通知 *** //
	// *** メッセージ送信 *** //
	r.ExecMessage(LC_MSG_DRAW_END, nil)
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲーム終了アニメーション
///	@fn				GameEndAnimExec() int
///	@return			ウェイト時間
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) GameEndAnimExec() int {
	var bCnt int
	var wCnt int
	ret := 0
	if r.Setting.GetmEndAnim() == DEF_END_ANIM_ON {
		bCnt = r.Reversi.GetBetCnt(REVERSI_STS_BLACK)
		wCnt = r.Reversi.GetBetCnt(REVERSI_STS_WHITE)
		// *** 色、コマ数表示消去 *** //
		// *** メッセージ送信 *** //
		r.ExecMessage(LC_MSG_CUR_COL_ERASE, nil)
		// *** メッセージ送信 *** //
		r.ExecMessage(LC_MSG_CUR_STS_ERASE, nil)
		r.WaitLocal(r.Setting.GetmEndInterVal())
		// *** マス消去 *** //
		for i := 0; i < r.Setting.GetmMasuCnt(); i++ {
			for j := 0; j < r.Setting.GetmMasuCnt(); j++ {
				r.Reversi.SetMasuStsForcibly(REVERSI_STS_NONE, i, j)
			}
		}
		// *** メッセージ送信 *** //
		r.ExecMessage(LC_MSG_ERASE_ALL, nil)
		// *** マス描画 *** //
		bCnt2 := 0
		wCnt2 := 0
		bEnd := 0
		wEnd := 0
		for i := 0; i < r.Setting.GetmMasuCnt(); i++ {
			for j := 0; j < r.Setting.GetmMasuCnt(); j++ {
				if bCnt2 < bCnt {
					bCnt2++
					r.Reversi.SetMasuStsForcibly(REVERSI_STS_BLACK, i, j)
					r.SendDrawMsg(i, j)
				} else {
					bEnd = 1
				}
				if wCnt2 < wCnt {
					wCnt2++
					r.Reversi.SetMasuStsForcibly(REVERSI_STS_WHITE, (r.Setting.GetmMasuCnt()-1)-i, (r.Setting.GetmMasuCnt()-1)-j)
					r.SendDrawMsg((r.Setting.GetmMasuCnt()-1)-i, (r.Setting.GetmMasuCnt()-1)-j)
				} else {
					wEnd = 1
				}
				if bEnd == 1 && wEnd == 1 {
					break
				} else {
					r.WaitLocal(r.Setting.GetmEndDrawInterVal())
				}
			}
		}
		ret = 0
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			描画メッセージ送信
///	@fn				SendDrawMsg(y int, x int)
///	@param[in]		y int			Y座標
///	@param[in]		x int			X座標
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SendDrawMsg(y int, x int) {
	mTmpPoint := NewReversiPoint()
	mTmpPoint.SetY(y)
	mTmpPoint.SetX(x)
	// *** メッセージ送信 *** //
	r.ExecMessage(LC_MSG_DRAW, mTmpPoint)
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			情報描画メッセージ送信
///	@fn				SendDrawInfoMsg(y int, x int)
///	@param[in]		y int			Y座標
///	@param[in]		x int			X座標
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) SendDrawInfoMsg(y int, x int) {
	mTmpPoint := NewReversiPoint()
	mTmpPoint.SetY(y)
	mTmpPoint.SetX(x)
	// *** メッセージ送信 *** //
	r.ExecMessage(LC_MSG_DRAW_INFO, mTmpPoint)
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			メッセージ
///	@fn				ExecMessage(what int, msgPoint *ReversiPoint)
///	@param[in]		what int
///	@param[in]		msgPoint *ReversiPoint
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) ExecMessage(what int, msgPoint *ReversiPoint) {
	var dMode int
	var dBack int
	var dCnt int
	if what == LC_MSG_DRAW {
		// *** マス描画 *** //
		dMode = r.Reversi.GetMasuSts(msgPoint.GetY(), msgPoint.GetX())
		dBack = r.Reversi.GetMasuStsEna(r.CurColor, msgPoint.GetY(), msgPoint.GetX())
		dCnt = r.Reversi.GetMasuStsCnt(r.CurColor, msgPoint.GetY(), msgPoint.GetX())
		r.DrawSingleLocal(msgPoint.GetY(), msgPoint.GetX(), dMode, dBack, strconv.Itoa(dCnt))
	} else if what == LC_MSG_ERASE {
		// *** マス消去 *** //
		r.DrawSingleLocal(msgPoint.GetY(), msgPoint.GetX(), 0, 0, "0")
	} else if what == LC_MSG_DRAW_INFO {
		// *** マス情報描画 *** //
		dMode = r.Reversi.GetMasuSts(msgPoint.GetY(), msgPoint.GetX())
		dBack = r.Reversi.GetMasuStsEna(r.CurColor, msgPoint.GetY(), msgPoint.GetX())
		dCnt = r.Reversi.GetMasuStsCnt(r.CurColor, msgPoint.GetY(), msgPoint.GetX())
		r.DrawSingleLocal(msgPoint.GetY(), msgPoint.GetX(), dMode, dBack, strconv.Itoa(dCnt))
	} else if what == LC_MSG_ERASE_INFO {
		// *** マス情報消去 *** //
		dMode = r.Reversi.GetMasuSts(msgPoint.GetY(), msgPoint.GetX())
		r.DrawSingleLocal(msgPoint.GetY(), msgPoint.GetX(), dMode, 0, "0")
	} else if what == LC_MSG_DRAW_ALL {
		// *** 全マス描画 *** //
		for i := 0; i < r.Setting.GetmMasuCnt(); i++ {
			for j := 0; j < r.Setting.GetmMasuCnt(); j++ {
				dMode = r.Reversi.GetMasuSts(i, j)
				dBack = r.Reversi.GetMasuStsEna(r.CurColor, i, j)
				dCnt = r.Reversi.GetMasuStsCnt(r.CurColor, i, j)
				r.DrawSingleLocal(i, j, dMode, dBack, strconv.Itoa(dCnt))
			}
		}
	} else if what == LC_MSG_ERASE_ALL {
		// *** 全マス消去 *** //
		for i := 0; i < r.Setting.GetmMasuCnt(); i++ {
			for j := 0; j < r.Setting.GetmMasuCnt(); j++ {
				r.DrawSingleLocal(i, j, 0, 0, "0")
			}
		}
	} else if what == LC_MSG_DRAW_INFO_ALL {
		// *** 全マス情報描画 *** //
		for i := 0; i < r.Setting.GetmMasuCnt(); i++ {
			for j := 0; j < r.Setting.GetmMasuCnt(); j++ {
				dMode = r.Reversi.GetMasuSts(i, j)
				dBack = r.Reversi.GetMasuStsEna(r.CurColor, i, j)
				dCnt = r.Reversi.GetMasuStsCnt(r.CurColor, i, j)
				r.DrawSingleLocal(i, j, dMode, dBack, strconv.Itoa(dCnt))
			}
		}
	} else if what == LC_MSG_ERASE_INFO_ALL {
		// *** 全マス情報消去 *** //
		for i := 0; i < r.Setting.GetmMasuCnt(); i++ {
			for j := 0; j < r.Setting.GetmMasuCnt(); j++ {
				dMode = r.Reversi.GetMasuSts(i, j)
				r.DrawSingleLocal(i, j, dMode, 0, "0")
			}
		}
	} else if what == LC_MSG_DRAW_END {
		r.PlayLock = 0
	} else if what == LC_MSG_CUR_COL {
		tmpStr := ""
		if r.Setting.GetmMode() == DEF_MODE_ONE {
			if r.CurColor == REVERSI_STS_BLACK {
				tmpStr = "あなたはプレイヤー1です "
			} else {
				tmpStr = "あなたはプレイヤー2です "
			}
		} else {
			if r.CurColor == REVERSI_STS_BLACK {
				tmpStr = "プレイヤー1の番です "
			} else {
				tmpStr = "プレイヤー2の番です "
			}
		}
		r.CurColMsgLocal(tmpStr)
	} else if what == LC_MSG_CUR_COL_ERASE {
		r.CurColMsgLocal("")
	} else if what == LC_MSG_CUR_STS {
		tmpStr := "プレイヤー1 = " + strconv.Itoa(r.Reversi.GetBetCnt(REVERSI_STS_BLACK)) + " プレイヤー2 = " + strconv.Itoa(r.Reversi.GetBetCnt(REVERSI_STS_WHITE))
		r.CurStsMsgLocal(tmpStr)
	} else if what == LC_MSG_CUR_STS_ERASE {
		r.CurStsMsgLocal("")
	} else if what == LC_MSG_MSG_DLG {
	} else if what == LC_MSG_DRAW_ALL_RESET {
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			メッセージダイアログ
///	@fn				ViewMsgDlgLocal(title string, msg string)
///	@param[in]		title string	タイトル
///	@param[in]		msg string		メッセージ
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) ViewMsgDlgLocal(title string, msg string) {
	if r.Delegate != nil {
		funcs := r.Callbacks.GetFuncs()
		funcs = append(funcs, r.Delegate.ViewMsgDlg(title, msg))
		r.Callbacks.SetFuncs(funcs)
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			1マス描画
///	@fn				void DrawSingleLocal(int y, int x, int sts, int bk, String text)
///	@param[in]		y int		Y座標
///	@param[in]		x int		X座標
///	@param[in]		sts int 	ステータス
///	@param[in]		bk int		背景
///	@param[in]		text string	テキスト
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) DrawSingleLocal(y int, x int, sts int, bk int, text string) {
	if r.Delegate != nil {
		funcs := r.Callbacks.GetFuncs()
		funcs = append(funcs, r.Delegate.DrawSingle(y, x, sts, bk, text))
		r.Callbacks.SetFuncs(funcs)
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			現在の色メッセージ
///	@fn				CurColMsgLocal(text string)
///	@param[in]		text string	テキスト
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) CurColMsgLocal(text string) {
	if r.Delegate != nil {
		funcs := r.Callbacks.GetFuncs()
		funcs = append(funcs, r.Delegate.CurColMsg(text))
		r.Callbacks.SetFuncs(funcs)
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			現在のステータスメッセージ
///	@fn				CurStsMsgLocal(text string)
///	@param[in]		text string	テキスト
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) CurStsMsgLocal(text string) {
	if r.Delegate != nil {
		funcs := r.Callbacks.GetFuncs()
		funcs = append(funcs, r.Delegate.CurStsMsg(text))
		r.Callbacks.SetFuncs(funcs)
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ウェイト
///	@fn				WaitLocal(time int)
///	@param[in]		time int	ウェイト時間(msec)
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiPlay) WaitLocal(time int) {
	if r.Delegate != nil {
		funcs := r.Callbacks.GetFuncs()
		funcs = append(funcs, r.Delegate.Wait(time))
		r.Callbacks.SetFuncs(funcs)
	}
}
