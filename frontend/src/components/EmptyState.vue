<script setup>
const props = defineProps({
  text: {
    type: String,
    default: '暂无数据',
  },
})
</script>

<template>
  <div class="design-card flex flex-col items-center justify-center px-6 py-12 text-center text-clay-muted">
    <div class="mb-5 flex h-20 w-20 items-center justify-center rounded-2xl border border-subtle bg-[#FAF8F6]">
      <svg class="h-10 w-10 opacity-50" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="1.5"
          d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
        />
      </svg>
    </div>
    <p class="text-base font-semibold text-clay-foreground">{{ props.text }}</p>
    <p class="mt-1 text-sm text-clay-muted">试试更换节次、日期或教学楼名称。</p>
  </div>
</template>
