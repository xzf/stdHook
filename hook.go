package stdX

import (
	"bufio"
	"os"
)

type hookStd struct {
	*os.File
	systemStd       *os.File
	hookStdReadFile *os.File
	callback        func([]byte)
}

func (std *hookStd) hookThread() {
	reader := bufio.NewReader(std.hookStdReadFile)
	for {
		//be careful
		//ReadLine doc mention one line too long ReadLine will do some extra job
		//but i never test this situation
		//
		//other situation fmt.Print MultiLine
		//expect output e.g.:
		//{line 1}
		//{line 2}
		//{line 3}
		//it‘s normal that insert some line in one fmt.Print
		//real output:
		//{line 1}
		//{other log line}
		//{line 2}
		//{line 3}
		line, _, err := reader.ReadLine()
		if err != nil {
			//get err should panic
			panic("[tbzai123s2] " + err.Error())
		}
		line = append(line, '\n')
		//ignore systemStd.Write error
		std.systemStd.Write(line)
		if std.callback != nil {
			//if callback be panic will exit program
			std.callback(line)
		}
	}
}
