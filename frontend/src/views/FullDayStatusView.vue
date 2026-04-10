<script setup>
import { reactive, ref } from 'vue'
import { getErrorMessage, queryFullDayStatus } from '@/api'
import { useSystemStatus } from '@/composables/useSystemStatus'
import { useSearchHistory } from '@/composables/useSearchHistory'
import { useTopBuildings } from '@/composables/useTopBuildings'
import AppFooter from '@/components/AppFooter.vue'
import AppHeader from '@/components/AppHeader.vue'
import DateSelector from '@/components/DateSelector.vue'
import EmptyState from '@/components/EmptyState.vue'
import LoadingSpinner from '@/components/LoadingSpinner.vue'
import QRCodeCard from '@/components/QRCodeCard.vue'
import StatusWarning from '@/components/StatusWarning.vue'

const { statusLoading, inTeachingCalendar, hasPermission } = useSystemStatus()
const { history, addToHistory, clearHistory } = useSearchHistory()
const { topQueries } = useTopBuildings()

function selectTopQuery(query) {
  form.building = query.building
  form.offset = query.date_offset
  showHistory.value = false
  search()
}

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
  <div class="min-h-screen bg-clay-canvas text-clay-foreground antialiased pb-10 relative overflow-hidden">
    <!-- Floating Clay Blobs -->
    <div class="pointer-events-none fixed inset-0 overflow-hidden -z-10">
      <div
        class="absolute h-[45vh] w-[45vh] rounded-full blur-3xl animate-clay-float"
        style="background: rgba(136, 79, 34, 0.06); top: -8%; left: -12%;"
      ></div>
      <div
        class="absolute h-[40vh] w-[40vh] rounded-full blur-3xl animate-clay-float-delayed animation-delay-2000"
        style="background: rgba(14, 165, 233, 0.05); bottom: 5%; right: -10%;"
      ></div>
    </div>

    <AppHeader title="教室全天状态" showBack />

    <main class="px-4 py-4 max-w-5xl mx-auto space-y-5 relative z-10">
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

      <!-- Search form -->
      <div v-else class="clay-card p-5 sm:p-7 space-y-5">
        <div class="relative z-10 space-y-5">
          <div>
            <label class="block text-sm font-bold text-clay-muted mb-2 ml-1">教学楼</label>
            <div class="relative">
              <input
                v-model="form.building"
                type="text"
                class="w-full clay-input py-3.5 px-5 text-[15px]"
                placeholder="例如：老文史楼"
                @focus="onInputFocus"
                @blur="onInputBlur"
                @input="onInputChange"
              />
              <!-- Search history dropdown -->
              <div
                v-if="inputFocused && showHistory && history.length > 0"
                class="absolute z-20 w-full mt-2 rounded-[20px] overflow-hidden"
                style="
                  background: rgba(255, 255, 255, 0.95);
                  backdrop-filter: blur(24px);
                  box-shadow:
                    16px 16px 32px rgba(136, 79, 34, 0.1),
                    -10px -10px 24px rgba(255, 255, 255, 0.9),
                    inset 4px 4px 8px rgba(255, 255, 255, 0.6),
                    inset -4px -4px 8px rgba(136, 79, 34, 0.02);
                "
              >
                <div class="flex items-center justify-between px-4 py-2.5 border-b border-primary-100/30">
                  <span class="text-xs text-clay-muted font-medium">搜索历史</span>
                  <button type="button" class="text-xs text-clay-muted hover:text-primary font-medium transition-colors" @click="clearHistory">清除</button>
                </div>
                <div class="max-h-48 overflow-y-auto">
                  <button
                    v-for="(item, index) in history"
                    :key="index"
                    type="button"
                    class="w-full px-4 py-3 text-left text-sm hover:bg-primary-50 flex items-center transition-colors font-medium"
                    @mousedown.prevent="selectHistoryItem(item)"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-clay-muted/60 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    {{ item }}
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Top queries quick select -->
          <div v-if="topQueries.length > 0" class="flex flex-wrap gap-2">
            <span class="text-xs text-clay-muted font-medium leading-7 mr-0.5">热搜</span>
            <button
              v-for="(query, idx) in topQueries"
              :key="idx"
              type="button"
              class="px-3 py-1 text-xs font-medium rounded-full transition-all duration-200 hover:-translate-y-0.5 text-primary"
              style="background: rgba(255, 255, 255, 0.6); box-shadow: 4px 4px 8px rgba(136, 79, 34, 0.06), -3px -3px 6px rgba(255, 255, 255, 0.8), inset 2px 2px 4px rgba(255, 255, 255, 0.5), inset -2px -2px 4px rgba(136, 79, 34, 0.02);"
              @click="selectTopQuery(query)"
            >
              {{ query.label }}
            </button>
          </div>

          <DateSelector v-model="form.offset" />

          <button
            type="button"
            :disabled="loading || !form.building.trim()"
            class="w-full btn-clay-primary h-14 text-base"
            @click="search"
          >
            <span v-if="!loading" style="font-family: 'Nunito', sans-serif;">查询全天状态</span>
            <span v-else class="flex items-center">
              <svg class="animate-spin -ml-1 mr-2 h-5 w-5 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
              </svg>
              查询中...
            </span>
          </button>
        </div>
      </div>

      <!-- Legend -->
      <div v-if="hasSearched" class="clay-card p-5 sm:p-6">
        <div class="relative z-10">
          <h3
            class="text-sm font-bold text-clay-foreground mb-3"
            style="font-family: 'Nunito', sans-serif;"
          >
            状态图例
          </h3>
          <div class="grid grid-cols-3 gap-2.5 text-xs">
            <div
              v-for="item in legendItems"
              :key="item.id"
              class="flex items-center space-x-2 py-1.5 px-2.5 rounded-[12px] transition-all duration-200 hover:bg-primary-50/50"
            >
              <span class="text-base">{{ item.emoji }}</span>
              <span class="text-clay-muted font-medium">{{ item.name }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Results table -->
      <div v-if="hasSearched && resultData" class="clay-card overflow-hidden">
        <div class="relative z-10">
          <div
            class="p-4 sm:p-5"
            style="
              border-bottom: 1px solid rgba(136, 79, 34, 0.08);
            "
          >
            <p class="text-sm text-clay-muted font-medium">
              {{ resultData.building }} -- {{ resultData.date }} -- {{ resultData.current_term }} 学期 -- 第{{ resultData.week }}周 -- 星期{{ resultData.day_of_week }}
            </p>
          </div>

          <div class="table-container">
            <table class="w-full text-sm">
              <thead>
                <tr style="background: rgba(136, 79, 34, 0.04);">
                  <th class="sticky-col px-3 py-3.5 text-left font-bold text-clay-foreground min-w-[100px] text-sm" style="font-family: 'Nunito', sans-serif;">
                    教室
                  </th>
                  <th
                    v-for="node in resultData.node_list"
                    :key="node.node_index"
                    class="px-1 py-3.5 text-center font-bold text-clay-foreground min-w-[40px] text-xs"
                    style="font-family: 'Nunito', sans-serif;"
                  >
                    {{ node.node_name }}
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="room in resultData.classrooms"
                  :key="room.room_name"
                  class="table-row-clay"
                >
                  <td class="sticky-col px-3 py-3 font-bold text-clay-foreground text-sm" style="font-family: 'Nunito', sans-serif;">
                    {{ room.room_name }}
                  </td>
                  <td
                    v-for="(status, idx) in room.status"
                    :key="`${room.room_name}-${idx}`"
                    class="px-1 py-2 text-center text-base"
                  >
                    {{ getEmoji(status.status_id) }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
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
  background: rgba(255, 255, 255, 0.92);
  backdrop-filter: blur(12px);
  -webkit-backdrop-filter: blur(12px);
  box-shadow: 3px 0 8px rgba(136, 79, 34, 0.04);
}

.table-row-clay {
  border-top: 1px solid rgba(136, 79, 34, 0.06);
  transition: background-color 0.2s ease;
}

.table-row-clay:hover {
  background-color: rgba(136, 79, 34, 0.03);
}
</style>
