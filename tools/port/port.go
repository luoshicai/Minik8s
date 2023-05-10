package port

import (
	"fmt"
	"net"
)

// GetFreePort 分配一个当前空闲端口
func GetFreePort() int {
	port := ":0" // 随机端口
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return -1
	}
	defer l.Close()

	// 获取随机分配的端口号
	address := l.Addr().(*net.TCPAddr)
	newport := address.Port
	fmt.Println("Listening on port:", newport)
	return newport
}
