package utils

import (
	"strings"

	"github.com/Sirupsen/logrus"
)

type Logger struct {
	*logrus.Entry
}

func (logger Logger) Output(calldepth int, s string) error {
	//l := logger.WithField("calldepth", calldepth)
	l := logger
	m := s[(strings.Index(s, "(")):]
	switch s[:3] {
	case "INF":
		l.Info(m)
	case "WRN":
		l.Warn(m)
	case "ERR":
		l.Error(m)
	case "DBG":
		l.Debug(m)
	}
	return nil
}
