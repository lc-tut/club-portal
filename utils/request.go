package utils

//
//import (
//	"context"
//	"io"
//	"net/http"
//)
//
//type req struct {
//	context context.Context
//	method  string
//	url     string
//	body    io.Reader
//	header  http.Header
//}
//
//type Request interface {
//	WithHeader() Request
//	WithBody() Request
//	Do() Request
//}
//
//func NewRequest(ctx context.Context, method string, url string) *req {
//	return &req{
//		context: ctx,
//		method:  method,
//		url:     url,
//	}
//}
//
//func (r *req) WithHeader() *req {
//	if r.header == nil {
//		r.header = make(http.Header)
//	}
//}
