import { onMounted, ref } from 'vue'
import { getTopBuildings } from '@/api'

const DATE_LABELS = ['今天', '明天', '后天']

/**
 * 将日期偏移转为可读标签
 */
function formatDateOffset(offset) {
  return DATE_LABELS[offset] ?? `${offset}天后`
}

/**
 * 将节次号 "01"-"11" 转为数字显示
 */
function formatNode(node) {
  return String(parseInt(node, 10))
}

/**
 * 为一个查询组合生成按钮展示文本
 * 有节次范围时: "格物楼 | 今天 | 1-4节"
 * 无节次范围时: "格物楼 | 今天"
 */
function formatLabel(query) {
  const parts = [query.building, formatDateOffset(query.date_offset)]
  if (query.start_node && query.end_node) {
    parts.push(`${formatNode(query.start_node)}-${formatNode(query.end_node)}节`)
  }
  return parts.join(' | ')
}

/**
 * 获取搜索排行前 N 的热门查询组合，用于快捷预选按钮
 */
export function useTopBuildings() {
  const topQueries = ref([])
  const loaded = ref(false)

  async function fetchTopQueries() {
    try {
      const data = await getTopBuildings()
      topQueries.value = (data.queries || []).map((item) => ({
        ...item,
        label: formatLabel(item),
      }))
      loaded.value = true
    } catch {
      topQueries.value = []
      loaded.value = true
    }
  }

  onMounted(fetchTopQueries)

  return { topQueries, loaded }
}
