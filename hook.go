package stdX

import (
	"bufio"
	"os"
)

type hookStd struct {
	systemStd       *os.File
	hookStdReadFile *os.File
	callback        func(byte)
}

func (std *hookStd) Write(data []byte) (int, error) {
	return std.systemStd.Write(data)
}

func (std *hookStd) hookThread() {
	reader := bufio.NewReader(std.hookStdReadFile)
	for {
		oneByte, err := reader.ReadByte()
		if err != nil {
			//get err should panic
			panic("[tbzai123s2] " + err.Error())
		}
		_, err = std.systemStd.Write([]byte{oneByte})
		if err != nil {
			//get err should panic
			panic("[6f0zsdxgqu] " + err.Error())
		}
		if std.callback != nil {
			std.callback(oneByte)
		}
	}
}
