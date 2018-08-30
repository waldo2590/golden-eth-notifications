package ethNotification

const (
	MethodIDTransferERC20Token MethodID = iota
)

const (
	Pending TxStatus = iota
	Success
	Failed
)

type TxStatus int8
type MethodID int8

func (ts TxStatus) String() string {
	return [...]string{"Pending", "Success", "Failed"}[ts]
}

func (m MethodID) String() string {
	return [...]string{"0xa9059cbb"}[m]
}
