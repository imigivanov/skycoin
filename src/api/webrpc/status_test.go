package webrpc

import (
	"reflect"
	"testing"
)

func Test_getStatusHandler(t *testing.T) {
	type args struct {
		req Request
		in1 Gatewayer
	}
	tests := []struct {
		name string
		args args
		want Response
	}{
		// TODO: Add test cases.
		{
			"normal",
			args{
				req: Request{
					ID:      "1",
					Method:  "get_status",
					Jsonrpc: jsonRPC,
				},
				in1: &fakeGateway{},
			},
			Response{
				ID:      ptrString("1"),
				Jsonrpc: jsonRPC,
				Result:  StatusResult{true},
			},
		},
		{
			"invalid params",
			args{
				req: Request{
					ID:      "1",
					Method:  "get_status",
					Jsonrpc: jsonRPC,
					Params:  map[string]string{"abc": "123"},
				},
				in1: &fakeGateway{},
			},
			makeErrorResponse(&RPCError{
				Code:    errCodeInvalidParams,
				Message: errMsgInvalidParams,
			}),
		},
	}
	for _, tt := range tests {
		if got := getStatusHandler(tt.args.req, tt.args.in1); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. getStatusHandler() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
