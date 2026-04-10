import { ref } from 'vue'

const STORAGE_KEY = 'building_alias_reminder_date'
const TARGET_ALIAS = '综合楼'
const TARGET_BUILDING = '综合教学楼'

function getTodayKey() {
  return new Date().toISOString().slice(0, 10)
}

function hasRemindedToday() {
  try {
    return localStorage.getItem(STORAGE_KEY) === getTodayKey()
  } catch {
    return false
  }
}

function markRemindedToday() {
  try {
    localStorage.setItem(STORAGE_KEY, getTodayKey())
  } catch {
    // localStorage 不可用时静默失败
  }
}

export function useBuildingAliasReminder() {
  const dialogOpen = ref(false)
  let resolver = null

  function openDialog() {
    dialogOpen.value = true

    return new Promise((resolve) => {
      resolver = resolve
    })
  }

  function closeDialog(shouldUseFullName) {
    dialogOpen.value = false
    resolver?.(shouldUseFullName)
    resolver = null
  }

  async function normalizeBuildingName(building) {
    const trimmed = building.trim()

    if (trimmed !== TARGET_ALIAS || hasRemindedToday()) {
      return trimmed
    }

    markRemindedToday()
    const shouldUseFullName = await openDialog()

    return shouldUseFullName ? TARGET_BUILDING : trimmed
  }

  return {
    dialogOpen,
    normalizeBuildingName,
    confirmReminder: () => closeDialog(true),
    cancelReminder: () => closeDialog(false),
  }
}
