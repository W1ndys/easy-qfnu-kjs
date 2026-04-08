/**
 * 公告配置文件
 *
 * 管理方式：直接编辑此文件即可发布/更新/下架公告。
 *
 * 字段说明：
 *   id        - 唯一标识，用于已读缓存判断（更新内容时请更改 id，这样用户会重新看到）
 *   date      - 发布日期，格式 YYYY-MM-DD
 *   title     - 公告标题
 *   content   - 公告正文，支持纯文本
 *   important - 可选，设为 true 时该条公告会以醒目样式高亮展示
 */
export const announcements = [
  {
    id: 'announcement-20260408-001',
    date: '2026-04-08',
    title: '需要可联系',
    content:
      '校园跑、U校园、Welearn、智慧树、学习通，病历，假条等生成可联系 QQ群 1087476180',
    important: false,
  },
  {
    id: 'announcement-20260407-002',
    date: '2026-04-07',
    title: '关于教学楼名称的说明',
    content:
      '本系统严格遵循教务系统内设定的教学楼名称，不支持简称，例如综合楼（错误）、综合教学楼（正确），以及日照校区的教学楼全部是大写字母J开头，例如JA101。',
    important: true,
  },
  {
    id: 'announcement-20260407-001',
    date: '2026-04-07',
    title: '系统上线通知',
    content:
      'QFNU 教室查询系统正式上线，支持空教室查询和教室全天占用状态查看，欢迎使用！',
    important: false,
  },
  // 发布新公告时，在此处添加新条目即可，最新的放在最前面
  // {
  //   id: 'announcement-20260408-001',
  //   date: '2026-04-08',
  //   title: '数据更新说明',
  //   content: '本系统数据每日自动更新，如遇数据延迟请耐心等待。',
  //   important: false,
  // },
]
