package service

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/W1ndys/easy-qfnu-empty-classrooms/internal/model"
	"github.com/W1ndys/easy-qfnu-empty-classrooms/pkg/logger"

	_ "modernc.org/sqlite"
)

// StatsService 查询统计服务
type StatsService struct {
	db *sql.DB
	mu sync.Mutex // 保护 SQLite 串行写入
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

	// 关键配置：启用 WAL 模式，允许读写并发
	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		db.Close()
		return nil, fmt.Errorf("设置 WAL 模式失败: %w", err)
	}

	// 设置 busy_timeout，当数据库被锁时等待 5 秒而非立即报错
	if _, err := db.Exec("PRAGMA busy_timeout=5000"); err != nil {
		db.Close()
		return nil, fmt.Errorf("设置 busy_timeout 失败: %w", err)
	}

	// 限制连接池：SQLite 是文件级数据库，多连接写入会导致锁冲突
	// 设置最大打开连接数为 1，确保写入串行化
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)

	// 执行迁移
	if err := migrateSchema(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("数据库迁移失败: %w", err)
	}

	logger.Info("统计服务已初始化，数据库路径: %s", dbPath)
	return &StatsService{db: db}, nil
}

// migrateSchema 检测并迁移表结构
func migrateSchema(db *sql.DB) error {
	// 检测表是否存在
	var tableExists int
	err := db.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='query_logs'").Scan(&tableExists)
	if err != nil {
		return fmt.Errorf("检测表存在失败: %w", err)
	}

	if tableExists == 0 {
		// 表不存在，直接创建新表
		return createNewTable(db)
	}

	// 检测表是否包含 result_count 列（v3 新结构标志）
	var hasResultCount int
	err = db.QueryRow("SELECT COUNT(*) FROM pragma_table_info('query_logs') WHERE name='result_count'").Scan(&hasResultCount)
	if err != nil {
		return fmt.Errorf("检测表结构失败: %w", err)
	}

	if hasResultCount == 0 {
		// 旧表结构（无论是 v1 classroom 版还是 v2 keyword-only 版），都需要迁移
		logger.Warn("检测到旧版 query_logs 表结构，开始迁移到 v3...")
		return migrateFromOldSchema(db)
	}

	// 新表结构，检查索引
	return createIndexesIfNotExist(db)
}

// createNewTable 创建新的表结构
func createNewTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS query_logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			keyword TEXT NOT NULL,
			date_offset INTEGER NOT NULL DEFAULT 0,
			start_node TEXT NOT NULL DEFAULT '',
			end_node TEXT NOT NULL DEFAULT '',
			result_count INTEGER NOT NULL DEFAULT 0,
			queried_at DATETIME DEFAULT (datetime('now', 'localtime'))
		);
		CREATE INDEX IF NOT EXISTS idx_queried_at ON query_logs(queried_at);
		CREATE INDEX IF NOT EXISTS idx_keyword ON query_logs(keyword);
		CREATE INDEX IF NOT EXISTS idx_combo ON query_logs(keyword, date_offset, start_node, end_node, result_count);
	`)
	if err != nil {
		return fmt.Errorf("创建表失败: %w", err)
	}
	return nil
}

// migrateFromOldSchema 从旧结构迁移到新结构（直接删除旧数据重建）
func migrateFromOldSchema(db *sql.DB) error {
	// 删除旧表
	_, err := db.Exec("DROP TABLE query_logs")
	if err != nil {
		return fmt.Errorf("删除旧表失败: %w", err)
	}

	// 创建新表
	err = createNewTable(db)
	if err != nil {
		return err
	}

	logger.Warn("旧版 query_logs 表已删除，已创建 v3 新表结构")
	return nil
}

// createIndexesIfNotExist 创建索引（如果不存在）
func createIndexesIfNotExist(db *sql.DB) error {
	_, err := db.Exec("CREATE INDEX IF NOT EXISTS idx_queried_at ON query_logs(queried_at)")
	if err != nil {
		return fmt.Errorf("创建 queried_at 索引失败: %w", err)
	}
	_, err = db.Exec("CREATE INDEX IF NOT EXISTS idx_keyword ON query_logs(keyword)")
	if err != nil {
		return fmt.Errorf("创建 keyword 索引失败: %w", err)
	}
	_, err = db.Exec("CREATE INDEX IF NOT EXISTS idx_combo ON query_logs(keyword, date_offset, start_node, end_node, result_count)")
	if err != nil {
		return fmt.Errorf("创建 combo 索引失败: %w", err)
	}
	return nil
}

// RecordQuery 记录一次搜索查询（异步调用），包含完整的搜索参数和结果数量
func (s *StatsService) RecordQuery(record model.QueryRecord) {
	if record.Keyword == "" {
		return
	}

	// 使用互斥锁确保写入串行化，防止 SQLite 并发写入冲突
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.db.Exec(
		"INSERT INTO query_logs (keyword, date_offset, start_node, end_node, result_count) VALUES (?, ?, ?, ?, ?)",
		record.Keyword, record.DateOffset, record.StartNode, record.EndNode, record.ResultCount,
	)
	if err != nil {
		logger.Warn("记录搜索查询失败: %v", err)
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
	if err := s.db.QueryRow("SELECT COUNT(*) FROM query_logs WHERE queried_at >= ?", todayStart).Scan(&resp.TodayCount); err != nil {
		logger.Error("查询今日统计失败: %v", err)
		return nil, fmt.Errorf("查询今日统计失败: %w", err)
	}

	// 本周查询次数
	if err := s.db.QueryRow("SELECT COUNT(*) FROM query_logs WHERE queried_at >= ?", weekStart).Scan(&resp.WeekCount); err != nil {
		logger.Error("查询本周统计失败: %v", err)
		return nil, fmt.Errorf("查询本周统计失败: %w", err)
	}

	// 本月查询次数
	if err := s.db.QueryRow("SELECT COUNT(*) FROM query_logs WHERE queried_at >= ?", monthStart).Scan(&resp.MonthCount); err != nil {
		logger.Error("查询本月统计失败: %v", err)
		return nil, fmt.Errorf("查询本月统计失败: %w", err)
	}

	// 今日最热搜索关键词
	if err := s.db.QueryRow(
		"SELECT keyword FROM query_logs WHERE queried_at >= ? GROUP BY keyword ORDER BY COUNT(*) DESC LIMIT 1",
		todayStart,
	).Scan(&resp.TodayTop); err != nil && err != sql.ErrNoRows {
		logger.Error("查询今日最热关键词失败: %v", err)
		return nil, fmt.Errorf("查询今日最热关键词失败: %w", err)
	}

	// 本周最热搜索关键词
	if err := s.db.QueryRow(
		"SELECT keyword FROM query_logs WHERE queried_at >= ? GROUP BY keyword ORDER BY COUNT(*) DESC LIMIT 1",
		weekStart,
	).Scan(&resp.WeekTop); err != nil && err != sql.ErrNoRows {
		logger.Error("查询本周最热关键词失败: %v", err)
		return nil, fmt.Errorf("查询本周最热关键词失败: %w", err)
	}

	// 本月最热搜索关键词
	if err := s.db.QueryRow(
		"SELECT keyword FROM query_logs WHERE queried_at >= ? GROUP BY keyword ORDER BY COUNT(*) DESC LIMIT 1",
		monthStart,
	).Scan(&resp.MonthTop); err != nil && err != sql.ErrNoRows {
		logger.Error("查询本月最热关键词失败: %v", err)
		return nil, fmt.Errorf("查询本月最热关键词失败: %w", err)
	}

	return resp, nil
}

// GetTopQueries 获取搜索排行前 N 的查询组合（仅统计结果非空的记录）
func (s *StatsService) GetTopQueries(limit int) ([]model.TopQueryItem, error) {
	if limit <= 0 {
		limit = 5
	}

	rows, err := s.db.Query(
		`SELECT keyword, date_offset, start_node, end_node, COUNT(*) AS cnt
		 FROM query_logs
		 WHERE result_count > 0
		 GROUP BY keyword, date_offset, start_node, end_node
		 ORDER BY cnt DESC
		 LIMIT ?`,
		limit,
	)
	if err != nil {
		return nil, fmt.Errorf("查询热门搜索组合失败: %w", err)
	}
	defer rows.Close()

	var queries []model.TopQueryItem
	for rows.Next() {
		var item model.TopQueryItem
		if err := rows.Scan(&item.Building, &item.DateOffset, &item.StartNode, &item.EndNode, &item.Count); err != nil {
			return nil, fmt.Errorf("扫描热门搜索组合数据失败: %w", err)
		}
		queries = append(queries, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历热门搜索组合结果失败: %w", err)
	}

	return queries, nil
}

// Close 关闭数据库连接
func (s *StatsService) Close() error {
	return s.db.Close()
}
