<script setup>
import { ref, onMounted } from 'vue'
import { getStats } from '@/api'

const stats = ref(null)
const loading = ref(true)

onMounted(async () => {
  try {
    stats.value = await getStats()
  } catch {
    // silent fail
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div v-if="!loading && stats" class="clay-card p-6 sm:p-8">
    <div class="relative z-10">
      <h3 class="mb-5 text-base font-semibold text-clay-foreground">
        查询统计
      </h3>

      <div class="grid grid-cols-3 gap-3 sm:gap-4 text-center">
        <div class="rounded-2xl border border-[#B7CBFF] bg-[#ECF3FF] p-4">
          <div class="text-2xl font-bold text-[#1D4ED8] sm:text-3xl">
            {{ stats.today_count }}
          </div>
          <div class="mt-1 text-xs font-medium text-clay-muted">今日查询</div>
          <div
            v-if="stats.today_top"
            class="mt-1.5 truncate text-xs font-medium text-[#1D4ED8]"
            :title="stats.today_top"
          >
            {{ stats.today_top }}
          </div>
        </div>

        <div class="rounded-2xl border border-[#A7DEC7] bg-[#EAF8F3] p-4">
          <div class="text-2xl font-bold text-[#156B52] sm:text-3xl">
            {{ stats.week_count }}
          </div>
          <div class="mt-1 text-xs font-medium text-clay-muted">本周查询</div>
          <div
            v-if="stats.week_top"
            class="mt-1.5 truncate text-xs font-medium text-[#156B52]"
            :title="stats.week_top"
          >
            {{ stats.week_top }}
          </div>
        </div>

        <div class="rounded-2xl border border-primary-200 bg-primary-100 p-4">
          <div class="text-2xl font-bold text-primary sm:text-3xl">
            {{ stats.month_count }}
          </div>
          <div class="mt-1 text-xs font-medium text-clay-muted">本月查询</div>
          <div
            v-if="stats.month_top"
            class="mt-1.5 truncate text-xs font-medium text-primary"
            :title="stats.month_top"
          >
            {{ stats.month_top }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
