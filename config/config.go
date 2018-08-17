package config

type Config struct {
	// 服务器
	Host	string
	// 端口号
	Port	int

	// DB 驱动
	DBDriver	string
	// DB 连接字符串
	DBConectionString	string
}