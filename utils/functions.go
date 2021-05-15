package utils

import (
	"fmt"
	"math"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

////////////////////////////////////////////////////////
// exception
func CaptureException(p ...interface{}) {
	if err := recover(); err != nil {
		stack := string(debug.Stack())
		log.Error().Caller(1).Interface("exception_param", p).Msgf("catch exception:%v, panic recovered with stack:%s", err, stack)
	}
}

////////////////////////////////////////////////////////
// relocate to project root path
func RelocatePath(filter ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("RelocatePath failed: %w", err)
	}

	fmt.Println("work directory: ", wd)

	var newPath string = wd

	for _, path := range filter {
		n := strings.LastIndex(wd, path)
		if n == -1 {
			continue
		}

		newPath = strings.Join([]string{wd[:n], path}, "")
		err = os.Chdir(newPath)
		if err != nil {
			return err
		}

		break
	}

	fmt.Println("relocate to new_path:", newPath)
	return nil
}

////////////////////////////////////////////////////////
// between [a, b)
func Between(n, a, b int) bool {
	return (n >= a && n < b)
}

func BetweenInt32(n, a, b int32) bool {
	return (n >= a && n < b)
}

////////////////////////////////////////////////////////
// Round
func Round(val float64) float64 {
	return math.Round((val*10 + 0.1) / 10)
}

////////////////////////////////////////////////////////
// weekday
func PrevWeekday(d time.Weekday) time.Weekday { return (d - 1) % 7 }
func NextWeekday(d time.Weekday) time.Weekday { return (d + 1) % 7 }

////////////////////////////////////////////////////////
// shift
func HighId(id int64) int32 {
	return int32(id >> 32)
}

func LowId(id int64) int32 {
	return int32((id) & 0xFFFFFFFF)
}

func PackId(high, low int32) int64 {
	return (int64(high)<<32 | (int64(low) & 0x00000000FFFFFFFF))
}
