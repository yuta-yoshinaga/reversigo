////////////////////////////////////////////////////////////////////////////////
///	@file			FuncsJson.go
///	@brief			ファンクションJSONクラス実装ファイル
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
///	@class		FuncsJson
///	@brief		ファンクションJSONクラス
///
////////////////////////////////////////////////////////////////////////////////
type FuncsJson struct {
	Function string //!< ファンクション
	Param1   string //!< パラメーター1
	Param2   string //!< パラメーター2
	Param3   string //!< パラメーター3
	Param4   string //!< パラメーター4
	Param5   string //!< パラメーター5
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			コンストラクタ
///	@fn				NewFuncsJson() *FuncsJson
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func NewFuncsJson() *FuncsJson {
	f := new(FuncsJson)
	f.Function = ""
	f.Param1 = ""
	f.Param2 = ""
	f.Param3 = ""
	f.Param4 = ""
	f.Param5 = ""
	return f
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetFunction() string
///	@return			function string
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f FuncsJson) GetFunction() string {
	return f.Function
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetFunction(function string)
///	@param[in]		function string
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f *FuncsJson) SetFunction(function string) {
	f.Function = function
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetParam1() string
///	@return			param1 string
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f FuncsJson) GetParam1() string {
	return f.Param1
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetParam1(param1 string)
///	@param[in]		param1 string
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f *FuncsJson) SetParam1(param1 string) {
	f.Param1 = param1
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetParam2() string
///	@return			param1 string
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f FuncsJson) GetParam2() string {
	return f.Param2
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetParam2(param2 string)
///	@param[in]		param2 string
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f *FuncsJson) SetParam2(param2 string) {
	f.Param2 = param2
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetParam3() string
///	@return			param3 string
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f FuncsJson) GetParam3() string {
	return f.Param3
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetParam3(param3 string)
///	@param[in]		param3 string
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f *FuncsJson) SetParam3(param3 string) {
	f.Param3 = param3
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetParam4() string
///	@return			param4 string
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f FuncsJson) GetParam4() string {
	return f.Param4
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetParam4(param4 string)
///	@param[in]		param4 string
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f *FuncsJson) SetParam4(param4 string) {
	f.Param4 = param4
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetParam5() string
///	@return			param5 string
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f FuncsJson) GetParam5() string {
	return f.Param5
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetParam5(param5 string)
///	@param[in]		param5 string
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (f *FuncsJson) SetParam5(param5 string) {
	f.Param5 = param5
}
