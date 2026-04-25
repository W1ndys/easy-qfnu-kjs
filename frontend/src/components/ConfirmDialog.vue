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
      <div class="absolute inset-0 bg-[rgba(31,27,24,0.48)] backdrop-blur-[6px]" @click="emit('cancel')"></div>

      <div class="relative w-full max-w-md rounded-[20px] border border-subtle bg-white px-5 py-6 shadow-claySurface sm:px-6 sm:py-7">
        <div class="flex items-start gap-4">
          <div class="flex h-12 w-12 shrink-0 items-center justify-center rounded-2xl bg-primary text-white">
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
            class="h-12 rounded-xl border border-subtle bg-[#F3EFEB] text-sm font-bold text-clay-muted transition hover:bg-[#E5DED7]"
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
