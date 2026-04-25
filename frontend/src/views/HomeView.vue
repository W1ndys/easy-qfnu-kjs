<script setup>
import AnnouncementCard from '@/components/AnnouncementCard.vue'
import QRCodeCard from '@/components/QRCodeCard.vue'
import AppFooter from '@/components/AppFooter.vue'
import StatsCard from '@/components/StatsCard.vue'

const featureCards = [
  {
    to: '/empty-classroom',
    title: '空教室查询',
    description: '输入教学楼、日期和节次，快速找到可用教室。',
    label: '高频查询',
    colorClass: 'border-[#A7DEC7] bg-[#EAF8F3] text-[#156B52]',
    iconPath: 'M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z',
  },
  {
    to: '/full-day-status',
    title: '教室全天状态',
    description: '按节次查看教学楼全天占用情况，适合对比和排查。',
    label: '状态矩阵',
    colorClass: 'border-primary-200 bg-primary-100 text-primary',
    iconPath: 'M9 17V7m0 10a2 2 0 01-2 2H5a2 2 0 01-2-2V7a2 2 0 012-2h2a2 2 0 012 2m0 10a2 2 0 002 2h2a2 2 0 002-2M9 7a2 2 0 012-2h2a2 2 0 012 2m0 10V7m0 10a2 2 0 002 2h2a2 2 0 002-2V7a2 2 0 00-2-2h-2a2 2 0 00-2 2',
  },
  {
    to: '/dashboard',
    title: '数据大屏',
    description: '查看查询趋势、热门关键词和结果分布。',
    label: '运营分析',
    colorClass: 'border-[#B7CBFF] bg-[#ECF3FF] text-[#1D4ED8]',
    iconPath: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z',
  },
]
</script>

<template>
  <div class="page-shell pb-10 antialiased">
    <main class="mx-auto max-w-5xl space-y-6 px-4 py-6 sm:px-6 lg:px-8">
      <section class="overflow-hidden rounded-[20px] border border-subtle bg-white shadow-clayCard">
        <div class="grid gap-6 p-6 sm:p-8 lg:grid-cols-[1.15fr_0.85fr] lg:items-end">
          <div>
            <div class="mb-4 inline-flex items-center rounded-full border border-primary-200 bg-primary-100 px-3 py-1 text-xs font-semibold text-primary">
              QFNU Campus Tool
            </div>
            <h1 class="font-heading text-[2rem] font-bold leading-tight tracking-[-0.02em] text-clay-foreground sm:text-[2.5rem]">
              教室查询系统
            </h1>
            <p class="mt-3 max-w-2xl text-base leading-7 text-clay-muted sm:text-lg">
              面向曲阜师范大学学生的轻量查询工具，把空教室、全天状态和查询热度组织在一个清晰入口里。
            </p>
          </div>
          <div class="rounded-2xl border border-subtle bg-[#FAF8F6] p-4">
            <div class="text-xs font-semibold uppercase tracking-[0.08em] text-clay-muted">Quick Path</div>
            <div class="mt-3 grid grid-cols-3 gap-2 text-center text-xs font-semibold text-clay-muted">
              <span class="rounded-xl bg-white px-2 py-3 text-primary shadow-sm">选楼</span>
              <span class="rounded-xl bg-white px-2 py-3 text-primary shadow-sm">选日期</span>
              <span class="rounded-xl bg-white px-2 py-3 text-primary shadow-sm">看结果</span>
            </div>
          </div>
        </div>
      </section>

      <AnnouncementCard />
      <StatsCard />

      <section class="grid gap-4 md:grid-cols-3">
        <RouterLink
          v-for="card in featureCards"
          :key="card.to"
          :to="card.to"
          class="clay-card group block p-5 focus:outline-none focus:ring-4 focus:ring-primary-200"
        >
          <div class="flex h-full flex-col gap-5">
            <div class="flex items-start justify-between gap-4">
              <div :class="['flex h-12 w-12 items-center justify-center rounded-2xl border', card.colorClass]">
                <svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" :d="card.iconPath" />
                </svg>
              </div>
              <span class="rounded-full border border-subtle bg-[#FAF8F6] px-2.5 py-1 text-xs font-semibold text-clay-muted">
                {{ card.label }}
              </span>
            </div>
            <div class="flex-1">
              <h2 class="text-lg font-semibold text-clay-foreground">{{ card.title }}</h2>
              <p class="mt-2 text-sm leading-6 text-clay-muted">{{ card.description }}</p>
            </div>
            <div class="flex items-center justify-between border-t border-subtle pt-4 text-sm font-semibold text-primary">
              <span>立即进入</span>
              <svg class="h-4 w-4 transition group-hover:translate-x-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
              </svg>
            </div>
          </div>
        </RouterLink>
      </section>

      <QRCodeCard />
      <AppFooter />
    </main>
  </div>
</template>
