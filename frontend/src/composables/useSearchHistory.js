import { ref } from 'vue'

const STORAGE_KEY = 'search_history'
const MAX_STORAGE = 10
const MAX_DISPLAY = 5

export function useSearchHistory() {
  const history = ref([])

  function loadHistory() {
    try {
      const stored = localStorage.getItem(STORAGE_KEY)
      if (stored) {
        const parsed = JSON.parse(stored)
        history.value = parsed.slice(0, MAX_DISPLAY)
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

  function addToHistory(keyword) {
    const trimmed = keyword.trim()
    if (!trimmed) return

    const items = history.value.filter((item) => item !== trimmed)
    items.unshift(trimmed)

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
