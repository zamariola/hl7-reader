package hl7reader

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMessage(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    *Message
		wantErr bool
	}{
		{"Empty (nil)", []byte(nil), nil, true},
		{"Empty (not nil)", []byte{}, nil, true},
		{"Too short", []byte(`MSH|^~\`), nil, true},
		{
			"Minimal example",
			[]byte(`MSH|^~\&`),
			&Message{
				reader:     bufio.NewReader(bytes.NewBuffer([]byte(`MSH|^~\&`))),
				fieldSep:   '|',
				compSep:    '^',
				subCompSep: '&',
				repeat:     '~',
				escape:     '\\',
			},
			false,
		},
		{
			"Custom separators",
			[]byte("MSH....."),
			&Message{
				reader:     bufio.NewReader(bytes.NewBuffer([]byte("MSH....."))),
				fieldSep:   '.',
				compSep:    '.',
				subCompSep: '.',
				repeat:     '.',
				escape:     '.',
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMessage(tt.data)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMessageReadSegment(t *testing.T) {
	tests := []struct {
		name  string
		data  []byte
		count int
	}{
		{"one segment", []byte("MSH|^~\\&"), 1},
		{"two segments", []byte("MSH|^~\\&\rMSH|^~\\&"), 2},
		{"two segments, extra whitespace", []byte("MSH|^~\\&\r\nMSH|^~\\&"), 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg, _ := NewMessage(tt.data)

			for i := 0; i < tt.count; i++ {
				_, err := msg.ReadSegment()

				assert.Nil(t, err)
			}
			_, err := msg.ReadSegment()

			assert.Error(t, err)
		})
	}
}
