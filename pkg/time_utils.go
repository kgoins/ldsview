package ldsview

import (
	"strconv"
	"strings"
	"time"
)

const adGeneralizedTimeFmt string = "20060102150405-0700"

func TimeFromADGeneralizedTime(adTime string) (time.Time, error) {
	adTime = strings.Split(adTime, ".")[0] + "-0000"

	convTime, err := time.Parse(adGeneralizedTimeFmt, adTime)
	if err != nil {
		return time.Time{}, err
	}

	return convTime, nil
}

func TimeFromADTimestamp(adTime string) time.Time {
	timeAsNum, _ := strconv.ParseInt(adTime, 10, 64)
	if timeAsNum == 0 {
		timeAsNum = 9223372036854775807
	}

	posixTime := int64(0)
	posixTime = timeAsNum - 11644473600*1000*10000
	posixTime = posixTime / 10000000

	timeResult := time.Unix(posixTime, 0)
	return timeResult
}

func ADIntervalToMins(intervalStr string) int {
	interval, _ := strconv.ParseInt(intervalStr, 10, 64)
	days := interval / (60 * 60 * 10000000 * -1)
	return int(days)
}

func ADIntervalToDays(intervalStr string) int {
	interval, _ := strconv.ParseInt(intervalStr, 10, 64)
	days := interval / (24 * 60 * 60 * 10000000 * -1)
	return int(days)
}
