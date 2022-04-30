package vr_test

import (
	"bytes"
	"context"
	"testing"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/shun159/vr"
)

func TestInet6Route(t *testing.T) {
    ctx := context.Background()
    trans := thrift.NewTMemoryBuffer()
    proto := vr.NewTSandeshProtocolTransport(trans)
    conf := vr.NewInet6RouteConfig()
    conf.Vrf = 1
    conf.NhIdx = 2
    conf.IPAddress = "fe80::2e6d:c1ff:fe13:4f61"
    conf.PrefixLen = 64
    req1, _ := vr.NewInet6Route(conf)
    if err := req1.Write(ctx, proto); err != nil {
        t.Errorf("Failed to encoding")
    }
 
    req2, _ := vr.NewInet6Route(&vr.Inet6RouteConfig{})
    b := trans.Bytes()
    trans.Buffer = bytes.NewBuffer(b)
    
    if err := req2.Read(ctx, proto); err != nil {
        t.Errorf("Failed to decoding")
    }

    if req2.GetRtrVrfID() != 1 {
        t.Errorf("RtrVrfID must be 1")
    }

    if req2.GetRtrNhID() != 2 {
        t.Errorf("RtrNhID must be 2")
    }

    if req2.GetPrefixString() != "fe80::2e6d:c1ff:fe13:4f61" {
        t.Errorf("RtrPrefix must be fe80::2e6d:c1ff:fe13:4f61")
    }

    if req2.GetRtrPrefixLen() != 64 {
        t.Errorf("RtrPrefix must be 64")
    }
}

func TestInet4Route(t *testing.T) {
    ctx := context.Background()
    trans := thrift.NewTMemoryBuffer()
    proto := vr.NewTSandeshProtocolTransport(trans)
    conf := vr.NewInetRouteConfig()
    conf.Vrf = 1
    conf.NhIdx = 2
    conf.IPAddress = "192.168.32.0"
    conf.PrefixLen = 24
    conf.MacAddress = "01:02:03:04:05:06"
    conf.Label = 3
    conf.LabelFlag = 4
    req1, _ := vr.NewInetRoute(conf)
    if err := req1.Write(ctx, proto); err != nil {
        t.Errorf("Failed to encoding")
    }
 
    req2, _ := vr.NewInetRoute(&vr.InetRouteConfig{})
    b := trans.Bytes()
    trans.Buffer = bytes.NewBuffer(b)
    
    if err := req2.Read(ctx, proto); err != nil {
        t.Errorf("Failed to decoding")
    }

    if req2.GetRtrVrfID() != 1 {
        t.Errorf("RtrVrfID must be 1")
    }

    if req2.GetRtrNhID() != 2 {
        t.Errorf("RtrNhID must be 2")
    }

    if req2.GetPrefixString() != "192.168.32.0" {
        t.Errorf("RtrPrefix must be 192.168.32.0 but %s", req2.GetPrefixString())
    }

    if req2.GetRtrPrefixLen() != 24 {
        t.Errorf("RtrPrefixLen must be 24")
    }

    if req2.GetMacAddressString() != "01:02:03:04:05:06" {
        t.Errorf("RtrMac must be 01:02:03:04:05:06 but %s", req2.GetMacAddressString())
    }

    if req2.GetRtrLabel() != 3 {
        t.Errorf("RtrLabel must be 3")
    } 

    if req2.GetRtrLabelFlags() != 4 {
        t.Errorf("RtrLabelFlags must be 4")
    }
}

func TestBridgeRoute(t *testing.T) {
    ctx := context.Background()
    trans := thrift.NewTMemoryBuffer()
    proto := vr.NewTSandeshProtocolTransport(trans)
    conf := vr.NewBridgeRouteConfig()
    conf.Vrf = 1
    conf.NhIdx = 2
    conf.MacAddress = "01:02:03:04:05:06"
    conf.Label = 3
    conf.LabelFlag = 4
    req1, _ := vr.NewBridgeRoute(conf)
    if err := req1.Write(ctx, proto); err != nil {
        t.Errorf("Failed to encoding")
    }
 
    req2, _ := vr.NewInetRoute(&vr.InetRouteConfig{})
    b := trans.Bytes()
    trans.Buffer = bytes.NewBuffer(b)
    
    if err := req2.Read(ctx, proto); err != nil {
        t.Errorf("Failed to decoding")
    }

    if req2.GetRtrVrfID() != 1 {
        t.Errorf("RtrVrfID must be 1")
    }

    if req2.GetRtrNhID() != 2 {
        t.Errorf("RtrNhID must be 2")
    }

    if req2.GetMacAddressString() != "01:02:03:04:05:06" {
        t.Errorf("RtrMac must be 01:02:03:04:05:06 but %s", req2.GetMacAddressString())
    }

    if req2.GetRtrLabel() != 3 {
        t.Errorf("RtrLabel must be 3")
    } 

    if req2.GetRtrLabelFlags() != 4 {
        t.Errorf("RtrLabelFlags must be 4")
    }
}
