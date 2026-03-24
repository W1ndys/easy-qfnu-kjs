package v1

import (
	"net/http"

	"github.com/W1ndys/easy-qfnu-empty-classrooms/internal/model"
	"github.com/W1ndys/easy-qfnu-empty-classrooms/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	classroomService *service.ClassroomService
	statsService     *service.StatsService
}

func NewHandler(cs *service.ClassroomService, ss *service.StatsService) *Handler {
	return &Handler{classroomService: cs, statsService: ss}
}

// GetStatus 返回系统状态，包括是否在教学周历内
func (h *Handler) GetStatus(c *gin.Context) {
	cal := service.GetCalendarService()
	if cal == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":                "日历服务未初始化",
			"in_teaching_calendar": false,
			"current_week":         0,
			"current_term":         "",
		})
		return
	}

	inCalendar := cal.IsInTeachingCalendar()
	c.JSON(http.StatusOK, gin.H{
		"in_teaching_calendar": inCalendar,
		"current_week":         cal.GetBaseWeek(),
		"current_term":         cal.GetCurrentYearStr(),
		"has_permission":       cal.HasPermission(),
	})
}

func (h *Handler) QueryClassrooms(c *gin.Context) {
	var req model.QueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数格式错误"})
		return
	}

	// 简单的校验
	if req.BuildingName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请输入教学楼名称"})
		return
	}
	if req.StartNode == "" || req.EndNode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请选择起始和终止节次"})
		return
	}

	resp, err := h.classroomService.GetEmptyClassrooms(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 异步记录搜索关键词统计
	if h.statsService != nil {
		go h.statsService.RecordQuery(req.BuildingName)
	}

	c.JSON(http.StatusOK, resp)
}

// QueryFullDayStatus 查询全天教室状态
func (h *Handler) QueryFullDayStatus(c *gin.Context) {
	var req model.FullDayQueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数格式错误"})
		return
	}

	if req.BuildingName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请输入教学楼名称"})
		return
	}

	resp, err := h.classroomService.GetFullDayStatus(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 异步记录搜索关键词统计
	if h.statsService != nil {
		go h.statsService.RecordQuery(req.BuildingName)
	}

	c.JSON(http.StatusOK, resp)
}

// GetStats 获取查询统计数据
func (h *Handler) GetStats(c *gin.Context) {
	if h.statsService == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "统计服务未初始化"})
		return
	}

	stats, err := h.statsService.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取统计数据失败"})
		return
	}

	c.JSON(http.StatusOK, stats)
}
