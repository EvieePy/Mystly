package routes

import (
	"mystly/internal/core"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"

	"github.com/gin-gonic/gin"
	"path/filepath"
	"runtime"
	"os"
)

type UsageStatistics struct {
	CpuPer     float64 `json:"cpu_per"`
	MemTotal   uint64  `json:"mem_total"`
	MemUsed    uint64  `json:"mem_used"`
	MemFree    uint64  `json:"mem_free"`
	MemUsedPer float64 `json:"mem_used_per"`
	PhysCore   int     `json:"physical_cores"`
	LogCore    int     `json:"logical_cores"`
	DiskTotal  uint64  `json:"disk_total"`
	DiskFree   uint64  `json:"disk_free"`
	DiskUsed   uint64  `json:"disk_used"`
	DiskPer    float64 `json:"disk_per"`
}

type OverviewHandler struct {
	Server *core.Server
	Path   string
}

func NewOverviewHandler(path string, server *core.Server) IHandler {
	return &OverviewHandler{Path: path, Server: server}
}

func (h *OverviewHandler) RegisterRoutes() {
	group := h.Server.Router.Group(h.Path)
	// auth := core.NewAuthMiddleware(h.Server)

	// Unauthed...

	// Authed...
	// group.Use(auth.Handler())
	group.GET("/overviewstats", h.apiSystemUsage())
}

func (h *OverviewHandler) apiSystemUsage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var hasErr bool
		stats := UsageStatistics{}

		cpuPer, err := cpu.Percent(0, false)
		if err != nil {
			hasErr = true
		}

		memStat, err := mem.VirtualMemory()
		if err != nil {
			hasErr = true
		}

		logCore, err := cpu.Counts(true)
		if err != nil {
			hasErr = true
		}

		physCore, err := cpu.Counts(false)
		if err != nil {
			hasErr = true
		}
		
		root := "/"
		system := runtime.GOOS

		if system == "windows" {
			cwd, err := os.Getwd()
			if err != nil {
				hasErr = true
			} else {
				root = filepath.VolumeName(cwd)
			}
		}

		diskStats, err := disk.Usage(root)
		if err != nil {
			hasErr = true
		}

		if hasErr {
			ctx.AbortWithStatus(500)
			return
		}

		stats.CpuPer = cpuPer[0]
		stats.MemTotal = memStat.Total
		stats.MemUsed = memStat.Used
		stats.MemFree = memStat.Free
		stats.MemUsedPer = memStat.UsedPercent
		stats.PhysCore = physCore
		stats.LogCore = logCore
		stats.DiskFree = diskStats.Free
		stats.DiskTotal = diskStats.Total
		stats.DiskUsed = diskStats.Used
		stats.DiskPer = diskStats.UsedPercent

		ctx.JSON(200, stats)
	}
}
