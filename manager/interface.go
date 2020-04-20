package manager

type VerifyCodeSender interface {
	Send(receiver string, codes ...string) error
}
