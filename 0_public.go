package stdX

import "os"

var (
	_stdoutHookObj *hookStd
	_stderrHookObj *hookStd
)

func InitHookStdout(callback func([]byte)) error {
	readFile, writeFile, err := os.Pipe()
	if err != nil {
		return err
	}
	oldStdout := os.Stdout
	os.Stdout = writeFile
	tmpStdout := &hookStd{
		systemStd:       oldStdout,
		hookStdReadFile: readFile,
		callback:        callback,
	}
	go tmpStdout.hookThread()
	_stdoutHookObj = tmpStdout
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
		systemStd:       oldStderr,
		hookStdReadFile: readFile,
		callback:        callback,
	}
	go tmpStderr.hookThread()
	_stderrHookObj = tmpStderr
	return nil
}
