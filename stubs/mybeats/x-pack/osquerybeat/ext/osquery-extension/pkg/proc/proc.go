package proc
import "io/fs"

type ProcStat struct {
	PID       string
	Name      string
	StartTime string
	Group     string
	Parent    string
}

func ListFS(root fs.FS) ([]ProcStat, error) { return nil, nil }
func ReadStatFS(root fs.FS, p ProcStat) (ProcStat, error) { return ProcStat{}, nil }
func ReadCmdLineFS(root fs.FS, p ProcStat) (string, error) { return "", nil }
func ReadUptimeFS(root fs.FS) (float64, error) { return 0, nil }