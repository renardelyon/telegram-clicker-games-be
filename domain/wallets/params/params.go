package params

type Params struct {
	PublicKey string `form:"phantom_encryption_public_key" json:"encryptPubKey" binding:"required"`
	Nonce     string `form:"nonce" json:"nonce" binding:"required"`
	Data      string `form:"data" json:"data" binding:"required"`
	WalletError
}

type WalletError struct {
	ErrorCode    *string `form:"errorCode,omitempty" json:"errorCode,omitempty"`
	ErrorMessage *string `form:"errorMessage,omitempty" json:"errorMessage,omitempty"`
}
