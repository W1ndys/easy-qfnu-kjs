<script setup>
const props = defineProps({
  open: {
    type: Boolean,
    default: false,
  },
  title: {
    type: String,
    required: true,
  },
  message: {
    type: String,
    required: true,
  },
  confirmText: {
    type: String,
    default: '确定',
  },
  cancelText: {
    type: String,
    default: '取消',
  },
  showCancel: {
    type: Boolean,
    default: true,
  },
})

const emit = defineEmits(['confirm', 'cancel'])
</script>

<template>
  <teleport to="body">
    <div v-if="props.open" class="fixed inset-0 z-[100] flex items-center justify-center p-4 sm:p-6">
      <div
        class="absolute inset-0 bg-[rgba(51,42,36,0.42)] backdrop-blur-[6px]"
        @click="emit('cancel')"
      ></div>

      <div class="relative w-full max-w-md rounded-[32px] px-5 py-6 sm:px-6 sm:py-7" style="background: rgba(255, 255, 255, 0.92); backdrop-filter: blur(24px); box-shadow: 20px 20px 40px rgba(136, 79, 34, 0.12), -12px -12px 28px rgba(255, 255, 255, 0.9), inset 6px 6px 12px rgba(255, 255, 255, 0.6), inset -6px -6px 12px rgba(136, 79, 34, 0.03);">
        <div class="flex items-start gap-4">
          <div
            class="flex h-12 w-12 shrink-0 items-center justify-center rounded-[18px]"
            style="background: linear-gradient(135deg, #F7C9A7 0%, #C4956A 45%, #884F22 100%); box-shadow: 8px 8px 16px rgba(136, 79, 34, 0.18), -4px -4px 10px rgba(255, 255, 255, 0.45);"
          >
            <svg class="h-6 w-6 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>

          <div class="min-w-0 flex-1 pt-1">
            <h3 class="text-lg font-bold text-clay-foreground font-heading leading-7">
              {{ props.title }}
            </h3>
            <p class="mt-2 text-sm leading-7 font-medium text-clay-muted">
              {{ props.message }}
            </p>
          </div>
        </div>

        <div :class="props.showCancel ? 'mt-6 grid grid-cols-2 gap-3' : 'mt-6'">
          <button
            v-if="props.showCancel"
            type="button"
            class="h-12 rounded-[18px] text-sm font-bold text-clay-muted transition-all duration-200 hover:-translate-y-0.5"
            style="background: #F0EAE3; box-shadow: inset 8px 8px 16px rgba(136, 79, 34, 0.05), inset -8px -8px 16px rgba(255, 255, 255, 0.95);"
            @click="emit('cancel')"
          >
            {{ props.cancelText }}
          </button>

          <button
            type="button"
            :class="props.showCancel ? 'btn-clay-primary h-12 text-sm' : 'btn-clay-primary h-12 w-full text-sm'"
            @click="emit('confirm')"
          >
            {{ props.confirmText }}
          </button>
        </div>
      </div>
    </div>
  </teleport>
</template>
