package process
import (
	"fmt"
	"client/message"
	"encoding/json"
)

func outputGroupMes(mes *message.Message) {

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes) 
		if err != nil {
			fmt.Println("json.Unmarshal err=", err.Error())
			return
		}

	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s", smsMes.UserId, smsMes.Content)
	fmt.Println(info+"\n")

}