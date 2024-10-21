package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[RequestParamError] = "参数错误"
	message[TokenExpireError] = "token失效，请重新登陆"
	message[TokenGenerateError] = "生成token失败"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[UserAddressNotExist] = "用户地址不存在"
	message[DbUpdateAffectedZeroError] = "更新数据影响行数为0"
	message[DataNoExistError] = "数据不存在"
	message[PaymentNotExist] = "订单记录不存在"
	message[PaymentFailToRefund] = "支付成功后才可退款"
	message[PaymentStatusNotSupport] = "支付状态不支持"
	message[VideoFailToCreateComment] = "视频评论失败"
	message[VideoDBErr] = "video数据库繁忙,请稍后再试"
	message[VideoFailToFavorite] = "视频收藏失败"
	message[VideoCommentNotExist] = "视频评论不存在"
	message[VideoCacheErr] = "video缓存繁忙,请稍后再试"
	message[VideoMQErr] = "video消息队列繁忙,请稍后再试"
	message[VideoUnfavoriteErr] = "视频取消收藏失败"
	message[StockNotEnough] = "库存不足"
	message[DtmError] = "分布式事务错误"
	message[MessageDataBaseError] = "评论数据库错误"
	message[MessageTransactionError] = "评论事务错误"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
