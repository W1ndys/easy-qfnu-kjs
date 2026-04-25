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
      containerClass: 'border-[#F5B3AE] bg-[#FDEEEE]',
      iconClass: 'bg-[#B42318] text-white',
      titleClass: 'text-[#B42318]',
      messageClass: 'text-[#7A271A]',
    }
  }

  return {
    containerClass: 'border-[#F3CF8D] bg-[#FFF6E8]',
    iconClass: 'bg-[#9A5A00] text-white',
    titleClass: 'text-[#9A5A00]',
    messageClass: 'text-[#6E4000]',
  }
})
</script>

<template>
  <div
    :class="['rounded-2xl border p-4 shadow-clayCard', styles.containerClass]"
  >
    <div class="flex items-start gap-3">
      <div
        :class="['flex h-9 w-9 flex-shrink-0 items-center justify-center rounded-xl', styles.iconClass]"
      >
        <svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"
          />
        </svg>
      </div>
      <div>
        <h3 :class="['font-semibold', styles.titleClass]">
          {{ props.title }}
        </h3>
        <p :class="['text-sm mt-1 font-medium leading-relaxed', styles.messageClass]">
          {{ props.message }}
        </p>
      </div>
    </div>
  </div>
</template>
