package hash_table

import (
	"net"
	"encoding/binary"
	"encoding/gob"
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/profzone/imblock/global"
	"github.com/profzone/imblock/core"
)

func packMessage(msg *core.Protocol) []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(msg)
	if err != nil {
		logrus.Fatal(err)
	}
	resultBytes := result.Bytes()
	dataLength := uint32(len(resultBytes))
	dataLengthBuffer := bytes.NewBuffer([]byte{})
	binary.Write(dataLengthBuffer, binary.BigEndian, dataLength)
	return append(dataLengthBuffer.Bytes(), resultBytes[:]...)
}

func unpackMessageFromPackage(source []byte) *core.Protocol {

	if len(source) <= int(global.HeaderLength) {
		logrus.Panicf("len(source) <= %d", global.HeaderLength)
	}

	source = source[global.HeaderLength:]

	var msg core.Protocol
	decoder := gob.NewDecoder(bytes.NewReader(source))
	err := decoder.Decode(&msg)
	if err != nil {
		logrus.Panic(err)
	}

	return &msg
}

func unpackMessage(source []byte, readedMessage chan *core.Protocol, remoteAddr net.Addr) []byte {
	length := len(source)
	if length == 0 {
		return source
	}

	// 分包
	var i uint32
	for i = 0; i < uint32(length); i = i + 1 {
		// 获取包长信息
		data := source[i : i+global.HeaderLength]
		packageLength := binary.BigEndian.Uint32(data)

		// 获取包数据
		offset := i + global.HeaderLength
		if offset+packageLength > uint32(length) {
			break
		}
		packageData := source[offset : offset+packageLength]

		var msg core.Protocol
		decoder := gob.NewDecoder(bytes.NewReader(packageData))
		err := decoder.Decode(&msg)
		if err != nil {
			logrus.Errorf("unpackMessage error: %v", err)
		} else {
			if remoteAddr != nil {
				msg.RemoteAddr = remoteAddr
			}
			readedMessage <- &msg
		}
		i += global.HeaderLength + packageLength - 1
	}

	if i == uint32(length) {
		return make([]byte, 0)
	}
	return source[i:]
}
