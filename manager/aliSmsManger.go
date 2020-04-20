package manager

import (
	"fmt"

	sms "github.com/aliyun-sdk/sms-go"
)

// SmsManager store info to send sms
type SmsManager struct {
	accessKey string
	secretKey string
	// 短信签名
	signName string
	// 短信发送模板
	template string
	// 模板变量
	templateVars []string

	*sms.Client
}

// NewSmsManager creates a SmsManager
func NewSmsManager(ak, sk, sn, tp string, tplVars ...string) (*SmsManager, error) {
	man := &SmsManager{
		accessKey: sk,
		secretKey: sk,
		signName:  sn,
		template:  tp,
	}
	client, err := sms.New(ak, sk, sms.SignName(sn), sms.Template(tp))
	if err != nil {
		return nil, err
	}
	man.Client = client

	man.templateVars = make([]string, len(tplVars))
	copy(man.templateVars, tplVars)

	return man, nil
}

// Send sends a sms
func (sm *SmsManager) Send(receiver string, codes ...string) error {
	if len(codes) != len(sm.templateVars) {
		return fmt.Errorf("code number not match template var")
	}
	data := make(map[string]interface{})

	for i, code := range codes {
		data[sm.templateVars[i]] = code
	}

	return sm.Client.Send(sms.Mobile(receiver), sms.Parameter(data))
}
