package vr

import "github.com/shun159/vr/vr"

// VXLAN interface
type Vxlan struct {
    *vr.VrVxlanReq
}

func NewVxlan(rid int16, vnid, nhid int32) (*Vxlan, error) {
	vxlan := &Vxlan{}
	vxlan.VrVxlanReq = vr.NewVrVxlanReq()
	vxlan.HOp = SANDESH_OPER_ADD
	vxlan.VxlanrRid = rid
	vxlan.VxlanrVnid = vnid
	vxlan.VxlanrNhid = nhid
	return vxlan, nil
}
