<script setup>
import { computed, reactive, ref } from 'vue'
import { getErrorMessage, queryClassrooms } from '@/api'
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

const { statusLoading, inTeachingCalendar, hasPermission, currentWeek, currentTerm } = useSystemStatus()
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
  if (query.start_node) form.start = query.start_node
  if (query.end_node) form.end = query.end_node
  showHistory.value = false
  search()
}

const loading = ref(false)
const hasSearched = ref(false)
const results = ref([])
const resultInfo = ref(null)
const displayLimit = ref(100)
const showHistory = ref(false)
const inputFocused = ref(false)

const form = reactive({
  building: '',
  offset: 0,
  start: '01',
  end: '11',
})

const nodeOptions = Array.from({ length: 11 }, (_, index) => String(index + 1).padStart(2, '0'))

const displayedResults = computed(() => results.value.slice(0, displayLimit.value))

const showHistoryList = computed(() => inputFocused.value && showHistory.value && history.value.length > 0)

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
  if (item.start) form.start = item.start
  if (item.end) form.end = item.end
  if (item.start_node) form.start = item.start_node
  if (item.end_node) form.end = item.end_node
  showHistory.value = false
}

function selectHistoryItem(item) {
  applySearchItem(item)
}

async function search() {
  const building = await normalizeBuildingName(form.building)

  if (!building) {
    showAlert('请输入教学楼', {
      title: '搜索条件不完整',
    })
    return
  }

  form.building = building

  loading.value = true
  displayLimit.value = 100
  hasSearched.value = false
  results.value = []
  resultInfo.value = null
  showHistory.value = false

  try {
    const data = await queryClassrooms({
      building,
      start_node: form.start,
      end_node: form.end,
      date_offset: form.offset,
    })

    results.value = data.classrooms || []
    resultInfo.value = {
      date: data.date,
      week: data.week,
      day: data.day_of_week,
    }
    hasSearched.value = true
    addToHistory({
      building,
      offset: form.offset,
      start: form.start,
      end: form.end,
    })
  } catch (error) {
    console.error(error)
    showAlert(getErrorMessage(error, '查询出错，请重试'), {
      title: '查询失败',
    })
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="page-shell pb-10 antialiased">
    <AppHeader title="空教室查询" showBack />

    <main class="relative z-10 mx-auto max-w-3xl space-y-5 px-4 py-5 sm:px-6">
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

      <div v-else class="clay-card space-y-5 p-5 sm:p-6">
        <div class="relative z-10 space-y-5">
          <div
            v-if="inTeachingCalendar"
            class="rounded-2xl border border-[#A7DEC7] bg-[#EAF8F3] p-3.5"
          >
            <div class="flex items-center gap-2.5 text-sm font-medium text-[#156B52]">
              <div class="flex h-6 w-6 flex-shrink-0 items-center justify-center rounded-full bg-[#156B52]">
                <svg class="h-3.5 w-3.5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
                </svg>
              </div>
              <span>当前：<strong>{{ currentTerm }}</strong> 第<strong>{{ currentWeek }}</strong>周</span>
            </div>
          </div>

          <!-- Building input -->
          <div>
            <label class="mb-2 block text-sm font-semibold text-clay-muted">教学楼</label>
            <div class="relative">
              <span class="absolute left-4 top-1/2 -translate-y-1/2 text-clay-muted/60">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"
                  />
                </svg>
              </span>
              <input
                v-model="form.building"
                type="text"
                class="w-full clay-input px-5 py-3.5 pl-11 text-[15px]"
                placeholder="例如：老文史楼"
                @focus="onInputFocus"
                @blur="onInputBlur"
                @input="onInputChange"
              />
              <!-- Search history dropdown -->
              <div
                v-if="showHistoryList"
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

          <!-- Date selector -->
          <DateSelector v-model="form.offset" />

          <!-- Node selectors -->
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="mb-2 block text-sm font-semibold text-clay-muted">起始节次</label>
              <select
                v-model="form.start"
                class="w-full clay-input appearance-none px-5 py-3.5 text-[15px]"
              >
                <option v-for="value in nodeOptions" :key="`start-${value}`" :value="value">{{ value }}</option>
              </select>
            </div>

            <div>
              <label class="mb-2 block text-sm font-semibold text-clay-muted">终止节次</label>
              <select
                v-model="form.end"
                class="w-full clay-input appearance-none px-5 py-3.5 text-[15px]"
              >
                <option v-for="value in nodeOptions" :key="`end-${value}`" :value="value">{{ value }}</option>
              </select>
            </div>
          </div>

          <!-- Search button -->
          <button
            type="button"
            :disabled="loading"
            class="btn-clay-primary h-12 w-full text-base"
            @click="search"
          >
            <span v-if="!loading">查询空闲教室</span>
            <span v-else class="flex items-center">
              <svg class="animate-spin -ml-1 mr-2 h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
              </svg>
              查询中...
            </span>
          </button>
        </div>
      </div>

      <!-- Result info bar -->
      <div v-if="resultInfo" class="flex items-center justify-between rounded-2xl border border-subtle bg-white px-4 py-3 text-xs font-medium text-clay-muted">
        <span>{{ resultInfo.date }} (第{{ resultInfo.week }}周 星期{{ resultInfo.day }})</span>
        <span class="rounded-full bg-primary-100 px-3 py-1 font-bold text-primary">
          共 {{ results.length }} 间
        </span>
      </div>

      <!-- Results grid -->
      <div v-if="results.length > 0" class="space-y-4">
        <div class="grid grid-cols-2 gap-2 sm:grid-cols-4 sm:gap-3 md:grid-cols-5">
          <div
            v-for="(room, index) in displayedResults"
            :key="`${room}-${index}`"
            class="flex min-h-11 items-center justify-center rounded-xl border border-subtle bg-white px-2 py-2.5 text-center transition hover:border-primary-200 hover:bg-primary-50"
          >
            <span class="text-sm font-bold text-primary sm:text-base">{{ room }}</span>
          </div>
        </div>

        <div v-if="results.length > displayLimit" class="mt-4 text-center">
          <button
            type="button"
            class="rounded-lg border border-primary-200 bg-white px-6 py-2.5 text-sm font-bold text-primary transition hover:bg-primary-50"
            @click="displayLimit += 100"
          >
            加载更多 (显示 {{ displayedResults.length }} / {{ results.length }})
          </button>
        </div>
      </div>

      <EmptyState v-if="hasSearched && results.length === 0 && !loading" text="该时间段暂无空闲教室" />

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
