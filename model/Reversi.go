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
	MasuSts       [][]int           //!< マスの状態
	MasuStsOld    [][]int           //!< 以前のマスの状態
	MasuStsEnaB   [][]int           //!< 黒の置ける場所
	MasuStsCntB   [][]int           //!< 黒の獲得コマ数
	MasuStsPassB  [][]int           //!< 黒が相手をパスさせる場所
	MasuStsAnzB   [][]*ReversiAnz   //!< 黒がその場所に置いた場合の解析結果
	MasuPointB    []*ReversiPoint   //!< 黒の置ける場所座標一覧
	MasuPointCntB int               //!< 黒の置ける場所座標一覧数
	MasuBetCntB   int               //!< 黒コマ数
	MasuStsEnaW   [][]int           //!< 白の置ける場所
	MasuStsCntW   [][]int           //!< 白の獲得コマ数
	MasuStsPassW  [][]int           //!< 白が相手をパスさせる場所
	MasuStsAnzW   [][]*ReversiAnz   //!< 白がその場所に置いた場合の解析結果
	MasuPointW    []*ReversiPoint   //!< 白の置ける場所座標一覧
	MasuPointCntW int               //!< 白の置ける場所座標一覧数
	MasuBetCntW   int               //!< 白コマ数
	MasuCnt       int               //!< 縦横マス数
	MasuCntMax    int               //!< 縦横マス最大数
	MasuHist      []*ReversiHistory //!< 履歴
	MasuHistCur   int               //!< 履歴現在位置
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			コンストラクタ
///	@fn				NewReversi(masuCnt int, masuMax int) *Reversi
///	@param[in]		masuCnt int		縦横マス数
///	@param[in]		masuMax int		縦横マス最大数
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func NewReversi(masuCnt int, masuMax int) *Reversi {
	r := new(Reversi)
	r.MasuCnt = masuCnt
	r.MasuCntMax = masuMax

	r.MasuSts = make([][]int, r.MasuCntMax)
	r.MasuStsOld = make([][]int, r.MasuCntMax)
	r.MasuStsEnaB = make([][]int, r.MasuCntMax)
	r.MasuStsCntB = make([][]int, r.MasuCntMax)
	r.MasuStsPassB = make([][]int, r.MasuCntMax)
	for i := 0; i < r.MasuCntMax; i++ {
		r.MasuSts[i] = make([]int, r.MasuCntMax)
		r.MasuStsOld[i] = make([]int, r.MasuCntMax)
		r.MasuStsEnaB[i] = make([]int, r.MasuCntMax)
		r.MasuStsCntB[i] = make([]int, r.MasuCntMax)
		r.MasuStsPassB[i] = make([]int, r.MasuCntMax)
	}
	r.MasuStsAnzB = make([][]*ReversiAnz, r.MasuCntMax)
	for i := 0; i < r.MasuCntMax; i++ {
		r.MasuStsAnzB[i] = make([]*ReversiAnz, r.MasuCntMax)
		for j := 0; j < r.MasuCntMax; j++ {
			r.MasuStsAnzB[i][j] = NewReversiAnz()
		}
	}
	r.MasuPointB = make([]*ReversiPoint, r.MasuCntMax*r.MasuCntMax)
	for i := 0; i < (r.MasuCntMax * r.MasuCntMax); i++ {
		r.MasuPointB[i] = NewReversiPoint()
	}
	r.MasuPointCntB = 0
	r.MasuStsEnaW = make([][]int, r.MasuCntMax)
	r.MasuStsCntW = make([][]int, r.MasuCntMax)
	r.MasuStsPassW = make([][]int, r.MasuCntMax)
	for i := 0; i < r.MasuCntMax; i++ {
		r.MasuStsEnaW[i] = make([]int, r.MasuCntMax)
		r.MasuStsCntW[i] = make([]int, r.MasuCntMax)
		r.MasuStsPassW[i] = make([]int, r.MasuCntMax)
	}
	r.MasuStsAnzW = make([][]*ReversiAnz, r.MasuCntMax)
	for i := 0; i < r.MasuCntMax; i++ {
		r.MasuStsAnzW[i] = make([]*ReversiAnz, r.MasuCntMax)
		for j := 0; j < r.MasuCntMax; j++ {
			r.MasuStsAnzW[i][j] = NewReversiAnz()
		}
	}
	r.MasuPointW = make([]*ReversiPoint, r.MasuCntMax*r.MasuCntMax)
	for i := 0; i < (r.MasuCntMax * r.MasuCntMax); i++ {
		r.MasuPointW[i] = NewReversiPoint()
	}
	r.MasuPointCntW = 0
	r.MasuBetCntB = 0
	r.MasuBetCntW = 0
	r.MasuHist = make([]*ReversiHistory, r.MasuCntMax*r.MasuCntMax)
	for i := 0; i < (r.MasuCntMax * r.MasuCntMax); i++ {
		r.MasuHist[i] = NewReversiHistory()
	}
	r.MasuHistCur = 0
	for i := 0; i < r.MasuCntMax; i++ {
		copy(r.MasuStsOld[i], r.MasuSts[i])
	}
	r.Reset()
	return r
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
	return r.MasuSts
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
	r.MasuSts = mMasuSts
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
	return r.MasuStsOld
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
	r.MasuStsOld = mMasuStsOld
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
	return r.MasuStsEnaB
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
	r.MasuStsEnaB = mMasuStsEnaB
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
	return r.MasuStsCntB
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
	r.MasuStsCntB = mMasuStsCntB
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
	return r.MasuStsPassB
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
	r.MasuStsPassB = mMasuStsPassB
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
	return r.MasuStsAnzB
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
	r.MasuStsAnzB = mMasuStsAnzB
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
	return r.MasuPointB
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
	r.MasuPointB = mMasuPointB
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
	return r.MasuPointCntB
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
	r.MasuPointCntB = mMasuPointCntB
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
	return r.MasuBetCntB
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
	r.MasuBetCntB = mMasuBetCntB
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
	return r.MasuStsEnaW
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
	r.MasuStsEnaW = mMasuStsEnaW
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
	return r.MasuStsCntW
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
	r.MasuStsCntW = mMasuStsCntW
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
	return r.MasuStsPassW
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
	r.MasuStsPassW = mMasuStsPassW
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
	return r.MasuStsAnzW
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
	r.MasuStsAnzW = mMasuStsAnzW
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
	return r.MasuPointW
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
	r.MasuPointW = mMasuPointW
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
	return r.MasuPointCntW
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
	r.MasuPointCntW = mMasuPointCntW
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
	return r.MasuBetCntW
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
	r.MasuBetCntW = mMasuBetCntW
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
	return r.MasuCnt
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
	r.MasuCnt = mMasuCnt
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
	return r.MasuCntMax
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
	r.MasuCntMax = mMasuCntMax
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
	return r.MasuHist
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
	r.MasuHist = mMasuHist
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
	return r.MasuHistCur
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
	r.MasuHistCur = mMasuHistCur
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
	for i := 0; i < r.MasuCnt; i++ {
		for j := 0; j < r.MasuCnt; j++ {
			r.MasuSts[i][j] = REVERSI_STS_NONE
			r.MasuStsPassB[i][j] = 0
			r.MasuStsAnzB[i][j].Reset()
			r.MasuStsPassW[i][j] = 0
			r.MasuStsAnzW[i][j].Reset()
		}
	}
	r.MasuSts[(r.MasuCnt>>1)-1][(r.MasuCnt>>1)-1] = REVERSI_STS_BLACK
	r.MasuSts[(r.MasuCnt>>1)-1][(r.MasuCnt >> 1)] = REVERSI_STS_WHITE
	r.MasuSts[(r.MasuCnt >> 1)][(r.MasuCnt>>1)-1] = REVERSI_STS_WHITE
	r.MasuSts[(r.MasuCnt >> 1)][(r.MasuCnt >> 1)] = REVERSI_STS_BLACK
	r.makeMasuSts(REVERSI_STS_BLACK)
	r.makeMasuSts(REVERSI_STS_WHITE)
	r.MasuHistCur = 0
	for i := 0; i < r.MasuCntMax; i++ {
		copy(r.MasuStsOld[i], r.MasuSts[i])
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
	for i := 0; i < r.MasuCnt; i++ {
		// 初期化
		for j := 0; j < r.MasuCnt; j++ {
			if color == REVERSI_STS_BLACK {
				r.MasuStsEnaB[i][j] = 0
				r.MasuStsCntB[i][j] = 0
			} else {
				r.MasuStsEnaW[i][j] = 0
				r.MasuStsCntW[i][j] = 0
			}
		}
	}

	loop = r.MasuCnt * r.MasuCnt
	for i := 0; i < loop; i++ {
		// 初期化
		if color == REVERSI_STS_BLACK {
			r.MasuPointB[i].SetX(0)
			r.MasuPointB[i].SetY(0)
		} else {
			r.MasuPointW[i].SetX(0)
			r.MasuPointW[i].SetY(0)
		}
	}
	if color == REVERSI_STS_BLACK {
		r.MasuPointCntB = 0
	} else {
		r.MasuPointCntW = 0
	}
	r.MasuBetCntB = 0
	r.MasuBetCntW = 0

	for i := 0; i < r.MasuCnt; i++ {
		for j := 0; j < r.MasuCnt; j++ {
			okflg = 0
			count2 = 0
			if r.MasuSts[i][j] == REVERSI_STS_NONE {
				// 何も置かれていないマスなら
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 上方向を調べる *** //
				for (cnt1 > 0) && (r.MasuSts[cnt1-1][j] != REVERSI_STS_NONE && r.MasuSts[cnt1-1][j] != color) {
					flg = 1
					cnt1--
					count1++
				}
				if (cnt1 > 0) && (flg == 1) && (r.MasuSts[cnt1-1][j] == color) {
					okflg = 1
					count2 += count1
				}
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 下方向を調べる *** //
				for (cnt1 < (r.MasuCnt - 1)) && (r.MasuSts[cnt1+1][j] != REVERSI_STS_NONE && r.MasuSts[cnt1+1][j] != color) {
					flg = 1
					cnt1++
					count1++
				}
				if (cnt1 < (r.MasuCnt - 1)) && (flg == 1) && (r.MasuSts[cnt1+1][j] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				count1 = 0
				flg = 0
				// *** 右方向を調べる *** //
				for (cnt2 < (r.MasuCnt - 1)) && (r.MasuSts[i][cnt2+1] != REVERSI_STS_NONE && r.MasuSts[i][cnt2+1] != color) {
					flg = 1
					cnt2++
					count1++
				}
				if (cnt2 < (r.MasuCnt - 1)) && (flg == 1) && (r.MasuSts[i][cnt2+1] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				count1 = 0
				flg = 0
				// *** 左方向を調べる *** //
				for (cnt2 > 0) && (r.MasuSts[i][cnt2-1] != REVERSI_STS_NONE && r.MasuSts[i][cnt2-1] != color) {
					flg = 1
					cnt2--
					count1++
				}
				if (cnt2 > 0) && (flg == 1) && (r.MasuSts[i][cnt2-1] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 右上方向を調べる *** //
				for ((cnt2 < (r.MasuCnt - 1)) && (cnt1 > 0)) && (r.MasuSts[cnt1-1][cnt2+1] != REVERSI_STS_NONE && r.MasuSts[cnt1-1][cnt2+1] != color) {
					flg = 1
					cnt1--
					cnt2++
					count1++
				}
				if ((cnt2 < (r.MasuCnt - 1)) && (cnt1 > 0)) && (flg == 1) && (r.MasuSts[cnt1-1][cnt2+1] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 左上方向を調べる *** //
				for ((cnt2 > 0) && (cnt1 > 0)) && (r.MasuSts[cnt1-1][cnt2-1] != REVERSI_STS_NONE && r.MasuSts[cnt1-1][cnt2-1] != color) {
					flg = 1
					cnt1--
					cnt2--
					count1++
				}
				if ((cnt2 > 0) && (cnt1 > 0)) && (flg == 1) && (r.MasuSts[cnt1-1][cnt2-1] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 右下方向を調べる *** //
				for ((cnt2 < (r.MasuCnt - 1)) && (cnt1 < (r.MasuCnt - 1))) && (r.MasuSts[cnt1+1][cnt2+1] != REVERSI_STS_NONE && r.MasuSts[cnt1+1][cnt2+1] != color) {
					flg = 1
					cnt1++
					cnt2++
					count1++
				}
				if ((cnt2 < (r.MasuCnt - 1)) && (cnt1 < (r.MasuCnt - 1))) && (flg == 1) && (r.MasuSts[cnt1+1][cnt2+1] == color) {
					okflg = 1
					count2 += count1
				}
				cnt2 = j
				cnt1 = i
				count1 = 0
				flg = 0
				// *** 左下方向を調べる *** //
				for ((cnt2 > 0) && (cnt1 < (r.MasuCnt - 1))) && (r.MasuSts[cnt1+1][cnt2-1] != REVERSI_STS_NONE && r.MasuSts[cnt1+1][cnt2-1] != color) {
					flg = 1
					cnt1++
					cnt2--
					count1++
				}
				if ((cnt2 > 0) && (cnt1 < (r.MasuCnt - 1))) && (flg == 1) && (r.MasuSts[cnt1+1][cnt2-1] == color) {
					okflg = 1
					count2 += count1
				}
				if okflg == 1 {
					if color == REVERSI_STS_BLACK {
						r.MasuStsEnaB[i][j] = 1
						r.MasuStsCntB[i][j] = count2
						// *** 置ける場所をリニアに保存、置けるポイント数も保存 *** //
						r.MasuPointB[r.MasuPointCntB].SetY(i)
						r.MasuPointB[r.MasuPointCntB].SetX(j)
						r.MasuPointCntB++
					} else {
						r.MasuStsEnaW[i][j] = 1
						r.MasuStsCntW[i][j] = count2
						// *** 置ける場所をリニアに保存、置けるポイント数も保存 *** //
						r.MasuPointW[r.MasuPointCntW].SetY(i)
						r.MasuPointW[r.MasuPointCntW].SetX(j)
						r.MasuPointCntW++
					}
					ret = 0
					if countMax < count2 {
						countMax = count2
					}
				}
			} else if r.MasuSts[i][j] == REVERSI_STS_BLACK {
				r.MasuBetCntB++
			} else if r.MasuSts[i][j] == REVERSI_STS_WHITE {
				r.MasuBetCntW++
			}
		}
	}

	// *** 一番枚数を獲得できるマスを設定 *** //
	for i := 0; i < r.MasuCnt; i++ {
		for j := 0; j < r.MasuCnt; j++ {
			if color == REVERSI_STS_BLACK {
				if r.MasuStsEnaB[i][j] != 0 && r.MasuStsCntB[i][j] == countMax {
					r.MasuStsEnaB[i][j] = 2
				}
			} else {
				if r.MasuStsEnaW[i][j] != 0 && r.MasuStsCntW[i][j] == countMax {
					r.MasuStsEnaW[i][j] = 2
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
		if r.MasuSts[cnt2][cnt1-1] != REVERSI_STS_NONE && r.MasuSts[cnt2][cnt1-1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt1--
		} else if r.MasuSts[cnt2][cnt1-1] == color {
			flg = 1
			break
		} else if r.MasuSts[cnt2][cnt1-1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		for rcnt1 = cnt1; rcnt1 < x; rcnt1++ {
			r.MasuSts[cnt2][rcnt1] = color
		}
	}

	// *** 右方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt1 < (r.MasuCnt - 1); {
		if r.MasuSts[cnt2][cnt1+1] != REVERSI_STS_NONE && r.MasuSts[cnt2][cnt1+1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt1++
		} else if r.MasuSts[cnt2][cnt1+1] == color {
			flg = 1
			break
		} else if r.MasuSts[cnt2][cnt1+1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		for rcnt1 = cnt1; rcnt1 > x; rcnt1-- {
			r.MasuSts[cnt2][rcnt1] = color
		}
	}

	// *** 上方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 > 0; {
		if r.MasuSts[cnt2-1][cnt1] != REVERSI_STS_NONE && r.MasuSts[cnt2-1][cnt1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2--
		} else if r.MasuSts[cnt2-1][cnt1] == color {
			flg = 1
			break
		} else if r.MasuSts[cnt2-1][cnt1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		for rcnt1 = cnt2; rcnt1 < y; rcnt1++ {
			r.MasuSts[rcnt1][cnt1] = color
		}
	}

	// *** 下方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 < (r.MasuCnt - 1); {
		if r.MasuSts[cnt2+1][cnt1] != REVERSI_STS_NONE && r.MasuSts[cnt2+1][cnt1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2++
		} else if r.MasuSts[cnt2+1][cnt1] == color {
			flg = 1
			break
		} else if r.MasuSts[cnt2+1][cnt1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		for rcnt1 = cnt2; rcnt1 > y; rcnt1-- {
			r.MasuSts[rcnt1][cnt1] = color
		}
	}

	// *** 左上方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 > 0 && cnt1 > 0; {
		if r.MasuSts[cnt2-1][cnt1-1] != REVERSI_STS_NONE && r.MasuSts[cnt2-1][cnt1-1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2--
			cnt1--
		} else if r.MasuSts[cnt2-1][cnt1-1] == color {
			flg = 1
			break
		} else if r.MasuSts[cnt2-1][cnt1-1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		rcnt2 = cnt1
		for rcnt1 = cnt2; (rcnt1 < y) && (rcnt2 < x); rcnt1++ {
			r.MasuSts[rcnt1][rcnt2] = color
			rcnt2++
		}
	}

	// *** 左下方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 < (r.MasuCnt-1) && cnt1 > 0; {
		if r.MasuSts[cnt2+1][cnt1-1] != REVERSI_STS_NONE && r.MasuSts[cnt2+1][cnt1-1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2++
			cnt1--
		} else if r.MasuSts[cnt2+1][cnt1-1] == color {
			flg = 1
			break
		} else if r.MasuSts[cnt2+1][cnt1-1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		rcnt2 = cnt1
		for rcnt1 = cnt2; (rcnt1 > y) && (rcnt2 < x); rcnt1-- {
			r.MasuSts[rcnt1][rcnt2] = color
			rcnt2++
		}
	}

	// *** 右上方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 > 0 && cnt1 < (r.MasuCnt-1); {
		if r.MasuSts[cnt2-1][cnt1+1] != REVERSI_STS_NONE && r.MasuSts[cnt2-1][cnt1+1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2--
			cnt1++
		} else if r.MasuSts[cnt2-1][cnt1+1] == color {
			flg = 1
			break
		} else if r.MasuSts[cnt2-1][cnt1+1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		rcnt2 = cnt1
		for rcnt1 = cnt2; (rcnt1 < y) && (rcnt2 > x); rcnt1++ {
			r.MasuSts[rcnt1][rcnt2] = color
			rcnt2--
		}
	}

	// *** 右下方向にある駒を調べる *** //
	cnt1 = x
	cnt2 = y
	for flg = 0; cnt2 < (r.MasuCnt-1) && cnt1 < (r.MasuCnt-1); {
		if r.MasuSts[cnt2+1][cnt1+1] != REVERSI_STS_NONE && r.MasuSts[cnt2+1][cnt1+1] != color {
			// *** プレイヤー以外の色の駒があるなら *** //
			cnt2++
			cnt1++
		} else if r.MasuSts[cnt2+1][cnt1+1] == color {
			flg = 1
			break
		} else if r.MasuSts[cnt2+1][cnt1+1] == REVERSI_STS_NONE {
			break
		}
	}
	if flg == 1 {
		// *** 駒をひっくり返す *** //
		rcnt2 = cnt1
		for rcnt1 = cnt2; (rcnt1 > y) && (rcnt2 > x); rcnt1-- {
			r.MasuSts[rcnt1][rcnt2] = color
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
	for cnt := 0; cnt < r.MasuPointCntB; cnt++ {
		// *** 現在ステータスを一旦コピー *** //
		tmpMasu := make([][]int, r.MasuCntMax)
		tmpMasuEnaB := make([][]int, r.MasuCntMax)
		tmpMasuEnaW := make([][]int, r.MasuCntMax)
		for i := 0; i < r.MasuCntMax; i++ {
			tmpMasu[i] = make([]int, r.MasuCntMax)
			tmpMasuEnaB[i] = make([]int, r.MasuCntMax)
			tmpMasuEnaW[i] = make([]int, r.MasuCntMax)
			copy(tmpMasu[i], r.MasuSts[i])
			copy(tmpMasuEnaB[i], r.MasuStsEnaB[i])
			copy(tmpMasuEnaW[i], r.MasuStsEnaW[i])
		}
		tmpY = r.MasuPointB[cnt].GetY()
		tmpX = r.MasuPointB[cnt].GetX()
		// 仮に置く
		r.MasuSts[tmpY][tmpX] = REVERSI_STS_BLACK
		// 仮にひっくり返す
		r.revMasuSts(REVERSI_STS_BLACK, tmpY, tmpX)
		r.makeMasuSts(REVERSI_STS_BLACK)
		r.makeMasuSts(REVERSI_STS_WHITE)
		// *** このマスに置いた場合の解析を行う *** //
		if r.GetColorEna(REVERSI_STS_WHITE) != 0 {
			// 相手がパス
			r.MasuStsPassB[tmpY][tmpX] = 1
		}
		if r.GetEdgeSideZero(tmpY, tmpX) == 0 {
			// 置いた場所が角
			r.MasuStsAnzB[tmpY][tmpX].SetOwnEdgeCnt(r.MasuStsAnzB[tmpY][tmpX].GetOwnEdgeCnt() + 1)
			r.MasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 10000*r.MasuStsCntB[tmpY][tmpX])
		} else if r.GetEdgeSideOne(tmpY, tmpX) == 0 {
			// 置いた場所が角の一つ手前
			r.MasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.MasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
			if r.CheckEdge(REVERSI_STS_BLACK, tmpY, tmpX) != 0 {
				// 角を取られない
				r.MasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 10*r.MasuStsCntB[tmpY][tmpX])
			} else {
				// 角を取られる
				r.MasuStsAnzB[tmpY][tmpX].SetBadPoint(r.MasuStsAnzB[tmpY][tmpX].GetBadPoint() + 100000)
			}
		} else if r.GetEdgeSideTwo(tmpY, tmpX) == 0 {
			// 置いた場所が角の二つ手前
			r.MasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.MasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
			r.MasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 1000*r.MasuStsCntB[tmpY][tmpX])
		} else if r.GetEdgeSideThree(tmpY, tmpX) == 0 {
			// 置いた場所が角の三つ手前
			r.MasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.MasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideThreeCnt() + 1)
			r.MasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 100*r.MasuStsCntB[tmpY][tmpX])
		} else {
			// 置いた場所がその他
			r.MasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideOtherCnt(r.MasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideOtherCnt() + 1)
			r.MasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 10*r.MasuStsCntB[tmpY][tmpX])
		}
		sum = 0
		sumOwn = 0
		for i := 0; i < r.MasuCnt; i++ {
			for j := 0; j < r.MasuCnt; j++ {
				tmpBadPoint = 0
				tmpGoodPoint = 0
				if r.GetMasuStsEna(REVERSI_STS_WHITE, i, j) != 0 {
					// 相手の獲得予定枚数
					sum += r.MasuStsCntW[i][j]
					// *** 相手の獲得予定の最大数保持 *** //
					if r.MasuStsAnzB[tmpY][tmpX].GetMax() < r.MasuStsCntW[i][j] {
						r.MasuStsAnzB[tmpY][tmpX].SetMax(r.MasuStsCntW[i][j])
					}
					// *** 相手の獲得予定の最小数保持 *** //
					if r.MasuStsCntW[i][j] < r.MasuStsAnzB[tmpY][tmpX].GetMin() {
						r.MasuStsAnzB[tmpY][tmpX].SetMin(r.MasuStsCntW[i][j])
					}
					// 相手の置ける場所の数
					r.MasuStsAnzB[tmpY][tmpX].SetPointCnt(r.MasuStsAnzB[tmpY][tmpX].GetPointCnt() + 1)
					if r.GetEdgeSideZero(i, j) == 0 {
						// 置く場所が角
						r.MasuStsAnzB[tmpY][tmpX].SetEdgeCnt(r.MasuStsAnzB[tmpY][tmpX].GetEdgeCnt() + 1)
						tmpBadPoint = 100000 * r.MasuStsCntW[i][j]
					} else if r.GetEdgeSideOne(i, j) == 0 {
						// 置く場所が角の一つ手前
						r.MasuStsAnzB[tmpY][tmpX].SetEdgeSideOneCnt(r.MasuStsAnzB[tmpY][tmpX].GetEdgeSideOneCnt() + 1)
						tmpBadPoint = 0
					} else if r.GetEdgeSideTwo(i, j) == 0 {
						// 置く場所が角の二つ手前
						r.MasuStsAnzB[tmpY][tmpX].SetEdgeSideTwoCnt(r.MasuStsAnzB[tmpY][tmpX].GetEdgeSideTwoCnt() + 1)
						tmpBadPoint = 1 * r.MasuStsCntW[i][j]
					} else if r.GetEdgeSideThree(i, j) == 0 {
						// 置く場所が角の三つ手前
						r.MasuStsAnzB[tmpY][tmpX].SetEdgeSideThreeCnt(r.MasuStsAnzB[tmpY][tmpX].GetEdgeSideThreeCnt() + 1)
						tmpBadPoint = 1 * r.MasuStsCntW[i][j]
					} else {
						// 置く場所がその他
						r.MasuStsAnzB[tmpY][tmpX].SetEdgeSideOtherCnt(r.MasuStsAnzB[tmpY][tmpX].GetEdgeSideOtherCnt() + 1)
						tmpBadPoint = 1 * r.MasuStsCntW[i][j]
					}
					if tmpMasuEnaW[i][j] != 0 {
						// ステータス変化していないなら
						tmpBadPoint = 0
					}
				}
				if r.GetMasuStsEna(REVERSI_STS_BLACK, i, j) != 0 {
					// 自分の獲得予定枚数
					sumOwn += r.MasuStsCntB[i][j]
					// *** 自分の獲得予定の最大数保持 *** //
					if r.MasuStsAnzB[tmpY][tmpX].GetOwnMax() < r.MasuStsCntB[i][j] {
						r.MasuStsAnzB[tmpY][tmpX].SetOwnMax(r.MasuStsCntB[i][j])
					}
					// *** 自分の獲得予定の最小数保持 *** //
					if r.MasuStsCntB[i][j] < r.MasuStsAnzB[tmpY][tmpX].GetOwnMin() {
						r.MasuStsAnzB[tmpY][tmpX].SetOwnMin(r.MasuStsCntB[i][j])
					}
					// 自分の置ける場所の数
					r.MasuStsAnzB[tmpY][tmpX].SetOwnPointCnt(r.MasuStsAnzB[tmpY][tmpX].GetOwnPointCnt() + 1)
					if r.GetEdgeSideZero(i, j) == 0 {
						// 置く場所が角
						r.MasuStsAnzB[tmpY][tmpX].SetOwnEdgeCnt(r.MasuStsAnzB[tmpY][tmpX].GetOwnEdgeCnt() + 1)
						tmpGoodPoint = 100 * r.MasuStsCntB[i][j]
					} else if r.GetEdgeSideOne(i, j) == 0 {
						// 置く場所が角の一つ手前
						r.MasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.MasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
						tmpGoodPoint = 0
					} else if r.GetEdgeSideTwo(i, j) == 0 {
						// 置く場所が角の二つ手前
						r.MasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.MasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
						tmpGoodPoint = 3 * r.MasuStsCntB[i][j]
					} else if r.GetEdgeSideThree(i, j) == 0 {
						// 置く場所が角の三つ手前
						r.MasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.MasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideThreeCnt() + 1)
						tmpGoodPoint = 2 * r.MasuStsCntB[i][j]
					} else {
						// 置く場所がその他
						r.MasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideOtherCnt(r.MasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideOtherCnt() + 1)
						tmpGoodPoint = 1 * r.MasuStsCntB[i][j]
					}
					if tmpMasuEnaB[i][j] != 0 {
						// ステータス変化していないなら
						tmpGoodPoint = 0
					}
				}
				if tmpBadPoint != 0 {
					r.MasuStsAnzB[tmpY][tmpX].SetBadPoint(r.MasuStsAnzB[tmpY][tmpX].GetBadPoint() + tmpBadPoint)
				}
				if tmpGoodPoint != 0 {
					r.MasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzB[tmpY][tmpX].GetGoodPoint() + tmpGoodPoint)
				}
			}
		}
		// *** 相手に取られる平均コマ数 *** //
		if r.GetPointCnt(REVERSI_STS_WHITE) != 0 {
			tmpD1 = float64(sum)
			tmpD2 = float64(r.GetPointCnt(REVERSI_STS_WHITE))
			r.MasuStsAnzB[tmpY][tmpX].SetAvg(tmpD1 / tmpD2)
		}
		// *** 自分が取れる平均コマ数 *** //
		if r.GetPointCnt(REVERSI_STS_BLACK) != 0 {
			tmpD1 = float64(sumOwn)
			tmpD2 = float64(r.GetPointCnt(REVERSI_STS_BLACK))
			r.MasuStsAnzB[tmpY][tmpX].SetOwnAvg(tmpD1 / tmpD2)
		}
		// *** 元に戻す *** //
		for i := 0; i < r.MasuCntMax; i++ {
			copy(r.MasuSts[i], tmpMasu[i])
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
	for cnt := 0; cnt < r.MasuPointCntW; cnt++ {
		// *** 現在ステータスを一旦コピー *** //
		tmpMasu := make([][]int, r.MasuCntMax)
		tmpMasuEnaB := make([][]int, r.MasuCntMax)
		tmpMasuEnaW := make([][]int, r.MasuCntMax)
		for i := 0; i < r.MasuCntMax; i++ {
			tmpMasu[i] = make([]int, r.MasuCntMax)
			tmpMasuEnaB[i] = make([]int, r.MasuCntMax)
			tmpMasuEnaW[i] = make([]int, r.MasuCntMax)
			copy(tmpMasu[i], r.MasuSts[i])
			copy(tmpMasuEnaB[i], r.MasuStsEnaB[i])
			copy(tmpMasuEnaW[i], r.MasuStsEnaW[i])
		}
		tmpY = r.MasuPointW[cnt].GetY()
		tmpX = r.MasuPointW[cnt].GetX()
		// 仮に置く
		r.MasuSts[tmpY][tmpX] = REVERSI_STS_WHITE
		// 仮にひっくり返す
		r.revMasuSts(REVERSI_STS_WHITE, tmpY, tmpX)
		r.makeMasuSts(REVERSI_STS_BLACK)
		r.makeMasuSts(REVERSI_STS_WHITE)
		// *** このマスに置いた場合の解析を行う *** //
		if r.GetColorEna(REVERSI_STS_BLACK) != 0 {
			// 相手がパス
			r.MasuStsPassW[tmpY][tmpX] = 1
		}
		if r.GetEdgeSideZero(tmpY, tmpX) == 0 {
			// 置いた場所が角
			r.MasuStsAnzW[tmpY][tmpX].SetOwnEdgeCnt(r.MasuStsAnzW[tmpY][tmpX].GetOwnEdgeCnt() + 1)
			r.MasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 10000*r.MasuStsCntW[tmpY][tmpX])
		} else if r.GetEdgeSideOne(tmpY, tmpX) == 0 {
			// 置いた場所が角の一つ手前
			r.MasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.MasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
			if r.CheckEdge(REVERSI_STS_WHITE, tmpY, tmpX) != 0 {
				// 角を取られない
				r.MasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 10*r.MasuStsCntW[tmpY][tmpX])
			} else {
				// 角を取られる
				r.MasuStsAnzW[tmpY][tmpX].SetBadPoint(r.MasuStsAnzW[tmpY][tmpX].GetBadPoint() + 100000)
			}
		} else if r.GetEdgeSideTwo(tmpY, tmpX) == 0 {
			// 置いた場所が角の二つ手前
			r.MasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.MasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
			r.MasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 1000*r.MasuStsCntW[tmpY][tmpX])
		} else if r.GetEdgeSideThree(tmpY, tmpX) == 0 {
			// 置いた場所が角の三つ手前
			r.MasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.MasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideThreeCnt() + 1)
			r.MasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 100*r.MasuStsCntW[tmpY][tmpX])
		} else {
			// 置いた場所がその他
			r.MasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOtherCnt(r.MasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOtherCnt() + 1)
			r.MasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 10*r.MasuStsCntW[tmpY][tmpX])
		}
		sum = 0
		sumOwn = 0
		for i := 0; i < r.MasuCnt; i++ {
			for j := 0; j < r.MasuCnt; j++ {
				tmpBadPoint = 0
				tmpGoodPoint = 0
				if r.GetMasuStsEna(REVERSI_STS_BLACK, i, j) != 0 {
					// 相手の獲得予定枚数
					sum += r.MasuStsCntB[i][j]
					// *** 相手の獲得予定の最大数保持 *** //
					if r.MasuStsAnzW[tmpY][tmpX].GetMax() < r.MasuStsCntB[i][j] {
						r.MasuStsAnzW[tmpY][tmpX].SetMax(r.MasuStsCntB[i][j])
					}
					// *** 相手の獲得予定の最小数保持 *** //
					if r.MasuStsCntB[i][j] < r.MasuStsAnzW[tmpY][tmpX].GetMin() {
						r.MasuStsAnzW[tmpY][tmpX].SetMin(r.MasuStsCntB[i][j])
					}
					// 相手の置ける場所の数
					r.MasuStsAnzW[tmpY][tmpX].SetPointCnt(r.MasuStsAnzW[tmpY][tmpX].GetPointCnt() + 1)
					if r.GetEdgeSideZero(i, j) == 0 {
						// 置く場所が角
						r.MasuStsAnzW[tmpY][tmpX].SetEdgeCnt(r.MasuStsAnzW[tmpY][tmpX].GetEdgeCnt() + 1)
						tmpBadPoint = 100000 * r.MasuStsCntB[i][j]
					} else if r.GetEdgeSideOne(i, j) == 0 {
						// 置く場所が角の一つ手前
						r.MasuStsAnzW[tmpY][tmpX].SetEdgeSideOneCnt(r.MasuStsAnzW[tmpY][tmpX].GetEdgeSideOneCnt() + 1)
						tmpBadPoint = 0
					} else if r.GetEdgeSideTwo(i, j) == 0 {
						// 置く場所が角の二つ手前
						r.MasuStsAnzW[tmpY][tmpX].SetEdgeSideTwoCnt(r.MasuStsAnzW[tmpY][tmpX].GetEdgeSideTwoCnt() + 1)
						tmpBadPoint = 1 * r.MasuStsCntB[i][j]
					} else if r.GetEdgeSideThree(i, j) == 0 {
						// 置く場所が角の三つ手前
						r.MasuStsAnzW[tmpY][tmpX].SetEdgeSideThreeCnt(r.MasuStsAnzW[tmpY][tmpX].GetEdgeSideThreeCnt() + 1)
						tmpBadPoint = 1 * r.MasuStsCntB[i][j]
					} else {
						// 置く場所がその他
						r.MasuStsAnzW[tmpY][tmpX].SetEdgeSideOtherCnt(r.MasuStsAnzW[tmpY][tmpX].GetEdgeSideOtherCnt() + 1)
						tmpBadPoint = 1 * r.MasuStsCntB[i][j]
					}
					if tmpMasuEnaB[i][j] != 0 {
						// ステータス変化していないなら
						tmpBadPoint = 0
					}
				}
				if r.GetMasuStsEna(REVERSI_STS_WHITE, i, j) != 0 {
					// 自分の獲得予定枚数
					sumOwn += r.MasuStsCntW[i][j]
					// *** 自分の獲得予定の最大数保持 *** //
					if r.MasuStsAnzW[tmpY][tmpX].GetOwnMax() < r.MasuStsCntW[i][j] {
						r.MasuStsAnzW[tmpY][tmpX].SetOwnMax(r.MasuStsCntW[i][j])
					}
					// *** 自分の獲得予定の最小数保持 *** //
					if r.MasuStsCntW[i][j] < r.MasuStsAnzW[tmpY][tmpX].GetOwnMin() {
						r.MasuStsAnzW[tmpY][tmpX].SetOwnMin(r.MasuStsCntW[i][j])
					}
					// 自分の置ける場所の数
					r.MasuStsAnzW[tmpY][tmpX].SetOwnPointCnt(r.MasuStsAnzW[tmpY][tmpX].GetOwnPointCnt() + 1)
					if r.GetEdgeSideZero(i, j) == 0 {
						// 置く場所が角
						r.MasuStsAnzW[tmpY][tmpX].SetOwnEdgeCnt(r.MasuStsAnzW[tmpY][tmpX].GetOwnEdgeCnt() + 1)
						tmpGoodPoint = 100 * r.MasuStsCntW[i][j]
					} else if r.GetEdgeSideOne(i, j) == 0 {
						// 置く場所が角の一つ手前
						r.MasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.MasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
						tmpGoodPoint = 0
					} else if r.GetEdgeSideTwo(i, j) == 0 {
						// 置く場所が角の二つ手前
						r.MasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.MasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
						tmpGoodPoint = 3 * r.MasuStsCntW[i][j]
					} else if r.GetEdgeSideThree(i, j) == 0 {
						// 置く場所が角の三つ手前
						r.MasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.MasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideThreeCnt() + 1)
						tmpGoodPoint = 2 * r.MasuStsCntW[i][j]
					} else {
						// 置く場所がその他
						r.MasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOtherCnt(r.MasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOtherCnt() + 1)
						tmpGoodPoint = 1 * r.MasuStsCntW[i][j]
					}
					if tmpMasuEnaW[i][j] != 0 {
						// ステータス変化していないなら
						tmpGoodPoint = 0
					}
				}
				if tmpBadPoint != 0 {
					r.MasuStsAnzW[tmpY][tmpX].SetBadPoint(r.MasuStsAnzW[tmpY][tmpX].GetBadPoint() + tmpBadPoint)
				}
				if tmpGoodPoint != 0 {
					r.MasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.MasuStsAnzW[tmpY][tmpX].GetGoodPoint() + tmpGoodPoint)
				}
			}
		}
		// *** 相手に取られる平均コマ数 *** //
		if r.GetPointCnt(REVERSI_STS_BLACK) != 0 {
			tmpD1 = float64(sum)
			tmpD2 = float64(r.GetPointCnt(REVERSI_STS_BLACK))
			r.MasuStsAnzW[tmpY][tmpX].SetAvg(tmpD1 / tmpD2)
		}
		// *** 自分が取れる平均コマ数 *** //
		if r.GetPointCnt(REVERSI_STS_WHITE) != 0 {
			tmpD1 = float64(sumOwn)
			tmpD2 = float64(r.GetPointCnt(REVERSI_STS_WHITE))
			r.MasuStsAnzW[tmpY][tmpX].SetOwnAvg(tmpD1 / tmpD2)
		}
		// *** 元に戻す *** //
		for i := 0; i < r.MasuCntMax; i++ {
			copy(r.MasuSts[i], tmpMasu[i])
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
	for i := 0; i < r.MasuCnt; i++ {
		// 初期化
		for j := 0; j < r.MasuCnt; j++ {
			r.MasuStsPassB[i][j] = 0
			r.MasuStsAnzB[i][j].Reset()
			r.MasuStsPassW[i][j] = 0
			r.MasuStsAnzW[i][j].Reset()
		}
	}
	// 黒解析
	r.analysisReversiBlack()
	// 白解析
	r.analysisReversiWhite()

	r.makeMasuSts(REVERSI_STS_BLACK)
	r.makeMasuSts(REVERSI_STS_WHITE)

	// *** パスマスを取得 *** //
	for i := 0; i < r.MasuCnt; i++ {
		for j := 0; j < r.MasuCnt; j++ {
			if r.MasuStsPassB[i][j] != 0 {
				if bPassEna != 0 {
					r.MasuStsEnaB[i][j] = 3
				}
			}
			if r.MasuStsPassW[i][j] != 0 {
				if wPassEna != 0 {
					r.MasuStsEnaW[i][j] = 3
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
	if r.checkPara(y, 0, r.MasuCnt) == 0 && r.checkPara(x, 0, r.MasuCnt) == 0 {
		ret = r.MasuSts[y][x]
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
	if r.checkPara(y, 0, r.MasuCnt) == 0 && r.checkPara(x, 0, r.MasuCnt) == 0 {
		ret = r.MasuStsOld[y][x]
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
	if r.checkPara(y, 0, r.MasuCnt) == 0 && r.checkPara(x, 0, r.MasuCnt) == 0 {
		if color == REVERSI_STS_BLACK {
			ret = r.MasuStsEnaB[y][x]
		} else {
			ret = r.MasuStsEnaW[y][x]
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
	if r.checkPara(y, 0, r.MasuCnt) == 0 && r.checkPara(x, 0, r.MasuCnt) == 0 {
		if color == REVERSI_STS_BLACK {
			ret = r.MasuStsCntB[y][x]
		} else {
			ret = r.MasuStsCntW[y][x]
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
	for i := 0; i < r.MasuCnt; i++ {
		for j := 0; j < r.MasuCnt; j++ {
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
		for i := 0; i < r.MasuCntMax; i++ {
			copy(r.MasuStsOld[i], r.MasuSts[i])
		}
		r.MasuSts[y][x] = color
		r.revMasuSts(color, y, x)
		r.makeMasuSts(REVERSI_STS_BLACK)
		r.makeMasuSts(REVERSI_STS_WHITE)
		// *** 履歴保存 *** //
		if r.MasuHistCur < (r.MasuCntMax * r.MasuCntMax) {
			r.MasuHist[r.MasuHistCur].SetColor(color)
			r.MasuHist[r.MasuHistCur].GetPoint().SetY(y)
			r.MasuHist[r.MasuHistCur].GetPoint().SetX(x)
			r.MasuHistCur++
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
	for i := 0; i < r.MasuCntMax; i++ {
		copy(r.MasuStsOld[i], r.MasuSts[i])
	}
	r.MasuSts[y][x] = color
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
	if r.checkPara(cnt, 0, r.MasuCntMax) == 0 {
		if r.MasuCnt != cnt {
			chg = 1
		}
		r.MasuCnt = cnt
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
	if r.checkPara(num, 0, (r.MasuCnt*r.MasuCnt)) == 0 {
		if color == REVERSI_STS_BLACK {
			ret = r.MasuPointB[num]
		} else {
			ret = r.MasuPointW[num]
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
		ret = r.MasuPointCntB
	} else {
		ret = r.MasuPointCntW
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
		ret = r.MasuBetCntB
	} else {
		ret = r.MasuBetCntW
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
	if r.checkPara(y, 0, r.MasuCnt) == 0 && r.checkPara(x, 0, r.MasuCnt) == 0 {
		if color == REVERSI_STS_BLACK {
			ret = r.MasuStsPassB[y][x]
		} else {
			ret = r.MasuStsPassW[y][x]
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
	if r.checkPara(num, 0, (r.MasuCnt*r.MasuCnt)) == 0 {
		ret = r.MasuHist[num]
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
	return r.MasuHistCur
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
	if r.checkPara(y, 0, r.MasuCnt) == 0 && r.checkPara(x, 0, r.MasuCnt) == 0 {
		if color == REVERSI_STS_BLACK {
			ret = r.MasuStsAnzB[y][x]
		} else {
			ret = r.MasuStsAnzW[y][x]
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
		for loop = x; loop < r.MasuCnt; loop++ {
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
		for loop = y; loop < r.MasuCnt; loop++ {
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
		for loop = y; loop < r.MasuCnt; loop++ {
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
	if y == 0 && x == (r.MasuCnt-2) {
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
	if y == 1 && x == (r.MasuCnt-1) {
		flg1 = 0
		flg2 = 0
		for loop = y; loop < r.MasuCnt; loop++ {
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
	if y == 1 && x == (r.MasuCnt-2) {
		loop2 = x
		flg1 = 0
		flg2 = 0
		for loop = y; loop < r.MasuCnt; loop++ {
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
	if y == (r.MasuCnt-2) && x == 0 {
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
	if y == (r.MasuCnt-1) && x == 1 {
		flg1 = 0
		flg2 = 0
		for loop = x; loop < r.MasuCnt; loop++ {
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
	if y == (r.MasuCnt-2) && x == 1 {
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
	if y == (r.MasuCnt-2) && x == (r.MasuCnt-1) {
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
	if y == (r.MasuCnt-1) && x == (r.MasuCnt-2) {
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
	if y == (r.MasuCnt-2) && x == (r.MasuCnt-2) {
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
	if (y == 0 && x == 0) || (y == 0 && x == (r.MasuCnt-1)) || (y == (r.MasuCnt-1) && x == 0) || (y == (r.MasuCnt-1) && x == (r.MasuCnt-1)) {
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
	if (y == 0 && x == 1) || (y == 0 && x == (r.MasuCnt-2)) || (y == 1 && x == 0) || (y == 1 && x == 1) || (y == 1 && x == (r.MasuCnt-2)) || (y == 1 && x == (r.MasuCnt-1)) || (y == (r.MasuCnt-2) && x == 0) || (y == (r.MasuCnt-2) && x == 1) || (y == (r.MasuCnt-2) && x == (r.MasuCnt-2)) || (y == (r.MasuCnt-2) && x == (r.MasuCnt-1)) || (y == (r.MasuCnt-1) && x == 1) || (y == (r.MasuCnt-1) && x == (r.MasuCnt-2)) {
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
	if (y == 0 && x == 2) || (y == 0 && x == (r.MasuCnt-3)) || (y == 2 && x == 0) || (y == 2 && x == 2) || (y == 2 && x == (r.MasuCnt-3)) || (y == 2 && x == (r.MasuCnt-1)) || (y == (r.MasuCnt-3) && x == 0) || (y == (r.MasuCnt-3) && x == 2) || (y == (r.MasuCnt-3) && x == (r.MasuCnt-3)) || (y == (r.MasuCnt-3) && x == (r.MasuCnt-1)) || (y == (r.MasuCnt-1) && x == 2) || (y == (r.MasuCnt-1) && x == (r.MasuCnt-3)) {
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
	if (y == 0 && (3 <= x && x <= (r.MasuCnt-4))) || ((3 <= y && y <= (r.MasuCnt-4)) && x == 0) || (y == (r.MasuCnt-1) && (3 <= x && x <= (r.MasuCnt-4))) || ((3 <= y && y <= (r.MasuCnt-4)) && x == (r.MasuCnt-1)) {
		ret = 0
	}
	return ret
}
