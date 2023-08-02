# stdX
golang stdout/stderr hook tool
in golang programming
write log to file is very common
log content write to os.Stdout/os.Stderr is easy
just use os.Stdout.Write
but os.Stdout/os.Stderr content copy and write to log file seems hard
this repository is designed to solve this problem

# example
var logFileWriter io.Writer
stdX.InitHookStdout(func(oneByte byte){
    logFileWriter.Write([]byte(oneByte))
})
