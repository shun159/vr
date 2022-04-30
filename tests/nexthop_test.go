package vr_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/shun159/vr"
)

func TestEncapNexthop(t *testing.T) {
    ctx := context.Background()
    trans := thrift.NewTMemoryBuffer()
    proto := vr.NewTSandeshProtocolTransport(trans)
    conf := vr.EncapNexthopConfig {
        Idx: 999,
        EncapOuterVifId: []int32{1, 2},
        Encap: []byte{3, 4},
        Flags: 5,
        Family: 2,
        EncapFamily: 9,
        Vrf: 6,
    }
    req1, _ := vr.NewEncapNexthop(&conf)
    if err := req1.Write(ctx, proto); err != nil {
        t.Errorf("Failed to encoding")
    }
 
    req2, _ := vr.NewEncapNexthop(&vr.EncapNexthopConfig{})
    b := trans.Bytes()
    trans.Buffer = bytes.NewBuffer(b)
    
    if err := req2.Read(ctx, proto); err != nil {
        t.Errorf("Failed to decoding")
    }

    if req2.GetNhrType() != vr.NH_TYPE_ENCAP {
        t.Errorf("NhrType must be NH_TYPE_ENCAP(2)")
    }

    if req2.GetNhrID() != 999 {
        t.Errorf("NhrID must be 999")
    }

    if req2.GetNhrEncapOifID()[0] != 1 && req2.GetNhrEncapOifID()[1] != 2 {
        t.Errorf("NhrEncapOifID must be [1, 2]")
    }

    if req2.GetNhrEncap()[0] != 3 && req2.GetNhrEncap()[1] != 4 {
        t.Errorf("NhrEncap must be [3, 4]")
    }

    if req2.GetNhrFlags() != 5 {
        t.Errorf("NhrFlags must be 5")
    }

    if req2.GetNhrFamily() != 2 {
        t.Errorf("NhrFamily must be 2")
    }

    if req2.GetNhrEncapFamily() != 9 {
        t.Errorf("NhrEncapFamily must be 9")
    }

    if req2.GetNhrVrf() != 6 {
        t.Errorf("NhrVrf must be 6")
    }
}

func TestRcvNexthop(t *testing.T) {
    ctx := context.Background()
    trans := thrift.NewTMemoryBuffer()
    proto := vr.NewTSandeshProtocolTransport(trans)
    conf := vr.NewReceiveNexthopConfig()
    conf.Idx = 1
    conf.EncapOifId = []int32{2, 3}
    conf.Encap = []byte{4, 5}
    conf.EncapFamily = 6
    conf.Vrf = 7
    req1, _ := vr.NewReceiveNexthop(conf)

    if err := req1.Write(ctx, proto); err != nil {
        t.Errorf("Failed to encoding")
    }

    req2, _ := vr.NewReceiveNexthop(&vr.ReceiveNexthopConfig{})
    b := trans.Bytes()
    trans.Buffer = bytes.NewBuffer(b)

    if err := req2.Read(ctx, proto); err != nil {
        t.Errorf("Failed to decoding")
    }

    if req2.GetNhrType() != vr.NH_TYPE_RCV {
        t.Errorf("NhrType must be NH_TYPE_RCV(1)")
    }

    if req2.GetNhrFamily() != 2 {
        t.Errorf("NhrFamily must be 2")
    }

    if req2.GetNhrFlags() != 1 {
        t.Errorf("NhrFlags must be 1")
    }

    if len(req2.GetNhrEncapOifID()) != 2 {
        t.Errorf("NhrEncapOifId must contains 2 elements")
    }

    if len(req2.GetNhrEncap()) != 2 {
        t.Errorf("NhrEncap must contais 2 elements")
    }

    if req2.GetNhrEncapFamily() != 6 {
        t.Errorf("NhrEncapFamily must be 6")
    }

    if req2.GetNhrVrf() != 7 {
        t.Errorf("NhrVrf must be 2")
    }
}
