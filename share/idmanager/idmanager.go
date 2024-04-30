package idmanager

import "errors"

type IDManager struct {
	nextID    uint16          // ค่า id ถัดไปที่จะใช้
	cancelled map[uint16]bool // เก็บข้อมูลเพื่อบอกว่า id ถูกยกเลิกไปแล้วหรือยัง
}

// NewIDManager สร้าง instance ใหม่ของ IDManager
func NewIDManager() *IDManager {
	return &IDManager{
		nextID:    1,
		cancelled: make(map[uint16]bool),
	}
}

// GetID คืนค่า unique id ใหม่
func (im *IDManager) GetID() (uint16, error) {
	// หา id ที่ถูกยกเลิกไว้เพื่อนำกลับมาใช้ใหม่
	for id := range im.cancelled {
		delete(im.cancelled, id)
		return id, nil
	}

	// ถ้าไม่มี id ถูกยกเลิกใช้ไปแล้ว จะสร้าง id ใหม่
	if im.nextID == 0 { // กรณี id เต็มแล้วให้คืน error
		return 0, errors.New("no available ID")
	}

	id := im.nextID
	im.nextID++

	if im.nextID == 0 { // ถ้าค่า nextID เกินขีดจำกัดของ uint16 ให้กลับไปเริ่มที่ 1
		im.nextID = 1
	}

	return id, nil
}

// CancelID ยกเลิกการใช้งาน id ที่กำหนด
func (im *IDManager) CancelID(id uint16) error {
	if id == 0 {
		return errors.New("invalid ID")
	}

	if _, exists := im.cancelled[id]; exists {
		return errors.New("ID already cancelled")
	}

	im.cancelled[id] = true
	return nil
}
