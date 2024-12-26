package enum

import "fmt"

// StatusEnum, Timing için durumları temsil eder
type StatusEnum int

// Enum değerleri
const (
	StatusPaused    StatusEnum = iota // 0
	StatusStarted                     // 1
	StatusStopped                     // 2
	StatusCompleted                   // 3
)

// statusStrings, StatusEnum değerlerinin string karşılıkları
var statusStrings = []string{
	"Paused",
	"Started",
	"Stopped",
	"Completed",
}

// String, StatusEnum için string karşılığını döndürür
func (s StatusEnum) String() string {
	if s < 0 || int(s) >= len(statusStrings) {
		return "Unknown" // Eğer geçersiz bir değer varsa
	}
	return statusStrings[s]
}

// ParseStatus, bir string değeri StatusEnum'a dönüştürür
func ParseStatus(value string) (StatusEnum, error) {
	for i, v := range statusStrings {
		if v == value {
			return StatusEnum(i), nil
		}
	}
	return -1, fmt.Errorf("geçersiz durum: %s", value)
}

// IsValid, StatusEnum'un geçerli bir değer olup olmadığını kontrol eder
func (s StatusEnum) IsValid() bool {
	return s >= StatusPaused && s <= StatusCompleted
}
