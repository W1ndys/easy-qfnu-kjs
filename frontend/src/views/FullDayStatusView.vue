<script setup>
import { reactive, ref } from 'vue'
import { getErrorMessage, queryFullDayStatus } from '@/api'
import { useSystemStatus } from '@/composables/useSystemStatus'
import { useSearchHistory } from '@/composables/useSearchHistory'
import { useTopBuildings } from '@/composables/useTopBuildings'
import { useBuildingAliasReminder } from '@/composables/useBuildingAliasReminder'
import { useAlertDialog } from '@/composables/useAlertDialog'
import AppFooter from '@/components/AppFooter.vue'
import AppHeader from '@/components/AppHeader.vue'
import ConfirmDialog from '@/components/ConfirmDialog.vue'
import DateSelector from '@/components/DateSelector.vue'
import EmptyState from '@/components/EmptyState.vue'
import LoadingSpinner from '@/components/LoadingSpinner.vue'
import QRCodeCard from '@/components/QRCodeCard.vue'
import StatusWarning from '@/components/StatusWarning.vue'

const { statusLoading, inTeachingCalendar, hasPermission } = useSystemStatus()
const { history, addToHistory, clearHistory } = useSearchHistory()
const { topQueries } = useTopBuildings()
const {
  dialogOpen: aliasDialogOpen,
  normalizeBuildingName,
  confirmReminder,
  cancelReminder,
} = useBuildingAliasReminder()
const { alertState, showAlert, closeAlert } = useAlertDialog()

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
  { id: 1, code: '课', name: '正常上课', class: 'bg-[#FDEEEE] text-[#B42318] border-[#F5B3AE]' },
  { id: 2, code: '借', name: '借用', class: 'bg-[#FFF6E8] text-[#9A5A00] border-[#F3CF8D]' },
  { id: 3, code: '锁', name: '锁定', class: 'bg-[#F2EEEA] text-[#4C433D] border-[#D1C7BE]' },
  { id: 4, code: '考', name: '考试', class: 'bg-[#ECF3FF] text-[#1D4ED8] border-[#B7CBFF]' },
  { id: 5, code: '空', name: '空闲', class: 'bg-[#EAF8F3] text-[#156B52] border-[#A7DEC7]' },
  { id: 6, code: '固', name: '固定调课', class: 'bg-[#ECF3FF] text-[#1D4ED8] border-[#B7CBFF]' },
  { id: 7, code: '临', name: '临时调课', class: 'bg-[#F3E5D8] text-[#5F3517] border-[#E7CFBA]' },
  { id: 8, code: '全', name: '完全空闲', class: 'bg-[#EAF8F3] text-[#156B52] border-[#A7DEC7]' },
  { id: 9, code: '混', name: '跨模式', class: 'bg-[#FAF8F6] text-[#4C433D] border-[#E5DED7]' },
]

function getStatusItem(statusId) {
  return legendItems.find((item) => item.id === statusId) || { code: '-', name: '未知', class: 'bg-white text-clay-muted border-subtle' }
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

function applySearchItem(item) {
  form.building = item.building || item
  if (item.offset !== undefined) form.offset = item.offset
  if (item.date_offset !== undefined) form.offset = item.date_offset
  showHistory.value = false
}

function selectHistoryItem(item) {
  applySearchItem(item)
}

async function search() {
  const building = await normalizeBuildingName(form.building)

  if (!building) {
    return
  }

  form.building = building

  loading.value = true
  hasSearched.value = false
  resultData.value = null
  showHistory.value = false

  try {
    const data = await queryFullDayStatus({
      building,
      date_offset: form.offset,
    })

    resultData.value = data
    hasSearched.value = true
    addToHistory({
      building,
      offset: form.offset,
    })
  } catch (error) {
    console.error(error)
    showAlert(getErrorMessage(error, '查询失败'), {
      title: '查询失败',
    })
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="page-shell pb-10 antialiased">
    <AppHeader title="教室全天状态" showBack />

    <main class="relative z-10 mx-auto max-w-6xl space-y-5 px-4 py-5 sm:px-6 lg:px-8">
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
      <div v-else class="clay-card space-y-5 p-5 sm:p-6">
        <div class="relative z-10 space-y-5">
          <div>
            <label class="mb-2 block text-sm font-semibold text-clay-muted">教学楼</label>
            <div class="relative">
              <input
                v-model="form.building"
                type="text"
                class="w-full clay-input px-5 py-3.5 text-[15px]"
                placeholder="例如：老文史楼"
                @focus="onInputFocus"
                @blur="onInputBlur"
                @input="onInputChange"
              />
              <!-- Search history dropdown -->
              <div
                v-if="inputFocused && showHistory && history.length > 0"
                class="absolute z-40 mt-2 w-full overflow-hidden rounded-2xl border border-subtle bg-white shadow-claySurface"
              >
                <div class="flex items-center justify-between border-b border-subtle px-4 py-2.5">
                  <span class="text-xs font-medium text-clay-muted">搜索历史</span>
                  <button type="button" class="text-xs font-medium text-clay-muted transition-colors hover:text-primary" @click="clearHistory">清除</button>
                </div>
                <div class="max-h-48 overflow-y-auto">
                  <button
                    v-for="(item, index) in history"
                    :key="index"
                    type="button"
                    class="flex w-full items-center px-4 py-3 text-left text-sm font-medium transition-colors hover:bg-primary-50"
                    @mousedown.prevent="selectHistoryItem(item)"
                  >
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-clay-muted/60 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                    <span>{{ item.label || item.building || item }}</span>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Top queries quick select -->
          <div v-if="topQueries.length > 0" class="flex flex-wrap gap-2">
            <span class="mr-0.5 text-xs font-medium leading-9 text-clay-muted">热搜</span>
            <button
              v-for="(query, idx) in topQueries"
              :key="idx"
              type="button"
              class="design-chip px-3 text-xs font-medium text-primary transition hover:border-primary-200 hover:bg-primary-50"
              @click="selectTopQuery(query)"
            >
              {{ query.label }}
            </button>
          </div>

          <DateSelector v-model="form.offset" />

          <button
            type="button"
            :disabled="loading || !form.building.trim()"
            class="btn-clay-primary h-12 w-full text-base"
            @click="search"
          >
            <span v-if="!loading">查询全天状态</span>
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
          <h3 class="mb-3 text-sm font-semibold text-clay-foreground">
            状态图例
          </h3>
          <div class="grid grid-cols-2 gap-2.5 text-xs sm:grid-cols-3 lg:grid-cols-5">
            <div
              v-for="item in legendItems"
              :key="item.id"
              class="flex items-center gap-2 rounded-xl border bg-white px-2.5 py-2"
            >
              <span :class="['inline-flex h-6 w-6 items-center justify-center rounded-md border text-[11px] font-bold', item.class]">{{ item.code }}</span>
              <span class="font-medium text-clay-muted">{{ item.name }}</span>
            </div>
          </div>
        </div>
      </div>

      <!-- Results table -->
      <div v-if="hasSearched && resultData" class="clay-card overflow-hidden">
        <div class="relative z-10">
          <div class="border-b border-subtle p-4 sm:p-5">
            <p class="text-sm font-medium text-clay-muted">
              {{ resultData.building }} -- {{ resultData.date }} -- {{ resultData.current_term }} 学期 -- 第{{ resultData.week }}周 -- 星期{{ resultData.day_of_week }}
            </p>
          </div>

          <div class="table-container">
            <table class="w-full text-sm">
              <thead>
                <tr class="bg-[#F3EFEB]">
                  <th class="sticky-col min-w-[100px] px-3 py-3.5 text-left text-sm font-bold text-clay-foreground">
                    教室
                  </th>
                  <th
                    v-for="node in resultData.node_list"
                    :key="node.node_index"
                    class="min-w-[44px] px-1 py-3.5 text-center text-xs font-bold text-clay-foreground"
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
                  <td class="sticky-col px-3 py-3 text-sm font-bold text-clay-foreground">
                    {{ room.room_name }}
                  </td>
                  <td
                    v-for="(status, idx) in room.status"
                    :key="`${room.room_name}-${idx}`"
                    class="px-1 py-2 text-center text-xs"
                  >
                    <span
                      :class="['inline-flex h-7 min-w-7 items-center justify-center rounded-md border px-1.5 font-bold', getStatusItem(status.status_id).class]"
                      :title="getStatusItem(status.status_id).name"
                    >
                      {{ getStatusItem(status.status_id).code }}
                    </span>
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

    <ConfirmDialog
      :open="aliasDialogOpen"
      title="教学楼名称提醒"
      message="你是否要搜索“综合教学楼”？注意老校区综合楼全称是“综合教学楼”哦！~"
      confirm-text="改为综合教学楼"
      cancel-text="继续搜索综合楼"
      @confirm="confirmReminder"
      @cancel="cancelReminder"
    />

    <ConfirmDialog
      :open="alertState.open"
      :title="alertState.title"
      :message="alertState.message"
      :confirm-text="alertState.buttonText"
      :show-cancel="false"
      @confirm="closeAlert"
      @cancel="closeAlert"
    />
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
  background: #FFFDFC;
  box-shadow: 3px 0 8px rgba(31, 27, 24, 0.04);
}

.table-row-clay {
  border-top: 1px solid #E5DED7;
  transition: background-color 0.2s ease;
}

.table-row-clay:hover {
  background-color: #F8F1EB;
}
</style>
