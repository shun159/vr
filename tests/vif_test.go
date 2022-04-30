package vr_test

import (
	"bytes"
	"context"
	"testing"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/shun159/vr"
)

func TestVirtualVif(t *testing.T) {
    ctx := context.Background()
    trans := thrift.NewTMemoryBuffer()
    proto := vr.NewTSandeshProtocolTransport(trans)
    conf := vr.NewVirtualVifConfig()
    conf.Idx = 1
    conf.Name = "virtual0"
    conf.MacAddr = "01:02:03:04:05:06"
    conf.IpAddr = "1.1.1.1"
    conf.Mtu = 1514
    conf.Flags = 3
    conf.Transport = vr.VIF_TRANSPORT_ETH
    conf.Vrf = 4
    conf.McastVrf = 0xffff

    req1, _ := vr.NewVirtualVif(conf)
    if err := req1.Write(ctx, proto); err != nil {
        t.Errorf("Failed to encoding")
    }

    req2, _ := vr.NewVirtualVif(vr.NewVirtualVifConfig())

    b := trans.Bytes()
    trans.Buffer = bytes.NewBuffer(b)

    if err := req2.Read(ctx, proto); err != nil {
        t.Errorf("Failed to decoding")
    }

    if req2.VifrType != vr.VIF_TYPE_VIRTUAL {
        t.Errorf("VifrType must be VIF_TYPE_VIRTUAL")
    }

    if req2.VifrName != "virtual0" {
        t.Errorf("VifrName must be virtual0")
    }

    if req2.VifrIdx != 1 {
        t.Errorf("VifrIdx must be 1")
    }

    if req2.GetIPAddressString() != "1.1.1.1" {
        t.Errorf("VifrIP must be 1.1.1.1")
    }

    if req2.GetMacAddressString() != "01:02:03:04:05:06" {
        t.Errorf("VifrMac must be 01:02:03:04:05:06")
    }

    if req2.VifrMtu != 1514 {
        t.Errorf("VifrMtu must be 1514")
    }

    if req2.VifrFlags != 3 {
        t.Errorf("VifrFlags must be 3")
    }

    if req2.VifrTransport != vr.VIF_TRANSPORT_ETH {
        t.Errorf("VifrTransport must be VIF_TRANSPORT_ETH")
    }

    if req2.VifrVrf != 4 {
        t.Errorf("VifrVrf must be 4")
    }

    if req2.VifrMcastVrf != 0xffff {
        t.Errorf("VifrMcastVrf must be 0xffff")
    }
}

func TestFabricVif(t *testing.T) {
    ctx := context.Background()
    trans := thrift.NewTMemoryBuffer()
    proto := vr.NewTSandeshProtocolTransport(trans)
    conf := vr.NewFabricVifConfig()
    conf.Idx = 1
    conf.Name = "fabric0"
    conf.MacAddr = "01:02:03:04:05:06"
    conf.IpAddr = "1.1.1.1"
    conf.Mtu = 1514
    conf.Flags = 322
    conf.Transport = vr.VIF_TRANSPORT_ETH
    conf.Vrf = 4
    conf.McastVrf = 0xffff

    req1, _ := vr.NewFabricVif(conf)
    if err := req1.Write(ctx, proto); err != nil {
        t.Errorf("Failed to encoding")
    }

    req2, _ := vr.NewFabricVif(vr.NewFabricVifConfig())

    b := trans.Bytes()
    trans.Buffer = bytes.NewBuffer(b)

    if err := req2.Read(ctx, proto); err != nil {
        t.Errorf("Failed to decoding")
    }

    if req2.VifrType != vr.VIF_TYPE_PHYSICAL {
        t.Errorf("VifrType must be VIF_TYPE_PHYSICAL")
    }

    if req2.VifrName != "fabric0" {
        t.Errorf("VifrName must be fabric0")
    }

    if req2.VifrIdx != 1 {
        t.Errorf("VifrIdx must be 1")
    }

    if req2.GetIPAddressString() != "1.1.1.1" {
        t.Errorf("VifrIP must be 1.1.1.1")
    }

    if req2.GetMacAddressString() != "01:02:03:04:05:06" {
        t.Errorf("VifrMac must be 01:02:03:04:05:06")
    }

    if req2.VifrMtu != 1514 {
        t.Errorf("VifrMtu must be 1514")
    }

    if req2.VifrFlags != 322 {
        t.Errorf("VifrFlags must be 322")
    }

    if req2.VifrTransport != vr.VIF_TRANSPORT_ETH {
        t.Errorf("VifrTransport must be VIF_TRANSPORT_ETH")
    }

    if req2.VifrVrf != 4 {
        t.Errorf("VifrVrf must be 4")
    }

    if req2.VifrMcastVrf != 0xffff {
        t.Errorf("VifrMcastVrf must be 0xffff")
    }
}

func TestVhostVif(t *testing.T) {
    ctx := context.Background()
    trans := thrift.NewTMemoryBuffer()
    proto := vr.NewTSandeshProtocolTransport(trans)
    conf := vr.NewVhostVifConfig()
    conf.Idx = 1
    conf.Name = "vhost0"
    conf.MacAddr = "01:02:03:04:05:06"
    conf.IpAddr = "1.1.1.1"
    conf.Mtu = 1514
    conf.Transport = vr.VIF_TRANSPORT_ETH
    conf.Vrf = 4
    conf.McastVrf = 0xffff

    req1, _ := vr.NewVhostVif(conf)
    if err := req1.Write(ctx, proto); err != nil {
        t.Errorf("Failed to encoding")
    }

    req2, _ := vr.NewVhostVif(vr.NewVhostVifConfig())

    b := trans.Bytes()
    trans.Buffer = bytes.NewBuffer(b)

    if err := req2.Read(ctx, proto); err != nil {
        t.Errorf("Failed to decoding")
    }

    if req2.VifrType != vr.VIF_TYPE_HOST {
        t.Errorf("VifrType must be VIF_TYPE_HOST")
    }

    if req2.VifrName != "vhost0" {
        t.Errorf("VifrName must be vhost0")
    }

    if req2.VifrIdx != 1 {
        t.Errorf("VifrIdx must be 1")
    }

    if req2.GetIPAddressString() != "1.1.1.1" {
        t.Errorf("VifrIP must be 1.1.1.1")
    }

    if req2.GetMacAddressString() != "01:02:03:04:05:06" {
        t.Errorf("VifrMac must be 01:02:03:04:05:06")
    }

    if req2.VifrMtu != 1514 {
        t.Errorf("VifrMtu must be 1514")
    }

    if req2.VifrFlags != 384 {
        t.Errorf("VifrFlags must be 384")
    }

    if req2.VifrTransport != vr.VIF_TRANSPORT_ETH {
        t.Errorf("VifrTransport must be VIF_TRANSPORT_ETH")
    }

    if req2.VifrVrf != 4 {
        t.Errorf("VifrVrf must be 4")
    }

    if req2.VifrMcastVrf != 0xffff {
        t.Errorf("VifrMcastVrf must be 0xffff")
    }
}


func TestAgentVif(t *testing.T) {
    ctx := context.Background()
    trans := thrift.NewTMemoryBuffer()
    proto := vr.NewTSandeshProtocolTransport(trans)
    conf := vr.NewAgentVifConfig()
    conf.Idx = 1
    conf.Name = "pkt0"
    conf.Mtu = 1514

    req1, _ := vr.NewAgentVif(conf)
    if err := req1.Write(ctx, proto); err != nil {
        t.Errorf("Failed to encoding")
    }

    req2, _ := vr.NewAgentVif(vr.NewAgentVifConfig())

    b := trans.Bytes()
    trans.Buffer = bytes.NewBuffer(b)

    if err := req2.Read(ctx, proto); err != nil {
        t.Errorf("Failed to decoding")
    }

    if req2.VifrType != vr.VIF_TYPE_AGENT {
        t.Errorf("VifrType must be VIF_TYPE_AGENT")
    }

    if req2.VifrName != "pkt0" {
        t.Errorf("VifrName must be pkt0")
    }

    if req2.VifrIdx != 1 {
        t.Errorf("VifrIdx must be 1")
    }

    if req2.VifrMtu != 1514 {
        t.Errorf("VifrMtu must be 1514")
    }

    if req2.VifrFlags != 320 {
        t.Errorf("VifrFlags must be 320")
    }

    if req2.VifrTransport != vr.VIF_TRANSPORT_SOCKET {
        t.Errorf("VifrTransport must be VIF_TRANSPORT_SOCKET")
    }

    if req2.VifrVrf != 0xffff {
        t.Errorf("VifrVrf must be 4")
    }

    if req2.VifrMcastVrf != 0xffff {
        t.Errorf("VifrMcastVrf must be 0xffff")
    }
}
