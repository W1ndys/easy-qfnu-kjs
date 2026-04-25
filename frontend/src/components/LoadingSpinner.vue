<script setup>
const props = defineProps({
  text: {
    type: String,
    default: '加载中...',
  },
})
</script>

<template>
  <div class="clay-card p-8 text-center">
    <div class="relative z-10 flex flex-col items-center">
      <div class="mb-4 h-9 w-9 rounded-full border-2 border-primary-100 border-t-primary animate-spin"></div>
      <p class="font-medium text-clay-muted">{{ props.text }}</p>
    </div>
  </div>
</template>
