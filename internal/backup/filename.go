package backup

import (
	"log"
	"strings"
	"time"
)

const timestampLayout = "2006-01-02_15-14-00-000MST"

func GenerateBackupFilePath(format string, timestamp time.Time) string {
	var sb strings.Builder

	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Println(err)
	}

	sb.WriteString("backup__")
	sb.WriteString(timestamp.In(tz).Format(timestampLayout))
	sb.WriteRune('.')
	sb.WriteString(format)

	return sb.String()
}
