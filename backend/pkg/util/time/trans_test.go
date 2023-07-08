package util

import (
	"fmt"
	"testing"
	"time"

	"kratos-uba/pkg/util/trans"
)

func TestUnixMilliToStringPtr(t *testing.T) {
	now := time.Now().UnixMilli()
	str := UnixMilliToStringPtr(&now)
	fmt.Println(now)
	fmt.Println(*str)

	fmt.Println(*UnixMilliToStringPtr(trans.Int64(1677135885288)))
	fmt.Println(*UnixMilliToStringPtr(trans.Int64(1677647430853)))
	fmt.Println(*UnixMilliToStringPtr(trans.Int64(1677647946234)))
	fmt.Println(*UnixMilliToStringPtr(trans.Int64(1678245350773)))

	fmt.Println("START: ", *StringToUnixMilliInt64Ptr(trans.String("2023-03-09 00:00:00")))
	fmt.Println("END: ", *StringToUnixMilliInt64Ptr(trans.String("2023-03-09 23:59:59")))

	fmt.Println(StringTimeToTime(trans.String("2023-03-01 00:00:00")))
	fmt.Println(*StringDateToTime(trans.String("2023-03-01")))

	fmt.Println(StringTimeToTime(trans.String("2023-03-08 00:00:00")).UnixMilli())
	fmt.Println(StringDateToTime(trans.String("2023-03-07")).UnixMilli())
}

func TestTimeToDateString(t *testing.T) {
	fmt.Println(*TimeToTimeString(trans.Time(time.Now())))
	fmt.Println(*TimeToDateString(trans.Time(time.Now())))
}
