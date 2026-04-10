<script setup>
import { computed, reactive, ref } from 'vue'
import { getErrorMessage, queryClassrooms } from '@/api'
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

const { statusLoading, inTeachingCalendar, hasPermission, currentWeek, currentTerm } = useSystemStatus()
const { history, addToHistory, clearHistory } = useSearchHistory()
const { topQueries } = useTopBuildings()

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

function selectHistoryItem(keyword) {
  form.building = keyword
  showHistory.value = false
}

async function search() {
  if (!form.building.trim()) {
    alert('请输入教学楼')
    return
  }

  loading.value = true
  displayLimit.value = 100
  hasSearched.value = false
  results.value = []
  resultInfo.value = null
  showHistory.value = false

  try {
    const data = await queryClassrooms({
      building: form.building,
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
    addToHistory(form.building)
  } catch (error) {
    console.error(error)
    alert(getErrorMessage(error, '查询出错，请重试'))
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
        class="absolute h-[50vh] w-[50vh] rounded-full blur-3xl animate-clay-float"
        style="background: rgba(136, 79, 34, 0.06); top: -5%; right: -15%;"
      ></div>
      <div
        class="absolute h-[40vh] w-[40vh] rounded-full blur-3xl animate-clay-float-delayed animation-delay-2000"
        style="background: rgba(16, 185, 129, 0.05); bottom: 10%; left: -10%;"
      ></div>
    </div>

    <AppHeader title="空教室查询" showBack />

    <main class="px-4 py-4 space-y-5 max-w-xl mx-auto relative z-10">
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

      <div v-else class="clay-card p-5 sm:p-7 space-y-5">
        <div class="relative z-10 space-y-5">
          <!-- Term info badge -->
          <div
            v-if="inTeachingCalendar"
            class="rounded-[20px] p-3.5"
            style="
              background: linear-gradient(135deg, rgba(16, 185, 129, 0.06) 0%, rgba(16, 185, 129, 0.12) 100%);
              box-shadow:
                8px 8px 16px rgba(16, 185, 129, 0.06),
                -6px -6px 12px rgba(255, 255, 255, 0.9),
                inset 3px 3px 6px rgba(255, 255, 255, 0.5),
                inset -3px -3px 6px rgba(16, 185, 129, 0.03);
            "
          >
            <div class="flex items-center space-x-2.5 text-emerald-700 text-sm font-medium">
              <div
                class="w-6 h-6 rounded-full flex items-center justify-center flex-shrink-0"
                style="
                  background: linear-gradient(135deg, #6EE7B7 0%, #10B981 100%);
                  box-shadow: 3px 3px 6px rgba(16, 185, 129, 0.2), -2px -2px 4px rgba(255, 255, 255, 0.3);
                "
              >
                <svg class="w-3.5 h-3.5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
                </svg>
              </div>
              <span>当前：<strong>{{ currentTerm }}</strong> 第<strong>{{ currentWeek }}</strong>周</span>
            </div>
          </div>

          <!-- Building input -->
          <div>
            <label class="block text-sm font-bold text-clay-muted mb-2 ml-1">教学楼</label>
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
                class="w-full clay-input py-3.5 pl-11 pr-5 text-[15px]"
                placeholder="例如：老文史楼"
                @focus="onInputFocus"
                @blur="onInputBlur"
                @input="onInputChange"
              />
              <!-- Search history dropdown -->
              <div
                v-if="showHistoryList"
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

          <!-- Date selector -->
          <DateSelector v-model="form.offset" />

          <!-- Node selectors -->
          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-bold text-clay-muted mb-2 ml-1">起始节次</label>
              <select
                v-model="form.start"
                class="w-full clay-input py-3.5 px-5 text-[15px] appearance-none"
              >
                <option v-for="value in nodeOptions" :key="`start-${value}`" :value="value">{{ value }}</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-bold text-clay-muted mb-2 ml-1">终止节次</label>
              <select
                v-model="form.end"
                class="w-full clay-input py-3.5 px-5 text-[15px] appearance-none"
              >
                <option v-for="value in nodeOptions" :key="`end-${value}`" :value="value">{{ value }}</option>
              </select>
            </div>
          </div>

          <!-- Search button -->
          <button
            type="button"
            :disabled="loading"
            class="w-full btn-clay-primary h-14 text-base"
            @click="search"
          >
            <span v-if="!loading" style="font-family: 'Nunito', sans-serif;">查询空闲教室</span>
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
      <div v-if="resultInfo" class="px-3 flex items-center justify-between text-xs text-clay-muted font-medium">
        <span>{{ resultInfo.date }} (第{{ resultInfo.week }}周 星期{{ resultInfo.day }})</span>
        <span
          class="px-3 py-1 rounded-full font-bold"
          style="
            background: rgba(255, 255, 255, 0.6);
            box-shadow:
              4px 4px 8px rgba(136, 79, 34, 0.04),
              -3px -3px 6px rgba(255, 255, 255, 0.8);
          "
        >
          共 {{ results.length }} 间
        </span>
      </div>

      <!-- Results grid -->
      <div v-if="results.length > 0" class="space-y-4">
        <div class="grid grid-cols-3 gap-2 sm:gap-3">
          <div
            v-for="(room, index) in displayedResults"
            :key="`${room}-${index}`"
            class="rounded-2xl py-2.5 px-2 flex items-center justify-center text-center transition-all duration-300 hover:-translate-y-0.5 hover:scale-105"
            style="
              background: rgba(255, 255, 255, 0.65);
              backdrop-filter: blur(12px);
              box-shadow:
                6px 6px 14px rgba(136, 79, 34, 0.05),
                -4px -4px 10px rgba(255, 255, 255, 0.9),
                inset 3px 3px 6px rgba(255, 255, 255, 0.7),
                inset -3px -3px 6px rgba(136, 79, 34, 0.02);
            "
          >
            <span class="text-primary font-bold text-sm sm:text-base" style="font-family: 'Nunito', sans-serif;">{{ room }}</span>
          </div>
        </div>

        <div v-if="results.length > displayLimit" class="mt-4 text-center">
          <button
            type="button"
            class="text-primary text-sm font-bold hover:underline py-2.5 px-6 rounded-[20px] transition-all duration-200 hover:-translate-y-0.5"
            style="
              background: rgba(255, 255, 255, 0.5);
              box-shadow:
                6px 6px 12px rgba(136, 79, 34, 0.05),
                -4px -4px 8px rgba(255, 255, 255, 0.8);
            "
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
  </div>
</template>
