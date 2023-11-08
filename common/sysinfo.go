// Package common provides utilities for system and process information.
package common

import (
	"fmt"
	"net"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	gopsutilNet "github.com/shirou/gopsutil/net"
)

// GetCurrentCPUInfo provides current CPU information using the gopsutil package.
func GetCurrentCPUInfo() ([]cpu.InfoStat, error) {
	infoStats, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	return infoStats, nil
}

// GetCurrentCPUUsage provides current CPU timing information using the gopsutil package.
func GetCurrentCPUUsage() ([]cpu.TimesStat, error) {
	infoStats, err := cpu.Times(true)
	if err != nil {
		return nil, err
	}
	return infoStats, nil
}

// GetCurrentMemoryInfo provides current memory information using the gopsutil package.
func GetCurrentMemoryInfo() (*mem.VirtualMemoryStat, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return vmStat, nil
}

// GetCurrentDiskUsage provides current disk usage information.
func GetCurrentDiskUsage(path string) (*disk.UsageStat, error) {
	usageStat, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}
	return usageStat, nil
}

// GetDiskPartitions lists all the disk partitions.
func GetDiskPartitions(all bool) ([]disk.PartitionStat, error) {
	partitions, err := disk.Partitions(all)
	if err != nil {
		return nil, err
	}
	return partitions, nil
}

// GetSystemLoadAverage provides the system load average.
func GetSystemLoadAverage() (*load.AvgStat, error) {
	avgStat, err := load.Avg()
	if err != nil {
		return nil, err
	}
	return avgStat, nil
}

// GetHostInfo provides detailed host information.
func GetHostInfo() (*host.InfoStat, error) {
	infoStat, err := host.Info()
	if err != nil {
		return nil, err
	}
	return infoStat, nil
}

// GetNetworkInterfaces lists all the network interfaces.
func GetNetworkInterfaces() ([]gopsutilNet.InterfaceStat, error) {
	interfaces, err := gopsutilNet.Interfaces()
	if err != nil {
		return nil, err
	}
	return interfaces, nil
}

// GetNetworkIOCounters provides network I/O counters.
func GetNetworkIOCounters(pernic bool) ([]gopsutilNet.IOCountersStat, error) {
	ioCounters, err := gopsutilNet.IOCounters(pernic)
	if err != nil {
		return nil, err
	}
	return ioCounters, nil
}

// GetInternalIPv4 returns the first internal IPv4 address it finds,
// typically one that starts with "192.168". If no such address is found,
// it returns an error.
func GetInternalIPv4() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ip := ipNet.IP.String()
			if ip[:7] == "192.168" {
				return ip, nil
			}
		}
	}

	return "", fmt.Errorf("no internal IPv4 address found")
}
