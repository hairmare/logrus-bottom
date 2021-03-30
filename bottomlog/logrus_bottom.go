package bottomlog

import (
	"bytes"

	"github.com/nihaals/bottom-go/bottom"
	"github.com/sirupsen/logrus"
)

// BottomFormatter formats messages, duh
type BottomFormatter struct {
}

// NewBottomFormatter gives you a new bottom formatter
func NewBottomFormatter() *BottomFormatter {
	return &BottomFormatter{}
}

// Format dominates your log message 👉👈
func (dom *BottomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	b.WriteString(bottom.Encode(entry.Message))
	b.WriteByte('\n')
	return b.Bytes(), nil
}
