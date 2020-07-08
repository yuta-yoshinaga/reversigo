////////////////////////////////////////////////////////////////////////////////
///	@file			Reversi.go
///	@brief			リバーシクラス実装ファイル
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///	$Version:		$
///	$Revision:		$
///
/// (c) 2020 Yuta Yoshinaga.
///
/// - 本ソフトウェアの一部又は全てを無断で複写複製（コピー）することは、
///   著作権侵害にあたりますので、これを禁止します。
/// - 本製品の使用に起因する侵害または特許権その他権利の侵害に関しては
///   当方は一切その責任を負いません。
///
////////////////////////////////////////////////////////////////////////////////
// Package model リバーシに関する構造体をまとめるパッケージ。
package model

////////////////////////////////////////////////////////////////////////////////
///	@class		Reversi
///	@brief		リバーシクラス
///
////////////////////////////////////////////////////////////////////////////////
type Reversi struct {
	mMasuSts       [][]int           //!< マスの状態
	mMasuStsOld    [][]int           //!< 以前のマスの状態
	mMasuStsEnaB   [][]int           //!< 黒の置ける場所
	mMasuStsCntB   [][]int           //!< 黒の獲得コマ数
	mMasuStsPassB  [][]int           //!< 黒が相手をパスさせる場所
	mMasuStsAnzB   [][]*ReversiAnz   //!< 黒がその場所に置いた場合の解析結果
	mMasuPointB    []*ReversiPoint   //!< 黒の置ける場所座標一覧
	mMasuPointCntB int               //!< 黒の置ける場所座標一覧数
	mMasuBetCntB   int               //!< 黒コマ数
	mMasuStsEnaW   [][]int           //!< 白の置ける場所
	mMasuStsCntW   [][]int           //!< 白の獲得コマ数
	mMasuStsPassW  [][]int           //!< 白が相手をパスさせる場所
	mMasuStsAnzW   [][]*ReversiAnz   //!< 白がその場所に置いた場合の解析結果
	mMasuPointW    []*ReversiPoint   //!< 白の置ける場所座標一覧
	mMasuPointCntW int               //!< 白の置ける場所座標一覧数
	mMasuBetCntW   int               //!< 白コマ数
	mMasuCnt       int               //!< 縦横マス数
	mMasuCntMax    int               //!< 縦横マス最大数
	mMasuHist      []*ReversiHistory //!< 履歴
	mMasuHistCur   int               //!< 履歴現在位置
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuSts() [][]int
///	@return			mMasuSts[][] int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuSts() [][]int {
	return r.mMasuSts
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuSts(mMasuSts [][]int)
///	@param[in]		mMasuSts [][]int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuSts(mMasuSts [][]int) {
	r.mMasuSts = mMasuSts
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuStsOld() [][]int
///	@return			mMasuStsOld [][]int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuStsOld() [][]int {
	return r.mMasuStsOld
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuStsOld(mMasuStsOld [][]int)
///	@param[in]		mMasuStsOld [][]int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuStsOld(mMasuStsOld [][]int) {
	r.mMasuStsOld = mMasuStsOld
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuStsEnaB() [][]int
///	@return			mMasuStsEnaB int[][]
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuStsEnaB() [][]int {
	return r.mMasuStsEnaB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuStsEnaB(mMasuStsEnaB [][]int)
///	@param[in]		mMasuStsEnaB [][]int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuStsEnaB(mMasuStsEnaB [][]int) {
	r.mMasuStsEnaB = mMasuStsEnaB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuStsCntB() [][]int
///	@return			mMasuStsCntB [][]int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuStsCntB() [][]int {
	return r.mMasuStsCntB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuStsCntB(mMasuStsCntB [][]int)
///	@param[in]		mMasuStsCntB [][]int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuStsCntB(mMasuStsCntB [][]int) {
	r.mMasuStsCntB = mMasuStsCntB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuStsPassB() [][]int
///	@return			mMasuStsPassB [][]int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuStsPassB() [][]int {
	return r.mMasuStsPassB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuStsPassB(mMasuStsPassB [][]int)
///	@param[in]		mMasuStsPassB [][]int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuStsPassB(mMasuStsPassB [][]int) {
	r.mMasuStsPassB = mMasuStsPassB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuStsAnzB() [][]*ReversiAnz
///	@return			mMasuStsAnzB [][]*ReversiAnz
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuStsAnzB() [][]*ReversiAnz {
	return r.mMasuStsAnzB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuStsAnzB(mMasuStsAnzB [][]*ReversiAnz)
///	@param[in]		mMasuStsAnzB [][]*ReversiAnz
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuStsAnzB(mMasuStsAnzB [][]*ReversiAnz) {
	r.mMasuStsAnzB = mMasuStsAnzB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuPointB() []*ReversiPoint
///	@return			mMasuPointB []*ReversiPoint
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuPointB() []*ReversiPoint {
	return r.mMasuPointB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuPointB(mMasuPointB []*ReversiPoint)
///	@param[in]		mMasuPointB []*ReversiPoint
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuPointB(mMasuPointB []*ReversiPoint) {
	r.mMasuPointB = mMasuPointB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuPointCntB() int
///	@return			mMasuPointCntB int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuPointCntB() int {
	return r.mMasuPointCntB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuPointCntB(mMasuPointCntB int)
///	@param[in]		mMasuPointCntB int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuPointCntB(mMasuPointCntB int) {
	r.mMasuPointCntB = mMasuPointCntB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuBetCntB() int
///	@return			mMasuBetCntB int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuBetCntB() int {
	return r.mMasuBetCntB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuBetCntB(mMasuBetCntB int)
///	@param[in]		mMasuBetCntB int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuBetCntB(mMasuBetCntB int) {
	r.mMasuBetCntB = mMasuBetCntB
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuStsEnaW() [][]int
///	@return			mMasuStsEnaW [][]int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuStsEnaW() [][]int {
	return r.mMasuStsEnaW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuStsEnaW(mMasuStsEnaW [][]int)
///	@param[in]		mMasuStsEnaW [][]int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuStsEnaW(mMasuStsEnaW [][]int) {
	r.mMasuStsEnaW = mMasuStsEnaW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuStsCntW() [][]int
///	@return			mMasuStsCntW [][]int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuStsCntW() [][]int {
	return r.mMasuStsCntW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuStsCntW(mMasuStsCntW [][]int)
///	@param[in]		mMasuStsCntW [][]int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuStsCntW(mMasuStsCntW [][]int) {
	r.mMasuStsCntW = mMasuStsCntW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuStsPassW() [][]int
///	@return			mMasuStsPassW [][]int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuStsPassW() [][]int {
	return r.mMasuStsPassW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuStsPassW(mMasuStsPassW [][]int)
///	@param[in]		mMasuStsPassW [][]int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuStsPassW(mMasuStsPassW [][]int) {
	r.mMasuStsPassW = mMasuStsPassW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuStsAnzW() [][]*ReversiAnz
///	@return			mMasuStsAnzW [][]*ReversiAnz
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuStsAnzW() [][]*ReversiAnz {
	return r.mMasuStsAnzW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuStsAnzW(mMasuStsAnzW [][]*ReversiAnz)
///	@param[in]		mMasuStsAnzW [][]*ReversiAnz
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuStsAnzW(mMasuStsAnzW [][]*ReversiAnz) {
	r.mMasuStsAnzW = mMasuStsAnzW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuPointW() []*ReversiPoint
///	@return			mMasuPointW []*ReversiPoint
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuPointW() []*ReversiPoint {
	return r.mMasuPointW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuPointW(mMasuPointW []*ReversiPoint)
///	@param[in]		mMasuPointW []*ReversiPoint
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuPointW(mMasuPointW []*ReversiPoint) {
	r.mMasuPointW = mMasuPointW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuPointCntW() int
///	@return			mMasuPointCntW int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuPointCntW() int {
	return r.mMasuPointCntW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuPointCntW(mMasuPointCntW int)
///	@param[in]		mMasuPointCntW int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuPointCntW(mMasuPointCntW int) {
	r.mMasuPointCntW = mMasuPointCntW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuBetCntW() int
///	@return			mMasuBetCntW int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuBetCntW() int {
	return r.mMasuBetCntW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuBetCntW(mMasuBetCntW int)
///	@param[in]		mMasuBetCntW int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuBetCntW(mMasuBetCntW int) {
	r.mMasuBetCntW = mMasuBetCntW
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuCnt() int
///	@return			mMasuCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuCnt() int {
	return r.mMasuCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuCnt(mMasuCnt int)
///	@param[in]		mMasuCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuCnt(mMasuCnt int) {
	r.mMasuCnt = mMasuCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuCntMax() int
///	@return			mMasuCntMax int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuCntMax() int {
	return r.mMasuCntMax
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuCntMax(mMasuCntMax int)
///	@param[in]		mMasuCntMax int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuCntMax(mMasuCntMax int) {
	r.mMasuCntMax = mMasuCntMax
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuHist() []*ReversiHistory
///	@return			mMasuHist []*ReversiHistory
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuHist() []*ReversiHistory {
	return r.mMasuHist
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuHist(mMasuHist []*ReversiHistory)
///	@param[in]		mMasuHist []*ReversiHistory
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuHist(mMasuHist []*ReversiHistory) {
	r.mMasuHist = mMasuHist
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuHistCur() int
///	@return			mMasuHistCur int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetmMasuHistCur() int {
	return r.mMasuHistCur
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuHistCur(mMasuHistCur int)
///	@param[in]		mMasuHistCur int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetmMasuHistCur(mMasuHistCur int) {
	r.mMasuHistCur = mMasuHistCur
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			コンストラクタ
///	@fn				public Reversi(int masuCnt,int masuMax)
///	@param[in]		int masuCnt		縦横マス数
///	@param[in]		int masuMax		縦横マス最大数
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func NewReversi(masuCnt int, masuMax int) *Reversi {
	r := new(Reversi)
	r.mMasuCnt = masuCnt
	r.mMasuCntMax = masuMax

	r.mMasuSts = make([][]int, r.mMasuCntMax)
	r.mMasuStsOld = make([][]int, r.mMasuCntMax)
	r.mMasuStsEnaB = make([][]int, r.mMasuCntMax)
	r.mMasuStsCntB = make([][]int, r.mMasuCntMax)
	r.mMasuStsPassB = make([][]int, r.mMasuCntMax)
	for i := 0; i < r.mMasuCntMax; i++ {
		r.mMasuSts[i] = make([]int, r.mMasuCntMax)
		r.mMasuStsOld[i] = make([]int, r.mMasuCntMax)
		r.mMasuStsEnaB[i] = make([]int, r.mMasuCntMax)
		r.mMasuStsCntB[i] = make([]int, r.mMasuCntMax)
		r.mMasuStsPassB[i] = make([]int, r.mMasuCntMax)
	}
	r.mMasuStsAnzB = make([][]*ReversiAnz, r.mMasuCntMax)
	for i := 0; i < r.mMasuCntMax; i++ {
		r.mMasuStsAnzB[i] = make([]*ReversiAnz, r.mMasuCntMax)
		for j := 0; j < r.mMasuCntMax; j++ {
			r.mMasuStsAnzB[i][j] = NewReversiAnz()
		}
	}
	r.mMasuPointB = make([]*ReversiPoint, r.mMasuCntMax*r.mMasuCntMax)
	for i := 0; i < (r.mMasuCntMax * r.mMasuCntMax); i++ {
		r.mMasuPointB[i] = NewReversiPoint()
	}
	r.mMasuPointCntB = 0
	r.mMasuStsEnaW = make([][]int, r.mMasuCntMax)
	r.mMasuStsCntW = make([][]int, r.mMasuCntMax)
	r.mMasuStsPassW = make([][]int, r.mMasuCntMax)
	for i := 0; i < r.mMasuCntMax; i++ {
		r.mMasuStsEnaW[i] = make([]int, r.mMasuCntMax)
		r.mMasuStsCntW[i] = make([]int, r.mMasuCntMax)
		r.mMasuStsPassW[i] = make([]int, r.mMasuCntMax)
	}
	r.mMasuStsAnzW = make([][]*ReversiAnz, r.mMasuCntMax)
	for i := 0; i < r.mMasuCntMax; i++ {
		r.mMasuStsAnzW[i] = make([]*ReversiAnz, r.mMasuCntMax)
		for j := 0; j < r.mMasuCntMax; j++ {
			r.mMasuStsAnzW[i][j] = NewReversiAnz()
		}
	}
	r.mMasuPointW = make([]*ReversiPoint, r.mMasuCntMax*r.mMasuCntMax)
	for i := 0; i < (r.mMasuCntMax * r.mMasuCntMax); i++ {
		r.mMasuPointW[i] = NewReversiPoint()
	}
	r.mMasuPointCntW = 0
	r.mMasuBetCntB = 0
	r.mMasuBetCntW = 0
	r.mMasuHist = make([]*ReversiHistory, r.mMasuCntMax*r.mMasuCntMax)
	for i := 0; i < (r.mMasuCntMax * r.mMasuCntMax); i++ {
		r.mMasuHist[i] = NewReversiHistory()
	}
	r.mMasuHistCur = 0
	for i := 0; i < r.mMasuCntMax; i++ {
		copy(r.mMasuStsOld[i], r.mMasuSts[i])
	}
	r.Reset()
	return r
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			リセット
///	@fn				Reset()
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) Reset() {
	for i := 0; i < r.mMasuCnt; i++ {
		for j := 0; j < r.mMasuCnt; j++ {
			r.mMasuSts[i][j] = REVERSI_STS_NONE
			r.mMasuStsPassB[i][j] = 0
			r.mMasuStsAnzB[i][j].Reset()
			r.mMasuStsPassW[i][j] = 0
			r.mMasuStsAnzW[i][j].Reset()
		}
	}
	r.mMasuSts[(r.mMasuCnt>>1)-1][(r.mMasuCnt>>1)-1] = REVERSI_STS_BLACK
	r.mMasuSts[(r.mMasuCnt>>1)-1][(r.mMasuCnt >> 1)] = REVERSI_STS_WHITE
	r.mMasuSts[(r.mMasuCnt >> 1)][(r.mMasuCnt>>1)-1] = REVERSI_STS_WHITE
	r.mMasuSts[(r.mMasuCnt >> 1)][(r.mMasuCnt >> 1)] = REVERSI_STS_BLACK
	r.makeMasuSts(REVERSI_STS_BLACK)
	r.makeMasuSts(REVERSI_STS_WHITE)
	r.mMasuHistCur = 0
	for i := 0; i < r.mMasuCntMax; i++ {
		copy(r.mMasuStsOld[i], r.mMasuSts[i])
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			各コマの置ける場所等のステータス作成
///	@fn				makeMasuSts(color int) int
///	@param[in]		color int		ステータスを作成するコマ
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) makeMasuSts(color int) int {
	var flg int
	var okflg int = 0
	var cnt1 int
	var cnt2 int
	var count1 int
	var count2 int = 0
	var ret int = -1
	var countMax int = 0
	var loop int
	for i := 0; i < r.mMasuCnt; i++ {
		// 初期化
		for j := 0; j < r.mMasuCnt; j++ {
			if color == REVERSI_STS_BLACK {
				r.mMasuStsEnaB[i][j] = 0
				r.mMasuStsCntB[i][j] = 0
			} else {
				r.mMasuStsEnaW[i][j] = 0
				r.mMasuStsCntW[i][j] = 0
			}
		}
	}

	loop = r.mMasuCnt * r.mMasuCnt
	for i := 0; i < loop; i++ {
		// 初期化
		if color == REVERSI_STS_BLACK {
			r.mMasuPointB[i].SetX(0)
			r.mMasuPointB[i].SetY(0)
		} else {
			r.mMasuPointW[i].SetX(0)
			r.mMasuPointW[i].SetY(0)
		}
	}
	if color == REVERSI_STS_BLACK {
		r.mMasuPointCntB = 0
	} else {
		r.mMasuPointCntW = 0
	}
	r.mMasuBetCntB = 0
	r.mMasuBetCntW = 0

	for i := 0; i < r.mMasuCnt; i++ {
		for j := 0; j < r.mMasuCnt; j++ {
			okflg = 0
			count2 = 0
			if r.mMasuSts[i][j] == REVERSI_STS_NONE {
				// 何も置かれていないマスなら
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 上方向を調べる *** //
				for (cnt1 > 0) && (r.mMasuSts[cnt1-1][j] != REVERSI_STS_NONE && r.mMasuSts[cnt1-1][j] != color) {
					flg = 1
					cnt1--
					count1++
				}
				if (cnt1 > 0) && (flg == 1) && (r.mMasuSts[cnt1-1][j] == color) {
					okflg = 1
					count2 += count1
				}
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 下方向を調べる *** //
				for (cnt1 < (r.mMasuCnt - 1)) && (r.mMasuSts[cnt1+1][j] != REVERSI_STS_NONE && r.mMasuSts[cnt1+1][j] != color) {
					flg = 1
					cnt1++
					count1++
				}
				if (cnt1 < (r.mMasuCnt - 1)) && (flg == 1) && (r.mMasuSts[cnt1+1][j] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				count1 = 0
				flg = 0
				// *** 右方向を調べる *** //
				for (cnt2 < (r.mMasuCnt - 1)) && (r.mMasuSts[i][cnt2+1] != REVERSI_STS_NONE && r.mMasuSts[i][cnt2+1] != color) {
					flg = 1
					cnt2++
					count1++
				}
				if (cnt2 < (r.mMasuCnt - 1)) && (flg == 1) && (r.mMasuSts[i][cnt2+1] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				count1 = 0
				flg = 0
				// *** 左方向を調べる *** //
				for (cnt2 > 0) && (r.mMasuSts[i][cnt2-1] != REVERSI_STS_NONE && r.mMasuSts[i][cnt2-1] != color) {
					flg = 1
					cnt2--
					count1++
				}
				if (cnt2 > 0) && (flg == 1) && (r.mMasuSts[i][cnt2-1] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 右上方向を調べる *** //
				for ((cnt2 < (r.mMasuCnt - 1)) && (cnt1 > 0)) && (r.mMasuSts[cnt1-1][cnt2+1] != REVERSI_STS_NONE && r.mMasuSts[cnt1-1][cnt2+1] != color) {
					flg = 1
					cnt1--
					cnt2++
					count1++
				}
				if ((cnt2 < (r.mMasuCnt - 1)) && (cnt1 > 0)) && (flg == 1) && (r.mMasuSts[cnt1-1][cnt2+1] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 左上方向を調べる *** //
				for ((cnt2 > 0) && (cnt1 > 0)) && (r.mMasuSts[cnt1-1][cnt2-1] != REVERSI_STS_NONE && r.mMasuSts[cnt1-1][cnt2-1] != color) {
					flg = 1
					cnt1--
					cnt2--
					count1++
				}
				if ((cnt2 > 0) && (cnt1 > 0)) && (flg == 1) && (r.mMasuSts[cnt1-1][cnt2-1] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 右下方向を調べる *** //
				for ((cnt2 < (r.mMasuCnt - 1)) && (cnt1 < (r.mMasuCnt - 1))) && (r.mMasuSts[cnt1+1][cnt2+1] != REVERSI_STS_NONE && r.mMasuSts[cnt1+1][cnt2+1] != color) {
					flg = 1
					cnt1++
					cnt2++
					count1++
				}
				if ((cnt2 < (r.mMasuCnt - 1)) && (cnt1 < (r.mMasuCnt - 1))) && (flg == 1) && (r.mMasuSts[cnt1+1][cnt2+1] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 左下方向を調べる *** //
				for ((cnt2 > 0) && (cnt1 < (r.mMasuCnt - 1))) && (r.mMasuSts[cnt1+1][cnt2-1] != REVERSI_STS_NONE && r.mMasuSts[cnt1+1][cnt2-1] != color) {
					flg = 1
					cnt1++
					cnt2--
					count1++
				}
				if ((cnt2 > 0) && (cnt1 < (r.mMasuCnt - 1))) && (flg == 1) && (r.mMasuSts[cnt1+1][cnt2-1] == color) {
					okflg = 1
					count2 += count1
				}
				if okflg == 1 {
					if color == REVERSI_STS_BLACK {
						r.mMasuStsEnaB[i][j] = 1
						r.mMasuStsCntB[i][j] = count2
						// *** 置ける場所をリニアに保存、置けるポイント数も保存 *** //
						r.mMasuPointB[r.mMasuPointCntB].SetY(i)
						r.mMasuPointB[r.mMasuPointCntB].SetX(j)
						r.mMasuPointCntB++
					} else {
						r.mMasuStsEnaW[i][j] = 1
						r.mMasuStsCntW[i][j] = count2
						// *** 置ける場所をリニアに保存、置けるポイント数も保存 *** //
						r.mMasuPointW[r.mMasuPointCntW].SetY(i)
						r.mMasuPointW[r.mMasuPointCntW].SetX(j)
						r.mMasuPointCntW++
					}
					ret = 0
					if countMax < count2 {
						countMax = count2
					}
				}
			} else if r.mMasuSts[i][j] == REVERSI_STS_BLACK {
				r.mMasuBetCntB++
			} else if r.mMasuSts[i][j] == REVERSI_STS_WHITE {
				r.mMasuBetCntW++
			}
		}
	}

	// *** 一番枚数を獲得できるマスを設定 *** //
	for i := 0; i < r.mMasuCnt; i++ {
		for j := 0; j < r.mMasuCnt; j++ {
			if color == REVERSI_STS_BLACK {
				if r.mMasuStsEnaB[i][j] != 0 && r.mMasuStsCntB[i][j] == countMax {
					r.mMasuStsEnaB[i][j] = 2
				}
			} else {
				if r.mMasuStsEnaW[i][j] != 0 && r.mMasuStsCntW[i][j] == countMax {
					r.mMasuStsEnaW[i][j] = 2
				}
			}
		}
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			コマをひっくり返す
///	@fn				revMasuSts(color int,y int,x int)
///	@param[in]		color int		ひっくり返す元コマ
///	@param[in]		y int			ひっくり返す元コマのY座標
///	@param[in]		x int			ひっくり返す元コマのX座標
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) revMasuSts(color int, y int, x int) {
	var cnt1 int
	var cnt2 int
	var rcnt1 int
	var rcnt2 int
	var flg int = 0

	// *** 左方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt1 > 0; {
		if r.mMasuSts[cnt2][cnt1-1] != REVERSI_STS_NONE && r.mMasuSts[cnt2][cnt1-1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt1--
		} else if r.mMasuSts[cnt2][cnt1-1] == color {
			flg = 1
			break
		} else if r.mMasuSts[cnt2][cnt1-1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		for rcnt1 = cnt1; rcnt1 < x; rcnt1++ {
			r.mMasuSts[cnt2][rcnt1] = color
		}
	}

	// *** 右方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt1 < (r.mMasuCnt - 1); {
		if r.mMasuSts[cnt2][cnt1+1] != REVERSI_STS_NONE && r.mMasuSts[cnt2][cnt1+1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt1++
		} else if r.mMasuSts[cnt2][cnt1+1] == color {
			flg = 1
			break
		} else if r.mMasuSts[cnt2][cnt1+1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		for rcnt1 = cnt1; rcnt1 > x; rcnt1-- {
			r.mMasuSts[cnt2][rcnt1] = color
		}
	}

	// *** 上方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 > 0; {
		if r.mMasuSts[cnt2-1][cnt1] != REVERSI_STS_NONE && r.mMasuSts[cnt2-1][cnt1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2--
		} else if r.mMasuSts[cnt2-1][cnt1] == color {
			flg = 1
			break
		} else if r.mMasuSts[cnt2-1][cnt1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		for rcnt1 = cnt2; rcnt1 < y; rcnt1++ {
			r.mMasuSts[rcnt1][cnt1] = color
		}
	}

	// *** 下方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 < (r.mMasuCnt - 1); {
		if r.mMasuSts[cnt2+1][cnt1] != REVERSI_STS_NONE && r.mMasuSts[cnt2+1][cnt1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2++
		} else if r.mMasuSts[cnt2+1][cnt1] == color {
			flg = 1
			break
		} else if r.mMasuSts[cnt2+1][cnt1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		for rcnt1 = cnt2; rcnt1 > y; rcnt1-- {
			r.mMasuSts[rcnt1][cnt1] = color
		}
	}

	// *** 左上方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 > 0 && cnt1 > 0; {
		if r.mMasuSts[cnt2-1][cnt1-1] != REVERSI_STS_NONE && r.mMasuSts[cnt2-1][cnt1-1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2--
			cnt1--
		} else if r.mMasuSts[cnt2-1][cnt1-1] == color {
			flg = 1
			break
		} else if r.mMasuSts[cnt2-1][cnt1-1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		rcnt2 = cnt1
		for rcnt1 = cnt2; (rcnt1 < y) && (rcnt2 < x); rcnt1++ {
			r.mMasuSts[rcnt1][rcnt2] = color
			rcnt2++
		}
	}

	// *** 左下方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 < (r.mMasuCnt-1) && cnt1 > 0; {
		if r.mMasuSts[cnt2+1][cnt1-1] != REVERSI_STS_NONE && r.mMasuSts[cnt2+1][cnt1-1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2++
			cnt1--
		} else if r.mMasuSts[cnt2+1][cnt1-1] == color {
			flg = 1
			break
		} else if r.mMasuSts[cnt2+1][cnt1-1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		rcnt2 = cnt1
		for rcnt1 = cnt2; (rcnt1 > y) && (rcnt2 < x); rcnt1-- {
			r.mMasuSts[rcnt1][rcnt2] = color
			rcnt2++
		}
	}

	// *** 右上方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 > 0 && cnt1 < (r.mMasuCnt-1); {
		if r.mMasuSts[cnt2-1][cnt1+1] != REVERSI_STS_NONE && r.mMasuSts[cnt2-1][cnt1+1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2--
			cnt1++
		} else if r.mMasuSts[cnt2-1][cnt1+1] == color {
			flg = 1
			break
		} else if r.mMasuSts[cnt2-1][cnt1+1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		rcnt2 = cnt1
		for rcnt1 = cnt2; (rcnt1 < y) && (rcnt2 > x); rcnt1++ {
			r.mMasuSts[rcnt1][rcnt2] = color
			rcnt2--
		}
	}

	// *** 右下方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 < (r.mMasuCnt-1) && cnt1 < (r.mMasuCnt-1); {
		if r.mMasuSts[cnt2+1][cnt1+1] != REVERSI_STS_NONE && r.mMasuSts[cnt2+1][cnt1+1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2++
			cnt1++
		} else if r.mMasuSts[cnt2+1][cnt1+1] == color {
			flg = 1
			break
		} else if r.mMasuSts[cnt2+1][cnt1+1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		rcnt2 = cnt1
		for rcnt1 = cnt2; (rcnt1 > y) && (rcnt2 > x); rcnt1-- {
			r.mMasuSts[rcnt1][rcnt2] = color
			rcnt2--
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			パラメーター範囲チェック
///	@fn				checkPara(para int,min int,max int) int
///	@param[in]		para int		チェックパラメーター
///	@param[in]		min int			パラメーター最小値
///	@param[in]		max int			パラメーター最大値
///	@return			0 : 成功 それ以外 : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) checkPara(para int, min int, max int) int {
	ret := -1
	if min <= para && para <= max {
		ret = 0
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			解析を行う(黒)
///	@fn				analysisReversiBlack()
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) analysisReversiBlack() {
	var tmpX int
	var tmpY int
	var sum int
	var sumOwn int
	var tmpGoodPoint int
	var tmpBadPoint int
	var tmpD1 float64
	var tmpD2 float64
	for cnt := 0; cnt < r.mMasuPointCntB; cnt++ {
		// *** 現在ステータスを一旦コピー *** //
		tmpMasu := make([][]int, r.mMasuCntMax)
		tmpMasuEnaB := make([][]int, r.mMasuCntMax)
		tmpMasuEnaW := make([][]int, r.mMasuCntMax)
		for i := 0; i < r.mMasuCntMax; i++ {
			tmpMasu[i] = make([]int, r.mMasuCntMax)
			tmpMasuEnaB[i] = make([]int, r.mMasuCntMax)
			tmpMasuEnaW[i] = make([]int, r.mMasuCntMax)
			copy(tmpMasu[i], r.mMasuSts[i])
			copy(tmpMasuEnaB[i], r.mMasuStsEnaB[i])
			copy(tmpMasuEnaW[i], r.mMasuStsEnaW[i])
		}
		tmpY = r.mMasuPointB[cnt].GetY()
		tmpX = r.mMasuPointB[cnt].GetX()
		// 仮に置く
		r.mMasuSts[tmpY][tmpX] = REVERSI_STS_BLACK
		// 仮にひっくり返す
		r.revMasuSts(REVERSI_STS_BLACK, tmpY, tmpX)
		r.makeMasuSts(REVERSI_STS_BLACK)
		r.makeMasuSts(REVERSI_STS_WHITE)
		// *** このマスに置いた場合の解析を行う *** //
		if r.GetColorEna(REVERSI_STS_WHITE) != 0 {
			// 相手がパス
			r.mMasuStsPassB[tmpY][tmpX] = 1
		}
		if r.GetEdgeSideZero(tmpY, tmpX) == 0 {
			// 置いた場所が角
			r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeCnt() + 1)
			r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 10000*r.mMasuStsCntB[tmpY][tmpX])
		} else if r.GetEdgeSideOne(tmpY, tmpX) == 0 {
			// 置いた場所が角の一つ手前
			r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
			if r.CheckEdge(REVERSI_STS_BLACK, tmpY, tmpX) != 0 {
				// 角を取られない
				r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 10*r.mMasuStsCntB[tmpY][tmpX])
			} else {
				// 角を取られる
				r.mMasuStsAnzB[tmpY][tmpX].SetBadPoint(r.mMasuStsAnzB[tmpY][tmpX].GetBadPoint() + 100000)
			}
		} else if r.GetEdgeSideTwo(tmpY, tmpX) == 0 {
			// 置いた場所が角の二つ手前
			r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
			r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 1000*r.mMasuStsCntB[tmpY][tmpX])
		} else if r.GetEdgeSideThree(tmpY, tmpX) == 0 {
			// 置いた場所が角の三つ手前
			r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideThreeCnt() + 1)
			r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 100*r.mMasuStsCntB[tmpY][tmpX])
		} else {
			// 置いた場所がその他
			r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideOtherCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideOtherCnt() + 1)
			r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 10*r.mMasuStsCntB[tmpY][tmpX])
		}
		sum = 0
		sumOwn = 0
		for i := 0; i < r.mMasuCnt; i++ {
			for j := 0; j < r.mMasuCnt; j++ {
				tmpBadPoint = 0
				tmpGoodPoint = 0
				if r.GetMasuStsEna(REVERSI_STS_WHITE, i, j) != 0 {
					// 相手の獲得予定枚数
					sum += r.mMasuStsCntW[i][j]
					// *** 相手の獲得予定の最大数保持 *** //
					if r.mMasuStsAnzB[tmpY][tmpX].GetMax() < r.mMasuStsCntW[i][j] {
						r.mMasuStsAnzB[tmpY][tmpX].SetMax(r.mMasuStsCntW[i][j])
					}
					// *** 相手の獲得予定の最小数保持 *** //
					if r.mMasuStsCntW[i][j] < r.mMasuStsAnzB[tmpY][tmpX].GetMin() {
						r.mMasuStsAnzB[tmpY][tmpX].SetMin(r.mMasuStsCntW[i][j])
					}
					// 相手の置ける場所の数
					r.mMasuStsAnzB[tmpY][tmpX].SetPointCnt(r.mMasuStsAnzB[tmpY][tmpX].GetPointCnt() + 1)
					if r.GetEdgeSideZero(i, j) == 0 {
						// 置く場所が角
						r.mMasuStsAnzB[tmpY][tmpX].SetEdgeCnt(r.mMasuStsAnzB[tmpY][tmpX].GetEdgeCnt() + 1)
						tmpBadPoint = 100000 * r.mMasuStsCntW[i][j]
					} else if r.GetEdgeSideOne(i, j) == 0 {
						// 置く場所が角の一つ手前
						r.mMasuStsAnzB[tmpY][tmpX].SetEdgeSideOneCnt(r.mMasuStsAnzB[tmpY][tmpX].GetEdgeSideOneCnt() + 1)
						tmpBadPoint = 0
					} else if r.GetEdgeSideTwo(i, j) == 0 {
						// 置く場所が角の二つ手前
						r.mMasuStsAnzB[tmpY][tmpX].SetEdgeSideTwoCnt(r.mMasuStsAnzB[tmpY][tmpX].GetEdgeSideTwoCnt() + 1)
						tmpBadPoint = 1 * r.mMasuStsCntW[i][j]
					} else if r.GetEdgeSideThree(i, j) == 0 {
						// 置く場所が角の三つ手前
						r.mMasuStsAnzB[tmpY][tmpX].SetEdgeSideThreeCnt(r.mMasuStsAnzB[tmpY][tmpX].GetEdgeSideThreeCnt() + 1)
						tmpBadPoint = 1 * r.mMasuStsCntW[i][j]
					} else {
						// 置く場所がその他
						r.mMasuStsAnzB[tmpY][tmpX].SetEdgeSideOtherCnt(r.mMasuStsAnzB[tmpY][tmpX].GetEdgeSideOtherCnt() + 1)
						tmpBadPoint = 1 * r.mMasuStsCntW[i][j]
					}
					if tmpMasuEnaW[i][j] != 0 {
						// ステータス変化していないなら
						tmpBadPoint = 0
					}
				}
				if r.GetMasuStsEna(REVERSI_STS_BLACK, i, j) != 0 {
					// 自分の獲得予定枚数
					sumOwn += r.mMasuStsCntB[i][j]
					// *** 自分の獲得予定の最大数保持 *** //
					if r.mMasuStsAnzB[tmpY][tmpX].GetOwnMax() < r.mMasuStsCntB[i][j] {
						r.mMasuStsAnzB[tmpY][tmpX].SetOwnMax(r.mMasuStsCntB[i][j])
					}
					// *** 自分の獲得予定の最小数保持 *** //
					if r.mMasuStsCntB[i][j] < r.mMasuStsAnzB[tmpY][tmpX].GetOwnMin() {
						r.mMasuStsAnzB[tmpY][tmpX].SetOwnMin(r.mMasuStsCntB[i][j])
					}
					// 自分の置ける場所の数
					r.mMasuStsAnzB[tmpY][tmpX].SetOwnPointCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnPointCnt() + 1)
					if r.GetEdgeSideZero(i, j) == 0 {
						// 置く場所が角
						r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeCnt() + 1)
						tmpGoodPoint = 100 * r.mMasuStsCntB[i][j]
					} else if r.GetEdgeSideOne(i, j) == 0 {
						// 置く場所が角の一つ手前
						r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
						tmpGoodPoint = 0
					} else if r.GetEdgeSideTwo(i, j) == 0 {
						// 置く場所が角の二つ手前
						r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
						tmpGoodPoint = 3 * r.mMasuStsCntB[i][j]
					} else if r.GetEdgeSideThree(i, j) == 0 {
						// 置く場所が角の三つ手前
						r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideThreeCnt() + 1)
						tmpGoodPoint = 2 * r.mMasuStsCntB[i][j]
					} else {
						// 置く場所がその他
						r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideOtherCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideOtherCnt() + 1)
						tmpGoodPoint = 1 * r.mMasuStsCntB[i][j]
					}
					if tmpMasuEnaB[i][j] != 0 {
						// ステータス変化していないなら
						tmpGoodPoint = 0
					}
				}
				if tmpBadPoint != 0 {
					r.mMasuStsAnzB[tmpY][tmpX].SetBadPoint(r.mMasuStsAnzB[tmpY][tmpX].GetBadPoint() + tmpBadPoint)
				}
				if tmpGoodPoint != 0 {
					r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + tmpGoodPoint)
				}
			}
		}
		// *** 相手に取られる平均コマ数 *** //
		if r.GetPointCnt(REVERSI_STS_WHITE) != 0 {
			tmpD1 = float64(sum)
			tmpD2 = float64(r.GetPointCnt(REVERSI_STS_WHITE))
			r.mMasuStsAnzB[tmpY][tmpX].SetAvg(tmpD1 / tmpD2)
		}
		// *** 自分が取れる平均コマ数 *** //
		if r.GetPointCnt(REVERSI_STS_BLACK) != 0 {
			tmpD1 = float64(sumOwn)
			tmpD2 = float64(r.GetPointCnt(REVERSI_STS_BLACK))
			r.mMasuStsAnzB[tmpY][tmpX].SetOwnAvg(tmpD1 / tmpD2)
		}
		// *** 元に戻す *** //
		for i := 0; i < r.mMasuCntMax; i++ {
			copy(r.mMasuSts[i], tmpMasu[i])
		}
		r.makeMasuSts(REVERSI_STS_BLACK)
		r.makeMasuSts(REVERSI_STS_WHITE)
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			解析を行う(白)
///	@fn				analysisReversiWhite()
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) analysisReversiWhite() {
	var tmpX int
	var tmpY int
	var sum int
	var sumOwn int
	var tmpGoodPoint int
	var tmpBadPoint int
	var tmpD1 float64
	var tmpD2 float64
	for cnt := 0; cnt < r.mMasuPointCntW; cnt++ {
		// *** 現在ステータスを一旦コピー *** //
		tmpMasu := make([][]int, r.mMasuCntMax)
		tmpMasuEnaB := make([][]int, r.mMasuCntMax)
		tmpMasuEnaW := make([][]int, r.mMasuCntMax)
		for i := 0; i < r.mMasuCntMax; i++ {
			tmpMasu[i] = make([]int, r.mMasuCntMax)
			tmpMasuEnaB[i] = make([]int, r.mMasuCntMax)
			tmpMasuEnaW[i] = make([]int, r.mMasuCntMax)
			copy(tmpMasu[i], r.mMasuSts[i])
			copy(tmpMasuEnaB[i], r.mMasuStsEnaB[i])
			copy(tmpMasuEnaW[i], r.mMasuStsEnaW[i])
		}
		tmpY = r.mMasuPointW[cnt].GetY()
		tmpX = r.mMasuPointW[cnt].GetX()
		// 仮に置く
		r.mMasuSts[tmpY][tmpX] = REVERSI_STS_WHITE
		// 仮にひっくり返す
		r.revMasuSts(REVERSI_STS_WHITE, tmpY, tmpX)
		r.makeMasuSts(REVERSI_STS_BLACK)
		r.makeMasuSts(REVERSI_STS_WHITE)
		// *** このマスに置いた場合の解析を行う *** //
		if r.GetColorEna(REVERSI_STS_BLACK) != 0 {
			// 相手がパス
			r.mMasuStsPassW[tmpY][tmpX] = 1
		}
		if r.GetEdgeSideZero(tmpY, tmpX) == 0 {
			// 置いた場所が角
			r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeCnt() + 1)
			r.mMasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 10000*r.mMasuStsCntW[tmpY][tmpX])
		} else if r.GetEdgeSideOne(tmpY, tmpX) == 0 {
			// 置いた場所が角の一つ手前
			r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
			if r.CheckEdge(REVERSI_STS_WHITE, tmpY, tmpX) != 0 {
				// 角を取られない
				r.mMasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 10*r.mMasuStsCntW[tmpY][tmpX])
			} else {
				// 角を取られる
				r.mMasuStsAnzW[tmpY][tmpX].SetBadPoint(r.mMasuStsAnzW[tmpY][tmpX].GetBadPoint() + 100000)
			}
		} else if r.GetEdgeSideTwo(tmpY, tmpX) == 0 {
			// 置いた場所が角の二つ手前
			r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
			r.mMasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 1000*r.mMasuStsCntW[tmpY][tmpX])
		} else if r.GetEdgeSideThree(tmpY, tmpX) == 0 {
			// 置いた場所が角の三つ手前
			r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideThreeCnt() + 1)
			r.mMasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 100*r.mMasuStsCntW[tmpY][tmpX])
		} else {
			// 置いた場所がその他
			r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOtherCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOtherCnt() + 1)
			r.mMasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 10*r.mMasuStsCntW[tmpY][tmpX])
		}
		sum = 0
		sumOwn = 0
		for i := 0; i < r.mMasuCnt; i++ {
			for j := 0; j < r.mMasuCnt; j++ {
				tmpBadPoint = 0
				tmpGoodPoint = 0
				if r.GetMasuStsEna(REVERSI_STS_BLACK, i, j) != 0 {
					// 相手の獲得予定枚数
					sum += r.mMasuStsCntB[i][j]
					// *** 相手の獲得予定の最大数保持 *** //
					if r.mMasuStsAnzW[tmpY][tmpX].GetMax() < r.mMasuStsCntB[i][j] {
						r.mMasuStsAnzW[tmpY][tmpX].SetMax(r.mMasuStsCntB[i][j])
					}
					// *** 相手の獲得予定の最小数保持 *** //
					if r.mMasuStsCntB[i][j] < r.mMasuStsAnzW[tmpY][tmpX].GetMin() {
						r.mMasuStsAnzW[tmpY][tmpX].SetMin(r.mMasuStsCntB[i][j])
					}
					// 相手の置ける場所の数
					r.mMasuStsAnzW[tmpY][tmpX].SetPointCnt(r.mMasuStsAnzW[tmpY][tmpX].GetPointCnt() + 1)
					if r.GetEdgeSideZero(i, j) == 0 {
						// 置く場所が角
						r.mMasuStsAnzW[tmpY][tmpX].SetEdgeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetEdgeCnt() + 1)
						tmpBadPoint = 100000 * r.mMasuStsCntB[i][j]
					} else if r.GetEdgeSideOne(i, j) == 0 {
						// 置く場所が角の一つ手前
						r.mMasuStsAnzW[tmpY][tmpX].SetEdgeSideOneCnt(r.mMasuStsAnzW[tmpY][tmpX].GetEdgeSideOneCnt() + 1)
						tmpBadPoint = 0
					} else if r.GetEdgeSideTwo(i, j) == 0 {
						// 置く場所が角の二つ手前
						r.mMasuStsAnzW[tmpY][tmpX].SetEdgeSideTwoCnt(r.mMasuStsAnzW[tmpY][tmpX].GetEdgeSideTwoCnt() + 1)
						tmpBadPoint = 1 * r.mMasuStsCntB[i][j]
					} else if r.GetEdgeSideThree(i, j) == 0 {
						// 置く場所が角の三つ手前
						r.mMasuStsAnzW[tmpY][tmpX].SetEdgeSideThreeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetEdgeSideThreeCnt() + 1)
						tmpBadPoint = 1 * r.mMasuStsCntB[i][j]
					} else {
						// 置く場所がその他
						r.mMasuStsAnzW[tmpY][tmpX].SetEdgeSideOtherCnt(r.mMasuStsAnzW[tmpY][tmpX].GetEdgeSideOtherCnt() + 1)
						tmpBadPoint = 1 * r.mMasuStsCntB[i][j]
					}
					if tmpMasuEnaB[i][j] != 0 {
						// ステータス変化していないなら
						tmpBadPoint = 0
					}
				}
				if r.GetMasuStsEna(REVERSI_STS_WHITE, i, j) != 0 {
					// 自分の獲得予定枚数
					sumOwn += r.mMasuStsCntW[i][j]
					// *** 自分の獲得予定の最大数保持 *** //
					if r.mMasuStsAnzW[tmpY][tmpX].GetOwnMax() < r.mMasuStsCntW[i][j] {
						r.mMasuStsAnzW[tmpY][tmpX].SetOwnMax(r.mMasuStsCntW[i][j])
					}
					// *** 自分の獲得予定の最小数保持 *** //
					if r.mMasuStsCntW[i][j] < r.mMasuStsAnzW[tmpY][tmpX].GetOwnMin() {
						r.mMasuStsAnzW[tmpY][tmpX].SetOwnMin(r.mMasuStsCntW[i][j])
					}
					// 自分の置ける場所の数
					r.mMasuStsAnzW[tmpY][tmpX].SetOwnPointCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnPointCnt() + 1)
					if r.GetEdgeSideZero(i, j) == 0 {
						// 置く場所が角
						r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeCnt() + 1)
						tmpGoodPoint = 100 * r.mMasuStsCntW[i][j]
					} else if r.GetEdgeSideOne(i, j) == 0 {
						// 置く場所が角の一つ手前
						r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
						tmpGoodPoint = 0
					} else if r.GetEdgeSideTwo(i, j) == 0 {
						// 置く場所が角の二つ手前
						r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
						tmpGoodPoint = 3 * r.mMasuStsCntW[i][j]
					} else if r.GetEdgeSideThree(i, j) == 0 {
						// 置く場所が角の三つ手前
						r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideThreeCnt() + 1)
						tmpGoodPoint = 2 * r.mMasuStsCntW[i][j]
					} else {
						// 置く場所がその他
						r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOtherCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOtherCnt() + 1)
						tmpGoodPoint = 1 * r.mMasuStsCntW[i][j]
					}
					if tmpMasuEnaW[i][j] != 0 {
						// ステータス変化していないなら
						tmpGoodPoint = 0
					}
				}
				if tmpBadPoint != 0 {
					r.mMasuStsAnzW[tmpY][tmpX].SetBadPoint(r.mMasuStsAnzW[tmpY][tmpX].GetBadPoint() + tmpBadPoint)
				}
				if tmpGoodPoint != 0 {
					r.mMasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzW[tmpY][tmpX].GetGoodPoint() + tmpGoodPoint)
				}
			}
		}
		// *** 相手に取られる平均コマ数 *** //
		if r.GetPointCnt(REVERSI_STS_BLACK) != 0 {
			tmpD1 = float64(sum)
			tmpD2 = float64(r.GetPointCnt(REVERSI_STS_BLACK))
			r.mMasuStsAnzW[tmpY][tmpX].SetAvg(tmpD1 / tmpD2)
		}
		// *** 自分が取れる平均コマ数 *** //
		if r.GetPointCnt(REVERSI_STS_WHITE) != 0 {
			tmpD1 = float64(sumOwn)
			tmpD2 = float64(r.GetPointCnt(REVERSI_STS_WHITE))
			r.mMasuStsAnzW[tmpY][tmpX].SetOwnAvg(tmpD1 / tmpD2)
		}
		// *** 元に戻す *** //
		for i := 0; i < r.mMasuCntMax; i++ {
			copy(r.mMasuSts[i], tmpMasu[i])
		}
		r.makeMasuSts(REVERSI_STS_BLACK)
		r.makeMasuSts(REVERSI_STS_WHITE)
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			解析を行う
///	@fn				AnalysisReversi bPassEna int,wPassEna int)
///	@param[in]		bPassEna int		1=黒パス有効
///	@param[in]		wPassEna int		1=白パス有効
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) AnalysisReversi(bPassEna int, wPassEna int) {
	// *** 相手をパスさせることができるマス検索 *** //
	for i := 0; i < r.mMasuCnt; i++ {
		// 初期化
		for j := 0; j < r.mMasuCnt; j++ {
			r.mMasuStsPassB[i][j] = 0
			r.mMasuStsAnzB[i][j].Reset()
			r.mMasuStsPassW[i][j] = 0
			r.mMasuStsAnzW[i][j].Reset()
		}
	}
	// 黒解析
	r.analysisReversiBlack()
	// 白解析
	r.analysisReversiWhite()

	r.makeMasuSts(REVERSI_STS_BLACK)
	r.makeMasuSts(REVERSI_STS_WHITE)

	// *** パスマスを取得 *** //
	for i := 0; i < r.mMasuCnt; i++ {
		for j := 0; j < r.mMasuCnt; j++ {
			if r.mMasuStsPassB[i][j] != 0 {
				if bPassEna != 0 {
					r.mMasuStsEnaB[i][j] = 3
				}
			}
			if r.mMasuStsPassW[i][j] != 0 {
				if wPassEna != 0 {
					r.mMasuStsEnaW[i][j] = 3
				}
			}
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			マスステータスを取得
///	@fn				GetMasuSts(y int,x int) int
///	@param[in]		y int			取得するマスのY座標
///	@param[in]		x int			取得するマスのX座標
///	@return			-1 : 失敗 それ以外 : マスステータス
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetMasuSts(y int, x int) int {
	ret := -1
	if r.checkPara(y, 0, r.mMasuCnt) == 0 && r.checkPara(x, 0, r.mMasuCnt) == 0 {
		ret = r.mMasuSts[y][x]
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			以前のマスステータスを取得
///	@fn				GetMasuStsOld(y int,x int) int
///	@param[in]		y int			取得するマスのY座標
///	@param[in]		x int			取得するマスのX座標
///	@return			-1 : 失敗 それ以外 : マスステータス
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetMasuStsOld(y int, x int) int {
	ret := -1
	if r.checkPara(y, 0, r.mMasuCnt) == 0 && r.checkPara(x, 0, r.mMasuCnt) == 0 {
		ret = r.mMasuStsOld[y][x]
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			指定座標に指定色を置けるかチェック
///	@fn				GetMasuStsEna(color int,y int,x int) int
///	@param[in]		color int		コマ色
///	@param[in]		y int			マスのY座標
///	@param[in]		x int			マスのX座標
///	@return			1 : 成功 それ以外 : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetMasuStsEna(color int, y int, x int) int {
	ret := 0
	if r.checkPara(y, 0, r.mMasuCnt) == 0 && r.checkPara(x, 0, r.mMasuCnt) == 0 {
		if color == REVERSI_STS_BLACK {
			ret = r.mMasuStsEnaB[y][x]
		} else {
			ret = r.mMasuStsEnaW[y][x]
		}
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			指定座標の獲得コマ数取得
///	@fn				GetMasuStsCnt(color int,y int,x int) int
///	@param[in]		color int		コマ色
///	@param[in]		y int			マスのY座標
///	@param[in]		x int			マスのX座標
///	@return			-1 : 失敗 それ以外 : 獲得コマ数
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetMasuStsCnt(color int, y int, x int) int {
	ret := -1
	if r.checkPara(y, 0, r.mMasuCnt) == 0 && r.checkPara(x, 0, r.mMasuCnt) == 0 {
		if color == REVERSI_STS_BLACK {
			ret = r.mMasuStsCntB[y][x]
		} else {
			ret = r.mMasuStsCntW[y][x]
		}
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			指定色が現在置ける場所があるかチェック
///	@fn				GetColorEna(color int) int
///	@param[in]		color int		コマ色
///	@return			0 : 成功 それ以外 : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetColorEna(color int) int {
	ret := -1
	for i := 0; i < r.mMasuCnt; i++ {
		for j := 0; j < r.mMasuCnt; j++ {
			if r.GetMasuStsEna(color, i, j) != 0 {
				ret = 0
				break
			}
		}
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲーム終了かチェック
///	@fn				GetGameEndSts() int
///	@return			0 : 続行 それ以外 : ゲーム終了
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetGameEndSts() int {
	ret := 1
	if r.GetColorEna(REVERSI_STS_BLACK) == 0 {
		ret = 0
	}
	if r.GetColorEna(REVERSI_STS_WHITE) == 0 {
		ret = 0
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			指定座標にコマを置く
///	@fn				SetMasuSts(color int,y int,x int) int
///	@param[in]		color int		コマ色
///	@param[in]		y int			マスのY座標
///	@param[in]		x int			マスのX座標
///	@return			0 : 成功 それ以外 : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetMasuSts(color int, y int, x int) int {
	ret := -1
	if r.GetMasuStsEna(color, y, x) != 0 {
		ret = 0
		for i := 0; i < r.mMasuCntMax; i++ {
			copy(r.mMasuStsOld[i], r.mMasuSts[i])
		}
		r.mMasuSts[y][x] = color
		r.revMasuSts(color, y, x)
		r.makeMasuSts(REVERSI_STS_BLACK)
		r.makeMasuSts(REVERSI_STS_WHITE)
		// *** 履歴保存 *** //
		if r.mMasuHistCur < (r.mMasuCntMax * r.mMasuCntMax) {
			r.mMasuHist[r.mMasuHistCur].SetColor(color)
			r.mMasuHist[r.mMasuHistCur].GetPoint().SetY(y)
			r.mMasuHist[r.mMasuHistCur].GetPoint().SetX(x)
			r.mMasuHistCur++
		}
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			指定座標にコマを強制的に置く
///	@fn				SetMasuStsForcibly(color int,y int,x int) int
///	@param[in]		color int		コマ色
///	@param[in]		y int			マスのY座標
///	@param[in]		x int			マスのX座標
///	@return			0 : 成功 それ以外 : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetMasuStsForcibly(color int, y int, x int) int {
	ret := 0
	for i := 0; i < r.mMasuCntMax; i++ {
		copy(r.mMasuStsOld[i], r.mMasuSts[i])
	}
	r.mMasuSts[y][x] = color
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			マスの数変更
///	@fn				SetMasuCnt(cnt int) int
///	@param[in]		cnt int		マスの数
///	@return			0 : 成功 それ以外 : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *Reversi) SetMasuCnt(cnt int) int {
	ret := -1
	chg := 0
	if r.checkPara(cnt, 0, r.mMasuCntMax) == 0 {
		if r.mMasuCnt != cnt {
			chg = 1
		}
		r.mMasuCnt = cnt
		ret = 0
		if chg == 1 {
			r.Reset()
		}
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ポイント座標取得
///	@fn				GetPoint(color int,num int) *ReversiPoint
///	@param[in]		color int		コマ色
///	@param[in]		num int			ポイント
///	@return			ポイント座標 nil : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetPoint(color int, num int) *ReversiPoint {
	var ret *ReversiPoint = nil
	if r.checkPara(num, 0, (r.mMasuCnt*r.mMasuCnt)) == 0 {
		if color == REVERSI_STS_BLACK {
			ret = r.mMasuPointB[num]
		} else {
			ret = r.mMasuPointW[num]
		}
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ポイント座標数取得
///	@fn				GetPointCnt(color int) int
///	@param[in]		color int		コマ色
///	@return			ポイント座標数取得
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetPointCnt(color int) int {
	ret := 0
	if color == REVERSI_STS_BLACK {
		ret = r.mMasuPointCntB
	} else {
		ret = r.mMasuPointCntW
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			コマ数取得
///	@fn				GetBetCnt(color int) int
///	@param[in]		color int		コマ色
///	@return			コマ数取得
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetBetCnt(color int) int {
	ret := 0
	if color == REVERSI_STS_BLACK {
		ret = r.mMasuBetCntB
	} else {
		ret = r.mMasuBetCntW
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			パス判定
///	@fn				GetPassEna(color int,y int,x int) int
///	@param[in]		color int		コマ色
///	@param[in]		y int			マスのY座標
///	@param[in]		x int			マスのX座標
///	@return			パス判定
///					- 0 : NOT PASS
///					- 1 : PASS
///
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetPassEna(color int, y int, x int) int {
	ret := 0
	if r.checkPara(y, 0, r.mMasuCnt) == 0 && r.checkPara(x, 0, r.mMasuCnt) == 0 {
		if color == REVERSI_STS_BLACK {
			ret = r.mMasuStsPassB[y][x]
		} else {
			ret = r.mMasuStsPassW[y][x]
		}
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			履歴取得
///	@fn				GetHistory(num int) *ReversiHistory
///	@param[in]		num	int		ポイント
///	@return			履歴 nil : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetHistory(num int) *ReversiHistory {
	var ret *ReversiHistory = nil
	if r.checkPara(num, 0, (r.mMasuCnt*r.mMasuCnt)) == 0 {
		ret = r.mMasuHist[num]
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			履歴数取得
///	@fn				GetHistoryCnt() int
///	@return			履歴数
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetHistoryCnt() int {
	return r.mMasuHistCur
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ポイント座標解析取得
///	@fn				GetPointAnz(color int,y int,x int) *ReversiAnz
///	@param[in]		color int		コマ色
///	@param[in]		y int			マスのY座標
///	@param[in]		x int			マスのX座標
///	@return			ポイント座標解析 nil : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetPointAnz(color int, y int, x int) *ReversiAnz {
	var ret *ReversiAnz = nil
	if r.checkPara(y, 0, r.mMasuCnt) == 0 && r.checkPara(x, 0, r.mMasuCnt) == 0 {
		if color == REVERSI_STS_BLACK {
			ret = r.mMasuStsAnzB[y][x]
		} else {
			ret = r.mMasuStsAnzW[y][x]
		}
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			角の隣に置いても角を取られないマス検索
///	@fn				CheckEdge(color int,y int,x int) int
///	@param[in]		color int		色
///	@param[in]		y int			Y座標
///	@param[in]		x int			X座標
///	@return			0 : 取られる それ以外 : 取られない
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) CheckEdge(color int, y int, x int) int {
	style := 0
	flg1 := 0
	flg2 := 0
	var loop int
	var loop2 int
	if y == 0 && x == 1 {
		flg1 = 0
		flg2 = 0
		for loop = x; loop < r.mMasuCnt; loop++ {
			if r.GetMasuSts(y, loop) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(y, loop) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(y, loop) != color) && (r.GetMasuSts(y, loop) != REVERSI_STS_NONE) {
				flg2 = 1
			}
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	if y == 1 && x == 0 {
		flg1 = 0
		flg2 = 0
		for loop = y; loop < r.mMasuCnt; loop++ {
			if r.GetMasuSts(loop, x) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(loop, x) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(loop, x) != color) && (r.GetMasuSts(loop, x) != REVERSI_STS_NONE) {
				flg2 = 1
			}
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	if y == 1 && x == 1 {
		flg1 = 0
		flg2 = 0
		for loop = y; loop < r.mMasuCnt; loop++ {
			if r.GetMasuSts(loop, loop) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(loop, loop) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(loop, loop) != color) && (r.GetMasuSts(loop, loop) != REVERSI_STS_NONE) {
				flg2 = 1
			}
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	if y == 0 && x == (r.mMasuCnt-2) {
		flg1 = 0
		flg2 = 0
		for loop = x; loop > 0; loop-- {
			if r.GetMasuSts(y, loop) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(y, loop) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(y, loop) != color) && (r.GetMasuSts(y, loop) != REVERSI_STS_NONE) {
				flg2 = 1
			}
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	if y == 1 && x == (r.mMasuCnt-1) {
		flg1 = 0
		flg2 = 0
		for loop = y; loop < r.mMasuCnt; loop++ {
			if r.GetMasuSts(loop, x) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(loop, x) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(loop, x) != color) && (r.GetMasuSts(loop, x) != REVERSI_STS_NONE) {
				flg2 = 1
			}
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	if y == 1 && x == (r.mMasuCnt-2) {
		loop2 = x
		flg1 = 0
		flg2 = 0
		for loop = y; loop < r.mMasuCnt; loop++ {
			if r.GetMasuSts(loop, loop2) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(loop, loop2) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(loop, loop2) != color) && (r.GetMasuSts(loop, loop2) != REVERSI_STS_NONE) {
				flg2 = 1
			}
			loop2--
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	if y == (r.mMasuCnt-2) && x == 0 {
		flg1 = 0
		flg2 = 0
		for loop = y; loop > 0; loop-- {
			if r.GetMasuSts(loop, x) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(loop, x) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(loop, x) != color) && (r.GetMasuSts(loop, x) != REVERSI_STS_NONE) {
				flg2 = 1
			}
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	if y == (r.mMasuCnt-1) && x == 1 {
		flg1 = 0
		flg2 = 0
		for loop = x; loop < r.mMasuCnt; loop++ {
			if r.GetMasuSts(y, loop) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(y, loop) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(y, loop) != color) && (r.GetMasuSts(y, loop) != REVERSI_STS_NONE) {
				flg2 = 1
			}
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	if y == (r.mMasuCnt-2) && x == 1 {
		loop2 = x
		flg1 = 0
		flg2 = 0
		for loop = y; loop > 0; loop-- {
			if r.GetMasuSts(loop, loop2) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(loop, loop2) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(loop, loop2) != color) && (r.GetMasuSts(loop, loop2) != REVERSI_STS_NONE) {
				flg2 = 1
			}
			loop2++
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	if y == (r.mMasuCnt-2) && x == (r.mMasuCnt-1) {
		flg1 = 0
		flg2 = 0
		for loop = y; loop > 0; loop-- {
			if r.GetMasuSts(loop, x) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(loop, x) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(loop, x) != color) && (r.GetMasuSts(loop, x) != REVERSI_STS_NONE) {
				flg2 = 1
			}
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	if y == (r.mMasuCnt-1) && x == (r.mMasuCnt-2) {
		flg1 = 0
		flg2 = 0
		for loop = x; loop > 0; loop-- {
			if r.GetMasuSts(y, loop) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(y, loop) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(y, loop) != color) && (r.GetMasuSts(y, loop) != REVERSI_STS_NONE) {
				flg2 = 1
			}
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	if y == (r.mMasuCnt-2) && x == (r.mMasuCnt-2) {
		loop2 = x
		flg1 = 0
		flg2 = 0
		for loop = y; loop > 0; loop-- {
			if r.GetMasuSts(loop, loop2) == color {
				flg1 = 1
			}
			if flg1 == 1 && r.GetMasuSts(loop, loop2) == REVERSI_STS_NONE {
				break
			}
			if (flg1 == 1) && (r.GetMasuSts(loop, loop2) != color) && (r.GetMasuSts(loop, loop2) != REVERSI_STS_NONE) {
				flg2 = 1
			}
			loop2--
		}
		if (flg1 == 1) && (flg2 == 0) {
			style = 1
		}
	}
	return style
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			指定座標が角か取得
///	@fn				GetEdgeSideZero(y int,x int) int
///	@param[in]		y int			Y座標
///	@param[in]		x int			X座標
///	@return			0 : 成功 それ以外 : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetEdgeSideZero(y int, x int) int {
	ret := -1
	if (y == 0 && x == 0) || (y == 0 && x == (r.mMasuCnt-1)) || (y == (r.mMasuCnt-1) && x == 0) || (y == (r.mMasuCnt-1) && x == (r.mMasuCnt-1)) {
		ret = 0
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			指定座標が角の一つ手前か取得
///	@fn				GetEdgeSideOne(int y,int x) int
///	@param[in]		y int			Y座標
///	@param[in]		x int			X座標
///	@return			0 : 成功 それ以外 : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetEdgeSideOne(y int, x int) int {
	ret := -1
	if (y == 0 && x == 1) || (y == 0 && x == (r.mMasuCnt-2)) || (y == 1 && x == 0) || (y == 1 && x == 1) || (y == 1 && x == (r.mMasuCnt-2)) || (y == 1 && x == (r.mMasuCnt-1)) || (y == (r.mMasuCnt-2) && x == 0) || (y == (r.mMasuCnt-2) && x == 1) || (y == (r.mMasuCnt-2) && x == (r.mMasuCnt-2)) || (y == (r.mMasuCnt-2) && x == (r.mMasuCnt-1)) || (y == (r.mMasuCnt-1) && x == 1) || (y == (r.mMasuCnt-1) && x == (r.mMasuCnt-2)) {
		ret = 0
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			指定座標が角の二つ手前か取得
///	@fn				GetEdgeSideTwo(int y,int x) int
///	@param[in]		y int			Y座標
///	@param[in]		x int			X座標
///	@return			0 : 成功 それ以外 : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetEdgeSideTwo(y int, x int) int {
	ret := -1
	if (y == 0 && x == 2) || (y == 0 && x == (r.mMasuCnt-3)) || (y == 2 && x == 0) || (y == 2 && x == 2) || (y == 2 && x == (r.mMasuCnt-3)) || (y == 2 && x == (r.mMasuCnt-1)) || (y == (r.mMasuCnt-3) && x == 0) || (y == (r.mMasuCnt-3) && x == 2) || (y == (r.mMasuCnt-3) && x == (r.mMasuCnt-3)) || (y == (r.mMasuCnt-3) && x == (r.mMasuCnt-1)) || (y == (r.mMasuCnt-1) && x == 2) || (y == (r.mMasuCnt-1) && x == (r.mMasuCnt-3)) {
		ret = 0
	}
	return ret
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			指定座標が角の三つ以上手前か取得
///	@fn				GetEdgeSideThree(int y,int x) int
///	@param[in]		y int			Y座標
///	@param[in]		x int			X座標
///	@return			0 : 成功 それ以外 : 失敗
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r Reversi) GetEdgeSideThree(y int, x int) int {
	ret := -1
	if (y == 0 && (3 <= x && x <= (r.mMasuCnt-4))) || ((3 <= y && y <= (r.mMasuCnt-4)) && x == 0) || (y == (r.mMasuCnt-1) && (3 <= x && x <= (r.mMasuCnt-4))) || ((3 <= y && y <= (r.mMasuCnt-4)) && x == (r.mMasuCnt-1)) {
		ret = 0
	}
	return ret
}
