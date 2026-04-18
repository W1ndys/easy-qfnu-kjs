<script setup>
import { ref, computed, onMounted, onUnmounted, watch, nextTick, shallowRef } from 'vue'
import * as echarts from 'echarts/core'
import { BarChart, LineChart, PieChart } from 'echarts/charts'
import {
  TitleComponent,
  TooltipComponent,
  GridComponent,
  LegendComponent,
} from 'echarts/components'
import { CanvasRenderer } from 'echarts/renderers'
import AppHeader from '@/components/AppHeader.vue'
import AppFooter from '@/components/AppFooter.vue'
import { getDashboard } from '@/api'

echarts.use([
  BarChart, LineChart, PieChart,
  TitleComponent, TooltipComponent, GridComponent, LegendComponent,
  CanvasRenderer,
])

// ---- 主色调 ----
const COLORS = {
  primary: '#884F22',
  primaryLight: '#A67C52',
  caramel: '#C4956A',
  cream: '#F5E6D3',
  success: '#10B981',
  info: '#0EA5E9',
  warning: '#F59E0B',
  rose: '#F43F5E',
  violet: '#8B5CF6',
  slate: '#64748B',
}

const CHART_PALETTE = [
  COLORS.primary, COLORS.success, COLORS.info,
  COLORS.caramel, COLORS.warning, COLORS.violet,
  COLORS.rose, COLORS.slate, COLORS.primaryLight, '#6366F1',
]

// ---- 状态 ----
const timeRange = ref('today')
const loading = ref(true)
const data = ref(null)
const error = ref(null)

const timeRangeOptions = [
  { value: 'today', label: '今天' },
  { value: 'week', label: '最近7天' },
  { value: 'month', label: '最近30天' },
]

const timeRangeLabel = computed(() =>
  timeRangeOptions.find((o) => o.value === timeRange.value)?.label || ''
)

// ---- ECharts 实例引用 ----
const trendChartRef = ref(null)
const keywordChartRef = ref(null)
const nodeChartRef = ref(null)
const resultChartRef = ref(null)
const hourlyChartRef = ref(null)

const trendChart = shallowRef(null)
const keywordChart = shallowRef(null)
const nodeChart = shallowRef(null)
const resultChart = shallowRef(null)
const hourlyChart = shallowRef(null)

// ---- 数据加载 ----
async function fetchData() {
  loading.value = true
  error.value = null
  try {
    data.value = await getDashboard(timeRange.value)
  } catch (e) {
    error.value = e?.response?.data?.error || '获取数据失败'
  } finally {
    loading.value = false
  }
}

// ---- 图表通用配置 ----
function baseTooltip() {
  return {
    backgroundColor: 'rgba(255,255,255,0.95)',
    borderColor: COLORS.cream,
    borderWidth: 1,
    textStyle: { color: COLORS.primary, fontSize: 13 },
    extraCssText: 'box-shadow: 0 4px 20px rgba(136,79,34,0.12); border-radius: 12px;',
  }
}

// ---- 渲染图表 ----
function renderTrendChart() {
  if (!trendChartRef.value || !data.value?.trend) return
  if (!trendChart.value) {
    trendChart.value = echarts.init(trendChartRef.value)
  }
  const trend = data.value.trend
  trendChart.value.setOption({
    tooltip: { ...baseTooltip(), trigger: 'axis' },
    grid: { left: 50, right: 20, top: 20, bottom: 30 },
    xAxis: {
      type: 'category',
      data: trend.map((t) => t.label),
      axisLabel: {
        color: COLORS.slate,
        fontSize: 11,
        rotate: timeRange.value === 'month' ? 45 : 0,
      },
      axisLine: { lineStyle: { color: COLORS.cream } },
      axisTick: { show: false },
    },
    yAxis: {
      type: 'value',
      minInterval: 1,
      axisLabel: { color: COLORS.slate, fontSize: 11 },
      splitLine: { lineStyle: { color: COLORS.cream, type: 'dashed' } },
    },
    series: [
      {
        type: 'line',
        data: trend.map((t) => t.count),
        smooth: true,
        symbol: 'circle',
        symbolSize: 6,
        lineStyle: { color: COLORS.primary, width: 3 },
        itemStyle: { color: COLORS.primary, borderColor: '#fff', borderWidth: 2 },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: 'rgba(136,79,34,0.25)' },
            { offset: 1, color: 'rgba(136,79,34,0.02)' },
          ]),
        },
      },
    ],
  }, true)
}

function renderKeywordChart() {
  if (!keywordChartRef.value || !data.value?.top_keywords) return
  if (!keywordChart.value) {
    keywordChart.value = echarts.init(keywordChartRef.value)
  }
  const kw = data.value.top_keywords.slice().reverse()
  keywordChart.value.setOption({
    tooltip: { ...baseTooltip(), trigger: 'axis', axisPointer: { type: 'shadow' } },
    grid: { left: 100, right: 30, top: 10, bottom: 20 },
    xAxis: {
      type: 'value',
      minInterval: 1,
      axisLabel: { color: COLORS.slate, fontSize: 11 },
      splitLine: { lineStyle: { color: COLORS.cream, type: 'dashed' } },
    },
    yAxis: {
      type: 'category',
      data: kw.map((k) => k.keyword),
      axisLabel: {
        color: COLORS.primary,
        fontSize: 12,
        fontWeight: 600,
        width: 80,
        overflow: 'truncate',
      },
      axisLine: { show: false },
      axisTick: { show: false },
    },
    series: [
      {
        type: 'bar',
        data: kw.map((k, i) => ({
          value: k.count,
          itemStyle: {
            color: new echarts.graphic.LinearGradient(0, 0, 1, 0, [
              { offset: 0, color: CHART_PALETTE[i % CHART_PALETTE.length] },
              { offset: 1, color: CHART_PALETTE[i % CHART_PALETTE.length] + '88' },
            ]),
            borderRadius: [0, 8, 8, 0],
          },
        })),
        barWidth: '60%',
        label: {
          show: true,
          position: 'right',
          color: COLORS.slate,
          fontSize: 12,
          fontWeight: 600,
        },
      },
    ],
  }, true)
}

function renderNodeChart() {
  if (!nodeChartRef.value || !data.value?.node_dist) return
  if (!nodeChart.value) {
    nodeChart.value = echarts.init(nodeChartRef.value)
  }
  const nd = data.value.node_dist
  const nodeLabels = {
    '01-02': '1-2节', '03-04': '3-4节', '05-06': '5-6节',
    '07-08': '7-8节', '09-10': '9-10节', '09-11': '9-11节',
    '01-04': '1-4节', '05-08': '5-8节', '01-11': '全天',
  }
  nodeChart.value.setOption({
    tooltip: {
      ...baseTooltip(),
      trigger: 'item',
      formatter: '{b}: {c} ({d}%)',
    },
    series: [
      {
        type: 'pie',
        radius: ['40%', '70%'],
        center: ['50%', '50%'],
        avoidLabelOverlap: true,
        itemStyle: { borderRadius: 8, borderColor: '#fff', borderWidth: 2 },
        label: {
          show: true,
          fontSize: 12,
          color: COLORS.slate,
          formatter: '{b}\n{d}%',
        },
        emphasis: {
          label: { fontSize: 14, fontWeight: 'bold' },
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(136,79,34,0.2)',
          },
        },
        data: nd.map((n, i) => ({
          name: nodeLabels[n.node] || n.node,
          value: n.count,
          itemStyle: { color: CHART_PALETTE[i % CHART_PALETTE.length] },
        })),
      },
    ],
  }, true)
}

function renderResultChart() {
  if (!resultChartRef.value || !data.value?.result_stats) return
  if (!resultChart.value) {
    resultChart.value = echarts.init(resultChartRef.value)
  }
  const dist = data.value.result_stats.distribution || []
  resultChart.value.setOption({
    tooltip: { ...baseTooltip(), trigger: 'axis', axisPointer: { type: 'shadow' } },
    grid: { left: 50, right: 20, top: 20, bottom: 30 },
    xAxis: {
      type: 'category',
      data: dist.map((d) => d.range),
      axisLabel: { color: COLORS.slate, fontSize: 11 },
      axisLine: { lineStyle: { color: COLORS.cream } },
      axisTick: { show: false },
    },
    yAxis: {
      type: 'value',
      minInterval: 1,
      axisLabel: { color: COLORS.slate, fontSize: 11 },
      splitLine: { lineStyle: { color: COLORS.cream, type: 'dashed' } },
    },
    series: [
      {
        type: 'bar',
        data: dist.map((d, i) => ({
          value: d.count,
          itemStyle: {
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
              { offset: 0, color: CHART_PALETTE[i % CHART_PALETTE.length] },
              { offset: 1, color: CHART_PALETTE[i % CHART_PALETTE.length] + '66' },
            ]),
            borderRadius: [8, 8, 0, 0],
          },
        })),
        barWidth: '50%',
        label: {
          show: true,
          position: 'top',
          color: COLORS.slate,
          fontSize: 11,
          fontWeight: 600,
        },
      },
    ],
  }, true)
}

function renderHourlyChart() {
  if (!hourlyChartRef.value || !data.value?.hourly_dist) return
  if (!hourlyChart.value) {
    hourlyChart.value = echarts.init(hourlyChartRef.value)
  }
  const hd = data.value.hourly_dist
  const maxCount = Math.max(...hd.map((h) => h.count), 1)
  hourlyChart.value.setOption({
    tooltip: {
      ...baseTooltip(),
      trigger: 'axis',
      axisPointer: { type: 'shadow' },
      formatter: (params) => {
        const p = params[0]
        return `${p.name}:00 - ${p.name}:59<br/>查询次数: <b>${p.value}</b>`
      },
    },
    grid: { left: 50, right: 20, top: 20, bottom: 30 },
    xAxis: {
      type: 'category',
      data: hd.map((h) => String(h.hour).padStart(2, '0')),
      axisLabel: { color: COLORS.slate, fontSize: 10 },
      axisLine: { lineStyle: { color: COLORS.cream } },
      axisTick: { show: false },
    },
    yAxis: {
      type: 'value',
      minInterval: 1,
      axisLabel: { color: COLORS.slate, fontSize: 11 },
      splitLine: { lineStyle: { color: COLORS.cream, type: 'dashed' } },
    },
    series: [
      {
        type: 'bar',
        data: hd.map((h) => ({
          value: h.count,
          itemStyle: {
            color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
              {
                offset: 0,
                color: h.count >= maxCount * 0.8
                  ? COLORS.rose
                  : h.count >= maxCount * 0.5
                    ? COLORS.warning
                    : COLORS.success,
              },
              {
                offset: 1,
                color: h.count >= maxCount * 0.8
                  ? COLORS.rose + '44'
                  : h.count >= maxCount * 0.5
                    ? COLORS.warning + '44'
                    : COLORS.success + '44',
              },
            ]),
            borderRadius: [6, 6, 0, 0],
          },
        })),
        barWidth: '60%',
      },
    ],
  }, true)
}

function renderAllCharts() {
  nextTick(() => {
    renderTrendChart()
    renderKeywordChart()
    renderNodeChart()
    renderResultChart()
    renderHourlyChart()
  })
}

function resizeAllCharts() {
  trendChart.value?.resize()
  keywordChart.value?.resize()
  nodeChart.value?.resize()
  resultChart.value?.resize()
  hourlyChart.value?.resize()
}

function disposeAllCharts() {
  trendChart.value?.dispose()
  keywordChart.value?.dispose()
  nodeChart.value?.dispose()
  resultChart.value?.dispose()
  hourlyChart.value?.dispose()
}

// ---- 生命周期 ----
watch(timeRange, async () => {
  await fetchData()
  renderAllCharts()
})

watch(data, () => {
  if (data.value) renderAllCharts()
})

onMounted(async () => {
  await fetchData()
  renderAllCharts()
  window.addEventListener('resize', resizeAllCharts)
})

onUnmounted(() => {
  window.removeEventListener('resize', resizeAllCharts)
  disposeAllCharts()
})

// ---- 计算属性 ----
const overview = computed(() => data.value?.overview || {})
const peakHour = computed(() => {
  if (!data.value?.hourly_dist) return '--'
  const max = data.value.hourly_dist.reduce(
    (a, b) => (b.count > a.count ? b : a),
    { hour: 0, count: 0 }
  )
  if (max.count === 0) return '--'
  return `${String(max.hour).padStart(2, '0')}:00`
})
const resultRate = computed(() => {
  const rs = data.value?.result_stats
  if (!rs || (rs.zero_count + rs.non_zero_count) === 0) return '--'
  return ((rs.non_zero_count / (rs.zero_count + rs.non_zero_count)) * 100).toFixed(1) + '%'
})
</script>

<template>
  <div class="min-h-screen bg-clay-canvas text-clay-foreground antialiased pb-10 relative overflow-hidden">
    <!-- Floating Clay Blobs -->
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

    <AppHeader title="数据大屏" show-back />

    <main class="px-4 py-6 space-y-5 max-w-3xl mx-auto relative z-10">
      <!-- 时间范围切换 -->
      <div class="clay-card p-4">
        <div class="relative z-10 flex items-center justify-center gap-2">
          <button
            v-for="opt in timeRangeOptions"
            :key="opt.value"
            @click="timeRange = opt.value"
            class="px-5 py-2.5 rounded-2xl text-sm font-bold transition-all duration-300"
            :class="timeRange === opt.value
              ? 'text-white'
              : 'text-clay-muted hover:text-clay-foreground'"
            :style="timeRange === opt.value
              ? {
                  background: 'linear-gradient(135deg, #C4956A 0%, #884F22 100%)',
                  boxShadow: '8px 8px 16px rgba(136,79,34,0.2), -4px -4px 8px rgba(255,255,255,0.3), inset 2px 2px 4px rgba(255,255,255,0.3), inset -2px -2px 4px rgba(0,0,0,0.05)',
                }
              : {
                  background: 'rgba(255,255,255,0.5)',
                  boxShadow: '4px 4px 8px rgba(136,79,34,0.04), -3px -3px 6px rgba(255,255,255,0.8)',
                }"
          >
            {{ opt.label }}
          </button>
        </div>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="clay-card p-12 text-center">
        <div class="relative z-10">
          <div class="inline-block w-10 h-10 border-4 border-primary-200 border-t-primary rounded-full animate-spin"></div>
          <p class="text-clay-muted mt-4 font-medium">加载中...</p>
        </div>
      </div>

      <!-- Error -->
      <div v-else-if="error" class="clay-card p-8 text-center">
        <div class="relative z-10">
          <p class="text-red-500 font-semibold">{{ error }}</p>
          <button @click="fetchData" class="mt-4 btn-clay-primary px-6 py-2 text-sm">重试</button>
        </div>
      </div>

      <!-- 数据内容 -->
      <template v-else-if="data">
        <!-- 总览数字 -->
        <div class="clay-card p-6">
          <div class="relative z-10">
            <h3 class="text-base font-bold text-clay-foreground mb-5 font-heading">
              数据总览
              <span class="text-xs font-medium text-clay-muted ml-2">{{ timeRangeLabel }}</span>
            </h3>
            <div class="grid grid-cols-2 sm:grid-cols-4 gap-3">
              <!-- 总查询 -->
              <div
                class="rounded-[20px] p-4 text-center transition-all duration-300 hover:scale-105 hover:-translate-y-1"
                style="
                  background: linear-gradient(135deg, rgba(136,79,34,0.08) 0%, rgba(136,79,34,0.15) 100%);
                  box-shadow: 8px 8px 16px rgba(136,79,34,0.08), -6px -6px 12px rgba(255,255,255,0.9), inset 3px 3px 6px rgba(255,255,255,0.6), inset -3px -3px 6px rgba(136,79,34,0.04);
                "
              >
                <div class="text-2xl sm:text-3xl font-black text-primary font-heading">
                  {{ overview.total_count || 0 }}
                </div>
                <div class="text-xs text-clay-muted mt-1 font-medium">总查询次数</div>
              </div>

              <!-- 独立搜索词 -->
              <div
                class="rounded-[20px] p-4 text-center transition-all duration-300 hover:scale-105 hover:-translate-y-1"
                style="
                  background: linear-gradient(135deg, rgba(14,165,233,0.08) 0%, rgba(14,165,233,0.15) 100%);
                  box-shadow: 8px 8px 16px rgba(14,165,233,0.08), -6px -6px 12px rgba(255,255,255,0.9), inset 3px 3px 6px rgba(255,255,255,0.6), inset -3px -3px 6px rgba(14,165,233,0.04);
                "
              >
                <div class="text-2xl sm:text-3xl font-black text-sky-600 font-heading">
                  {{ overview.unique_keywords || 0 }}
                </div>
                <div class="text-xs text-clay-muted mt-1 font-medium">搜索词数</div>
              </div>

              <!-- 查询成功率 -->
              <div
                class="rounded-[20px] p-4 text-center transition-all duration-300 hover:scale-105 hover:-translate-y-1"
                style="
                  background: linear-gradient(135deg, rgba(16,185,129,0.08) 0%, rgba(16,185,129,0.15) 100%);
                  box-shadow: 8px 8px 16px rgba(16,185,129,0.08), -6px -6px 12px rgba(255,255,255,0.9), inset 3px 3px 6px rgba(255,255,255,0.6), inset -3px -3px 6px rgba(16,185,129,0.04);
                "
              >
                <div class="text-2xl sm:text-3xl font-black text-emerald-600 font-heading">
                  {{ resultRate }}
                </div>
                <div class="text-xs text-clay-muted mt-1 font-medium">有结果率</div>
              </div>

              <!-- 高峰时段 -->
              <div
                class="rounded-[20px] p-4 text-center transition-all duration-300 hover:scale-105 hover:-translate-y-1"
                style="
                  background: linear-gradient(135deg, rgba(245,158,11,0.08) 0%, rgba(245,158,11,0.15) 100%);
                  box-shadow: 8px 8px 16px rgba(245,158,11,0.08), -6px -6px 12px rgba(255,255,255,0.9), inset 3px 3px 6px rgba(255,255,255,0.6), inset -3px -3px 6px rgba(245,158,11,0.04);
                "
              >
                <div class="text-2xl sm:text-3xl font-black text-amber-600 font-heading">
                  {{ peakHour }}
                </div>
                <div class="text-xs text-clay-muted mt-1 font-medium">高峰时段</div>
              </div>
            </div>

            <!-- 二级指标 -->
            <div class="grid grid-cols-3 gap-3 mt-4">
              <div class="rounded-[16px] p-3 text-center" style="background: rgba(255,255,255,0.5); box-shadow: 4px 4px 8px rgba(136,79,34,0.04), -3px -3px 6px rgba(255,255,255,0.8);">
                <div class="text-lg font-bold text-primary font-heading">{{ overview.today_count || 0 }}</div>
                <div class="text-xs text-clay-muted font-medium">今日</div>
              </div>
              <div class="rounded-[16px] p-3 text-center" style="background: rgba(255,255,255,0.5); box-shadow: 4px 4px 8px rgba(136,79,34,0.04), -3px -3px 6px rgba(255,255,255,0.8);">
                <div class="text-lg font-bold text-emerald-600 font-heading">{{ overview.week_count || 0 }}</div>
                <div class="text-xs text-clay-muted font-medium">本周</div>
              </div>
              <div class="rounded-[16px] p-3 text-center" style="background: rgba(255,255,255,0.5); box-shadow: 4px 4px 8px rgba(136,79,34,0.04), -3px -3px 6px rgba(255,255,255,0.8);">
                <div class="text-lg font-bold text-sky-600 font-heading">{{ overview.month_count || 0 }}</div>
                <div class="text-xs text-clay-muted font-medium">本月</div>
              </div>
            </div>
          </div>
        </div>

        <!-- 查询次数趋势 -->
        <div class="clay-card p-6">
          <div class="relative z-10">
            <h3 class="text-base font-bold text-clay-foreground mb-4 font-heading">
              查询次数趋势
            </h3>
            <div ref="trendChartRef" class="w-full" style="height: 280px;"></div>
          </div>
        </div>

        <!-- 搜索词排行榜 -->
        <div class="clay-card p-6">
          <div class="relative z-10">
            <h3 class="text-base font-bold text-clay-foreground mb-4 font-heading">
              搜索词排行榜
            </h3>
            <div v-if="data.top_keywords && data.top_keywords.length > 0">
              <div ref="keywordChartRef" class="w-full" style="height: 300px;"></div>
            </div>
            <div v-else class="text-center py-8 text-clay-muted font-medium">
              暂无搜索数据
            </div>
          </div>
        </div>

        <!-- 节次分布 + 高峰时段 (并排) -->
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-5">
          <!-- 节次分布 -->
          <div class="clay-card p-6">
            <div class="relative z-10">
              <h3 class="text-base font-bold text-clay-foreground mb-4 font-heading">
                节次分布
              </h3>
              <div v-if="data.node_dist && data.node_dist.length > 0">
                <div ref="nodeChartRef" class="w-full" style="height: 260px;"></div>
              </div>
              <div v-else class="text-center py-8 text-clay-muted font-medium">
                暂无节次数据
              </div>
            </div>
          </div>

          <!-- 查询结果分布 -->
          <div class="clay-card p-6">
            <div class="relative z-10">
              <h3 class="text-base font-bold text-clay-foreground mb-4 font-heading">
                结果数量分布
              </h3>
              <div v-if="data.result_stats">
                <div ref="resultChartRef" class="w-full" style="height: 260px;"></div>
                <!-- 结果摘要 -->
                <div class="grid grid-cols-3 gap-2 mt-3">
                  <div class="text-center">
                    <div class="text-sm font-bold text-primary font-heading">
                      {{ data.result_stats.avg_count?.toFixed(1) || '0' }}
                    </div>
                    <div class="text-xs text-clay-muted">平均结果数</div>
                  </div>
                  <div class="text-center">
                    <div class="text-sm font-bold text-emerald-600 font-heading">
                      {{ data.result_stats.max_count || 0 }}
                    </div>
                    <div class="text-xs text-clay-muted">最多结果</div>
                  </div>
                  <div class="text-center">
                    <div class="text-sm font-bold text-sky-600 font-heading">
                      {{ data.result_stats.non_zero_count || 0 }}
                    </div>
                    <div class="text-xs text-clay-muted">有效查询</div>
                  </div>
                </div>
              </div>
              <div v-else class="text-center py-8 text-clay-muted font-medium">
                暂无结果数据
              </div>
            </div>
          </div>
        </div>

        <!-- 高峰时段分析 -->
        <div class="clay-card p-6">
          <div class="relative z-10">
            <h3 class="text-base font-bold text-clay-foreground mb-4 font-heading">
              高峰时段分析
              <span class="text-xs font-medium text-clay-muted ml-2">24小时查询分布</span>
            </h3>
            <div ref="hourlyChartRef" class="w-full" style="height: 260px;"></div>
            <!-- 高峰时段提示 -->
            <div v-if="data.hourly_dist" class="flex items-center justify-center gap-4 mt-3 text-xs text-clay-muted">
              <span class="flex items-center gap-1">
                <span class="inline-block w-3 h-3 rounded-sm" style="background: #10B981;"></span>
                低峰
              </span>
              <span class="flex items-center gap-1">
                <span class="inline-block w-3 h-3 rounded-sm" style="background: #F59E0B;"></span>
                中峰
              </span>
              <span class="flex items-center gap-1">
                <span class="inline-block w-3 h-3 rounded-sm" style="background: #F43F5E;"></span>
                高峰
              </span>
            </div>
          </div>
        </div>
      </template>

      <AppFooter />
    </main>
  </div>
</template>
