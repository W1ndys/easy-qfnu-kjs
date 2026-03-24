<script setup>
import { computed } from 'vue'

const props = defineProps({
  type: {
    type: String,
    default: 'warning',
    validator: (value) => ['error', 'warning'].includes(value),
  },
  title: {
    type: String,
    required: true,
  },
  message: {
    type: String,
    required: true,
  },
})

const styles = computed(() => {
  if (props.type === 'error') {
    return {
      bg: 'linear-gradient(135deg, rgba(239, 68, 68, 0.06) 0%, rgba(239, 68, 68, 0.12) 100%)',
      shadow: '12px 12px 24px rgba(239, 68, 68, 0.08), -8px -8px 16px rgba(255, 255, 255, 0.9), inset 4px 4px 8px rgba(255, 255, 255, 0.5), inset -4px -4px 8px rgba(239, 68, 68, 0.04)',
      iconBg: 'linear-gradient(135deg, #FCA5A5 0%, #EF4444 100%)',
      iconShadow: '6px 6px 12px rgba(239, 68, 68, 0.2), -3px -3px 6px rgba(255, 255, 255, 0.3), inset 2px 2px 4px rgba(255, 255, 255, 0.4), inset -2px -2px 4px rgba(0, 0, 0, 0.05)',
      titleClass: 'text-red-800',
      messageClass: 'text-red-600',
    }
  }

  return {
    bg: 'linear-gradient(135deg, rgba(245, 158, 11, 0.06) 0%, rgba(245, 158, 11, 0.12) 100%)',
    shadow: '12px 12px 24px rgba(245, 158, 11, 0.08), -8px -8px 16px rgba(255, 255, 255, 0.9), inset 4px 4px 8px rgba(255, 255, 255, 0.5), inset -4px -4px 8px rgba(245, 158, 11, 0.04)',
    iconBg: 'linear-gradient(135deg, #FCD34D 0%, #F59E0B 100%)',
    iconShadow: '6px 6px 12px rgba(245, 158, 11, 0.2), -3px -3px 6px rgba(255, 255, 255, 0.3), inset 2px 2px 4px rgba(255, 255, 255, 0.4), inset -2px -2px 4px rgba(0, 0, 0, 0.05)',
    titleClass: 'text-amber-800',
    messageClass: 'text-amber-700',
  }
})
</script>

<template>
  <div
    class="rounded-[24px] p-5"
    :style="{
      background: styles.bg,
      boxShadow: styles.shadow,
    }"
  >
    <div class="flex items-center space-x-4">
      <div
        class="w-10 h-10 rounded-2xl flex items-center justify-center flex-shrink-0"
        :style="{
          background: styles.iconBg,
          boxShadow: styles.iconShadow,
        }"
      >
        <svg class="w-5 h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
          />
        </svg>
      </div>
      <div>
        <h3 :class="['font-bold', styles.titleClass]" style="font-family: 'Nunito', sans-serif;">
          {{ props.title }}
        </h3>
        <p :class="['text-sm mt-1 font-medium leading-relaxed', styles.messageClass]">
          {{ props.message }}
        </p>
      </div>
    </div>
  </div>
</template>
