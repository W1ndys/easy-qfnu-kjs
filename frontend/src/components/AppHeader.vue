<script setup>
import { useRouter } from 'vue-router'

const props = defineProps({
  title: {
    type: String,
    required: true,
  },
  showBack: {
    type: Boolean,
    default: false,
  },
})

const router = useRouter()

function goBack() {
  if (window.history.length > 1) {
    router.back()
    return
  }
  router.push('/')
}
</script>

<template>
  <header class="sticky top-0 z-50 px-4 pt-[env(safe-area-inset-top)]">
    <div
      class="mx-auto mt-2 max-w-xl rounded-[32px] sm:rounded-[40px] px-4 sm:px-8"
      style="
        background: rgba(255, 255, 255, 0.65);
        backdrop-filter: blur(24px);
        -webkit-backdrop-filter: blur(24px);
        box-shadow:
          16px 16px 32px rgba(136, 79, 34, 0.06),
          -10px -10px 24px rgba(255, 255, 255, 0.9),
          inset 6px 6px 12px rgba(136, 79, 34, 0.02),
          inset -6px -6px 12px rgba(255, 255, 255, 1);
      "
    >
      <div class="h-14 sm:h-16 flex items-center justify-between relative">
        <button
          v-if="props.showBack"
          type="button"
          class="w-10 h-10 rounded-[20px] flex items-center justify-center text-clay-muted hover:text-clay-foreground transition-all duration-200 hover:-translate-y-0.5"
          style="
            background: rgba(255, 255, 255, 0.6);
            box-shadow:
              6px 6px 12px rgba(136, 79, 34, 0.06),
              -4px -4px 8px rgba(255, 255, 255, 0.8),
              inset 2px 2px 4px rgba(255, 255, 255, 0.6),
              inset -2px -2px 4px rgba(0, 0, 0, 0.03);
          "
          @click="goBack"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M15 19l-7-7 7-7" />
          </svg>
        </button>
        <div v-else class="w-10" aria-hidden="true"></div>

        <h1
          class="text-lg font-extrabold absolute left-1/2 -translate-x-1/2 text-clay-foreground font-heading"
          style="font-family: 'Nunito', sans-serif;"
        >
          {{ props.title }}
        </h1>
        <div class="w-10" aria-hidden="true"></div>
      </div>
    </div>
  </header>
</template>
