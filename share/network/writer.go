package network

import (
	"math"
	"reflect"

	"github.com/ubis/Freya/share/log"
)

const DEFAULT_BUFFER_SIZE = 1024

type Writer struct {
	buffer []byte
	index  int

	Type int
}

// Attempts to create a new packet writer and write initial packet header
func NewWriter(code uint16, magic ...uint16) *Writer {
	var w = &Writer{}

	w.buffer = make([]byte, DEFAULT_BUFFER_SIZE)
	w.index = 0

	w.WriteUint16(code) // packet type
	w.WriteUint16(0x00) // size

	w.Type = int(code)

	return w
}

func NewEmptyWriter() *Writer {
	var w = &Writer{}
	w.buffer = make([]byte, DEFAULT_BUFFER_SIZE)
	w.index = 0
	return w
}

// Checks buffer length and if it's too small, it will resize it
func (w *Writer) checkLength(length int) {
	if len(w.buffer) < w.index+length {
		// resize...
		var tmp = make([]byte, len(w.buffer)+DEFAULT_BUFFER_SIZE)
		copy(tmp[:len(w.buffer)], w.buffer)
		w.buffer = tmp

		// recursion
		w.checkLength(length)
	}
}

// Attempts to read specified interface type and serializes it into byte array.
// It has a length parameter, which tells the required size of interface type.
// If interface type length is smaller or higher than required, the correct
// length will be returned
func (w *Writer) getType(obj interface{}, length int) []byte {
	// check length
	w.checkLength(length)
	var tmp = make([]byte, 8)

	switch objType := obj.(type) {
	case int8:
		tmp[0] = byte(objType)
	case uint8:
		tmp[0] = byte(objType)
	case int16:
		tmp[0] = byte(objType)
		tmp[1] = byte(objType >> 8)
	case uint16:
		tmp[0] = byte(objType)
		tmp[1] = byte(objType >> 8)
	case int32:
		tmp[0] = byte(objType)
		tmp[1] = byte(objType >> 8)
		tmp[2] = byte(objType >> 16)
		tmp[3] = byte(objType >> 24)
	case uint32:
		tmp[0] = byte(objType)
		tmp[1] = byte(objType >> 8)
		tmp[2] = byte(objType >> 16)
		tmp[3] = byte(objType >> 24)
	case int:
		tmp[0] = byte(objType)
		tmp[1] = byte(objType >> 8)
		tmp[2] = byte(objType >> 16)
		tmp[3] = byte(objType >> 24)
	case int64:
		tmp[0] = byte(objType)
		tmp[1] = byte(objType >> 8)
		tmp[2] = byte(objType >> 16)
		tmp[3] = byte(objType >> 24)
		tmp[4] = byte(objType >> 32)
		tmp[5] = byte(objType >> 40)
		tmp[6] = byte(objType >> 48)
		tmp[7] = byte(objType >> 56)
	case uint64:
		tmp[0] = byte(objType)
		tmp[1] = byte(objType >> 8)
		tmp[2] = byte(objType >> 16)
		tmp[3] = byte(objType >> 24)
		tmp[4] = byte(objType >> 32)
		tmp[5] = byte(objType >> 40)
		tmp[6] = byte(objType >> 48)
		tmp[7] = byte(objType >> 56)
	case float32:
		bits := uint32(math.Float32bits(objType))
		tmp[0] = byte(bits)
		tmp[1] = byte(bits >> 8)
		tmp[2] = byte(bits >> 16)
		tmp[3] = byte(bits >> 24)
	case float64:
		bits := math.Float64bits(objType)
		tmp[0] = byte(bits)
		tmp[1] = byte(bits >> 8)
		tmp[2] = byte(bits >> 16)
		tmp[3] = byte(bits >> 24)
		tmp[4] = byte(bits >> 32)
		tmp[5] = byte(bits >> 40)
		tmp[6] = byte(bits >> 48)
		tmp[7] = byte(bits >> 56)
	default:
		log.Error("Unknown data type:", reflect.TypeOf(obj))
		return nil
	}

	return tmp[:length]
}

// Writes a bool
func (w *Writer) WriteBool(data bool) {
	value := 0
	if data {
		value = 1
	}

	w.buffer[w.index] = byte(value)
	w.index++
}

// Writes an signed byte
func (w *Writer) WriteSbyte(data interface{}) {
	var t = w.getType(data, 1)

	w.buffer[w.index] = byte(t[0])
	w.index++
}

// Writes an unsigned byte
func (w *Writer) WriteByte(data interface{}) {
	var t = w.getType(data, 1)

	w.buffer[w.index] = byte(t[0])
	w.index++
}

// Writes an signed 16-bit integer
func (w *Writer) WriteInt16(data interface{}) {
	var t = w.getType(data, 2)

	w.buffer[w.index] = byte(t[0])
	w.buffer[w.index+1] = byte(t[1])
	w.index += 2
}

// Writes an unsigned 16-bit integer
func (w *Writer) WriteUint16(data interface{}) {
	var t = w.getType(data, 2)

	w.buffer[w.index] = byte(t[0])
	w.buffer[w.index+1] = byte(t[1])
	w.index += 2
}

// Writes an signed 32-bit integer
func (w *Writer) WriteInt32(data interface{}) {
	var t = w.getType(data, 4)

	w.buffer[w.index] = byte(t[0])
	w.buffer[w.index+1] = byte(t[1])
	w.buffer[w.index+2] = byte(t[2])
	w.buffer[w.index+3] = byte(t[3])
	w.index += 4
}

// Writes an unsigned 32-bit integer
func (w *Writer) WriteUint32(data interface{}) {
	var t = w.getType(data, 4)

	w.buffer[w.index] = byte(t[0])
	w.buffer[w.index+1] = byte(t[1])
	w.buffer[w.index+2] = byte(t[2])
	w.buffer[w.index+3] = byte(t[3])
	w.index += 4
}

// Writes an signed 64-bit integer
func (w *Writer) WriteInt64(data interface{}) {
	var t = w.getType(data, 8)

	w.buffer[w.index] = byte(t[0])
	w.buffer[w.index+1] = byte(t[1])
	w.buffer[w.index+2] = byte(t[2])
	w.buffer[w.index+3] = byte(t[3])
	w.buffer[w.index+4] = byte(t[4])
	w.buffer[w.index+5] = byte(t[5])
	w.buffer[w.index+6] = byte(t[6])
	w.buffer[w.index+7] = byte(t[7])
	w.index += 8
}

// Writes an unsigned 64-bit integer
func (w *Writer) WriteUint64(data interface{}) {
	var t = w.getType(data, 8)

	w.buffer[w.index] = byte(t[0])
	w.buffer[w.index+1] = byte(t[1])
	w.buffer[w.index+2] = byte(t[2])
	w.buffer[w.index+3] = byte(t[3])
	w.buffer[w.index+4] = byte(t[4])
	w.buffer[w.index+5] = byte(t[5])
	w.buffer[w.index+6] = byte(t[6])
	w.buffer[w.index+7] = byte(t[7])
	w.index += 8
}

// Writes a double-precision floating point number (float64)
func (w *Writer) WriteDouble(data float64) {
	bits := math.Float64bits(data)
	w.WriteUint64(bits) // Use existing WriteUint64 to write the bits
}

// Writes a single-precision floating point number (float32)
func (w *Writer) WriteFloat(data float32) {
	bits := math.Float32bits(data)
	w.WriteUint32(bits) // Use existing WriteUint32 to write the bits
}

// Writes a string with length prefixed
func (w *Writer) WriteString(data string) {
	// คำนวณความยาวของข้อมูล string
	dataLength := len(data)

	// เช็คความยาวของข้อมูล
	w.checkLength(w.index + 2 + dataLength) // เพิ่ม 2 ไบต์สำหรับขนาดของข้อมูล

	// เขียนขนาดข้อมูล (data length) ด้วยรูปแบบ Little-endian 2 ไบต์
	w.buffer[w.index] = byte(dataLength & 0xFF)
	w.buffer[w.index+1] = byte((dataLength >> 8) & 0xFF)
	w.index += 2

	// เขียนข้อมูล string ต่อจากขนาดข้อมูล
	copy(w.buffer[w.index:], []byte(data))
	w.index += dataLength
}

// Writes a string with length prefixed
func (w *Writer) AppendString(data string, length int) {
	if length <= 0 {
		length = len(data)
	}

	src := []byte(data)
	dst := make([]byte, length)
	copy(dst, src)
	w.WriteBytes(dst)
}

// Writes an byte array
func (w *Writer) WriteBytes(data []byte) {
	// check length
	w.checkLength(len(data))

	copy(w.buffer[w.index:], data)
	w.index += len(data)
}

/*
Updates packet length and returns byte array
@return byte array of packet
*/
func (w *Writer) Finalize() []byte {
	// update size
	var length = w.index
	newLength := int16(length - 4) // Calculate the new length as int16

	// Convert newLength to bytes (little-endian)
	w.buffer[2] = byte(newLength)
	w.buffer[3] = byte(newLength >> 8)

	// create a new slice and copy the data
	result := make([]byte, length)  // Create a new byte slice with the original length
	copy(result, w.buffer[:length]) // Copy the first 'length' bytes from w.buffer to result

	return result // Return the resulting byte array
}

func (w *Writer) RawBytes() []byte {
	var length = w.index

	result := make([]byte, length)  // Create a new byte slice with the original length
	copy(result, w.buffer[:length]) // Copy the first 'length' bytes from w.buffer to result
	return result
}
