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
	mMasuSts [][]int											//!< マスの状態
	mMasuStsOld [][]int										//!< 以前のマスの状態
	mMasuStsEnaB [][]int										//!< 黒の置ける場所
	mMasuStsCntB [][]int										//!< 黒の獲得コマ数
	mMasuStsPassB [][]int										//!< 黒が相手をパスさせる場所
	mMasuStsAnzB[][]*ReversiAnz								//!< 黒がその場所に置いた場合の解析結果
	mMasuPointB []*ReversiPoint									//!< 黒の置ける場所座標一覧
	mMasuPointCntB int											//!< 黒の置ける場所座標一覧数
	mMasuBetCntB int											//!< 黒コマ数
	mMasuStsEnaW [][]int										//!< 白の置ける場所
	mMasuStsCntW [][]int										//!< 白の獲得コマ数
	mMasuStsPassW [][]int										//!< 白が相手をパスさせる場所
	mMasuStsAnzW [][] *ReversiAnz								//!< 白がその場所に置いた場合の解析結果
	mMasuPointW []	*ReversiPoint								//!< 白の置ける場所座標一覧
	mMasuPointCntW int											//!< 白の置ける場所座標一覧数
	mMasuBetCntW int											//!< 白コマ数
	mMasuCnt int												//!< 縦横マス数
	mMasuCntMax int											//!< 縦横マス最大数
	mMasuHist [] *ReversiHistory									//!< 履歴
	mMasuHistCur int											//!< 履歴現在位置
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
	func NewReversi(masuCnt int,masuMax int) *Reversi {
		r := new(Reversi)
		r.mMasuCnt = masuCnt;
		r.mMasuCntMax = masuMax;

		r.mMasuSts = make([][]int, r.mMasuCntMax)
		r.mMasuStsOld = make([][]int, r.mMasuCntMax)
		r.mMasuStsEnaB = make([][]int, r.mMasuCntMax)
		r.mMasuStsCntB = make([][]int, r.mMasuCntMax)
		r.mMasuStsPassB = make([][]int, r.mMasuCntMax)
		for i:=0; i<r.mMasuCntMax; i++{
			r.mMasuSts[i] = make([]int, r.mMasuCntMax)
			r.mMasuStsOld[i] = make([]int, r.mMasuCntMax)
			r.mMasuStsEnaB[i] = make([]int, r.mMasuCntMax)
			r.mMasuStsCntB[i] = make([]int, r.mMasuCntMax)
			r.mMasuStsPassB[i] = make([]int, r.mMasuCntMax)
		}
		r.mMasuStsAnzB = make([][]*ReversiAnz, r.mMasuCntMax)
		for i:=0; i<r.mMasuCntMax; i++{
			r.mMasuStsAnzB[i] = make([]*ReversiAnz, r.mMasuCntMax)
			for j:=0; j<r.mMasuCntMax; j++{
				r.mMasuStsAnzB[i][j] = NewReversiAnz()
			}
		}
		r.mMasuPointB = make([]*ReversiPoint, r.mMasuCntMax * r.mMasuCntMax)
		for i:=0; i<(r.mMasuCntMax * r.mMasuCntMax); i++{
			r.mMasuPointB[i] = NewReversiPoint()
		}
		r.mMasuPointCntB = 0;
		r.mMasuStsEnaW = make([][]int, r.mMasuCntMax)
		r.mMasuStsCntW = make([][]int, r.mMasuCntMax)
		r.mMasuStsPassW = make([][]int, r.mMasuCntMax)
		for i:=0; i<r.mMasuCntMax; i++{
			r.mMasuStsEnaW[i] = make([]int, r.mMasuCntMax)
			r.mMasuStsCntW[i] = make([]int, r.mMasuCntMax)
			r.mMasuStsPassW[i] = make([]int, r.mMasuCntMax)
		}
		r.mMasuStsAnzW = make([][]*ReversiAnz, r.mMasuCntMax)
		for i:=0; i<r.mMasuCntMax; i++{
			r.mMasuStsAnzW[i] = make([]*ReversiAnz, r.mMasuCntMax)
			for j:=0; j<r.mMasuCntMax; j++{
				r.mMasuStsAnzW[i][j] = NewReversiAnz()
			}
		}
		r.mMasuPointW = make([]*ReversiPoint, r.mMasuCntMax * r.mMasuCntMax)
		for i:=0; i<(r.mMasuCntMax * r.mMasuCntMax); i++{
			r.mMasuPointW[i] = NewReversiPoint()
		}
		r.mMasuPointCntW	= 0;
		r.mMasuBetCntB = 0;
		r.mMasuBetCntW = 0;
		r.mMasuHist = make([]*ReversiHistory, r.mMasuCntMax * r.mMasuCntMax)
		for i:=0; i<(r.mMasuCntMax * r.mMasuCntMax); i++{
			r.mMasuHist[i] = NewReversiHistory()
		}

		r.mMasuHistCur = make([]*ReversiHistory, r.mMasuCntMax * r.mMasuCntMax)
		for i:=0; i<(r.mMasuCntMax * r.mMasuCntMax); i++{
			r.mMasuHistCur[i] = NewReversiHistory()
		}
		r.mMasuHistCur = 0;
		copy(r.mMasuStsOld,r.mMasuSts)
		r.Reset();
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
		for i:=0; i<r.mMasuCnt; i++ {
			for j:=0; j<r.mMasuCnt; j++ {
				r.mMasuSts[i][j] = REVERSI_STS_NONE;
				r.mMasuStsPassB[i][j] = 0;
				r.mMasuStsAnzB[i][j].Reset();
				r.mMasuStsPassW[i][j] = 0;
				r.mMasuStsAnzW[i][j].Reset();
			}
		}
		r.mMasuSts[(r.mMasuCnt >> 1) - 1][(r.mMasuCnt >> 1) - 1]	= REVERSI_STS_BLACK;
		r.mMasuSts[(r.mMasuCnt >> 1) - 1][(r.mMasuCnt >> 1)]		= REVERSI_STS_WHITE;
		r.mMasuSts[(r.mMasuCnt >> 1)][(r.mMasuCnt >> 1) - 1]		= REVERSI_STS_WHITE;
		r.mMasuSts[(r.mMasuCnt >> 1)][(r.mMasuCnt >> 1)]			= REVERSI_STS_BLACK;
		r.makeMasuSts(REVERSI_STS_BLACK);
		r.makeMasuSts(REVERSI_STS_WHITE);
		r.mMasuHistCur = 0;
		copy(r.mMasuStsOld,r.mMasuSts)
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
		for  i := 0;i < r.mMasuCnt; i++ {								// 初期化
			for j := 0;j < r.mMasuCnt; j++ {
				if color == REVERSI_STS_BLACK {
					r.mMasuStsEnaB[i][j] = 0
					r.mMasuStsCntB[i][j] = 0
				}else{
					r.mMasuStsEnaW[i][j] = 0
					r.mMasuStsCntW[i][j] = 0
				}
			}
		}

		loop = r.mMasuCnt * r.mMasuCnt;
		for i := 0;i < loop;i++ {										// 初期化
			if color == REVERSI_STS_BLACK {
				r.mMasuPointB[i].SetX(0)
				r.mMasuPointB[i].SetY(0)
			}else{
				r.mMasuPointW[i].SetX(0)
				r.mMasuPointW[i].SetY(0)
			}
		}
		if color == REVERSI_STS_BLACK {
			r.mMasuPointCntB	= 0
		}else{
			r.mMasuPointCntW	= 0
		}
		r.mMasuBetCntB	= 0
		r.mMasuBetCntW	= 0

		for i := 0;i < r.mMasuCnt;i++ {
			for j := 0;j < r.mMasuCnt;j++ {
				okflg = 0
				count2 = 0
				if r.mMasuSts[i][j] == REVERSI_STS_NONE {	// 何も置かれていないマスなら
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
					for (cnt1 < (r.mMasuCnt-1)) && (r.mMasuSts[cnt1+1][j] != REVERSI_STS_NONE && r.mMasuSts[cnt1+1][j] != color) {
						flg = 1
						cnt1++
						count1++
					}
					if (cnt1 < (r.mMasuCnt-1)) && (flg == 1) && (r.mMasuSts[cnt1+1][j] == color) {
						okflg = 1
						count2 += count1
					}
					cnt2 = j
					count1 = 0
					flg = 0
					// *** 右方向を調べる *** //
					for (cnt2 < (r.mMasuCnt-1)) && (r.mMasuSts[i][cnt2+1] != REVERSI_STS_NONE && r.mMasuSts[i][cnt2+1] != color) {
						flg = 1
						cnt2++
						count1++
					}
					if (cnt2 < (r.mMasuCnt-1)) && (flg == 1) && (r.mMasuSts[i][cnt2+1] == color) {
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
					for ((cnt2 < (r.mMasuCnt-1)) && (cnt1 > 0)) && (r.mMasuSts[cnt1-1][cnt2+1] != REVERSI_STS_NONE && r.mMasuSts[cnt1-1][cnt2+1] != color) {
						flg = 1
						cnt1--
						cnt2++
						count1++
					}
					if ((cnt2 < (r.mMasuCnt-1)) && (cnt1 > 0)) && (flg == 1) && (r.mMasuSts[cnt1-1][cnt2+1] == color) {
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
					for ((cnt2 < (r.mMasuCnt-1)) && (cnt1 < (r.mMasuCnt-1))) && (r.mMasuSts[cnt1+1][cnt2+1] != REVERSI_STS_NONE && r.mMasuSts[cnt1+1][cnt2+1] != color) {
						flg = 1
						cnt1++
						cnt2++
						count1++
					}
					if ((cnt2 < (r.mMasuCnt-1)) && (cnt1 < (r.mMasuCnt-1))) && (flg == 1) && (r.mMasuSts[cnt1+1][cnt2+1] == color) {
						okflg = 1
						count2 += count1
					}
					cnt2 = j
					cnt1 = i
					count1 = 0
					flg = 0
					// *** 左下方向を調べる *** //
					for ((cnt2 > 0) && (cnt1 < (r.mMasuCnt-1))) && (r.mMasuSts[cnt1+1][cnt2-1] != REVERSI_STS_NONE && r.mMasuSts[cnt1+1][cnt2-1] != color) {
						flg = 1
						cnt1++
						cnt2--
						count1++
					}
					if ((cnt2 > 0) && (cnt1 < (r.mMasuCnt-1))) && (flg == 1) && (r.mMasuSts[cnt1+1][cnt2-1] == color) {
						okflg = 1
						count2 += count1
					}
					if okflg == 1 {
						if color == REVERSI_STS_BLACK {
							r.mMasuStsEnaB[i][j] = 1
							r.mMasuStsCntB[i][j] = count2
							// *** 置ける場所をリニアに保存、置けるポイント数も保存 *** //
							r.mMasuPointB[this.mMasuPointCntB].SetY(i)
							r.mMasuPointB[this.mMasuPointCntB].SetX(j)
							r.mMasuPointCntB++
						}else{
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
				}else if r.mMasuSts[i][j] == REVERSI_STS_BLACK {
					r.mMasuBetCntB++
				}else if r.mMasuSts[i][j] == REVERSI_STS_WHITE {
					r.mMasuBetCntW++
				}
			}
		}

		// *** 一番枚数を獲得できるマスを設定 *** //
		for i := 0;i < r.mMasuCnt;i++ {
			for j := 0;j < r.mMasuCnt;j++ {
				if color == REVERSI_STS_BLACK {
					if r.mMasuStsEnaB[i][j] != 0 && r.mMasuStsCntB[i][j] == countMax {
						r.mMasuStsEnaB[i][j] = 2
					}
				}else{
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
	func (r *Reversi) revMasuSts(color int,y int,x int)	{
		var cnt1 int
		var cnt2 int
		var rcnt1 int
		var rcnt2 int
		var flg int = 0

		// *** 左方向にある駒を調べる *** //
		cnt1 = x
		cnt2 = y
		for flg = 0;cnt1 > 0; {
			if r.mMasuSts[cnt2][cnt1-1] != REVERSI_STS_NONE && r.mMasuSts[cnt2][cnt1-1] != color {
				// *** プレイヤー以外の色の駒があるなら *** //
				cnt1--
			}else if r.mMasuSts[cnt2][cnt1-1] == color {
				flg = 1
				break
			}else if r.mMasuSts[cnt2][cnt1-1] == REVERSI_STS_NONE {
				break
			}
		}
		if flg == 1 {
			// *** 駒をひっくり返す *** //
			for rcnt1 = cnt1;rcnt1 < x;rcnt1++ {
				r.mMasuSts[cnt2][rcnt1] = color
			}
		}

		// *** 右方向にある駒を調べる *** //
		cnt1 = x
		cnt2 = y
		for flg = 0;cnt1 < (r.mMasuCnt-1); {
			if r.mMasuSts[cnt2][cnt1+1] != REVERSI_STS_NONE && r.mMasuSts[cnt2][cnt1+1] != color {
				// *** プレイヤー以外の色の駒があるなら *** //
				cnt1++
			}else if r.mMasuSts[cnt2][cnt1+1] == color {
				flg = 1
				break
			}else if r.mMasuSts[cnt2][cnt1+1] == REVERSI_STS_NONE {
				break
			}
		}
		if flg == 1 {
			// *** 駒をひっくり返す *** //
			for rcnt1 = cnt1;rcnt1 > x;rcnt1-- {
				r.mMasuSts[cnt2][rcnt1] = color
			}
		}

		// *** 上方向にある駒を調べる *** //
		cnt1 = x
		cnt2 = y
		for flg = 0;cnt2 > 0; {
			if r.mMasuSts[cnt2-1][cnt1] != REVERSI_STS_NONE && r.mMasuSts[cnt2-1][cnt1] != color {
				// *** プレイヤー以外の色の駒があるなら *** //
				cnt2--
			}else if r.mMasuSts[cnt2-1][cnt1] == color {
				flg = 1
				break
			}else if r.mMasuSts[cnt2-1][cnt1] == REVERSI_STS_NONE {
				break
			}
		}
		if flg == 1 {
			// *** 駒をひっくり返す *** //
			for rcnt1 = cnt2;rcnt1 < y;rcnt1++ {
				r.mMasuSts[rcnt1][cnt1] = color
			}
		}

		// *** 下方向にある駒を調べる *** //
		cnt1 = x
		cnt2 = y
		for flg = 0;cnt2 < (r.mMasuCnt-1); {
			if r.mMasuSts[cnt2+1][cnt1] != REVERSI_STS_NONE && r.mMasuSts[cnt2+1][cnt1] != color {
				// *** プレイヤー以外の色の駒があるなら *** //
				cnt2++
			}else if r.mMasuSts[cnt2+1][cnt1] == color {
				flg = 1
				break
			}else if r.mMasuSts[cnt2+1][cnt1] == REVERSI_STS_NONE {
				break
			}
		}
		if flg == 1 {
			// *** 駒をひっくり返す *** //
			for rcnt1 = cnt2;rcnt1 > y;rcnt1-- {
				r.mMasuSts[rcnt1][cnt1] = color
			}
		}

		// *** 左上方向にある駒を調べる *** //
		cnt1 = x
		cnt2 = y
		for flg = 0;cnt2 > 0 && cnt1 > 0; {
			if r.mMasuSts[cnt2-1][cnt1-1] != REVERSI_STS_NONE && r.mMasuSts[cnt2-1][cnt1-1] != color {
				// *** プレイヤー以外の色の駒があるなら *** //
				cnt2--
				cnt1--
			}else if r.mMasuSts[cnt2-1][cnt1-1] == color {
				flg = 1
				break
			}else if r.mMasuSts[cnt2-1][cnt1-1] == REVERSI_STS_NONE {
				break
			}
		}
		if flg == 1 {
			// *** 駒をひっくり返す *** //
			rcnt2 = cnt1
			for rcnt1 = cnt2;(rcnt1 < y) && (rcnt2 < x);rcnt1++ {
				r.mMasuSts[rcnt1][rcnt2] = color
				rcnt2++
			}
		}

		// *** 左下方向にある駒を調べる *** //
		cnt1 = x
		cnt2 = y
		for flg = 0;cnt2 < (r.mMasuCnt-1) && cnt1 > 0; {
			if r.mMasuSts[cnt2+1][cnt1-1] != REVERSI_STS_NONE && r.mMasuSts[cnt2+1][cnt1-1] != color {
				// *** プレイヤー以外の色の駒があるなら *** //
				cnt2++
				cnt1--
			}else if r.mMasuSts[cnt2+1][cnt1-1] == color {
				flg = 1
				break
			}else if r.mMasuSts[cnt2+1][cnt1-1] == REVERSI_STS_NONE {
				break
			}
		}
		if flg == 1 {
			// *** 駒をひっくり返す *** //
			rcnt2 = cnt1
			for rcnt1 = cnt2;(rcnt1 > y) && (rcnt2 < x);rcnt1-- {
				r.mMasuSts[rcnt1][rcnt2] = color
				rcnt2++
			}
		}

		// *** 右上方向にある駒を調べる *** //
		cnt1 = x
		cnt2 = y
		for flg = 0;cnt2 > 0 && cnt1 < (r.mMasuCnt-1); {
			if r.mMasuSts[cnt2-1][cnt1+1] != REVERSI_STS_NONE && r.mMasuSts[cnt2-1][cnt1+1] != color {
				// *** プレイヤー以外の色の駒があるなら *** //
				cnt2--
				cnt1++
			}else if r.mMasuSts[cnt2-1][cnt1+1] == color {
				flg = 1
				break
			}else if r.mMasuSts[cnt2-1][cnt1+1] == REVERSI_STS_NONE {
				break
			}
		}
		if flg == 1 {
			// *** 駒をひっくり返す *** //
			rcnt2 = cnt1
			for rcnt1 = cnt2;(rcnt1 < y) && (rcnt2 > x);rcnt1++ {
				r.mMasuSts[rcnt1][rcnt2] = color
				rcnt2--
			}
		}

		// *** 右下方向にある駒を調べる *** //
		cnt1 = x
		cnt2 = y
		for flg = 0;cnt2 < (r.mMasuCnt-1) && cnt1 < (r.mMasuCnt-1); {
			if r.mMasuSts[cnt2+1][cnt1+1] != REVERSI_STS_NONE && r.mMasuSts[cnt2+1][cnt1+1] != color {
				// *** プレイヤー以外の色の駒があるなら *** //
				cnt2++
				cnt1++
			}else if r.mMasuSts[cnt2+1][cnt1+1] == color {
				flg = 1
				break
			}else if r.mMasuSts[cnt2+1][cnt1+1] == REVERSI_STS_NONE {
				break
			}
		}
		if flg == 1 {
			// *** 駒をひっくり返す *** //
			rcnt2 = cnt1
			for rcnt1 = cnt2;(rcnt1 > y) && (rcnt2 > x);rcnt1-- {
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
	func (r Reversi) checkPara(para int,min int,max int) int {
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
		for cnt := 0;cnt < r.mMasuPointCntB;cnt++ {
			// *** 現在ステータスを一旦コピー *** //
			var tmpMasu [][]int
			var tmpMasuEnaB [][]int
			var tmpMasuEnaW [][]int
			copy(tmpMasu,r.mMasuSts)
			copy(tmpMasuEnaB,r.mMasuStsEnaB)
			copy(tmpMasuEnaW,r.mMasuStsEnaW)
			tmpY = r.mMasuPointB[cnt].GetY()
			tmpX = r.mMasuPointB[cnt].GetX()
			// 仮に置く
			r.mMasuSts[tmpY][tmpX] = REVERSI_STS_BLACK
			// 仮にひっくり返す
			r.revMasuSts(REVERSI_STS_BLACK,tmpY,tmpX)
			r.makeMasuSts(REVERSI_STS_BLACK)
			r.makeMasuSts(REVERSI_STS_WHITE)
			// *** このマスに置いた場合の解析を行う *** //
			if r.getColorEna(REVERSI_STS_WHITE) != 0 {
				// 相手がパス
				r.mMasuStsPassB[tmpY][tmpX] = 1
			}
			if r.getEdgeSideZero(tmpY,tmpX) == 0 {
				// 置いた場所が角
				r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeCnt() + 1)
				r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 10000 * r.mMasuStsCntB[tmpY][tmpX])
			}else if r.getEdgeSideOne(tmpY,tmpX) == 0 {
				// 置いた場所が角の一つ手前
				r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
				if r.checkEdge(REVERSI_STS_BLACK,tmpY,tmpX) != 0 {
					// 角を取られない
					r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 10 * r.mMasuStsCntB[tmpY][tmpX])
				}else{
					// 角を取られる
					r.mMasuStsAnzB[tmpY][tmpX].SetBadPoint(r.mMasuStsAnzB[tmpY][tmpX].GetBadPoint() + 100000)
				}
			}else if r.getEdgeSideTwo(tmpY,tmpX) == 0 {
				// 置いた場所が角の二つ手前
				r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
				r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 1000 * r.mMasuStsCntB[tmpY][tmpX])
			}else if r.getEdgeSideThree(tmpY,tmpX) == 0 {
				// 置いた場所が角の三つ手前
				r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.mMasuStsAnzB[tmpY][tmpX].getOwnEdgeSideThreeCnt() + 1)
				r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 100 * r.mMasuStsCntB[tmpY][tmpX])
			}else{
				// 置いた場所がその他
				r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideOtherCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideOtherCnt() + 1)
				r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + 10 * r.mMasuStsCntB[tmpY][tmpX])
			}
			sum = 0
			sumOwn = 0
			for i := 0;i < r.mMasuCnt;i++ {
				for j := 0;j < r.mMasuCnt;j++ {
					tmpBadPoint = 0
					tmpGoodPoint = 0
					if r.getMasuStsEna(REVERSI_STS_WHITE,i,j) != 0 {
						// 相手の獲得予定枚数
						sum += this.mMasuStsCntW[i][j]
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
						if r.getEdgeSideZero(i,j) == 0 {
							// 置く場所が角
							r.mMasuStsAnzB[tmpY][tmpX].SetEdgeCnt(r.mMasuStsAnzB[tmpY][tmpX].GetEdgeCnt() + 1)
							tmpBadPoint = 100000 * r.mMasuStsCntW[i][j]
						}else if r.getEdgeSideOne(i,j) == 0 {
							// 置く場所が角の一つ手前
							r.mMasuStsAnzB[tmpY][tmpX].SetEdgeSideOneCnt(r.mMasuStsAnzB[tmpY][tmpX].GetEdgeSideOneCnt() + 1)
							tmpBadPoint = 0
						}else if r.getEdgeSideTwo(i,j) == 0 {
							// 置く場所が角の二つ手前
							r.mMasuStsAnzB[tmpY][tmpX].SetEdgeSideTwoCnt(r.mMasuStsAnzB[tmpY][tmpX].GetEdgeSideTwoCnt() + 1)
							tmpBadPoint = 1 * r.mMasuStsCntW[i][j]
						}else if r.getEdgeSideThree(i,j) == 0 {
							// 置く場所が角の三つ手前
							r.mMasuStsAnzB[tmpY][tmpX].SetEdgeSideThreeCnt(r.mMasuStsAnzB[tmpY][tmpX].GetEdgeSideThreeCnt() + 1)
							tmpBadPoint = 1 * r.mMasuStsCntW[i][j]
						}else{
							// 置く場所がその他
							r.mMasuStsAnzB[tmpY][tmpX].SetEdgeSideOtherCnt(r.mMasuStsAnzB[tmpY][tmpX].GetEdgeSideOtherCnt() + 1)
							tmpBadPoint = 1 * r.mMasuStsCntW[i][j]
						}
						if tmpMasuEnaW[i][j] != 0 {
							// ステータス変化していないなら
							tmpBadPoint = 0
						}
					}
					if r.getMasuStsEna(REVERSI_STS_BLACK,i,j) != 0 {
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
						if r.getEdgeSideZero(i,j) == 0 {
							// 置く場所が角
							r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeCnt() + 1)
							tmpGoodPoint = 100 * this.mMasuStsCntB[i][j]
						}else if r.getEdgeSideOne(i,j) == 0 {
							// 置く場所が角の一つ手前
							r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
							tmpGoodPoint = 0
						}else if r.getEdgeSideTwo(i,j) == 0 {
							// 置く場所が角の二つ手前
							r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
							tmpGoodPoint = 3 * r.mMasuStsCntB[i][j]
						}else if r.getEdgeSideThree(i,j) == 0 {
							// 置く場所が角の三つ手前
							r.mMasuStsAnzB[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.mMasuStsAnzB[tmpY][tmpX].GetOwnEdgeSideThreeCnt() + 1)
							tmpGoodPoint = 2 * r.mMasuStsCntB[i][j];
						}else{
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
						r.mMasuStsAnzB[tmpY][tmpX].SetBadPoint(r.mMasuStsAnzB[tmpY][tmpX].GetBadPoint()	+ tmpBadPoint)
					}
					if tmpGoodPoint != 0 {
						r.mMasuStsAnzB[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzB[tmpY][tmpX].GetGoodPoint() + tmpGoodPoint)
					}
				}
			}
			// *** 相手に取られる平均コマ数 *** //
			if r.getPointCnt(REVERSI_STS_WHITE) != 0 {
				tmpD1 = float64(sum)
				tmpD2 = float64(r.getPointCnt(REVERSI_STS_WHITE))
				R.mMasuStsAnzB[tmpY][tmpX].setAvg(tmpD1 / tmpD2)
			}
			// *** 自分が取れる平均コマ数 *** //
			if r.getPointCnt(REVERSI_STS_BLACK) != 0 {
				tmpD1 = float64(sumOwn)
				tmpD2 = float64(r.getPointCnt(REVERSI_STS_BLACK))
				r.mMasuStsAnzB[tmpY][tmpX].SetOwnAvg(tmpD1 / tmpD2)
			}
			// *** 元に戻す *** //
			copy(r.mMasuSts,tmpMasu)
			r.makeMasuSts(REVERSI_STS_BLACK);
			r.makeMasuSts(REVERSI_STS_WHITE);
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
		for cnt := 0;cnt < r.mMasuPointCntW;cnt++ {
			// *** 現在ステータスを一旦コピー *** //
			var tmpMasu [][]int
			var tmpMasuEnaB [][]int
			var tmpMasuEnaW [][]int
			copy(tmpMasu,r.mMasuSts)
			copy(tmpMasuEnaB,r.mMasuStsEnaB)
			copy(tmpMasuEnaW,r.mMasuStsEnaW)
			tmpY = r.mMasuPointW[cnt].GetY()
			tmpX = r.mMasuPointW[cnt].GetX()
			// 仮に置く
			r.mMasuSts[tmpY][tmpX] = REVERSI_STS_WHITE
			// 仮にひっくり返す
			r.revMasuSts(REVERSI_STS_WHITE,tmpY,tmpX)
			r.makeMasuSts(REVERSI_STS_BLACK)
			r.makeMasuSts(REVERSI_STS_WHITE)
			// *** このマスに置いた場合の解析を行う *** //
			if r.getColorEna(REVERSI_STS_BLACK) != 0 {
				// 相手がパス
				r.mMasuStsPassW[tmpY][tmpX] = 1
			}
			if r.getEdgeSideZero(tmpY,tmpX) == 0 {
				// 置いた場所が角
				r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeCnt() + 1)
				r.mMasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 10000 * r.mMasuStsCntW[tmpY][tmpX])
			}else if r.getEdgeSideOne(tmpY,tmpX) == 0 {
				// 置いた場所が角の一つ手前
				r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
				if r.checkEdge(REVERSI_STS_WHITE,tmpY,tmpX) != 0 {
					// 角を取られない
					r.mMasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 10 * r.mMasuStsCntW[tmpY][tmpX])
				}else{
					// 角を取られる
					r.mMasuStsAnzW[tmpY][tmpX].SetBadPoint(r.mMasuStsAnzW[tmpY][tmpX].GetBadPoint() + 100000)
				}
			}else if r.getEdgeSideTwo(tmpY,tmpX) == 0{
				// 置いた場所が角の二つ手前
				r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
				r.mMasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 1000 * r.mMasuStsCntW[tmpY][tmpX])
			}else if(this.getEdgeSideThree(tmpY,tmpX) == 0){
				// 置いた場所が角の三つ手前
				r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideThreeCnt() + 1)
				r.mMasuStsAnzW[tmpY][tmpX].SetGoodPoint(r.mMasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 100 * r.mMasuStsCntW[tmpY][tmpX])
			}else{
				// 置いた場所がその他
				r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOtherCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOtherCnt() + 1)
				r.mMasuStsAnzW[tmpY][tmpX].SetGoodPoint(thris.mMasuStsAnzW[tmpY][tmpX].GetGoodPoint() + 10 * r.mMasuStsCntW[tmpY][tmpX])
			}
			sum = 0
			sumOwn = 0
			for i := 0;i < r.mMasuCnt;i++ {
				for j := 0;j < r.mMasuCnt;j++ {
					tmpBadPoint = 0
					tmpGoodPoint = 0
					if r.getMasuStsEna(REVERSI_STS_BLACK,i,j) != 0 {
						// 相手の獲得予定枚数
						sum += r.mMasuStsCntB[i][j];
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
						if r.getEdgeSideZero(i,j) == 0 {
							// 置く場所が角
							r.mMasuStsAnzW[tmpY][tmpX].SetEdgeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetEdgeCnt() + 1)
							tmpBadPoint = 100000 * r.mMasuStsCntB[i][j]
						}else if r.getEdgeSideOne(i,j) == 0 {
							// 置く場所が角の一つ手前
							r.mMasuStsAnzW[tmpY][tmpX].SetEdgeSideOneCnt(r.mMasuStsAnzW[tmpY][tmpX].GetEdgeSideOneCnt() + 1)
							tmpBadPoint = 0
						}else if r.getEdgeSideTwo(i,j) == 0 {
							// 置く場所が角の二つ手前
							r.mMasuStsAnzW[tmpY][tmpX].SetEdgeSideTwoCnt(r.mMasuStsAnzW[tmpY][tmpX].GetEdgeSideTwoCnt() + 1)
							tmpBadPoint = 1 * r.mMasuStsCntB[i][j]
						}else if r.getEdgeSideThree(i,j) == 0 {
							// 置く場所が角の三つ手前
							r.mMasuStsAnzW[tmpY][tmpX].SetEdgeSideThreeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetEdgeSideThreeCnt() + 1)
							tmpBadPoint = 1 * r.mMasuStsCntB[i][j]
						}else{
							// 置く場所がその他
							r.mMasuStsAnzW[tmpY][tmpX].SetEdgeSideOtherCnt(r.mMasuStsAnzW[tmpY][tmpX].GetEdgeSideOtherCnt() + 1)
							tmpBadPoint = 1 * r.mMasuStsCntB[i][j]
						}
						if tmpMasuEnaB[i][j] != 0 {
							// ステータス変化していないなら
							tmpBadPoint = 0
						}
					}
					if r.getMasuStsEna(REVERSI_STS_WHITE,i,j) != 0 {
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
						if r.getEdgeSideZero(i,j) == 0 {
							// 置く場所が角
							r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeCnt() + 1)
							tmpGoodPoint = 100 * r.mMasuStsCntW[i][j]
						}else if r.getEdgeSideOne(i,j) == 0 {
							// 置く場所が角の一つ手前
							r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOneCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOneCnt() + 1)
							tmpGoodPoint = 0
						}else if r.getEdgeSideTwo(i,j) == 0 {
							// 置く場所が角の二つ手前
							r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideTwoCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideTwoCnt() + 1)
							tmpGoodPoint = 3 * r.mMasuStsCntW[i][j]
						}else if r.getEdgeSideThree(i,j) == 0 {
							// 置く場所が角の三つ手前
							r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideThreeCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideThreeCnt() + 1)
							tmpGoodPoint = 2 * r.mMasuStsCntW[i][j]
						}else{
							// 置く場所がその他
							r.mMasuStsAnzW[tmpY][tmpX].SetOwnEdgeSideOtherCnt(r.mMasuStsAnzW[tmpY][tmpX].GetOwnEdgeSideOtherCnt() + 1)
							tmpGoodPoint = 1 * r.mMasuStsCntW[i][j];
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
			if r.getPointCnt(REVERSI_STS_BLACK) != 0 {
				tmpD1 = float64(sum)
				tmpD2 = float64(r.getPointCnt(REVERSI_STS_BLACK))
				mMasuStsAnzW[tmpY][tmpX].SetAvg(tmpD1 / tmpD2)
			}
			// *** 自分が取れる平均コマ数 *** //
			if r.getPointCnt(REVERSI_STS_WHITE) != 0 {
				tmpD1 = float64(sumOwn)
				tmpD2 = float64(r.getPointCnt(REVERSI_STS_WHITE))
				r.mMasuStsAnzW[tmpY][tmpX].SetOwnAvg(tmpD1 / tmpD2)
			}
			// *** 元に戻す *** //
			copy(r.mMasuSts,tmpMasu)
			r.makeMasuSts(REVERSI_STS_BLACK);
			r.makeMasuSts(REVERSI_STS_WHITE);
		}
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			解析を行う
	///	@fn				public void AnalysisReversi(int bPassEna,int wPassEna)
	///	@param[in]		int bPassEna		1=黒パス有効
	///	@param[in]		int wPassEna		1=白パス有効
	///	@return			ありません
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public void AnalysisReversi(int bPassEna,int wPassEna)
	{
		// *** 相手をパスさせることができるマス検索 *** //
		for(int i = 0;i < this.mMasuCnt;i++){						// 初期化
			for(int j = 0;j < this.mMasuCnt;j++){
				this.mMasuStsPassB[i][j] = 0;
				this.mMasuStsAnzB[i][j].reset();
				this.mMasuStsPassW[i][j] = 0;
				this.mMasuStsAnzW[i][j].reset();
			}
		}
		AnalysisReversiBlack();										// 黒解析
		AnalysisReversiWhite();										// 白解析

		this.makeMasuSts(ReversiConst.REVERSI_STS_BLACK);
		this.makeMasuSts(ReversiConst.REVERSI_STS_WHITE);

		// *** パスマスを取得 *** //
		for(int i = 0;i < this.mMasuCnt;i++){
			for(int j = 0;j < this.mMasuCnt;j++){
				if(this.mMasuStsPassB[i][j] != 0){
					if(bPassEna != 0) this.mMasuStsEnaB[i][j] = 3;
				}
				if(this.mMasuStsPassW[i][j] != 0){
					if(wPassEna != 0) this.mMasuStsEnaW[i][j] = 3;
				}
			}
		}
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			マスステータスを取得
	///	@fn				public int getMasuSts(int y,int x)
	///	@param[in]		int y			取得するマスのY座標
	///	@param[in]		int x			取得するマスのX座標
	///	@return			-1 : 失敗 それ以外 : マスステータス
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getMasuSts(int y,int x)
	{
		int ret = -1;
		if(this.checkPara(y,0,this.mMasuCnt) == 0 && this.checkPara(x,0,this.mMasuCnt) == 0) ret = this.mMasuSts[y][x];
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			以前のマスステータスを取得
	///	@fn				int getMasuStsOld(int y, int x)
	///	@param[in]		int y			取得するマスのY座標
	///	@param[in]		int x			取得するマスのX座標
	///	@return			-1 : 失敗 それ以外 : マスステータス
	///	@author			Yuta Yoshinaga
	///	@date			2017.10.20
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getMasuStsOld(int y, int x)
	{
		int ret = -1;
		if (this.checkPara(y, 0, this.mMasuCnt) == 0 && this.checkPara(x, 0, this.mMasuCnt) == 0) ret = this.mMasuStsOld[y][x];
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			指定座標に指定色を置けるかチェック
	///	@fn				public int getMasuStsEna(int color,int y,int x)
	///	@param[in]		int color		コマ色
	///	@param[in]		int y			マスのY座標
	///	@param[in]		int x			マスのX座標
	///	@return			1 : 成功 それ以外 : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getMasuStsEna(int color,int y,int x)
	{
		int ret = 0;
		if(this.checkPara(y,0,this.mMasuCnt) == 0 && this.checkPara(x,0,this.mMasuCnt) == 0){
			if(color == ReversiConst.REVERSI_STS_BLACK)	ret = this.mMasuStsEnaB[y][x];
			else										ret = this.mMasuStsEnaW[y][x];
		}
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			指定座標の獲得コマ数取得
	///	@fn				public int getMasuStsCnt(int color,int y,int x)
	///	@param[in]		int color		コマ色
	///	@param[in]		int y			マスのY座標
	///	@param[in]		int x			マスのX座標
	///	@return			-1 : 失敗 それ以外 : 獲得コマ数
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getMasuStsCnt(int color,int y,int x)
	{
		int ret = -1;
		if(this.checkPara(y,0,this.mMasuCnt) == 0 && this.checkPara(x,0,this.mMasuCnt) == 0){
			if(color == ReversiConst.REVERSI_STS_BLACK)	ret = this.mMasuStsCntB[y][x];
			else										ret = this.mMasuStsCntW[y][x];
		}
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			指定色が現在置ける場所があるかチェック
	///	@fn				public int getColorEna(int color)
	///	@param[in]		int color		コマ色
	///	@return			0 : 成功 それ以外 : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getColorEna(int color)
	{
		int ret = -1;
		for(int i=0;i<this.mMasuCnt;i++){
			for(int j=0;j<this.mMasuCnt;j++){
				if(this.getMasuStsEna(color,i,j) != 0){
					ret = 0;
					break;
				}
			}
		}
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			ゲーム終了かチェック
	///	@fn				public int getGameEndSts()
	///	@return			0 : 続行 それ以外 : ゲーム終了
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getGameEndSts()
	{
		int ret = 1;
		if(getColorEna(ReversiConst.REVERSI_STS_BLACK) == 0) ret = 0;
		if(getColorEna(ReversiConst.REVERSI_STS_WHITE) == 0) ret = 0;
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			指定座標にコマを置く
	///	@fn				public int setMasuSts(int color,int y,int x)
	///	@param[in]		int color		コマ色
	///	@param[in]		int y			マスのY座標
	///	@param[in]		int x			マスのX座標
	///	@return			0 : 成功 それ以外 : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int setMasuSts(int color,int y,int x)
	{
		int ret = -1;
		if(this.getMasuStsEna(color,y,x) != 0){
			ret = 0;
			for (int i = 0; i < mMasuCntMax; i++){
				System.arraycopy(this.mMasuSts[i], 0, this.mMasuStsOld[i], 0, this.mMasuSts[i].length);
			}
			this.mMasuSts[y][x] = color;
			this.revMasuSts(color,y,x);
			this.makeMasuSts(ReversiConst.REVERSI_STS_BLACK);
			this.makeMasuSts(ReversiConst.REVERSI_STS_WHITE);
			// *** 履歴保存 *** //
			if(this.mMasuHistCur < (this.mMasuCntMax * this.mMasuCntMax)){
				this.mMasuHist[this.mMasuHistCur].setColor(color);
				this.mMasuHist[this.mMasuHistCur].getPoint().setY(y);
				this.mMasuHist[this.mMasuHistCur].getPoint().setX(x);
				this.mMasuHistCur++;
			}
		}
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			指定座標にコマを強制的に置く
	///	@fn				public int setMasuStsForcibly(int color,int y,int x)
	///	@param[in]		int color		コマ色
	///	@param[in]		int y			マスのY座標
	///	@param[in]		int x			マスのX座標
	///	@return			0 : 成功 それ以外 : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int setMasuStsForcibly(int color,int y,int x)
	{
		int ret = -1;
		ret = 0;
		for (int i = 0; i < mMasuCntMax; i++){
			System.arraycopy(this.mMasuSts[i], 0, this.mMasuStsOld[i], 0, this.mMasuSts[i].length);
		}
		this.mMasuSts[y][x] = color;
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			マスの数変更
	///	@fn				public int setMasuCnt(int cnt)
	///	@param[in]		int cnt		マスの数
	///	@return			0 : 成功 それ以外 : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int setMasuCnt(int cnt)
	{
		int ret = -1,chg = 0;

		if(checkPara(cnt,0,this.mMasuCntMax) == 0){
			if(this.mMasuCnt != cnt) chg = 1;
			this.mMasuCnt = cnt;
			ret = 0;
			if(chg == 1) this.reset();
		}

		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			ポイント座標取得
	///	@fn				public ReversiPoint getPoint(int color,int num)
	///	@param[in]		int color		コマ色
	///	@param[in]		int num			ポイント
	///	@return			ポイント座標 null : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public ReversiPoint getPoint(int color,int num)
	{
		ReversiPoint ret = null;
		if(this.checkPara(num,0,(this.mMasuCnt * this.mMasuCnt)) == 0){
			if(color == ReversiConst.REVERSI_STS_BLACK)	ret = this.mMasuPointB[num];
			else										ret = this.mMasuPointW[num];
		}
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			ポイント座標数取得
	///	@fn				public int getPointCnt(int color)
	///	@param[in]		int color		コマ色
	///	@return			ポイント座標数取得
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getPointCnt(int color)
	{
		int ret = 0;
		if(color == ReversiConst.REVERSI_STS_BLACK)	ret = this.mMasuPointCntB;
		else										ret = this.mMasuPointCntW;
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			コマ数取得
	///	@fn				public int getBetCnt(int color)
	///	@param[in]		int color		コマ色
	///	@return			コマ数取得
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getBetCnt(int color)
	{
		int ret = 0;
		if(color == ReversiConst.REVERSI_STS_BLACK)	ret = this.mMasuBetCntB;
		else										ret = this.mMasuBetCntW;
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			パス判定
	///	@fn				public int getPassEna(int color,int y,int x)
	///	@param[in]		int color		コマ色
	///	@param[in]		int y			マスのY座標
	///	@param[in]		int x			マスのX座標
	///	@return			パス判定
	///					- 0 : NOT PASS
	///					- 1 : PASS
	///
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getPassEna(int color,int y,int x)
	{
		int ret = 0;
		if(this.checkPara(y,0,this.mMasuCnt) == 0 && this.checkPara(x,0,this.mMasuCnt) == 0){
			if(color == ReversiConst.REVERSI_STS_BLACK)	ret = this.mMasuStsPassB[y][x];
			else										ret = this.mMasuStsPassW[y][x];
		}
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			履歴取得
	///	@fn				public ReversiHistory getHistory(int num)
	///	@param[in]		int num			ポイント
	///	@return			履歴 null : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public ReversiHistory getHistory(int num)
	{
		ReversiHistory ret = null;
		if(this.checkPara(num,0,(this.mMasuCnt * this.mMasuCnt)) == 0){
			ret = this.mMasuHist[num];
		}
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			履歴数取得
	///	@fn				public int getHistoryCnt(int color)
	///	@return			履歴数
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getHistoryCnt()
	{
		int ret = 0;
		ret = this.mMasuHistCur;
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			ポイント座標解析取得
	///	@fn				public ReversiAnz getPointAnz(int color,int y,int x)
	///	@param[in]		int color		コマ色
	///	@param[in]		int y			マスのY座標
	///	@param[in]		int x			マスのX座標
	///	@return			ポイント座標解析 null : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public ReversiAnz getPointAnz(int color,int y,int x)
	{
		ReversiAnz ret = null;
		if(this.checkPara(y,0,this.mMasuCnt) == 0 && this.checkPara(x,0,this.mMasuCnt) == 0){
			if(color == ReversiConst.REVERSI_STS_BLACK)	ret = this.mMasuStsAnzB[y][x];
			else										ret = this.mMasuStsAnzW[y][x];
		}
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			角の隣に置いても角を取られないマス検索
	///	@fn				public int checkEdge(int y,int x)
	///	@param[in]		int color		色
	///	@param[in]		int y			Y座標
	///	@param[in]		int x			X座標
	///	@return			0 : 取られる それ以外 : 取られない
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int checkEdge(int color,int y,int x)
	{
		int style = 0,flg1 = 0,flg2 = 0;
		int loop,loop2;
		if(y == 0 && x == 1){
			for(loop = x,flg1 = 0,flg2 = 0;loop < this.mMasuCnt;loop++){
				if(this.getMasuSts(y,loop) == color) flg1 = 1;
				if(flg1 == 1 && this.getMasuSts(y,loop) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(y,loop) != color) && (this.getMasuSts(y,loop) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}
		if(y == 1 && x == 0){
			for(loop = y,flg1 = 0,flg2 = 0;loop < this.mMasuCnt;loop++){
				if(this.getMasuSts(loop,x) == color) flg1 = 1;
				if(flg1 == 1 && this.getMasuSts(loop,x) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(loop,x) != color) && (this.getMasuSts(loop,x) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}
		if(y == 1 && x == 1){
			for(loop = y,flg1 = 0,flg2 = 0;loop < this.mMasuCnt;loop++){
				if(this.getMasuSts(loop,loop) == color) flg1 = 1;
				if(flg1 == 1 && this.getMasuSts(loop,loop) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(loop,loop) != color) && (this.getMasuSts(loop,loop) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}
		if(y == 0 && x == (this.mMasuCnt-2)){
			for(loop = x,flg1 = 0,flg2 = 0;loop > 0;loop--){
				if(this.getMasuSts(y,loop) == color) flg1 = 1;
				if(flg1 == 1 && this.getMasuSts(y,loop) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(y,loop) != color) && (this.getMasuSts(y,loop) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}
		if(y == 1 && x == (this.mMasuCnt-1)){
			for(loop = y,flg1 = 0,flg2 = 0;loop < this.mMasuCnt;loop++){
				if(this.getMasuSts(loop,x) == color) flg1 = 1;
				if(flg1 == 1 && this.getMasuSts(loop,x) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(loop,x) != color) && (this.getMasuSts(loop,x) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}
		if(y == 1 && x == (this.mMasuCnt-2)){
			for(loop = y,loop2 = x,flg1 = 0,flg2 = 0;loop < this.mMasuCnt;loop++,loop2--){
				if(this.getMasuSts(loop,loop2) == color) flg1 = 1;
				if(flg1 == 1 && this.getMasuSts(loop,loop2) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(loop,loop2) != color) && (this.getMasuSts(loop,loop2) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}
		if(y == (this.mMasuCnt-2) && x == 0){
			for(loop = y,flg1 = 0,flg2 = 0;loop > 0;loop--){
				if(this.getMasuSts(loop,x) == color) flg1=1;
				if(flg1 == 1 && this.getMasuSts(loop,x) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(loop,x) != color) && (this.getMasuSts(loop,x) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}
		if(y == (this.mMasuCnt-1) && x == 1){
			for(loop = x,flg1 = 0,flg2 = 0;loop < this.mMasuCnt;loop++){
				if(this.getMasuSts(y,loop) == color) flg1 = 1;
				if(flg1 == 1 && this.getMasuSts(y,loop) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(y,loop) != color) && (this.getMasuSts(y,loop) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}
		if(y == (this.mMasuCnt-2) && x == 1){
			for(loop = y,loop2 = x,flg1 = 0,flg2 = 0;loop > 0;loop--,loop2++){
				if(this.getMasuSts(loop,loop2) == color) flg1 = 1;
				if(flg1 == 1 && this.getMasuSts(loop,loop2) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(loop,loop2) != color) && (this.getMasuSts(loop,loop2) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}
		if(y == (this.mMasuCnt-2) && x == (this.mMasuCnt-1)){
			for(loop = y,flg1 = 0,flg2 = 0;loop > 0;loop--){
				if(this.getMasuSts(loop,x) == color) flg1 = 1;
				if(flg1 == 1 && this.getMasuSts(loop,x) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(loop,x) != color) && (this.getMasuSts(loop,x) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}
		if(y == (this.mMasuCnt-1) && x == (this.mMasuCnt-2)){
			for(loop = x,flg1 = 0,flg2 = 0;loop > 0;loop--){
				if(this.getMasuSts(y,loop) == color) flg1 = 1;
				if(flg1 == 1 && this.getMasuSts(y,loop) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(y,loop) != color) && (this.getMasuSts(y,loop) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}
		if(y == (this.mMasuCnt-2) && x == (this.mMasuCnt-2)){
			for(loop = y,loop2 = x,flg1 = 0,flg2 = 0;loop > 0;loop--,loop2--){
				if(this.getMasuSts(loop,loop2) == color) flg1 = 1;
				if(flg1 == 1 && this.getMasuSts(loop,loop2) == ReversiConst.REVERSI_STS_NONE) break;
				if((flg1 == 1) && (this.getMasuSts(loop,loop2) != color) && (this.getMasuSts(loop,loop2) != ReversiConst.REVERSI_STS_NONE)) flg2 = 1;
			}
			if((flg1 == 1) && (flg2 == 0)) style = 1;
		}

		return style;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			指定座標が角か取得
	///	@fn				public int getEdgeSideZero(int y,int x)
	///	@param[in]		int y			Y座標
	///	@param[in]		int x			X座標
	///	@return			0 : 成功 それ以外 : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getEdgeSideZero(int y,int x)
	{
		int ret = -1;
		if(
				(y == 0 && x == 0)
			||	(y == 0 && x == (this.mMasuCnt-1))
			||	(y == (this.mMasuCnt-1) && x == 0)
			||	(y == (this.mMasuCnt-1) && x == (this.mMasuCnt-1))
		){
			ret = 0;
		}
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			指定座標が角の一つ手前か取得
	///	@fn				public int getEdgeSideOne(int y,int x)
	///	@param[in]		int y			Y座標
	///	@param[in]		int x			X座標
	///	@return			0 : 成功 それ以外 : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getEdgeSideOne(int y,int x)
	{
		int ret = -1;
		if(
				(y == 0 && x == 1)
			||	(y == 0 && x == (this.mMasuCnt-2))
			||	(y == 1 && x == 0)
			||	(y == 1 && x == 1)
			||	(y == 1 && x == (this.mMasuCnt-2))
			||	(y == 1 && x == (this.mMasuCnt-1))
			||	(y == (this.mMasuCnt-2) && x == 0)
			||	(y == (this.mMasuCnt-2) && x == 1)
			||	(y == (this.mMasuCnt-2) && x == (this.mMasuCnt-2))
			||	(y == (this.mMasuCnt-2) && x == (this.mMasuCnt-1))
			||	(y == (this.mMasuCnt-1) && x == 1)
			||	(y == (this.mMasuCnt-1) && x == (this.mMasuCnt-2))
		){
			ret = 0;
		}
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			指定座標が角の二つ手前か取得
	///	@fn				public int getEdgeSideTwo(int y,int x)
	///	@param[in]		int y			Y座標
	///	@param[in]		int x			X座標
	///	@return			0 : 成功 それ以外 : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getEdgeSideTwo(int y,int x)
	{
		int ret = -1;
		if(
				(y == 0 && x == 2)
			||	(y == 0 && x == (this.mMasuCnt-3))
			||	(y == 2 && x == 0)
			||	(y == 2 && x == 2)
			||	(y == 2 && x == (this.mMasuCnt-3))
			||	(y == 2 && x == (this.mMasuCnt-1))
			||	(y == (this.mMasuCnt-3) && x == 0)
			||	(y == (this.mMasuCnt-3) && x == 2)
			||	(y == (this.mMasuCnt-3) && x == (this.mMasuCnt-3))
			||	(y == (this.mMasuCnt-3) && x == (this.mMasuCnt-1))
			||	(y == (this.mMasuCnt-1) && x == 2)
			||	(y == (this.mMasuCnt-1) && x == (this.mMasuCnt-3))
		){
			ret = 0;
		}
		return ret;
	}

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			指定座標が角の三つ以上手前か取得
	///	@fn				public int getEdgeSideThree(int y,int x)
	///	@param[in]		int y			Y座標
	///	@param[in]		int x			X座標
	///	@return			0 : 成功 それ以外 : 失敗
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	public int getEdgeSideThree(int y,int x){
		int ret = -1;
		if(
				(y == 0 && (3 <= x && x <= (this.mMasuCnt-4)))
			||	((3 <= y && y <= (this.mMasuCnt-4)) && x == 0)
			||	(y == (this.mMasuCnt-1) && (3 <= x && x <= (this.mMasuCnt-4)))
			||	((3 <= y && y <= (this.mMasuCnt-4)) && x == (this.mMasuCnt-1))
		){
			ret = 0;
		}
		return ret;
	}

