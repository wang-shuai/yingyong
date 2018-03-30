package tool

import (
	"strings"
	"regexp"
	"strconv"
	"math"
	//"fmt"
	"time"
)

//const IpPattern  = "((25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))\\.){3}(25[0-5]|2[0-4]\\d|((1\\d{2})|([1-9]?\\d)))"

// 公式：字符型IP为202.192.13.32
// 转换：202 * 256 * 256 * 256 + 192 * 256 * 256 + 13 * 256 + 32 = 3401583904

const IpPattern = `((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))`

func HandIP(ip string) string {
	if t, _ := regexp.MatchString(IpPattern, ip); t {
		nums := strings.Split(strings.TrimSpace(ip), ".")
		if len(nums) == 4 {
			var total float64 = 0
			for i, v := range nums {
				num, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return ``
				} else {
					total += num * math.Pow(256, float64(4-i-1))
				}
			}
			return strconv.FormatFloat(total, 'f', 0, 64)
		}
	} else {
		//fmt.Println("IP格式错误")
	}
	return ``
}

func HandTimeStr(tm string) string {
	if len(strings.TrimSpace(tm)) == 0{
		return ``
	}
	the_time, err := time.Parse(time.RFC3339, tm)
	if err == nil {
		return HandTime(the_time)
	} else {
		return `` //return strconv.FormatInt(time.Now().Unix(),10)
	}
}

func HandTime(tm time.Time) string {
	return strconv.FormatInt(tm.Unix(), 10)
}
