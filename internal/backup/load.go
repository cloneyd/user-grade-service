package backup

import (
	"compress/gzip"
	"net/http"
	"os"
	"sync"
	"time"
)

func LoadBackup(datasource *sync.Map) (time.Time, error) {
	resp, err := http.Get(os.Getenv("MASTER_BACKUP_ADDRESS"))
	if err != nil {
		return time.Now(), err
	}

	zipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return time.Now(), err
	}

	userGrades, err := ReadDump(zipReader)
	if err != nil {
		return time.Now(), err
	}

	for _, userGrade := range userGrades {
		datasource.Store(userGrade.UserId, &userGrade)
	}

	if err = resp.Body.Close(); err != nil {
		return time.Now(), err
	}

	filename := resp.Header.Get("Content-Disposition")

	return GetTimestampFromFilename(filename)
}
