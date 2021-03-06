////////////////////////////////////////////////////////////////////////////////
///	@file			ReversiAnz.go
///	@brief			リバーシ解析クラス実装ファイル
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
package model

////////////////////////////////////////////////////////////////////////////////
///	@class		ReversiAnz
///	@brief		リバーシ解析クラス
///
////////////////////////////////////////////////////////////////////////////////
type ReversiAnz struct {
	Min                 int     //!< 最小値
	Max                 int     //!< 最大値
	Avg                 float64 //!< 平均
	PointCnt            int     //!< 置けるポイント数
	EdgeCnt             int     //!< 角を取れるポイント数
	EdgeSideOneCnt      int     //!< 角一つ前を取れるポイント数
	EdgeSideTwoCnt      int     //!< 角二つ前を取れるポイント数
	EdgeSideThreeCnt    int     //!< 角三つ前を取れるポイント数
	EdgeSideOtherCnt    int     //!< それ以外を取れるポイント数
	OwnMin              int     //!< 最小値
	OwnMax              int     //!< 最大値
	OwnAvg              float64 //!< 平均
	OwnPointCnt         int     //!< 置けるポイント数
	OwnEdgeCnt          int     //!< 角を取れるポイント数
	OwnEdgeSideOneCnt   int     //!< 角一つ前を取れるポイント数
	OwnEdgeSideTwoCnt   int     //!< 角二つ前を取れるポイント数
	OwnEdgeSideThreeCnt int     //!< 角三つ前を取れるポイント数
	OwnEdgeSideOtherCnt int     //!< それ以外を取れるポイント数
	BadPoint            int     //!< 悪手ポイント
	GoodPoint           int     //!< 良手ポイント
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			コンストラクタ
///	@fn				NewReversiAnz() *ReversiAnz
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func NewReversiAnz() *ReversiAnz {
	r := new(ReversiAnz)
	r.Reset()
	return r
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetMin() int
///	@return			min int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetMin() int {
	return r.Min
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetMin(min int)
///	@param[in]		min int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetMin(min int) {
	r.Min = min
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetMax() int
///	@return			max int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetMax() int {
	return r.Max
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetMax(max int)
///	@param[in]		max int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetMax(max int) {
	r.Max = max
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetAvg() float64
///	@return			avg float64
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetAvg() float64 {
	return r.Avg
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetAvg(avg float64)
///	@param[in]		avg float64
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetAvg(avg float64) {
	r.Avg = avg
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetPointCnt() int
///	@return			pointCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetPointCnt() int {
	return r.PointCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetPointCnt(pointCnt int)
///	@param[in]		pointCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetPointCnt(pointCnt int) {
	r.PointCnt = pointCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetEdgeCnt() int
///	@return			edgeCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetEdgeCnt() int {
	return r.EdgeCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetEdgeCnt(edgeCnt int)
///	@param[in]		edgeCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetEdgeCnt(edgeCnt int) {
	r.EdgeCnt = edgeCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetEdgeSideOneCnt() int
///	@return			edgeSideOneCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetEdgeSideOneCnt() int {
	return r.EdgeSideOneCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetEdgeSideOneCnt(edgeSideOneCnt int)
///	@param[in]		edgeSideOneCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetEdgeSideOneCnt(edgeSideOneCnt int) {
	r.EdgeSideOneCnt = edgeSideOneCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetEdgeSideTwoCnt() int
///	@return			edgeSideTwoCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetEdgeSideTwoCnt() int {
	return r.EdgeSideTwoCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetEdgeSideTwoCnt(edgeSideTwoCnt int)
///	@param[in]		edgeSideTwoCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetEdgeSideTwoCnt(edgeSideTwoCnt int) {
	r.EdgeSideTwoCnt = edgeSideTwoCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetEdgeSideThreeCnt() int
///	@return			edgeSideThreeCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetEdgeSideThreeCnt() int {
	return r.EdgeSideThreeCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetEdgeSideThreeCnt(edgeSideThreeCnt int)
///	@param[in]		edgeSideThreeCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetEdgeSideThreeCnt(edgeSideThreeCnt int) {
	r.EdgeSideThreeCnt = edgeSideThreeCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetEdgeSideOtherCnt() int
///	@return			edgeSideOtherCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetEdgeSideOtherCnt() int {
	return r.EdgeSideOtherCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetEdgeSideOtherCnt(edgeSideOtherCnt int)
///	@param[in]		edgeSideOtherCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetEdgeSideOtherCnt(edgeSideOtherCnt int) {
	r.EdgeSideOtherCnt = edgeSideOtherCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetOwnMin() int
///	@return			ownMin int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetOwnMin() int {
	return r.OwnMin
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetOwnMin(ownMin int)
///	@param[in]		ownMin int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetOwnMin(ownMin int) {
	r.OwnMin = ownMin
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetOwnMax() int
///	@return			ownMax int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetOwnMax() int {
	return r.OwnMax
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetOwnMax(ownMax int)
///	@param[in]		ownMax int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetOwnMax(ownMax int) {
	r.OwnMax = ownMax
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetOwnAvg() float64
///	@return			ownAvg float64
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetOwnAvg() float64 {
	return r.OwnAvg
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetOwnAvg(ownAvg float64)
///	@param[in]		ownAvg float64
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetOwnAvg(ownAvg float64) {
	r.OwnAvg = ownAvg
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetOwnPointCnt() int
///	@return			ownPointCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetOwnPointCnt() int {
	return r.OwnPointCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetOwnPointCnt(ownPointCnt int)
///	@param[in]		ownPointCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetOwnPointCnt(ownPointCnt int) {
	r.OwnPointCnt = ownPointCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetOwnEdgeCnt() int
///	@return			ownEdgeCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetOwnEdgeCnt() int {
	return r.OwnEdgeCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetOwnEdgeCnt(ownEdgeCnt int)
///	@param[in]		ownEdgeCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetOwnEdgeCnt(ownEdgeCnt int) {
	r.OwnEdgeCnt = ownEdgeCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetOwnEdgeSideOneCnt() int
///	@return			ownEdgeSideOneCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetOwnEdgeSideOneCnt() int {
	return r.OwnEdgeSideOneCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetOwnEdgeSideOneCnt(ownEdgeSideOneCnt int)
///	@param[in]		ownEdgeSideOneCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetOwnEdgeSideOneCnt(ownEdgeSideOneCnt int) {
	r.OwnEdgeSideOneCnt = ownEdgeSideOneCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetOwnEdgeSideTwoCnt() int
///	@return			ownEdgeSideTwoCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetOwnEdgeSideTwoCnt() int {
	return r.OwnEdgeSideTwoCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetOwnEdgeSideTwoCnt(ownEdgeSideTwoCnt int)
///	@param[in]		ownEdgeSideTwoCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetOwnEdgeSideTwoCnt(ownEdgeSideTwoCnt int) {
	r.OwnEdgeSideTwoCnt = ownEdgeSideTwoCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetOwnEdgeSideThreeCnt() int
///	@return			ownEdgeSideThreeCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetOwnEdgeSideThreeCnt() int {
	return r.OwnEdgeSideThreeCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetOwnEdgeSideThreeCnt(ownEdgeSideThreeCnt int)
///	@param[in]		ownEdgeSideThreeCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetOwnEdgeSideThreeCnt(ownEdgeSideThreeCnt int) {
	r.OwnEdgeSideThreeCnt = ownEdgeSideThreeCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetOwnEdgeSideOtherCnt() int
///	@return			ownEdgeSideOtherCnt int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetOwnEdgeSideOtherCnt() int {
	return r.OwnEdgeSideOtherCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetOwnEdgeSideOtherCnt(ownEdgeSideOtherCnt int)
///	@param[in]		ownEdgeSideOtherCnt int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetOwnEdgeSideOtherCnt(ownEdgeSideOtherCnt int) {
	r.OwnEdgeSideOtherCnt = ownEdgeSideOtherCnt
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetBadPoint() int
///	@return			badPoint int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetBadPoint() int {
	return r.BadPoint
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetBadPoint(badPoint int)
///	@param[in]		badPoint int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetBadPoint(badPoint int) {
	r.BadPoint = badPoint
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			ゲッター
///	@fn				GetGoodPoint() int
///	@return			goodPoint int
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r ReversiAnz) GetGoodPoint() int {
	return r.GoodPoint
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			セッター
///	@fn				SetGoodPoint(goodPoint int)
///	@param[in]		goodPoint int
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) SetGoodPoint(goodPoint int) {
	r.GoodPoint = goodPoint
}

////////////////////////////////////////////////////////////////////////////////
///	@brief			リセット
///	@fn				Reset()
///	@return			ありません
///	@author			Yuta Yoshinaga
///	@date			2020.07.06
///
////////////////////////////////////////////////////////////////////////////////
func (r *ReversiAnz) Reset() {
	r.Min = 0
	r.Max = 0
	r.Avg = 0.0
	r.PointCnt = 0
	r.EdgeCnt = 0
	r.EdgeSideOneCnt = 0
	r.EdgeSideTwoCnt = 0
	r.EdgeSideThreeCnt = 0
	r.EdgeSideOtherCnt = 0
	r.OwnMin = 0
	r.OwnMax = 0
	r.OwnAvg = 0.0
	r.OwnPointCnt = 0
	r.OwnEdgeCnt = 0
	r.OwnEdgeSideOneCnt = 0
	r.OwnEdgeSideTwoCnt = 0
	r.OwnEdgeSideThreeCnt = 0
	r.OwnEdgeSideOtherCnt = 0
	r.BadPoint = 0
	r.GoodPoint = 0
}
