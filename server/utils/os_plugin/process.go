package os_plugin

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const query = "ps -ef | grep rp1 | grep -v grep | grep -v rp2 | awk '{print $2}'"

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Server struct {
	Os   Os   `json:"os"`
	Cpu  Cpu  `json:"cpu"`
	Ram  Ram  `json:"ram"`
	Disk Disk `json:"disk"`
}

type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"numCpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"goVersion"`
	NumGoroutine int    `json:"numGoroutine"`
}

type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

type Ram struct {
	UsedMB      int `json:"usedMb"`
	TotalMB     int `json:"totalMb"`
	UsedPercent int `json:"usedPercent"`
}

type Disk struct {
	UsedMB      int `json:"usedMb"`
	UsedGB      int `json:"usedGb"`
	TotalMB     int `json:"totalMb"`
	TotalGB     int `json:"totalGb"`
	UsedPercent int `json:"usedPercent"`
}

func InitOS() (o Os) {
	o.GOOS = runtime.GOOS
	o.NumCPU = runtime.NumCPU()
	o.Compiler = runtime.Compiler
	o.GoVersion = runtime.Version()
	o.NumGoroutine = runtime.NumGoroutine()
	return o
}

func InitCPU() (c Cpu, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return c, err
	} else {
		c.Cpus = cpus
	}
	return c, nil
}

func InitRAM() (r Ram, err error) {
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		r.UsedMB = int(u.Used) / MB
		r.TotalMB = int(u.Total) / MB
		r.UsedPercent = int(u.UsedPercent)
	}
	return r, nil
}

func InitDisk() (d Disk, err error) {
	if u, err := disk.Usage("/"); err != nil {
		return d, err
	} else {
		d.UsedMB = int(u.Used) / MB
		d.UsedGB = int(u.Used) / GB
		d.TotalMB = int(u.Total) / MB
		d.TotalGB = int(u.Total) / GB
		d.UsedPercent = int(u.UsedPercent)
	}
	return d, nil
}

func StopProcess(name string) {
	text := strings.Replace(query, "rp1", name, -1)
	text = strings.Replace(text, "rp2", strconv.Itoa(os.Getpid()), -1)
	os.WriteFile("temp.sh", []byte(text), 0777)
	defer os.Remove("temp.sh")

	out, err := exec.Command("bash", "temp.sh").Output()
	if err != nil {
		return
	}
	exec.Command("kill", "-2", strings.TrimSpace(string(out))).Run()
}
