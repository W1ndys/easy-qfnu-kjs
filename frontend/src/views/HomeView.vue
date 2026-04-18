<script setup>
import AnnouncementCard from '@/components/AnnouncementCard.vue'
import QRCodeCard from '@/components/QRCodeCard.vue'
import AppFooter from '@/components/AppFooter.vue'
import StatsCard from '@/components/StatsCard.vue'
</script>

<template>
  <div class="min-h-screen bg-clay-canvas text-clay-foreground antialiased pb-10 relative overflow-hidden">
    <!-- Floating Clay Blobs - ambient background lighting -->
    <div class="pointer-events-none fixed inset-0 overflow-hidden -z-10">
      <div
        class="absolute h-[60vh] w-[60vh] rounded-full blur-3xl animate-clay-float"
        style="background: rgba(136, 79, 34, 0.08); top: -10%; left: -10%;"
      ></div>
      <div
        class="absolute h-[50vh] w-[50vh] rounded-full blur-3xl animate-clay-float-delayed animation-delay-2000"
        style="background: rgba(196, 149, 106, 0.08); top: 20%; right: -10%;"
      ></div>
      <div
        class="absolute h-[45vh] w-[45vh] rounded-full blur-3xl animate-clay-float-slow animation-delay-4000"
        style="background: rgba(16, 185, 129, 0.06); bottom: -5%; left: 20%;"
      ></div>
    </div>

    <main class="px-4 py-6 space-y-6 max-w-xl mx-auto relative z-10">
      <!-- Hero Title -->
      <div class="text-center py-8 sm:py-12">
        <h2
          class="text-3xl sm:text-4xl md:text-5xl font-black tracking-tight text-clay-foreground leading-[1.1]"
          style="font-family: 'Nunito', sans-serif;"
        >
          QFNU
          <span
            class="block mt-1 bg-gradient-to-r from-clay-foreground via-primary to-primary-light bg-clip-text text-transparent"
          >
            教室查询系统
          </span>
        </h2>
      </div>

      <!-- Announcement Card -->
      <AnnouncementCard />

      <!-- Stats Card -->
      <StatsCard />

      <!-- Navigation Cards -->
      <div class="space-y-5">
        <!-- Empty Classroom Query Card -->
        <RouterLink
          to="/empty-classroom"
          class="block clay-card p-6 sm:p-8"
        >
          <div class="relative z-10 flex items-center space-x-4">
            <div
              class="w-14 h-14 rounded-2xl flex items-center justify-center flex-shrink-0"
              style="
                background: linear-gradient(135deg, #6EE7B7 0%, #10B981 100%);
                box-shadow:
                  8px 8px 16px rgba(16, 185, 129, 0.25),
                  -4px -4px 8px rgba(255, 255, 255, 0.3),
                  inset 2px 2px 4px rgba(255, 255, 255, 0.4),
                  inset -2px -2px 4px rgba(0, 0, 0, 0.05);
              "
            >
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                />
              </svg>
            </div>
            <div class="flex-1">
              <h3 class="text-lg font-bold text-clay-foreground" style="font-family: 'Nunito', sans-serif;">空教室查询</h3>
              <p class="text-sm text-clay-muted mt-1 leading-relaxed font-medium">查询指定时间段内的空闲教室</p>
            </div>
            <div
              class="w-8 h-8 rounded-full flex items-center justify-center flex-shrink-0"
              style="
                background: rgba(255, 255, 255, 0.6);
                box-shadow:
                  4px 4px 8px rgba(136, 79, 34, 0.06),
                  -3px -3px 6px rgba(255, 255, 255, 0.8);
              "
            >
              <svg class="w-4 h-4 text-clay-muted" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </div>
        </RouterLink>

        <!-- Full Day Status Card -->
        <RouterLink
          to="/full-day-status"
          class="block clay-card p-6 sm:p-8"
        >
          <div class="relative z-10 flex items-center space-x-4">
            <div
              class="w-14 h-14 rounded-2xl flex items-center justify-center flex-shrink-0"
              style="
                background: linear-gradient(135deg, #C4956A 0%, #884F22 100%);
                box-shadow:
                  8px 8px 16px rgba(136, 79, 34, 0.25),
                  -4px -4px 8px rgba(255, 255, 255, 0.3),
                  inset 2px 2px 4px rgba(255, 255, 255, 0.4),
                  inset -2px -2px 4px rgba(0, 0, 0, 0.05);
              "
            >
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 17V7m0 10a2 2 0 01-2 2H5a2 2 0 01-2-2V7a2 2 0 012-2h2a2 2 0 012 2m0 10a2 2 0 002 2h2a2 2 0 002-2M9 7a2 2 0 012-2h2a2 2 0 012 2m0 10V7m0 10a2 2 0 002 2h2a2 2 0 002-2V7a2 2 0 00-2-2h-2a2 2 0 00-2 2"
                />
              </svg>
            </div>
            <div class="flex-1">
              <h3 class="text-lg font-bold text-clay-foreground" style="font-family: 'Nunito', sans-serif;">教室全天状态</h3>
              <p class="text-sm text-clay-muted mt-1 leading-relaxed font-medium">查看指定教学楼全天各节次的占用情况</p>
            </div>
            <div
              class="w-8 h-8 rounded-full flex items-center justify-center flex-shrink-0"
              style="
                background: rgba(255, 255, 255, 0.6);
                box-shadow:
                  4px 4px 8px rgba(136, 79, 34, 0.06),
                  -3px -3px 6px rgba(255, 255, 255, 0.8);
              "
            >
              <svg class="w-4 h-4 text-clay-muted" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </div>
        </RouterLink>

        <!-- Dashboard Data Screen Card -->
        <RouterLink
          to="/dashboard"
          class="block clay-card p-6 sm:p-8"
        >
          <div class="relative z-10 flex items-center space-x-4">
            <div
              class="w-14 h-14 rounded-2xl flex items-center justify-center flex-shrink-0"
              style="
                background: linear-gradient(135deg, #93C5FD 0%, #3B82F6 100%);
                box-shadow:
                  8px 8px 16px rgba(59, 130, 246, 0.25),
                  -4px -4px 8px rgba(255, 255, 255, 0.3),
                  inset 2px 2px 4px rgba(255, 255, 255, 0.4),
                  inset -2px -2px 4px rgba(0, 0, 0, 0.05);
              "
            >
              <svg class="w-7 h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
                />
              </svg>
            </div>
            <div class="flex-1">
              <h3 class="text-lg font-bold text-clay-foreground" style="font-family: 'Nunito', sans-serif;">数据大屏</h3>
              <p class="text-sm text-clay-muted mt-1 leading-relaxed font-medium">查看详细的查询统计与趋势分析</p>
            </div>
            <div
              class="w-8 h-8 rounded-full flex items-center justify-center flex-shrink-0"
              style="
                background: rgba(255, 255, 255, 0.6);
                box-shadow:
                  4px 4px 8px rgba(136, 79, 34, 0.06),
                  -3px -3px 6px rgba(255, 255, 255, 0.8);
              "
            >
              <svg class="w-4 h-4 text-clay-muted" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </div>
        </RouterLink>
      </div>

      <QRCodeCard />
      <AppFooter />
    </main>
  </div>
</template>
