import { ref } from 'vue'

const STORAGE_KEY = 'search_history'
const MAX_STORAGE = 10
const MAX_DISPLAY = 5

const DATE_LABELS = ['今天', '明天', '后天']

function formatDateOffset(offset = 0) {
  return DATE_LABELS[offset] ?? `${offset}天后`
}

function formatNode(node) {
  return String(parseInt(node, 10))
}

function normalizeItem(item) {
  if (typeof item === 'string') {
    return {
      building: item,
      label: item,
    }
  }

  const normalized = {
    building: item.building || '',
    offset: Number(item.offset ?? item.date_offset ?? 0),
    start: item.start || item.start_node,
    end: item.end || item.end_node,
  }

  const parts = [normalized.building, formatDateOffset(normalized.offset)]
  if (normalized.start && normalized.end) {
    parts.push(`${formatNode(normalized.start)}-${formatNode(normalized.end)}节`)
  }

  return {
    ...normalized,
    label: parts.filter(Boolean).join(' | '),
  }
}

function getItemKey(item) {
  return [item.building, item.offset ?? '', item.start ?? '', item.end ?? ''].join('|')
}

export function useSearchHistory() {
  const history = ref([])

  function loadHistory() {
    try {
      const stored = localStorage.getItem(STORAGE_KEY)
      if (stored) {
        const parsed = JSON.parse(stored)
        history.value = parsed.map(normalizeItem).slice(0, MAX_DISPLAY)
      }
    } catch {
      history.value = []
    }
  }

  function saveHistory(items) {
    try {
      localStorage.setItem(STORAGE_KEY, JSON.stringify(items))
    } catch {
      // localStorage 不可用时静默失败
    }
  }

  function addToHistory(keyword, options = {}) {
    const baseItem = typeof keyword === 'string'
      ? { building: keyword.trim(), ...options }
      : keyword
    const normalized = normalizeItem(baseItem)
    if (!normalized.building) return

    const items = history.value.filter((item) => getItemKey(item) !== getItemKey(normalized))
    items.unshift(normalized)

    if (items.length > MAX_STORAGE) {
      items.pop()
    }

    history.value = items.slice(0, MAX_DISPLAY)
    saveHistory(items)
  }

  function clearHistory() {
    history.value = []
    try {
      localStorage.removeItem(STORAGE_KEY)
    } catch {
      // localStorage 不可用时静默失败
    }
  }

  loadHistory()

  return {
    history,
    addToHistory,
    clearHistory,
  }
}
