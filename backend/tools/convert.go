package tools

const base62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func UintToBase62(n uint64) string {
	if n == 0 {
		return string(base62[0]) 								// 处理特殊情况：n为0时直接返回'0'
	}

	var result []byte 											// 创建一个空的字节切片用于存储编码结果
	for n > 0 {
		result = append(result, base62[n%62])					// 计算余数并添加到结果切片中
		n /= 62													// 更新n为商，准备处理下一个更高位的数字
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {		// 反转结果切片
		result[i], result[j] = result[j], result[i]
	}

	return string(result)										// 将反转后的结果切片转换为字符串并返回
}