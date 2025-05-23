package entity

import "time"

type Traffic struct {
	Interface   Interface `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	InterfaceID uint64
	Date        time.Time
	Bandwidth   uint64
	In          uint64
	Out         uint64
	BytesIn     uint64
	BytesOut    uint64
}

type TrafficOnt struct {
	Interface        Interface `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	InterfaceID      uint64
	Idx              uint64
	Despt            *string
	SerialNumber     *string
	LineProfName     *string
	ControlMacCount  *int64
	OltDistance      *int64
	ControlRunStatus *int8
	InBps            *uint64
	OutBps           *uint64
	BytesIn          *uint64
	BytesOut         *uint64
	Date             time.Time
}

type TrafficResponse struct {
	Date      time.Time `gorm:"column:date"`
	In        float64   `gorm:"column:mbps_in"`
	Out       float64   `gorm:"column:mbps_out"`
	Bandwidth float64   `gorm:"column:bandwidth"`
	BytesIn   float64   `gorm:"column:mbytes_in"`
	BytesOut  float64   `gorm:"column:mbytes_out"`
}
