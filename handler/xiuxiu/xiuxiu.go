package xiuxiu

/*建议：
如果消息体有 JSON 参数需要传递，针对每一个 API 接口定义独立的 go struct 来接收，比如 GetRequest 和 GetResponse，并将这些结构体统一放在一个 Go 文件中，以方便后期维护和修改。
*/

type GetRequest struct {
	Image string `json:"image"`
	Top   string `json:"top"`
}

type GetResponse struct {
	Image string `json:"image"`
}
