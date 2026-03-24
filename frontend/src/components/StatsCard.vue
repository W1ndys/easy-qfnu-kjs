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
      <h3
        class="text-base font-bold text-clay-foreground mb-5"
        style="font-family: 'Nunito', sans-serif;"
      >
        查询统计
      </h3>

      <div class="grid grid-cols-3 gap-3 sm:gap-4 text-center">
        <!-- Today -->
        <div
          class="rounded-[24px] p-4 transition-all duration-300 hover:scale-105 hover:-translate-y-1"
          style="
            background: linear-gradient(135deg, rgba(14, 165, 233, 0.08) 0%, rgba(14, 165, 233, 0.15) 100%);
            box-shadow:
              8px 8px 16px rgba(14, 165, 233, 0.08),
              -6px -6px 12px rgba(255, 255, 255, 0.9),
              inset 3px 3px 6px rgba(255, 255, 255, 0.6),
              inset -3px -3px 6px rgba(14, 165, 233, 0.04);
          "
        >
          <div class="text-2xl sm:text-3xl font-black text-sky-600" style="font-family: 'Nunito', sans-serif;">
            {{ stats.today_count }}
          </div>
          <div class="text-xs text-clay-muted mt-1 font-medium">今日查询</div>
          <div
            v-if="stats.today_top"
            class="text-xs text-sky-500 mt-1.5 truncate font-medium"
            :title="stats.today_top"
          >
            {{ stats.today_top }}
          </div>
        </div>

        <!-- Week -->
        <div
          class="rounded-[24px] p-4 transition-all duration-300 hover:scale-105 hover:-translate-y-1"
          style="
            background: linear-gradient(135deg, rgba(16, 185, 129, 0.08) 0%, rgba(16, 185, 129, 0.15) 100%);
            box-shadow:
              8px 8px 16px rgba(16, 185, 129, 0.08),
              -6px -6px 12px rgba(255, 255, 255, 0.9),
              inset 3px 3px 6px rgba(255, 255, 255, 0.6),
              inset -3px -3px 6px rgba(16, 185, 129, 0.04);
          "
        >
          <div class="text-2xl sm:text-3xl font-black text-emerald-600" style="font-family: 'Nunito', sans-serif;">
            {{ stats.week_count }}
          </div>
          <div class="text-xs text-clay-muted mt-1 font-medium">本周查询</div>
          <div
            v-if="stats.week_top"
            class="text-xs text-emerald-500 mt-1.5 truncate font-medium"
            :title="stats.week_top"
          >
            {{ stats.week_top }}
          </div>
        </div>

        <!-- Month -->
        <div
          class="rounded-[24px] p-4 transition-all duration-300 hover:scale-105 hover:-translate-y-1"
          style="
            background: linear-gradient(135deg, rgba(136, 79, 34, 0.08) 0%, rgba(136, 79, 34, 0.15) 100%);
            box-shadow:
              8px 8px 16px rgba(136, 79, 34, 0.08),
              -6px -6px 12px rgba(255, 255, 255, 0.9),
              inset 3px 3px 6px rgba(255, 255, 255, 0.6),
              inset -3px -3px 6px rgba(136, 79, 34, 0.04);
          "
        >
          <div class="text-2xl sm:text-3xl font-black text-primary" style="font-family: 'Nunito', sans-serif;">
            {{ stats.month_count }}
          </div>
          <div class="text-xs text-clay-muted mt-1 font-medium">本月查询</div>
          <div
            v-if="stats.month_top"
            class="text-xs text-primary-light mt-1.5 truncate font-medium"
            :title="stats.month_top"
          >
            {{ stats.month_top }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
