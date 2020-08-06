package coder

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

const msgHeader = "tcp send"

/**
 *  编码
 *  定义消息的格式： msgHeader + contentLength + content
 *  conn 本身实现了 io.Writer 接口
 */
func Encode(conn io.Writer, content string) (err error) {
	//写入消息头
	if err = binary.Write(conn, binary.BigEndian, []byte(msgHeader)); err != nil {
		fmt.Printf("Fail to write magHeader to conn,err:[%v]", err)
	}
	//写入消息体长度
	contentLength := int32(len([]byte(content)))
	if err = binary.Write(conn, binary.BigEndian, contentLength); err != nil {
		fmt.Printf("Fail to write contentLength to conn,err:[%v]", err)
	}
	// 写入消息
	if err = binary.Write(conn, binary.BigEndian, []byte(content)); err != nil {
		fmt.Printf("Fail to write content to conn,err:[%v]", err)
	}
	return err
}

/**
 * 	解码：
 */
func Decode(reader io.Reader) (bytes []byte, err error) {
	// 先把消息头读出来
	headerBuf := make([]byte, len(msgHeader))
	if _, err = io.ReadFull(reader, headerBuf); err != nil {
		fmt.Printf("Fail ro read header from conn error:[%v]", err)
		return nil, err
	}

	//检验消息头
	if string(headerBuf) != msgHeader {
		err = errors.New("msgHeader error")
		return nil, err
	}

	//读取实际肉容的长度
	lengthBuf := make([]byte, 4)
	if _, err = io.ReadFull(reader, lengthBuf); err != nil {
		return nil, err
	}
	contentLength := binary.BigEndian.Uint32(lengthBuf)
	contentBuf := make([]byte, contentLength)

	//读出消息体
	if _, err := io.ReadFull(reader, contentBuf); err != nil {
		return nil, err
	}

	return contentBuf, err
}
