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
    <label class="block text-sm font-bold text-clay-muted mb-2 ml-1">日期</label>

    <!-- Clay segmented control -->
    <div
      class="p-1.5 rounded-[20px] flex mb-3"
      style="
        background: #F0EAE3;
        box-shadow:
          inset 6px 6px 12px rgba(136, 79, 34, 0.06),
          inset -6px -6px 12px rgba(255, 255, 255, 0.9);
      "
    >
      <button
        v-for="(label, idx) in quickDateLabels"
        :key="idx"
        type="button"
        class="flex-1 py-2 text-[13px] font-bold rounded-[16px] transition-all duration-300"
        :class="dateOffset === idx && !useCustomDate
          ? 'text-clay-foreground'
          : 'text-clay-muted hover:text-clay-foreground'"
        :style="dateOffset === idx && !useCustomDate
          ? 'background: white; box-shadow: 8px 8px 16px rgba(136, 79, 34, 0.08), -6px -6px 12px rgba(255, 255, 255, 0.9), inset 2px 2px 4px rgba(255, 255, 255, 0.6), inset -2px -2px 4px rgba(0, 0, 0, 0.02);'
          : ''"
        @click="setQuickDate(idx)"
      >
        {{ label }}
      </button>

      <button
        type="button"
        class="flex-1 py-2 text-[13px] font-bold rounded-[16px] transition-all duration-300"
        :class="useCustomDate
          ? 'text-clay-foreground'
          : 'text-clay-muted hover:text-clay-foreground'"
        :style="useCustomDate
          ? 'background: white; box-shadow: 8px 8px 16px rgba(136, 79, 34, 0.08), -6px -6px 12px rgba(255, 255, 255, 0.9), inset 2px 2px 4px rgba(255, 255, 255, 0.6), inset -2px -2px 4px rgba(0, 0, 0, 0.02);'
          : ''"
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
      <p class="text-xs text-clay-muted ml-1 font-medium">{{ customDatePreview }}</p>
    </div>
  </div>
</template>
