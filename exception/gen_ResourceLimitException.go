// Code generated by gotemplate. DO NOT EDIT.

package exception

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/eosspark/eos-go/log"
)

// template type Exception(PARENT,CODE,WHAT)

var ResourceLimitExceptionName = reflect.TypeOf(ResourceLimitException{}).Name()

type ResourceLimitException struct {
	_ResourceLimitException
	Elog log.Messages
}

func NewResourceLimitException(parent _ResourceLimitException, message log.Message) *ResourceLimitException {
	return &ResourceLimitException{parent, log.Messages{message}}
}

func (e ResourceLimitException) Code() int64 {
	return 3210000
}

func (e ResourceLimitException) Name() string {
	return ResourceLimitExceptionName
}

func (e ResourceLimitException) What() string {
	return "Resource limit exception"
}

func (e *ResourceLimitException) AppendLog(l log.Message) {
	e.Elog = append(e.Elog, l)
}

func (e ResourceLimitException) GetLog() log.Messages {
	return e.Elog
}

func (e ResourceLimitException) TopMessage() string {
	for _, l := range e.Elog {
		if msg := l.GetMessage(); msg != "" {
			return msg
		}
	}
	return e.String()
}

func (e ResourceLimitException) DetailMessage() string {
	var buffer bytes.Buffer
	buffer.WriteString(strconv.Itoa(int(e.Code())))
	buffer.WriteString(" ")
	buffer.WriteString(e.Name())
	buffer.WriteString(": ")
	buffer.WriteString(e.What())
	buffer.WriteString("\n")
	for _, l := range e.Elog {
		buffer.WriteString("[")
		buffer.WriteString(l.GetMessage())
		buffer.WriteString("]")
		buffer.WriteString("\n")
		buffer.WriteString(l.GetContext().String())
		buffer.WriteString("\n")
	}
	return buffer.String()
}

func (e ResourceLimitException) String() string {
	return e.DetailMessage()
}

func (e ResourceLimitException) MarshalJSON() ([]byte, error) {
	type Exception struct {
		Code int64  `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
	}

	except := Exception{
		Code: 3210000,
		Name: ResourceLimitExceptionName,
		What: "Resource limit exception",
	}

	return json.Marshal(except)
}

func (e ResourceLimitException) Callback(f interface{}) bool {
	switch callback := f.(type) {
	case func(*ResourceLimitException):
		callback(&e)
		return true
	case func(ResourceLimitException):
		callback(e)
		return true
	default:
		return false
	}
}