package vr_test

import (
	"testing"
    "context"
    "bytes"
	"github.com/shun159/vr"
    "github.com/apache/thrift/lib/go/thrift"
)

func TestVxlan(t *testing.T) {
    ctx := context.Background()
    trans := thrift.NewTMemoryBuffer()
    proto := vr.NewTSandeshProtocolTransport(trans)

    req1, _ := vr.NewVxlan(1, 2, 3)
    if err := req1.Write(ctx, proto); err != nil {
        t.Errorf("Failed to encoding")
    }

    req2, _ := vr.NewVxlan(-1, -1, -1)
    b := trans.Bytes()
    trans.Buffer = bytes.NewBuffer(b)

    if err := req2.Read(ctx, proto); err != nil {
        t.Errorf("Failed to decoding")
    }

    if req2.GetVxlanrRid() !=  1 {
        t.Errorf("VxlanrRid must be 1")
    }

    if req2.GetVxlanrVnid() !=  2 {
        t.Errorf("VxlanrVnid must be 2")
    }

    if req2.GetVxlanrNhid() !=  3 {
        t.Errorf("VxlanrNhid must be 3")
    }
}

