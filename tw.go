package tw

import (
	"bytes"
	"text/tabwriter"
)

const (
	defaultMinWidth = 0
	defaultTabWidth = 8
	defaultPadding  = 1
	defaultPadchar  = ' '
)

func defaultFlag() uint {
	return tabwriter.Debug
}

type DefaultTabWriter struct {
	buf    *bytes.Buffer
	writer *tabwriter.Writer
}

func Default() *DefaultTabWriter {
	buf := bytes.NewBuffer([]byte{})
	return &DefaultTabWriter{
		buf:    buf,
		writer: tabwriter.NewWriter(buf, defaultMinWidth, defaultTabWidth, defaultPadding, defaultPadchar, defaultFlag()),
	}
}

func (tw *DefaultTabWriter) WriteString(s string) {
	tw.writer.Write([]byte(s + "\t"))
}

func (tw *DefaultTabWriter) FlushAndString() (string, error) {
	if err := tw.writer.Flush(); err != nil {
		return "", err
	}

	return tw.buf.String(), nil
}
