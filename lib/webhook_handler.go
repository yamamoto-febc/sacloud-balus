package lib

import (
	balus "github.com/yamamoto-febc/sacloud-delete-all/lib"
	sakura "github.com/yamamoto-febc/sakura-iot-go"
)

const (
	SAKURA_IOT_CHANNEL = 0 // 利用チャンネル

	SAKURA_IOT_START_CODE       = 0 // 処理開始
	SAKURA_IOT_EXIT_NORMAL_CODE = 1 // 正常終了
	SAKURA_IOT_EXIT_ERROR_CODE  = 2 // 異常終了
)

type WebhookHandler struct {
	option *Option
	out    func(string, ...interface{})
}

func NewWebhookHandler(option *Option, logFunc func(string, ...interface{})) *WebhookHandler {
	w := &WebhookHandler{
		option: option,
		out:    logFunc,
	}
	if w.out == nil {
		w.out = func(f string, v ...interface{}) {}
	}
	return w
}

func (w *WebhookHandler) HandleRequest(p sakura.Payload) {

	w.debuglog("[DEBUG] Start handle request\n")

	if w.hasStartRequest(p) { // 処理開始リクエストを受けたら
		w.debuglog("[DEBUG] Received a start job request \n")

		// sacloud-delete-allを実行
		o := w.createBalusOption()
		err := balus.Run(o)

		// 処理結果をメッセージ表示
		res := (err == nil) //エラーがなければ正常終了
		if res {
			// 正常終了
			w.infolog("[INFO] Magical spell `balus` is accepted and done. All resources are gone...")
		} else {
			// 異常終了
			w.infolog("[ERROR] Failed on deleting all resources : %s ", err)
		}

		// 処理結果をさくらのIoT Platformへ送信
		sender := sakura.NewWebhookSender(w.option.SendToken, w.option.SendSecret)

		// ペイロード準備
		p := sakura.NewPayload(w.option.TargetModuleID)
		v := SAKURA_IOT_EXIT_NORMAL_CODE
		if !res {
			v = SAKURA_IOT_EXIT_ERROR_CODE
		}
		p.AddValueByInt(SAKURA_IOT_CHANNEL, int32(v))

		// 送信
		err = sender.Send(p)
		if err == nil {
			// 正常終了
			w.infolog("[INFO] Message was sended to Sakura-IoT-Platform")
		} else {
			// 送信エラー
			w.infolog("[ERROR] Failed on sending message to Sakura-IoT-Platform :%s ", err)
		}
	}

}

func (w *WebhookHandler) createBalusOption() *balus.Option {
	o := balus.NewOption()
	o.AccessToken = w.option.AccessToken
	o.AccessTokenSecret = w.option.AccessTokenSecret
	o.Zones = w.option.Zones
	o.TraceMode = w.option.TraceMode
	o.ForceMode = true

	o.JobQueueOption.InfoLog = true
	o.JobQueueOption.WarnLog = true
	o.JobQueueOption.ErrorLog = true

	return o
}

func (w *WebhookHandler) infolog(f string, v ...interface{}) {
	if len(v) == 0 {
		w.out(f)
	} else {
		w.out(f, v)
	}
}
func (w *WebhookHandler) debuglog(f string, v ...interface{}) {
	if w.option.Debug {
		if len(v) == 0 {
			w.out(f)
		} else {
			w.out(f, v)
		}
	}

}

func (w *WebhookHandler) hasStartRequest(p sakura.Payload) bool {
	for _, c := range p.Payload.Channels {
		if c.Channel == SAKURA_IOT_CHANNEL {
			value, err := c.GetInt()
			if err != nil {
				return false
			}
			if value == int32(SAKURA_IOT_START_CODE) {
				return true
			}
		}
	}
	return false

}
