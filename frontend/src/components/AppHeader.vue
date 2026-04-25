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
  <header class="sticky top-0 z-20 border-b border-subtle/80 bg-white/90 px-4 pt-[env(safe-area-inset-top)] backdrop-blur-xl">
    <div class="mx-auto max-w-5xl px-0 sm:px-2">
      <div class="flex h-16 items-center justify-between sm:h-[72px]">
        <button
          v-if="props.showBack"
          type="button"
          class="flex h-10 w-10 items-center justify-center rounded-[10px] border border-subtle bg-white text-clay-muted transition hover:border-primary-200 hover:bg-primary-50 hover:text-primary"
          @click="goBack"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M15 19l-7-7 7-7" />
          </svg>
        </button>
        <div v-else class="w-10" aria-hidden="true"></div>

        <h1
          class="absolute left-1/2 -translate-x-1/2 text-lg font-bold tracking-[-0.01em] text-clay-foreground font-heading"
        >
          {{ props.title }}
        </h1>
        <div class="w-10" aria-hidden="true"></div>
      </div>
    </div>
  </header>
</template>
