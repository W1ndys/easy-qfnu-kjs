<script setup>
import { ref, onMounted } from 'vue'
import { getStats } from '@/api'

const stats = ref(null)
const loading = ref(true)

onMounted(async () => {
  try {
    stats.value = await getStats()
  } catch {
    // 静默失败，不影响首页其他功能
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div v-if="!loading && stats" class="bg-white rounded-2xl p-5 shadow-sm border border-gray-100">
    <h3 class="text-base font-semibold text-gray-700 mb-4">查询统计</h3>

    <div class="grid grid-cols-3 gap-3 text-center">
      <div class="bg-blue-50 rounded-xl p-3">
        <div class="text-2xl font-bold text-blue-600">{{ stats.today_count }}</div>
        <div class="text-xs text-gray-500 mt-1">今日查询</div>
        <div v-if="stats.today_top" class="text-xs text-blue-500 mt-1 truncate" :title="stats.today_top">
          {{ stats.today_top }}
        </div>
      </div>

      <div class="bg-green-50 rounded-xl p-3">
        <div class="text-2xl font-bold text-green-600">{{ stats.week_count }}</div>
        <div class="text-xs text-gray-500 mt-1">本周查询</div>
        <div v-if="stats.week_top" class="text-xs text-green-500 mt-1 truncate" :title="stats.week_top">
          {{ stats.week_top }}
        </div>
      </div>

      <div class="bg-purple-50 rounded-xl p-3">
        <div class="text-2xl font-bold text-purple-600">{{ stats.month_count }}</div>
        <div class="text-xs text-gray-500 mt-1">本月查询</div>
        <div v-if="stats.month_top" class="text-xs text-purple-500 mt-1 truncate" :title="stats.month_top">
          {{ stats.month_top }}
        </div>
      </div>
    </div>
  </div>
</template>
