<script setup>
const props = defineProps({
  text: {
    type: String,
    default: '暂无数据',
  },
})
</script>

<template>
  <div class="py-12 flex flex-col items-center justify-center text-clay-muted">
    <div
      class="w-20 h-20 rounded-full flex items-center justify-center mb-5 animate-clay-breathe"
      style="
        background: rgba(255, 255, 255, 0.6);
        box-shadow:
          12px 12px 24px rgba(136, 79, 34, 0.06),
          -8px -8px 16px rgba(255, 255, 255, 0.9),
          inset 4px 4px 8px rgba(255, 255, 255, 0.6),
          inset -4px -4px 8px rgba(136, 79, 34, 0.03);
      "
    >
      <svg class="w-10 h-10 opacity-40" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="1.5"
          d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
        />
      </svg>
    </div>
    <p class="font-medium text-sm">{{ props.text }}</p>
  </div>
</template>
