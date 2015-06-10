package main

import (
	"fmt"
	"github.com/brunohenrique/gcm"
)

func main() {
	ios := "nM599PoA1dI:APA91bFnObCGNKDHYedS6X0bXE6AUFk7O-hnNF1FlcBljIHHvD7NRUpVkKNsh4qUYeB8l9SgSD2peNmv4-OPgwdab1_mhDrg1jsaeUqAmmi7yXwiv4MW5NHBBmDsaofuLl-egEPb4llV"
	ios2 := "mf2-RQA9hzk:APA91bHbpGh8StEPEfm2TlCE2jTxLwVI5LhCDGLHrkprWP-d92iT9Ud3r2cgyesMeakihiXU2egmsO00EYd_byWkh2ebdnj44tk-qeQQsjHLwsVDSKornozvqD8wtBpN0_d9Wmc4rejt"
	android := "ebIT3rptXz8:APA91bGG1Z1LhOrvPo6-YrhiHBM1SKFsgVStWfSrSn8pGke2zVfQInS9NaJpOGL0EWoBnd-6vOIfzDDCC54HjWzlKzgE0cHYqFqxQN3gxie7sguUdPmQ7ZIlQzxsJKs9bkpZvvZZ3Dm2"

	msg := &gcm.Message{
		RegistrationIDs: []string{ios, ios2, android},
		Notification: map[string]interface{}{
			"body":  "Pretty Ygor from body",
			"badge": 1,
		},
	}

	// Create a Sender to send the message.
	sender := &gcm.Sender{ApiKey: "AIzaSyCA3NgMjPU9HASWygpW17luwjY3ZIf0hZw"}

	// Send the message and receive the response after at most two retries.
	response, err := sender.Send(msg, 0)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Printf("Response: %+v", response)
}
