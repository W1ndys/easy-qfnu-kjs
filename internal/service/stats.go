package service

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/W1ndys/easy-qfnu-empty-classrooms/internal/model"
	"github.com/W1ndys/easy-qfnu-empty-classrooms/pkg/logger"

	_ "modernc.org/sqlite"
)

// StatsService 查询统计服务
type StatsService struct {
	db *sql.DB
}

// NewStatsService 创建统计服务，打开或创建 SQLite 数据库
func NewStatsService(dbPath string) (*StatsService, error) {
	// 确保目录存在
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("创建数据目录失败: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("打开数据库失败: %w", err)
	}

	// 创建表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS query_logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			classroom TEXT NOT NULL,
			building TEXT NOT NULL,
			query_type TEXT NOT NULL,
			queried_at DATETIME DEFAULT (datetime('now', 'localtime'))
		);
		CREATE INDEX IF NOT EXISTS idx_queried_at ON query_logs(queried_at);
		CREATE INDEX IF NOT EXISTS idx_classroom ON query_logs(classroom);
	`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("创建表失败: %w", err)
	}

	logger.Info("统计服务已初始化，数据库路径: %s", dbPath)
	return &StatsService{db: db}, nil
}

// RecordQuery 记录一次查询（异步调用，记录被查询的教室）
func (s *StatsService) RecordQuery(building string, classrooms []string, queryType string) {
	if len(classrooms) == 0 {
		// 即使没有查到教室，也记录一次查询（用 building 作为标记）
		_, err := s.db.Exec(
			"INSERT INTO query_logs (classroom, building, query_type) VALUES (?, ?, ?)",
			building, building, queryType,
		)
		if err != nil {
			logger.Warn("记录查询统计失败: %v", err)
		}
		return
	}

	tx, err := s.db.Begin()
	if err != nil {
		logger.Warn("开始事务失败: %v", err)
		return
	}

	stmt, err := tx.Prepare("INSERT INTO query_logs (classroom, building, query_type) VALUES (?, ?, ?)")
	if err != nil {
		tx.Rollback()
		logger.Warn("准备语句失败: %v", err)
		return
	}
	defer stmt.Close()

	for _, classroom := range classrooms {
		if _, err := stmt.Exec(classroom, building, queryType); err != nil {
			tx.Rollback()
			logger.Warn("插入记录失败: %v", err)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		logger.Warn("提交事务失败: %v", err)
	}
}

// GetStats 获取统计数据
func (s *StatsService) GetStats() (*model.StatsResponse, error) {
	now := time.Now()
	todayStart := now.Format("2006-01-02") + " 00:00:00"

	// 本周一
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	mondayDate := now.AddDate(0, 0, -(weekday - 1))
	weekStart := mondayDate.Format("2006-01-02") + " 00:00:00"

	// 本月1号
	monthStart := now.Format("2006-01") + "-01 00:00:00"

	resp := &model.StatsResponse{}

	// 今日查询次数
	s.db.QueryRow("SELECT COUNT(*) FROM query_logs WHERE queried_at >= ?", todayStart).Scan(&resp.TodayCount)

	// 本周查询次数
	s.db.QueryRow("SELECT COUNT(*) FROM query_logs WHERE queried_at >= ?", weekStart).Scan(&resp.WeekCount)

	// 本月查询次数
	s.db.QueryRow("SELECT COUNT(*) FROM query_logs WHERE queried_at >= ?", monthStart).Scan(&resp.MonthCount)

	// 今日最热教室
	s.db.QueryRow(
		"SELECT classroom FROM query_logs WHERE queried_at >= ? GROUP BY classroom ORDER BY COUNT(*) DESC LIMIT 1",
		todayStart,
	).Scan(&resp.TodayTop)

	// 本周最热教室
	s.db.QueryRow(
		"SELECT classroom FROM query_logs WHERE queried_at >= ? GROUP BY classroom ORDER BY COUNT(*) DESC LIMIT 1",
		weekStart,
	).Scan(&resp.WeekTop)

	// 本月最热教室
	s.db.QueryRow(
		"SELECT classroom FROM query_logs WHERE queried_at >= ? GROUP BY classroom ORDER BY COUNT(*) DESC LIMIT 1",
		monthStart,
	).Scan(&resp.MonthTop)

	return resp, nil
}

// Close 关闭数据库连接
func (s *StatsService) Close() error {
	return s.db.Close()
}
