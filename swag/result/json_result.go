package result

// 分页统一json返回
type PageJsonResult struct {
	Code    int         `json:"code" `   //Http状态码
	Message interface{} `json:"message"` //消息内容
	Count   int         `json:"count"`   //数据总数
	Data    interface{} `json:"data"`    //数据内容
	Skip    uint        `json:"skip"`    //跳过记录数
}

// 普通数据json返回
type NormalJsonResult struct {
	Code    int         `json:"code" `   //Http状态码
	Message interface{} `json:"message"` //消息内容
	Data    interface{} `json:"data"`    //数据内容
	TaskId  string      `json:"task_id"` //任务id
	Events  interface{} `json:"events"`
}

// 普通数据不含taskId json返回
type NormalJsonResultWithoutTask struct {
	Code    int         `json:"code" `   //Http状态码
	Message interface{} `json:"message"` //消息内容
	Data    interface{} `json:"data"`    //数据内容
}
