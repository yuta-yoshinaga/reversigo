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
	mReversi    *Reversi             //!< リバーシクラス
	mSetting    *ReversiSetting      //!< リバーシ設定クラス
	mCurColor   int                  //!< 現在の色
	mCpu        []*ReversiPoint      //!< CPU用ワーク
	mEdge       []*ReversiPoint      //!< CPU用角マスワーク
	mPassEnaB   int                  //!< 黒のパス有効フラグ
	mPassEnaW   int                  //!< 白のパス有効フラグ
	mGameEndSts int                  //!< ゲーム終了ステータス
	mPlayLock   int                  //!< プレイロック
	mDelegate   *ReversiPlayDelegate //!< デリゲート
	mCallbacks  *CallbacksJson       //!< コールバック
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
	r.mReversi = NewReversi(DEF_MASU_CNT_MAX_VAL, DEF_MASU_CNT_MAX_VAL)
	r.mSetting = NewReversiSetting()
	r.mCpu = make([]*ReversiPoint, DEF_MASU_CNT_MAX_VAL*DEF_MASU_CNT_MAX_VAL)
	r.mEdge = make([]*ReversiPoint, DEF_MASU_CNT_MAX_VAL*DEF_MASU_CNT_MAX_VAL)
	for i := 0; i < (DEF_MASU_CNT_MAX_VAL * DEF_MASU_CNT_MAX_VAL); i++ {
		r.mCpu[i] = NewReversiPoint()
		r.mEdge[i] = NewReversiPoint()
	}
	r.mCurColor = 0
	r.mPassEnaB = 0
	r.mPassEnaW = 0
	r.mGameEndSts = 0
	r.mPlayLock = 0
	r.mDelegate = nil
	r.mCallbacks = nil
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
	return r.mReversi
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
	r.mReversi = mReversi
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
	return r.mSetting
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
	r.mSetting = mSetting
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
	return r.mCurColor
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
	r.mCurColor = mCurColor
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
	return r.mCpu
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
	r.mCpu = mCpu
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
	return r.mEdge
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
	r.mEdge = mEdge
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
	return r.mPassEnaB
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
	r.mPassEnaB = mPassEnaB
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
	return r.mPassEnaW
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
	r.mPassEnaW = mPassEnaW
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
	return r.mGameEndSts
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
	r.mGameEndSts = mGameEndSts
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
	return r.mPlayLock
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
	r.mPlayLock = mPlayLock
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
	return r.mDelegate
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
	r.mDelegate = mDelegate
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
	return r.mCallbacks
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
	r.mCallbacks = mCallbacks
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
	tmpCol := r.mCurColor
	pass := 0
	if r.mPlayLock == 1 {
		return
	}
	r.mPlayLock = 1
	if r.mReversi.GetColorEna(r.mCurColor) == 0 {
		if r.mReversi.SetMasuSts(r.mCurColor, y, x) == 0 {
			if r.mSetting.GetmType() == DEF_TYPE_HARD {
				r.mReversi.AnalysisReversi(r.mPassEnaB, r.mPassEnaW)
			}
			if r.mSetting.GetmAssist() == DEF_ASSIST_ON {
				// *** メッセージ送信 *** //
				r.ExecMessage(LC_MSG_ERASE_INFO_ALL, nil)
			}
			// 描画
			r.SendDrawMsg(y, x)
			// その他コマ描画
			r.DrawUpdate(DEF_ASSIST_OFF)
			if r.mReversi.GetGameEndSts() == 0 {
				if tmpCol == REVERSI_STS_BLACK {
					tmpCol = REVERSI_STS_WHITE
				} else {
					tmpCol = REVERSI_STS_BLACK
				}
				if r.mReversi.GetColorEna(tmpCol) == 0 {
					if r.mSetting.GetmMode() == DEF_MODE_ONE {
						// CPU対戦
						cpuEna = 1
					} else {
						// 二人対戦
						r.mCurColor = tmpCol
						// 次のプレイヤーコマ描画
						r.DrawUpdate(r.mSetting.GetmAssist())
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
		if r.mReversi.GetGameEndSts() == 0 {
			if tmpCol == REVERSI_STS_BLACK {
				tmpCol = REVERSI_STS_WHITE
			} else {
				tmpCol = REVERSI_STS_BLACK
			}
			if r.mReversi.GetColorEna(tmpCol) == 0 {
				if r.mSetting.GetmMode() == DEF_MODE_ONE {
					// CPU対戦
					update = 1
					cpuEna = 1
				} else {
					// 二人対戦
					r.mCurColor = tmpCol
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
		if r.mSetting.GetmMode() == DEF_MODE_ONE {
			// CPU対戦
			if r.mSetting.GetmAssist() == DEF_ASSIST_ON {
				// *** メッセージ送信 *** //
				r.ExecMessage(LC_MSG_DRAW_INFO_ALL, nil)
			}
		}
	}
	if update == 1 {
		waitTime := 0
		if cpuEna == 1 {
			waitTime = r.mSetting.GetmPlayCpuInterVal()
		}
		r.WaitLocal(waitTime)
		r.ReversiPlaySub(cpuEna, tmpCol)
		r.mPlayLock = 0
	} else {
		r.mPlayLock = 0
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
			if r.mReversi.GetGameEndSts() == 0 {
				if r.mReversi.GetColorEna(r.mCurColor) != 0 {
					// *** パスメッセージ *** //
					r.ReversiPlayPass(r.mCurColor)
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
	if r.mGameEndSts == 0 {
		r.mGameEndSts = 1
		// 終了アニメ実行
		waitTime := r.GameEndAnimExec()
		r.mPlayLock = 1
		r.WaitLocal(waitTime)
		// *** ゲーム終了メッセージ *** //
		var tmpMsg1 string
		var tmpMsg2 string
		var msgStr string
		var blk int
		var whi int
		blk = r.mReversi.GetBetCnt(REVERSI_STS_BLACK)
		whi = r.mReversi.GetBetCnt(REVERSI_STS_WHITE)
		tmpMsg1 = "プレイヤー1 = " + strconv.Itoa(blk) + " プレイヤー2 = " + strconv.Itoa(whi)
		if r.mSetting.GetmMode() == DEF_MODE_ONE {
			if whi == blk {
				tmpMsg2 = "引き分けです。"
			} else if whi < blk {
				if r.mCurColor == REVERSI_STS_BLACK {
					tmpMsg2 = "あなたの勝ちです。"
				} else {
					tmpMsg2 = "あなたの負けです。"
				}
			} else {
				if r.mCurColor == REVERSI_STS_WHITE {
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
		if r.mSetting.GetmEndAnim() == DEF_END_ANIM_ON {
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
	if r.mSetting.GetmMode() == DEF_MODE_ONE {
		if color == r.mCurColor {
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
			pCnt := r.mReversi.GetPointCnt(color)
			pInfo := r.mReversi.GetPoint(color, rand.Intn(pCnt))
			if pInfo != nil {
				setY = pInfo.GetY()
				setX = pInfo.GetX()
				if r.mSetting.GetmType() != DEF_TYPE_EASY {
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
					loop := r.mSetting.GetmMasuCnt() * r.mSetting.GetmMasuCnt()
					pcnt := 0
					passCnt := 0
					othColor := 0
					if color == REVERSI_STS_BLACK {
						othColor = REVERSI_STS_WHITE
					} else {
						othColor = REVERSI_STS_BLACK
					}
					// 対戦相手のコマ数
					othBet := r.mReversi.GetBetCnt(othColor)
					// 自分のコマ数
					ownBet := r.mReversi.GetBetCnt(color)
					endZone := 0
					if (loop - (othBet + ownBet)) <= 16 {
						// ゲーム終盤フラグON
						endZone = 1
					}
					for i := 0; i < loop; i++ {
						r.mCpu[i].SetX(0)
						r.mCpu[i].SetY(0)
						r.mEdge[i].SetX(0)
						r.mEdge[i].SetY(0)
					}

					for i := 0; i < r.mSetting.GetmMasuCnt(); i++ {
						for j := 0; j < r.mSetting.GetmMasuCnt(); j++ {
							if r.mReversi.GetMasuStsEna(color, i, j) != 0 {
								// *** 角の一つ手前なら別のとこに格納 *** //
								if r.mReversi.GetEdgeSideOne(i, j) == 0 {
									r.mEdge[kadocnt].SetX(j)
									r.mEdge[kadocnt].SetY(i)
									kadocnt++
								} else {
									r.mCpu[rcnt1].SetX(j)
									r.mCpu[rcnt1].SetY(i)
									rcnt1++
								}
								if r.mSetting.GetmType() == DEF_TYPE_NOR {
									// *** 角に置けるなら優先的にとらせるため場所を記憶させる *** //
									if r.mReversi.GetEdgeSideZero(i, j) == 0 {
										cpuflg1 = 1
										rcnt2 = (rcnt1 - 1)
									}
									// *** 角の二つ手前も優先的にとらせるため場所を記憶させる *** //
									if cpuflg1 == 0 {
										if r.mReversi.GetEdgeSideTwo(i, j) == 0 {
											cpuflg2 = 1
											rcnt2 = (rcnt1 - 1)
										}
									}
									// *** 角の三つ手前も優先的にとらせるため場所を記憶させる *** //
									if cpuflg1 == 0 && cpuflg2 == 0 {
										if r.mReversi.GetEdgeSideThree(i, j) == 0 {
											cpuflg0 = 1
											rcnt2 = (rcnt1 - 1)
										}
									}
								}
								// *** パーフェクトゲームなら *** //
								if r.mReversi.GetMasuStsCnt(color, i, j) == othBet {
									setY = i
									setX = j
									pcnt = 1
								}
								// *** 相手をパスさせるなら *** //
								if pcnt == 0 {
									if r.mReversi.GetPassEna(color, i, j) != 0 {
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
								if r.mSetting.GetmType() == DEF_TYPE_HARD {
									tmpAnz = r.mReversi.GetPointAnz(color, r.mCpu[i].GetY(), r.mCpu[i].GetX())
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
								if r.mReversi.GetMasuStsEna(color, r.mCpu[i].GetY(), r.mCpu[i].GetX()) == 2 {
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
								if r.mSetting.GetmType() == DEF_TYPE_HARD {
									tmpAnz = r.mReversi.GetPointAnz(color, r.mEdge[i].GetY(), r.mEdge[i].GetX())
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
								if r.mReversi.GetMasuStsEna(color, r.mEdge[i].GetY(), r.mEdge[i].GetX()) == 2 {
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
								if r.mReversi.CheckEdge(color, r.mEdge[i].GetY(), r.mEdge[i].GetX()) != 0 {
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
							setY = r.mCpu[rcnt2].GetY()
							setX = r.mCpu[rcnt2].GetX()
						} else if kadocnt != 0 {
							setY = r.mEdge[rcnt2].GetY()
							setX = r.mEdge[rcnt2].GetX()
						}
					}
				}
				if r.mReversi.SetMasuSts(color, setY, setX) == 0 {
					if r.mSetting.GetmType() == DEF_TYPE_HARD {
						r.mReversi.AnalysisReversi(r.mPassEnaB, r.mPassEnaW)
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
		if r.mSetting.GetmAssist() == DEF_ASSIST_ON {
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
		for i := 0; i < r.mSetting.GetmMasuCnt(); i++ {
			for j := 0; j < r.mSetting.GetmMasuCnt(); j++ {
				r.SendDrawInfoMsg(i, j)
			}
		}
	}
	waitTime := r.mSetting.GetmPlayDrawInterVal()
	for i := 0; i < r.mSetting.GetmMasuCnt(); i++ {
		for j := 0; j < r.mSetting.GetmMasuCnt(); j++ {
			if r.mReversi.GetMasuSts(i, j) != r.mReversi.GetMasuStsOld(i, j) {
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
	r.mPassEnaB = 0
	r.mPassEnaW = 0
	if r.mSetting.GetmGameSpd() == DEF_GAME_SPD_FAST {
		r.mSetting.SetmPlayDrawInterVal(DEF_GAME_SPD_FAST_VAL) // 描画のインターバル(msec)
		r.mSetting.SetmPlayCpuInterVal(DEF_GAME_SPD_FAST_VAL2) // CPU対戦時のインターバル(msec)
	} else if r.mSetting.GetmGameSpd() == DEF_GAME_SPD_MID {
		r.mSetting.SetmPlayDrawInterVal(DEF_GAME_SPD_MID_VAL) // 描画のインターバル(msec)
		r.mSetting.SetmPlayCpuInterVal(DEF_GAME_SPD_MID_VAL2) // CPU対戦時のインターバル(msec)
	} else {
		r.mSetting.SetmPlayDrawInterVal(DEF_GAME_SPD_SLOW_VAL) // 描画のインターバル(msec)
		r.mSetting.SetmPlayCpuInterVal(DEF_GAME_SPD_SLOW_VAL2) // CPU対戦時のインターバル(msec)
	}
	r.mCurColor = r.mSetting.GetmPlayer()
	if r.mSetting.GetmMode() == DEF_MODE_TWO {
		r.mCurColor = REVERSI_STS_BLACK
	}
	r.mReversi.SetMasuCnt(r.mSetting.GetmMasuCnt()) // マスの数設定
	r.mReversi.Reset()
	if r.mSetting.GetmMode() == DEF_MODE_ONE {
		if r.mCurColor == REVERSI_STS_WHITE {
			pCnt := r.mReversi.GetPointCnt(REVERSI_STS_BLACK)
			pInfo := r.mReversi.GetPoint(REVERSI_STS_BLACK, rand.Intn(pCnt))
			if pInfo != nil {
				r.mReversi.SetMasuSts(REVERSI_STS_BLACK, pInfo.GetY(), pInfo.GetX())
				if r.mSetting.GetmType() == DEF_TYPE_HARD {
					r.mReversi.AnalysisReversi(r.mPassEnaB, r.mPassEnaW)
				}
			}
		}
	}
	r.mPlayLock = 1
	r.mGameEndSts = 0
	r.DrawUpdateForcibly(r.mSetting.GetmAssist())
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
	if r.mSetting.GetmEndAnim() == DEF_END_ANIM_ON {
		bCnt = r.mReversi.GetBetCnt(REVERSI_STS_BLACK)
		wCnt = r.mReversi.GetBetCnt(REVERSI_STS_WHITE)
		// *** 色、コマ数表示消去 *** //
		// *** メッセージ送信 *** //
		r.ExecMessage(LC_MSG_CUR_COL_ERASE, nil)
		// *** メッセージ送信 *** //
		r.ExecMessage(LC_MSG_CUR_STS_ERASE, nil)
		r.WaitLocal(r.mSetting.GetmEndInterVal())
		// *** マス消去 *** //
		for i := 0; i < r.mSetting.GetmMasuCnt(); i++ {
			for j := 0; j < r.mSetting.GetmMasuCnt(); j++ {
				r.mReversi.SetMasuStsForcibly(REVERSI_STS_NONE, i, j)
			}
		}
		// *** メッセージ送信 *** //
		r.ExecMessage(LC_MSG_ERASE_ALL, nil)
		// *** マス描画 *** //
		bCnt2 := 0
		wCnt2 := 0
		bEnd := 0
		wEnd := 0
		for i := 0; i < r.mSetting.GetmMasuCnt(); i++ {
			for j := 0; j < r.mSetting.GetmMasuCnt(); j++ {
				if bCnt2 < bCnt {
					bCnt2++
					r.mReversi.SetMasuStsForcibly(REVERSI_STS_BLACK, i, j)
					r.SendDrawMsg(i, j)
				} else {
					bEnd = 1
				}
				if wCnt2 < wCnt {
					wCnt2++
					r.mReversi.SetMasuStsForcibly(REVERSI_STS_WHITE, (r.mSetting.GetmMasuCnt()-1)-i, (r.mSetting.GetmMasuCnt()-1)-j)
					r.SendDrawMsg((r.mSetting.GetmMasuCnt()-1)-i, (r.mSetting.GetmMasuCnt()-1)-j)
				} else {
					wEnd = 1
				}
				if bEnd == 1 && wEnd == 1 {
					break
				} else {
					r.WaitLocal(r.mSetting.GetmEndDrawInterVal())
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
		dMode = r.mReversi.GetMasuSts(msgPoint.GetY(), msgPoint.GetX())
		dBack = r.mReversi.GetMasuStsEna(r.mCurColor, msgPoint.GetY(), msgPoint.GetX())
		dCnt = r.mReversi.GetMasuStsCnt(r.mCurColor, msgPoint.GetY(), msgPoint.GetX())
		r.DrawSingleLocal(msgPoint.GetY(), msgPoint.GetX(), dMode, dBack, strconv.Itoa(dCnt))
	} else if what == LC_MSG_ERASE {
		// *** マス消去 *** //
		r.DrawSingleLocal(msgPoint.GetY(), msgPoint.GetX(), 0, 0, "0")
	} else if what == LC_MSG_DRAW_INFO {
		// *** マス情報描画 *** //
		dMode = r.mReversi.GetMasuSts(msgPoint.GetY(), msgPoint.GetX())
		dBack = r.mReversi.GetMasuStsEna(r.mCurColor, msgPoint.GetY(), msgPoint.GetX())
		dCnt = r.mReversi.GetMasuStsCnt(r.mCurColor, msgPoint.GetY(), msgPoint.GetX())
		r.DrawSingleLocal(msgPoint.GetY(), msgPoint.GetX(), dMode, dBack, strconv.Itoa(dCnt))
	} else if what == LC_MSG_ERASE_INFO {
		// *** マス情報消去 *** //
		dMode = r.mReversi.GetMasuSts(msgPoint.GetY(), msgPoint.GetX())
		r.DrawSingleLocal(msgPoint.GetY(), msgPoint.GetX(), dMode, 0, "0")
	} else if what == LC_MSG_DRAW_ALL {
		// *** 全マス描画 *** //
		for i := 0; i < r.mSetting.GetmMasuCnt(); i++ {
			for j := 0; j < r.mSetting.GetmMasuCnt(); j++ {
				dMode = r.mReversi.GetMasuSts(i, j)
				dBack = r.mReversi.GetMasuStsEna(r.mCurColor, i, j)
				dCnt = r.mReversi.GetMasuStsCnt(r.mCurColor, i, j)
				r.DrawSingleLocal(i, j, dMode, dBack, strconv.Itoa(dCnt))
			}
		}
	} else if what == LC_MSG_ERASE_ALL {
		// *** 全マス消去 *** //
		for i := 0; i < r.mSetting.GetmMasuCnt(); i++ {
			for j := 0; j < r.mSetting.GetmMasuCnt(); j++ {
				r.DrawSingleLocal(i, j, 0, 0, "0")
			}
		}
	} else if what == LC_MSG_DRAW_INFO_ALL {
		// *** 全マス情報描画 *** //
		for i := 0; i < r.mSetting.GetmMasuCnt(); i++ {
			for j := 0; j < r.mSetting.GetmMasuCnt(); j++ {
				dMode = r.mReversi.GetMasuSts(i, j)
				dBack = r.mReversi.GetMasuStsEna(r.mCurColor, i, j)
				dCnt = r.mReversi.GetMasuStsCnt(r.mCurColor, i, j)
				r.DrawSingleLocal(i, j, dMode, dBack, strconv.Itoa(dCnt))
			}
		}
	} else if what == LC_MSG_ERASE_INFO_ALL {
		// *** 全マス情報消去 *** //
		for i := 0; i < r.mSetting.GetmMasuCnt(); i++ {
			for j := 0; j < r.mSetting.GetmMasuCnt(); j++ {
				dMode = r.mReversi.GetMasuSts(i, j)
				r.DrawSingleLocal(i, j, dMode, 0, "0")
			}
		}
	} else if what == LC_MSG_DRAW_END {
		r.mPlayLock = 0
	} else if what == LC_MSG_CUR_COL {
		tmpStr := ""
		if r.mSetting.GetmMode() == DEF_MODE_ONE {
			if r.mCurColor == REVERSI_STS_BLACK {
				tmpStr = "あなたはプレイヤー1です "
			} else {
				tmpStr = "あなたはプレイヤー2です "
			}
		} else {
			if r.mCurColor == REVERSI_STS_BLACK {
				tmpStr = "プレイヤー1の番です "
			} else {
				tmpStr = "プレイヤー2の番です "
			}
		}
		r.CurColMsgLocal(tmpStr)
	} else if what == LC_MSG_CUR_COL_ERASE {
		r.CurColMsgLocal("")
	} else if what == LC_MSG_CUR_STS {
		tmpStr := "プレイヤー1 = " + strconv.Itoa(r.mReversi.GetBetCnt(REVERSI_STS_BLACK)) + " プレイヤー2 = " + strconv.Itoa(r.mReversi.GetBetCnt(REVERSI_STS_WHITE))
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
	if r.mDelegate != nil {
		funcs := r.mCallbacks.GetFuncs()
		funcs = append(funcs, r.mDelegate.ViewMsgDlg(title, msg))
		r.mCallbacks.SetFuncs(funcs)
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
	if r.mDelegate != nil {
		funcs := r.mCallbacks.GetFuncs()
		funcs = append(funcs, r.mDelegate.DrawSingle(y, x, sts, bk, text))
		r.mCallbacks.SetFuncs(funcs)
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
	if r.mDelegate != nil {
		funcs := r.mCallbacks.GetFuncs()
		funcs = append(funcs, r.mDelegate.CurColMsg(text))
		r.mCallbacks.SetFuncs(funcs)
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
	if r.mDelegate != nil {
		funcs := r.mCallbacks.GetFuncs()
		funcs = append(funcs, r.mDelegate.CurStsMsg(text))
		r.mCallbacks.SetFuncs(funcs)
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
	if r.mDelegate != nil {
		funcs := r.mCallbacks.GetFuncs()
		funcs = append(funcs, r.mDelegate.Wait(time))
		r.mCallbacks.SetFuncs(funcs)
	}
}
