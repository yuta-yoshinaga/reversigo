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
	Mode             int    //!< 現在のモード
	Type             int    //!< 現在のタイプ
	Player           int    //!< プレイヤーの色
	Assist           int    //!< アシスト
	GameSpd          int    //!< ゲームスピード
	EndAnim          int    //!< ゲーム終了アニメーション
	MasuCntMenu      int    //!< マスの数
	MasuCnt          int    //!< マスの数
	PlayCpuInterVal  int    //!< CPU対戦時のインターバル(msec)
	PlayDrawInterVal int    //!< 描画のインターバル(msec)
	EndDrawInterVal  int    //!< 終了アニメーション描画のインターバル(msec)
	EndInterVal      int    //!< 終了アニメーションのインターバル(msec)
	PlayerColor1     string //!< プレイヤー1の色
	PlayerColor2     string //!< プレイヤー2の色
	BackGroundColor  string //!< 背景の色
	BorderColor      string //!< 枠線の色
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			コンストラクタ
///	@fn				NewReversiSetting() *ReversiSetting
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
///	@brief			ゲッター
///	@fn				GetmMode() int
///	@return			mMode int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiSetting) GetmMode() int {
	return r.Mode
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
	r.Mode = mMode
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
	return r.Type
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
	r.Type = mType
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
	return r.Player
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
	r.Player = mPlayer
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
	return r.Assist
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
	r.Assist = mAssist
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
	return r.GameSpd
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
	r.GameSpd = mGameSpd
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
	return r.EndAnim
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
	r.EndAnim = mEndAnim
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
	return r.MasuCntMenu
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
	r.MasuCntMenu = mMasuCntMenu
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
func (r *ReversiSetting) SetmMasuCnt(mMasuCnt int) {
	r.MasuCnt = mMasuCnt
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
	return r.PlayCpuInterVal
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
	r.PlayCpuInterVal = mPlayCpuInterVal
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
	return r.PlayDrawInterVal
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
	r.PlayDrawInterVal = mPlayDrawInterVal
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
	return r.EndDrawInterVal
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
	r.EndDrawInterVal = mEndDrawInterVal
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
	return r.EndInterVal
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
	r.EndInterVal = mEndInterVal
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
	return r.PlayerColor1
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
	r.PlayerColor1 = mPlayerColor1
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
	return r.PlayerColor2
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
func (r *ReversiSetting) SetmPlayerColor2(mPlayerColor2 string) {
	r.PlayerColor2 = mPlayerColor2
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
	return r.BackGroundColor
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
	r.BackGroundColor = mBackGroundColor
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
	return r.BorderColor
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
	r.BorderColor = mBorderColor
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
	r.Mode = DEF_MODE_ONE                     // 現在のモード
	r.Type = DEF_TYPE_HARD                    // 現在のタイプ
	r.Player = REVERSI_STS_BLACK              // プレイヤーの色
	r.Assist = DEF_ASSIST_ON                  // アシスト
	r.GameSpd = DEF_GAME_SPD_MID              // ゲームスピード
	r.EndAnim = DEF_END_ANIM_ON               // ゲーム終了アニメーション
	r.MasuCntMenu = DEF_MASU_CNT_8            // マスの数
	r.MasuCnt = DEF_MASU_CNT_8_VAL            // マスの数
	r.PlayCpuInterVal = DEF_GAME_SPD_MID_VAL2 // CPU対戦時のインターバル(msec)
	r.PlayDrawInterVal = DEF_GAME_SPD_MID_VAL // 描画のインターバル(msec)
	r.EndDrawInterVal = 100                   // 終了アニメーション描画のインターバル(msec)
	r.EndInterVal = 500                       // 終了アニメーションのインターバル(msec)
	r.PlayerColor1 = "#000000"                // プレイヤー1の色
	r.PlayerColor2 = "#FFFFFF"                // プレイヤー2の色
	r.BackGroundColor = "#00FF00"             // 背景の色
	r.BorderColor = "#000000"                 // 枠線の色
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			リセット
///	@fn				Reset()
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiSetting) Copy(s *ReversiSetting) {
	r.Mode = s.Mode                         // 現在のモード
	r.Type = s.Type                         // 現在のタイプ
	r.Player = s.Player                     // プレイヤーの色
	r.Assist = s.Assist                     // アシスト
	r.GameSpd = s.GameSpd                   // ゲームスピード
	r.EndAnim = s.EndAnim                   // ゲーム終了アニメーション
	r.MasuCntMenu = s.MasuCntMenu           // マスの数
	r.MasuCnt = s.MasuCnt                   // マスの数
	r.PlayCpuInterVal = s.PlayCpuInterVal   // CPU対戦時のインターバル(msec)
	r.PlayDrawInterVal = s.PlayDrawInterVal // 描画のインターバル(msec)
	r.EndDrawInterVal = s.EndDrawInterVal   // 終了アニメーション描画のインターバル(msec)
	r.EndInterVal = s.EndInterVal           // 終了アニメーションのインターバル(msec)
	r.PlayerColor1 = s.PlayerColor1         // プレイヤー1の色
	r.PlayerColor2 = s.PlayerColor2         // プレイヤー2の色
	r.BackGroundColor = s.BackGroundColor   // 背景の色
	r.BorderColor = s.BorderColor           // 枠線の色
}
