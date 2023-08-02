package stdX

import "os"

var (
	Stdout *hookStd
	Stderr *hookStd
)

func InitHookStdout(callback func([]byte)) error {
	readFile, writeFile, err := os.Pipe()
	if err != nil {
		return err
	}
	oldStdout := os.Stdout
	os.Stdout = writeFile
	tmpStdout := &hookStd{
		File:            writeFile,
		systemStd:       oldStdout,
		hookStdReadFile: readFile,
		callback:        callback,
	}
	go tmpStdout.hookThread()
	Stdout = tmpStdout
	return nil
}

func InitHookStderr(callback func([]byte)) error {
	readFile, writeFile, err := os.Pipe()
	if err != nil {
		return err
	}
	oldStderr := os.Stderr
	os.Stderr = writeFile
	tmpStderr := &hookStd{
		File:            writeFile,
		systemStd:       oldStderr,
		hookStdReadFile: readFile,
		callback:        callback,
	}
	go tmpStderr.hookThread()
	Stderr = tmpStderr
	return nil
}
