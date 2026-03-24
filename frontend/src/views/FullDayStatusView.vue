<script setup>
import { reactive, ref } from 'vue'
import { getErrorMessage, queryFullDayStatus } from '@/api'
import { useSystemStatus } from '@/composables/useSystemStatus'
import { useSearchHistory } from '@/composables/useSearchHistory'
import AppFooter from '@/components/AppFooter.vue'
import AppHeader from '@/components/AppHeader.vue'
import DateSelector from '@/components/DateSelector.vue'
import EmptyState from '@/components/EmptyState.vue'
import LoadingSpinner from '@/components/LoadingSpinner.vue'
import QRCodeCard from '@/components/QRCodeCard.vue'
import StatusWarning from '@/components/StatusWarning.vue'

const { statusLoading, inTeachingCalendar, hasPermission } = useSystemStatus()
const { history, addToHistory, clearHistory } = useSearchHistory()

const loading = ref(false)
const hasSearched = ref(false)
const resultData = ref(null)
const showHistory = ref(false)
const inputFocused = ref(false)

const form = reactive({
  building: '',
  offset: 0,
})

const legendItems = [
  { id: 1, emoji: '🔴', name: '正常上课' },
  { id: 2, emoji: '🟠', name: '借用' },
  { id: 3, emoji: '🔒', name: '锁定' },
  { id: 4, emoji: '🟣', name: '考试' },
  { id: 5, emoji: '🟢', name: '空闲' },
  { id: 6, emoji: '🔵', name: '固定调课' },
  { id: 7, emoji: '💠', name: '临时调课' },
  { id: 8, emoji: '🌳', name: '完全空闲' },
  { id: 9, emoji: '🌸', name: '跨模式' },
]

function getEmoji(statusId) {
  return legendItems.find((item) => item.id === statusId)?.emoji || ''
}

function onInputFocus() {
  inputFocused.value = true
  showHistory.value = true
}

function onInputBlur() {
  inputFocused.value = false
  setTimeout(() => {
    showHistory.value = false
  }, 200)
}

function onInputChange() {
  showHistory.value = false
}

function selectHistoryItem(keyword) {
  form.building = keyword
  showHistory.value = false
}

async function search() {
  if (!form.building.trim()) {
    return
  }

  loading.value = true
  hasSearched.value = false
  resultData.value = null
  showHistory.value = false

  try {
    const data = await queryFullDayStatus({
      building: form.building,
      date_offset: form.offset,
    })

    resultData.value = data
    hasSearched.value = true
    addToHistory(form.building)
  } catch (error) {
    console.error(error)
    alert(getErrorMessage(error, '查询失败'))
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen bg-gray-50 font-sans antialiased pb-10">
    <AppHeader title="教室全天状态" showBack />

    <main class="px-4 py-4 max-w-5xl mx-auto space-y-4">
      <StatusWarning
        v-if="!hasPermission && !statusLoading"
        type="error"
        title="权限不足"
        message="当前账号无权限访问教务系统查询接口，请检查账号状态。"
      />

      <StatusWarning
        v-if="!inTeachingCalendar && !statusLoading"
        type="warning"
        title="提示"
        message="当前日期不在教学周历内，查询结果可能不准确。"
      />

      <LoadingSpinner v-if="statusLoading" text="正在检查系统状态..." />

      <div v-else class="bg-white rounded-2xl p-4 shadow-sm space-y-4">
        <div>
          <label class="block text-sm font-medium text-gray-500 mb-1.5 ml-1">教学楼</label>
          <div class="relative">
            <input
              v-model="form.building"
              type="text"
              class="w-full bg-gray-100 rounded-xl py-3 px-4 text-[15px] focus:outline-none focus:ring-2 focus:ring-primary/20"
              placeholder="例如：老文史楼"
              @focus="onInputFocus"
              @blur="onInputBlur"
              @input="onInputChange"
            />
            <div
              v-if="inputFocused && showHistory && history.length > 0"
              class="absolute z-10 w-full mt-1 bg-white rounded-xl shadow-lg border border-gray-100 overflow-hidden"
            >
              <div class="flex items-center justify-between px-3 py-2 border-b border-gray-100">
                <span class="text-xs text-gray-400">搜索历史</span>
                <button type="button" class="text-xs text-gray-400 hover:text-gray-600" @click="clearHistory">清除</button>
              </div>
              <div class="max-h-48 overflow-y-auto">
                <button
                  v-for="(item, index) in history"
                  :key="index"
                  type="button"
                  class="w-full px-3 py-2.5 text-left text-sm hover:bg-gray-50 flex items-center"
                  @mousedown.prevent="selectHistoryItem(item)"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-gray-400 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  {{ item }}
                </button>
              </div>
            </div>
          </div>
        </div>

        <DateSelector v-model="form.offset" />

        <button
          type="button"
          :disabled="loading || !form.building.trim()"
          class="w-full bg-primary text-white font-semibold py-3.5 rounded-xl disabled:opacity-70 flex items-center justify-center"
          @click="search"
        >
          <span v-if="!loading">查询全天状态</span>
          <span v-else class="flex items-center">
            <svg class="animate-spin -ml-1 mr-2 h-5 w-5" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
            </svg>
            查询中...
          </span>
        </button>
      </div>

      <div v-if="hasSearched" class="bg-white rounded-2xl p-4 shadow-sm">
        <h3 class="text-sm font-medium text-gray-700 mb-3">状态图例</h3>
        <div class="grid grid-cols-3 gap-2 text-xs">
          <div v-for="item in legendItems" :key="item.id" class="flex items-center space-x-2">
            <span class="text-base">{{ item.emoji }}</span>
            <span class="text-gray-600">{{ item.name }}</span>
          </div>
        </div>
      </div>

      <div v-if="hasSearched && resultData" class="bg-white rounded-2xl shadow-sm overflow-hidden">
        <div class="p-4 border-b border-gray-100">
          <p class="text-sm text-gray-600">
            {{ resultData.building }} · {{ resultData.date }} · {{ resultData.current_term }} 学期 · 第{{ resultData.week }}周 · 星期{{ resultData.day_of_week }}
          </p>
        </div>

        <div class="table-container">
          <table class="w-full text-sm">
            <thead>
              <tr class="bg-gray-50">
                <th class="sticky-col px-3 py-3 text-left font-medium text-gray-700 min-w-[100px]">教室</th>
                <th
                  v-for="node in resultData.node_list"
                  :key="node.node_index"
                  class="px-1 py-3 text-center font-medium text-gray-700 min-w-[40px]"
                >
                  {{ node.node_name }}
                </th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="room in resultData.classrooms" :key="room.room_name" class="border-t border-gray-100">
                <td class="sticky-col px-3 py-3 font-medium text-gray-800 bg-white">{{ room.room_name }}</td>
                <td v-for="(status, idx) in room.status" :key="`${room.room_name}-${idx}`" class="px-1 py-2 text-center text-base">
                  {{ getEmoji(status.status_id) }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <EmptyState v-if="hasSearched && !loading && !resultData" text="暂无数据" />

      <QRCodeCard />
    </main>

    <AppFooter />
  </div>
</template>

<style scoped>
.table-container {
  overflow-x: auto;
  -webkit-overflow-scrolling: touch;
}

.sticky-col {
  position: sticky;
  left: 0;
  z-index: 10;
  box-shadow: 2px 0 4px rgba(0, 0, 0, 0.05);
}
</style>
