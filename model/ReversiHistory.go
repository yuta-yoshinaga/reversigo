////////////////////////////////////////////////////////////////////////////////
///	@file			ReversiHistory.go
///	@brief			リバーシ履歴クラス実装ファイル
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
///	@class		ReversiHistory
///	@brief		リバーシ履歴クラス
///
////////////////////////////////////////////////////////////////////////////////
type ReversiHistory struct {
	point *ReversiPoint //!< リバーシポイント
	color int           //!< カラー
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetPoint() ReversiPoint
///	@return			point ReversiPoint
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiHistory) GetPoint() *ReversiPoint {
	return r.point
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetPoint(point *ReversiPoint)
///	@param[in]		point *ReversiPoint
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiHistory) SetPoint(point *ReversiPoint) {
	r.point = point
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetColor() int
///	@return			color int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiHistory) GetColor() int {
	return r.color
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetColor(color int)
///	@param[in]		color int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiHistory) SetColor(color int) {
	r.color = color
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			コンストラクタ
///	@fn				NewReversiHistory()
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func NewReversiHistory() *ReversiHistory {
	r := new(ReversiHistory)
	r.point = NewReversiPoint()
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
func (r *ReversiHistory) Reset() {
	r.point.SetX(-1)
	r.point.SetY(-1)
	r.color = -1
}