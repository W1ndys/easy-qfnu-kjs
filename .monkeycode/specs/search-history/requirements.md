# Requirements Document

## Introduction

为搜索框增加搜索历史记录缓存功能，提升用户搜索体验。用户查询教学楼时，系统自动记录搜索关键词，用户下次搜索时可以快速选择历史记录。

## Glossary

- **搜索历史缓存**: 本地存储的用户历史搜索关键词列表
- **localStorage**: 浏览器本地存储 API，用于持久化搜索历史

## Requirements

### Requirement 1: 历史记录存储

**User Story:** 作为用户，我希望搜索关键词被自动保存，这样下次搜索时可以快速选择。

#### Acceptance Criteria

1. WHEN 用户执行教学楼搜索操作，system SHALL 将搜索关键词存储到 localStorage
2. WHEN 用户执行搜索时，system SHALL 将搜索关键词添加到历史记录列表头部
3. IF 历史记录数量超过 10 条，system SHALL 删除最旧的记录

### Requirement 2: 历史记录展示

**User Story:** 作为用户，我希望在搜索时能看到历史记录，这样可以快速选择之前的搜索词。

#### Acceptance Criteria

1. WHEN 用户聚焦教学楼输入框，system SHALL 显示最近 5 条搜索历史记录
2. WHEN 用户点击历史记录项，system SHALL 将该关键词填充到搜索框
3. WHEN 用户开始输入且输入内容与历史记录不匹配，system SHALL 隐藏历史记录列表

### Requirement 3: 历史记录管理

**User Story:** 作为用户，我希望能够清除历史记录，这样可以保护隐私。

#### Acceptance Criteria

1. IF 用户点击历史记录列表的清除按钮，system SHALL 清空所有搜索历史
2. THE system SHALL 在页面加载时自动恢复历史记录到内存

## Data Model

### SearchHistoryItem

```typescript
interface SearchHistoryItem {
  keyword: string      // 搜索关键词
  timestamp: number    // 搜索时间戳
}
```

### Storage Structure

localStorage key: `search_history`
Max items: 10
Display items: 5
