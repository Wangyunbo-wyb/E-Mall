package xerr

// success
const OK uint32 = 200

/*
	wrong code
*/

const ServerCommonError uint32 = 100001

const RequestParamError uint32 = 100002

const TokenExpireError uint32 = 100003

const TokenGenerateError uint32 = 100004

const DbError uint32 = 100005

const DbUpdateAffectedZeroError uint32 = 100006

const DataNoExistError uint32 = 100007

// user service
const UserAddressNotExist uint32 = 200001

//order service

// product service
const StockNotEnough uint32 = 300001
const ProductNotFound uint32 = 300002
const DtmError = 300003

// video service
const VideoFailToCreateComment uint32 = 400001
const VideoDBErr uint32 = 400002
const VideoFailToFavorite uint32 = 400003
const VideoCommentNotExist uint32 = 400004
const VideoCacheErr uint32 = 400005
const VideoMQErr uint32 = 400006
const VideoUnfavoriteErr uint32 = 400007

// message service
const MessageDataBaseError = 500001
const MessageTransactionError = 500002

// payment service
const PaymentNotExist uint32 = 600001
const PaymentFailToRefund uint32 = 600002
const PaymentStatusNotSupport uint32 = 600003
