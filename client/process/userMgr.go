package process 
import (
	"fmt"
	"client/message"
	"client/model"

)

//客户端要维护的map，
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //我们在用户登录成功后，完成对CurUser初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
//遍历一把 onlineUsers
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers{
		//如果不显示自己.
		fmt.Println("用户id:\t", id)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//尝试把用户列表里的ID直接赋值给user
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有此用户，则创建一个新的数据结构
		user = &message.User{
			UserId : notifyUserStatusMes.UserId,
		}
	}
	//update the user's status
	user.UserStatus = notifyUserStatusMes.Status
	//update onlineUsers
	onlineUsers[notifyUserStatusMes.UserId] = user
	//打印出新在线列表
	outputOnlineUser()
}