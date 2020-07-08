////////////////////////////////////////////////////////////////////////////////
///	@file			ReversiPlayInterface.go
///	@brief			リバーシインターフェースファイル
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
///	@interface	ReversiPlayInterface
///	@brief		リバーシプレイインターフェース
///
////////////////////////////////////////////////////////////////////////////////
type ReversiPlayInterface interface {
	////////////////////////////////////////////////////////////////////////////////
	///	@brief			メッセージダイアログ
	///	@fn				ViewMsgDlg(title string , msg string) *FuncsJson
	///	@param[in]		title string	タイトル
	///	@param[in]		msg string	 	メッセージ
	///	@return			FuncsJson
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	ViewMsgDlg(title string, msg string) *FuncsJson

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			1マス描画
	///	@fn				DrawSingle(y int, x int, sts int, bk int, text string) *FuncsJson
	///	@param[in]		y int		Y座標
	///	@param[in]		x int		X座標
	///	@param[in]		sts int		ステータス
	///	@param[in]		bk int		背景
	///	@param[in]		text string	テキスト
	///	@return			FuncsJson
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	DrawSingle(y int, x int, sts int, bk int, text string) *FuncsJson

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			現在の色メッセージ
	///	@fn				CurColMsg(text string) *FuncsJson
	///	@param[in]		text string	テキスト
	///	@return			FuncsJson
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	CurColMsg(text string) *FuncsJson

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			現在のステータスメッセージ
	///	@fn				CurStsMsg(text string) *FuncsJson
	///	@param[in]		text string	テキスト
	///	@return			FuncsJson
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	CurStsMsg(text string) *FuncsJson

	////////////////////////////////////////////////////////////////////////////////
	///	@brief			ウェイト
	///	@fn				Wait(time int) *FuncsJson
	///	@param[in]		time int	ウェイト時間(msec)
	///	@return			FuncsJson
	///	@author			Yuta Yoshinaga
	///	@date			2020.07.06
	///
	////////////////////////////////////////////////////////////////////////////////
	Wait(time int) *FuncsJson
}