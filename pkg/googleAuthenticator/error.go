package googleAuthenticator

//错误定义
import "errors"

var (
	ErrSecretLengthLss     = errors.New("secret length lss 6 error")
	ErrSecretLength        = errors.New("secret length error")
	ErrPaddingCharCount    = errors.New("padding char count error")
	ErrPaddingCharLocation = errors.New("padding char Location error")
	ErrParam               = errors.New("param error")
)
