package timeTool

import (
	"strconv"
	"strings"
	"time"
)

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d) //最前面和最后面的ASCII定义的空格去掉
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	//如果出错开始判断
	if strings.Contains(d, "d") { //判断字符串s中是否包含个子串str。包含或者str为空则返回true
		index := strings.Index(d, "d") //获取指定子字符串的第一个实例的索引值

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour) //天数转换成了time.Duration类型
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}
