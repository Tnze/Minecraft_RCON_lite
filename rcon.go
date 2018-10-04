package mcrcon

import (
	"fmt"
	"github.com/micvbang/pocketmine-rcon"
)

//RCON 是一个rcon连接管理对象
type RCON struct {
	conn *rcon.Connection
}

//Connect 连接至服务器
func (r *RCON) Connect(address, password string) (err error) {
	r.conn, err = rcon.NewConnection(address, password)
	if err != nil {
		err = fmt.Errorf("登录失败: %v", err)
	}
	return
}

//Do 执行一个命令
func (r *RCON) Do(command string) {
	r.conn.SendCommand(command)
}
