<script setup>
import { watch } from 'vue'
import { useDateSelection } from '@/composables/useDateSelection'

const props = defineProps({
  modelValue: {
    type: Number,
    default: 0,
  },
})

const emit = defineEmits(['update:modelValue'])

const {
  quickDateLabels,
  useCustomDate,
  customOffset,
  dateOffset,
  customDatePreview,
  setOffset,
  setQuickDate,
  toggleCustomDate,
  updateCustomOffset,
} = useDateSelection(props.modelValue)

watch(
  () => props.modelValue,
  (nextValue) => {
    if (nextValue !== dateOffset.value) {
      setOffset(nextValue)
    }
  },
)

watch(dateOffset, (nextValue) => {
  emit('update:modelValue', nextValue)
})

function handleCustomOffsetInput() {
  updateCustomOffset(customOffset.value)
}
</script>

<template>
  <div>
    <label class="mb-2 block text-sm font-semibold text-clay-muted">日期</label>

    <div class="mb-3 flex rounded-xl bg-[#F3EFEB] p-1">
      <button
        v-for="(label, idx) in quickDateLabels"
        :key="idx"
        type="button"
        class="min-h-10 flex-1 rounded-lg px-2 py-2 text-[13px] font-semibold transition"
        :class="dateOffset === idx && !useCustomDate
          ? 'bg-white text-primary shadow-sm'
          : 'text-clay-muted hover:text-clay-foreground'"
        @click="setQuickDate(idx)"
      >
        {{ label }}
      </button>

      <button
        type="button"
        class="min-h-10 flex-1 rounded-lg px-2 py-2 text-[13px] font-semibold transition"
        :class="useCustomDate
          ? 'bg-white text-primary shadow-sm'
          : 'text-clay-muted hover:text-clay-foreground'"
        @click="toggleCustomDate"
      >
        自定义
      </button>
    </div>

    <!-- Custom date input -->
    <div v-if="useCustomDate" class="space-y-2">
      <div class="flex items-center space-x-3">
        <div class="flex-1">
          <input
            v-model.number="customOffset"
            type="number"
            min="0"
            max="180"
            class="w-full clay-input py-3 px-5 text-[15px] text-clay-foreground"
            placeholder="输入天数"
            @input="handleCustomOffsetInput"
          />
        </div>
        <span class="text-clay-muted text-sm font-medium whitespace-nowrap">天后</span>
      </div>
      <p class="ml-1 text-xs font-medium text-clay-muted">{{ customDatePreview }}</p>
    </div>
  </div>
</template>
