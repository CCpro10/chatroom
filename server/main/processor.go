package main
import (
	"fmt"
	"net"
	"server/message"
	"server/utils"
	"server/process"
	"io"
)

//先创建一个Processor 的结构体体
type Processor struct {
	Conn net.Conn
}

func (this *Processor) process2() (err error) {

	//循环读取客户端发送的信息
	for {
		//创建一个Transfer 实例完成读包任务
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出..")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}

		}
		//处理读取后的消息
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}


//编写一个ServerProcessMes 函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {

	//看看是否能接收到客户端发送的群发的消息
	fmt.Println("mes=", mes)

	switch mes.Type {
		case message.LoginMesType :
		   //处理登录登录
		   //创建一个UserProcess实例
			up := &process2.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		case message.RegisterMesType :
		   //处理注册
		   up := &process2.UserProcess{
				Conn : this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		case message.SmsMesType :
			//创建一个SmsProcess实例完成转发群聊消息.
			smsProcess := &process2.SmsProcess{}
			smsProcess.SendGroupMes(mes)
		default :
		   fmt.Println("消息类型不存在，无法处理...")
	}
	return 
}

