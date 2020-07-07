////////////////////////////////////////////////////////////////////////////////
///	@file			ReversiSetting.go
///	@brief			アプリ設定クラス
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

////////////////////////////////////////////////////////////////////////////////
///	@class		ReversiSetting
///	@brief		アプリ設定クラス
///
////////////////////////////////////////////////////////////////////////////////
type ReversiSetting struct {
	mMode             int    //!< 現在のモード
	mType             int    //!< 現在のタイプ
	mPlayer           int    //!< プレイヤーの色
	mAssist           int    //!< アシスト
	mGameSpd          int    //!< ゲームスピード
	mEndAnim          int    //!< ゲーム終了アニメーション
	mMasuCntMenu      int    //!< マスの数
	mMasuCnt          int    //!< マスの数
	mPlayCpuInterVal  int    //!< CPU対戦時のインターバル(msec)
	mPlayDrawInterVal int    //!< 描画のインターバル(msec)
	mEndDrawInterVal  int    //!< 終了アニメーション描画のインターバル(msec)
	mEndInterVal      int    //!< 終了アニメーションのインターバル(msec)
	mPlayerColor1     string //!< プレイヤー1の色
	mPlayerColor2     string //!< プレイヤー2の色
	mBackGroundColor  string //!< 背景の色
	mBorderColor      string //!< 枠線の色
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMode() int
///	@return			mMode int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmMode() int {
	return r.mMode
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMode(mMode int)
///	@param[in]		mMode int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmMode(mMode int) {
	r.mMode = mMode
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmType() int
///	@return			mType
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmType() int {
	return r.mType
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmType(mType int)
///	@param[in]		mType int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmType(mType int) {
	r.mType = mType
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmPlayer() int
///	@return			mPlayer int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmPlayer() int {
	return r.mPlayer
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmPlayer(mPlayer int)
///	@param[in]		mPlayer int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmPlayer(mPlayer int) {
	r.mPlayer = mPlayer
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmAssist() int
///	@return			mAssist int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmAssist() int {
	return r.mAssist
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmAssist(mAssist int)
///	@param[in]		mAssist int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmAssist(mAssist int) {
	r.mAssist = mAssist
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmGameSpd() int
///	@return			mGameSpd int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmGameSpd() int {
	return r.mGameSpd
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmGameSpd(mGameSpd int)
///	@param[in]		mGameSpd
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmGameSpd(mGameSpd int) {
	r.mGameSpd = mGameSpd
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmEndAnim() int
///	@return			mEndAnim int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmEndAnim() int {
	return r.mEndAnim
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmEndAnim(mEndAnim int)
///	@param[in]		mEndAnim int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmEndAnim(mEndAnim int) {
	r.mEndAnim = mEndAnim
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuCntMenu() int
///	@return			mMasuCntMenu int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmMasuCntMenu() int {
	return r.mMasuCntMenu
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmMasuCntMenu(mMasuCntMenu int)
///	@param[in]		mMasuCntMenu int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmMasuCntMenu(mMasuCntMenu int) {
	r.mMasuCntMenu = mMasuCntMenu
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmMasuCnt() int
///	@return			mMasuCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmMasuCnt() int {
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
func (r *ReversiSetting) SetmMasuCnt(mMasuCnt int) {
	r.mMasuCnt = mMasuCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmPlayCpuInterVal() int
///	@return			mPlayCpuInterVal int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmPlayCpuInterVal() int {
	return r.mPlayCpuInterVal
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmPlayCpuInterVal(mPlayCpuInterVal int)
///	@param[in]		mPlayCpuInterVal int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmPlayCpuInterVal(mPlayCpuInterVal int) {
	r.mPlayCpuInterVal = mPlayCpuInterVal
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmPlayDrawInterVal() int
///	@return			mPlayDrawInterVal int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmPlayDrawInterVal() int {
	return r.mPlayDrawInterVal
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmPlayDrawInterVal(mPlayDrawInterVal int)
///	@param[in]		mPlayDrawInterVal int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmPlayDrawInterVal(mPlayDrawInterVal int) {
	r.mPlayDrawInterVal = mPlayDrawInterVal
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmEndDrawInterVal() int
///	@return			mEndDrawInterVal int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmEndDrawInterVal() int {
	return r.mEndDrawInterVal
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmEndDrawInterVal(mEndDrawInterVal int)
///	@param[in]		mEndDrawInterVal int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmEndDrawInterVal(mEndDrawInterVal int) {
	r.mEndDrawInterVal = mEndDrawInterVal
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmEndInterVal() int
///	@return			mEndInterVal int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmEndInterVal() int {
	return r.mEndInterVal
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmEndInterVal(mEndInterVal int)
///	@param[in]		mEndInterVal int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmEndInterVal(mEndInterVal int) {
	r.mEndInterVal = mEndInterVal
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmPlayerColor1() string
///	@return			mPlayerColor1 string
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmPlayerColor1() string {
	return r.mPlayerColor1
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmPlayerColor1(mPlayerColor1 string)
///	@param[in]		mPlayerColor1 string
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmPlayerColor1(mPlayerColor1 string) {
	r.mPlayerColor1 = mPlayerColor1
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmPlayerColor2() string
///	@return			mPlayerColor2 string
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmPlayerColor2() string {
	return r.mPlayerColor2
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmPlayerColor2(mPlayerColor2 string)
///	@param[in]		mPlayerColor2 string
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmPlayerColor2(mPlayerColor2 string) string {
	r.mPlayerColor2 = mPlayerColor2
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmBackGroundColor() string
///	@return			mBackGroundColor string
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmBackGroundColor() string {
	return r.mBackGroundColor
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmBackGroundColor(mBackGroundColor string)
///	@param[in]		mBackGroundColor string
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmBackGroundColor(mBackGroundColor string) {
	r.mBackGroundColor = mBackGroundColor
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetmBorderColor() string
///	@return			mBorderColor string
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmBorderColor() string {
	return r.mBorderColor
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetmBorderColor(mBorderColor string)
///	@param[in]		mBorderColor string
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) SetmBorderColor(mBorderColor string) {
	r.mBorderColor = mBorderColor
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			コンストラクタ
///	@fn				NewReversiSetting()
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func NewReversiSetting() *ReversiSetting {
	r := new(ReversiSetting)
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
func (r *ReversiSetting) Reset() {
	r.mMode = DEF_MODE_ONE                     // 現在のモード
	r.mType = DEF_TYPE_HARD                    // 現在のタイプ
	r.mPlayer = REVERSI_STS_BLACK              // プレイヤーの色
	r.mAssist = DEF_ASSIST_ON                  // アシスト
	r.mGameSpd = DEF_GAME_SPD_MID              // ゲームスピード
	r.mEndAnim = DEF_END_ANIM_ON               // ゲーム終了アニメーション
	r.mMasuCntMenu = DEF_MASU_CNT_8            // マスの数
	r.mMasuCnt = DEF_MASU_CNT_8_VAL            // マスの数
	r.mPlayCpuInterVal = DEF_GAME_SPD_MID_VAL2 // CPU対戦時のインターバル(msec)
	r.mPlayDrawInterVal = DEF_GAME_SPD_MID_VAL // 描画のインターバル(msec)
	r.mEndDrawInterVal = 100                   // 終了アニメーション描画のインターバル(msec)
	r.mEndInterVal = 500                       // 終了アニメーションのインターバル(msec)
	r.mPlayerColor1 = "#000000"                // プレイヤー1の色
	r.mPlayerColor2 = "#FFFFFF"                // プレイヤー2の色
	r.mBackGroundColor = "#00FF00"             // 背景の色
	r.mBorderColor = "#000000"                 // 枠線の色
}
