////////////////////////////////////////////////////////////////////////////////
///	@file			ReversiPlayInterfaceImpl.go
///	@brief			リバーシプレイインタフェース実装ファイル
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

import "strconv"

////////////////////////////////////////////////////////////////////////////////
///	@interface	ReversiPlayInterfaceImpl
///	@brief		リバーシプレイデリゲート
///
////////////////////////////////////////////////////////////////////////////////
type ReversiPlayInterfaceImpl struct {
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			コンストラクタ
///	@fn				NewReversiPlayInterfaceImpl() *ReversiPlayInterfaceImpl
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func NewReversiPlayInterfaceImpl() *ReversiPlayInterfaceImpl {
	r := new(ReversiPlayInterfaceImpl)
	return r
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			メッセージダイアログ
///	@fn				ViewMsgDlg(title string, msg string) *FuncsJson
///	@param[in]		title string	タイトル
///	@param[in]		msg string		メッセージ
///	@return			FuncsJson
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlayInterfaceImpl) ViewMsgDlg(title string, msg string) *FuncsJson {
	funcs := NewFuncsJson()
	funcs.SetFunction("ViewMsgDlg")
	funcs.SetParam1(title)
	funcs.SetParam2(msg)
	return funcs
}

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
func (r ReversiPlayInterfaceImpl) DrawSingle(y int, x int, sts int, bk int, text string) *FuncsJson {
	funcs := NewFuncsJson()
	funcs.SetFunction("DrawSingle")
	funcs.SetParam1(strconv.Itoa(y))
	funcs.SetParam2(strconv.Itoa(x))
	funcs.SetParam3(strconv.Itoa(sts))
	funcs.SetParam4(strconv.Itoa(bk))
	funcs.SetParam5(text)
	return funcs
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			現在の色メッセージ
///	@fn				CurColMsg(text string) *FuncsJson
///	@param[in]		text string	テキスト
///	@return			FuncsJson
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlayInterfaceImpl) CurColMsg(text string) *FuncsJson {
	funcs := NewFuncsJson()
	funcs.SetFunction("CurColMsg")
	funcs.SetParam1(text)
	return funcs
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			現在のステータスメッセージ
///	@fn				CurStsMsg(text string) *FuncsJson
///	@param[in]		text string
///	@return			FuncsJson
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlayInterfaceImpl) CurStsMsg(text string) *FuncsJson {
	funcs := NewFuncsJson()
	funcs.SetFunction("CurStsMsg")
	funcs.SetParam1(text)
	return funcs
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ウェイト
///	@fn				Wait(time int) *FuncsJson
///	@param[in]		time int	ウェイト時間(msec)
///	@return			FuncsJson
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiPlayInterfaceImpl) Wait(time int) *FuncsJson {
	funcs := NewFuncsJson()
	funcs.SetFunction("Wait")
	funcs.SetParam1(strconv.Itoa(time))
	return funcs
}
