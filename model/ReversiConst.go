////////////////////////////////////////////////////////////////////////////////
///	@file			ReversiConst.go
///	@brief			アプリ定数クラス
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

var DEF_MODE_ONE int = 0 //!< 一人対戦
var DEF_MODE_TWO int = 1 //!< 二人対戦

var DEF_TYPE_EASY int = 0 //!< CPU 弱い
var DEF_TYPE_NOR int = 1  //!< CPU 普通
var DEF_TYPE_HARD int = 2 //!< CPU 強い

var DEF_COLOR_BLACK int = 0 //!< コマ色 黒
var DEF_COLOR_WHITE int = 1 //!< コマ色 白

var DEF_ASSIST_OFF int = 0 //!< アシスト OFF
var DEF_ASSIST_ON int = 1  //!< アシスト ON

var DEF_GAME_SPD_FAST int = 0 //!< ゲームスピード 早い
var DEF_GAME_SPD_MID int = 1  //!< ゲームスピード 普通
var DEF_GAME_SPD_SLOW int = 2 //!< ゲームスピード 遅い

var DEF_GAME_SPD_FAST_VAL int = 0   //!< ゲームスピード 早い
var DEF_GAME_SPD_MID_VAL int = 50   //!< ゲームスピード 普通
var DEF_GAME_SPD_SLOW_VAL int = 100 //!< ゲームスピード 遅い

var DEF_GAME_SPD_FAST_VAL2 int = 0    //!< ゲームスピード 早い
var DEF_GAME_SPD_MID_VAL2 int = 500   //!< ゲームスピード 普通
var DEF_GAME_SPD_SLOW_VAL2 int = 1000 //!< ゲームスピード 遅い

var DEF_END_ANIM_OFF int = 0 //!< 終了アニメーション OFF
var DEF_END_ANIM_ON int = 1  //!< 終了アニメーション ON

var DEF_MASU_CNT_6 int = 0  //!< マス縦横6
var DEF_MASU_CNT_8 int = 1  //!< マス縦横8
var DEF_MASU_CNT_10 int = 2 //!< マス縦横10
var DEF_MASU_CNT_12 int = 3 //!< マス縦横12
var DEF_MASU_CNT_14 int = 4 //!< マス縦横14
var DEF_MASU_CNT_16 int = 5 //!< マス縦横16
var DEF_MASU_CNT_18 int = 6 //!< マス縦横18
var DEF_MASU_CNT_20 int = 7 //!< マス縦横20

var DEF_MASU_CNT_6_VAL int = 6                     //!< マス縦横6
var DEF_MASU_CNT_8_VAL int = 8                     //!< マス縦横8
var DEF_MASU_CNT_10_VAL int = 10                   //!< マス縦横10
var DEF_MASU_CNT_12_VAL int = 12                   //!< マス縦横12
var DEF_MASU_CNT_14_VAL int = 14                   //!< マス縦横14
var DEF_MASU_CNT_16_VAL int = 16                   //!< マス縦横16
var DEF_MASU_CNT_18_VAL int = 18                   //!< マス縦横18
var DEF_MASU_CNT_20_VAL int = 20                   //!< マス縦横20
var DEF_MASU_CNT_MAX_VAL int = DEF_MASU_CNT_20_VAL //!< マス縦横最大

var REVERSI_STS_NONE int = 0  //!< コマ無し
var REVERSI_STS_BLACK int = 1 //!< 黒
var REVERSI_STS_WHITE int = 2 //!< 白
var REVERSI_STS_MIN int = 0   //!< ステータス最小値
var REVERSI_STS_MAX int = 2   //!< ステータス最大値
var REVERSI_MASU_CNT int = 8  //!< 縦横マス数

var LC_MSG_DRAW int = 0            //!< マス描画
var LC_MSG_ERASE int = 1           //!< マス消去
var LC_MSG_DRAW_INFO int = 2       //!< マス情報描画
var LC_MSG_ERASE_INFO int = 3      //!< マス情報消去
var LC_MSG_DRAW_ALL int = 4        //!< 全マス描画
var LC_MSG_ERASE_ALL int = 5       //!< 全マス消去
var LC_MSG_DRAW_INFO_ALL int = 6   //!< 全マス情報描画
var LC_MSG_ERASE_INFO_ALL int = 7  //!< 全マス情報消去
var LC_MSG_DRAW_END int = 8        //!< 描画終わり
var LC_MSG_CUR_COL int = 9         //!< 現在の色
var LC_MSG_CUR_COL_ERASE int = 10  //!< 現在の色消去
var LC_MSG_CUR_STS int = 11        //!< 現在のステータス
var LC_MSG_CUR_STS_ERASE int = 12  //!< 現在のステータス消去
var LC_MSG_MSG_DLG int = 13        //!< メッセージダイアログ
var LC_MSG_DRAW_ALL_RESET int = 14 //!< 全マスビットマップインスタンスクリア
