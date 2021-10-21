package utils
import (
	"fmt"
	"net"
	"go_code/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
)
//这里将这些方法关联到结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf [8096]byte //这时传输时，使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("读取客户端发送的数据...")
	//先读4个字节接收长度
	_, err = this.Conn.Read(this.Buf[:4])
		if err != nil {
			//err = errors.New("read pkg header error")
			return
		}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])  //根据buf[:4]，把[]byte转成一个 uint32类型
	n, err := this.Conn.Read(this.Buf[:pkgLen]) //根据 pkgLen 读取消息内容
	if n != int(pkgLen) || err != nil {
		//err = errors.New("read pkg body error")
		return 
	}

	// 解包要把[]byte切片传入到结构体指针里 &mes
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
		if err != nil {
			fmt.Println("json.Unmarsha err=", err)
			return
		}
	return 
}


func (this *Transfer) WritePkg(data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	// 发送长度
	n, err := this.Conn.Write(this.Buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Write(bytes) fail", err)
			return
		}

	//发送data本身
	n, err = this.Conn.Write(data)
		if n != int(pkgLen) || err != nil {
			fmt.Println("conn.Write(bytes) fail", err)
			return
		}
	return 
}