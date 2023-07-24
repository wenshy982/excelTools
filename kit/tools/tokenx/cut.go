package tokenx

import (
	"strconv"

	"go.uber.org/zap"

	"kit/tools/jsonx"
)

// CutByToken2 删除x条以上的数据，并持续压缩消息到合适token大小的数据组,同时返回token数据
// 数据缺失， token计算不准确
func CutByToken2(msgData string, max int) (string, int) {
	// ChatCompletionMessage 消息本身的结构体
	type ChatCompletionMessage struct {
		Role    string `json:"role"`
		Content string `json:"content"`
		Name    string `json:"name,omitempty"`
	}
	var inMsg []ChatCompletionMessage
	jsonx.UnmarshalString(msgData, &inMsg)

	var result []ChatCompletionMessage
	n := len(inMsg)
	zap.L().Info("【聊天websocket接口】【输入】" + strconv.Itoa(n) + "组数据")
	var count = 0 // 有效token累加长度
	var s = jsonx.MarshalString(inMsg[0])
	zap.L().Info("【聊天websocket接口】【单项】" + strconv.Itoa(0) + "组数据长度:" + strconv.Itoa(CalToken(s)))
	count = CalToken(jsonx.MarshalString(inMsg[0])) // 先把系统指令token长度进行赋值
	if n > 2 {
		for k := n - 1; k > 0; k-- {
			rowCount := CalToken(jsonx.MarshalString(inMsg[k]))
			// zap.L().Info("【聊天websocket接口】【单项】" + strconv.Itoa(k) + "组数据长度:" + strconv.Itoa(rowCount))
			if rowCount+count >= max {
				break
			} else {
				count += rowCount + 1
				result = append(result, inMsg[k])
			}
		}
		result = append(result, inMsg[0])
		// zap.L().Info("【聊天websocket接口】【删减后】" + strconv.Itoa(x) + "组数据," + "总长度:" + strconv.Itoa(count))
		x := len(result)
		left, right := 0, x-1
		for left < right {
			// 交换两个指针所指的元素
			result[left], result[right] = result[right], result[left]
			left++
			right--
		}
		// 重新封装
		msgByte := jsonx.MarshalString(result)
		zap.L().Info("【聊天websocket接口】【删减后】" + strconv.Itoa(x) + "组数据," + "总长度:" + strconv.Itoa(count-1))
		return msgByte, count + 1
	}
	return msgData, count + 1
}

// CutByToken 删除x条以上的数据，并持续压缩消息到合适token大小的数据组,同时返回token数据
func CutByToken(data string, max int) (string, int) {
	var messages []Message
	jsonx.UnmarshalString(data, &messages)
	l := len(messages)
	if l == 0 || l%2 != 0 { // 对话成对出现，2，4，6......
		return data, 0
	}
	var i = l - 1                                    // 最后一条消息的索引
	var str = messages[0].Str() + messages[i].Str2() // 第一条消息和最后一条消息+空数组
	var curToken = CalToken(str) + 2                 // 2是空数组的token
	for ; i >= 0; i -= 2 {
		if i == 1 { // 新消息只有2条，并且不超过最大token, 则退出
			return jsonx.MarshalString(messages), curToken
		}
		str = messages[i-2].Str() + messages[i-1].Str() // 最近的两条消息
		token := CalToken(str) - 1                      // 计算最近的两条消息token, -1 是去掉""的token
		if curToken+token > max {                       // 如果超过最大token, 则退出
			messages = append(messages[:1], messages[i:]...) // 删除第 1 到 i 条消息（1，i 为数组索引）
			break
		}
		curToken += token // 累加token
	}

	return jsonx.MarshalString(messages), curToken
}

func (m *Messages) Cut(max int) (string, int) {
	l := len(*m)
	if l == 0 || l%2 != 0 { // 对话成对出现，2，4，6......
		return jsonx.MarshalString(*m), 0
	}
	var i = l - 1                            // 最后一条消息的索引
	var str = (*m)[0].Str() + (*m)[i].Str2() // 第一条消息和最后一条消息+空数组
	var curToken = CalToken(str) + 2         // 2是空数组的token
	for ; i >= 0; i -= 2 {
		if i == 1 { // 新消息只有2条，并且不超过最大token, 则退出
			return jsonx.MarshalString(*m), curToken
		}
		str = (*m)[i-1].Str() + (*m)[i-2].Str() // 最近的两条消息
		token := CalToken(str) - 1              // 计算最近的两条消息token, -1 是去掉""的token
		if curToken+token > max {               // 如果超过最大token, 则退出
			*m = append((*m)[:1], (*m)[i:]...) // 删除第 1 到 i 条消息（1，i 为数组索引）
			break
		}
		curToken += token // 累加token
	}
	return jsonx.MarshalString(*m), curToken
}

func CheckToken(m Messages, max int) (string, bool) {
	if m == nil || len(m) == 0 {
		return "", false
	}
	result := jsonx.MarshalString(m)
	token := CalToken(jsonx.MarshalString(m))
	if token > max {
		result, token = m.Cut(max)
		if token > max || token == 0 {
			return result, false
		}
		return result, true
	}
	return result, true
}
