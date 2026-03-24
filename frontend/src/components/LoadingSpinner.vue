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
      <div
        class="w-14 h-14 rounded-full flex items-center justify-center mb-4 animate-clay-breathe"
        style="
          background: linear-gradient(135deg, #C4956A 0%, #884F22 100%);
          box-shadow:
            8px 8px 16px rgba(136, 79, 34, 0.2),
            -4px -4px 8px rgba(255, 255, 255, 0.3),
            inset 2px 2px 4px rgba(255, 255, 255, 0.4),
            inset -2px -2px 4px rgba(0, 0, 0, 0.05);
        "
      >
        <svg class="animate-spin h-7 w-7 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
          <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
          <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z"></path>
        </svg>
      </div>
      <p class="text-clay-muted font-medium">{{ props.text }}</p>
    </div>
  </div>
</template>
