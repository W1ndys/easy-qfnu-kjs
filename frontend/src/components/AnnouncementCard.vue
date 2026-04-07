<script setup>
import { ref } from 'vue'
import { useAnnouncements } from '@/composables/useAnnouncements'

const { allAnnouncements, unreadCount, hasUnread, isRead, markAllAsRead } =
  useAnnouncements()

const expanded = ref(true)

function handleToggle() {
  if (expanded.value && hasUnread.value) {
    markAllAsRead()
  }
  expanded.value = !expanded.value
}
</script>

<template>
  <!-- 无公告时不渲染 -->
  <div v-if="allAnnouncements.length > 0">
    <!-- 折叠状态：仅显示摘要条 -->
    <div
      v-if="!expanded"
      class="clay-card cursor-pointer transition-all duration-300"
      :class="hasUnread ? 'p-4 sm:p-5' : 'p-3 sm:p-4'"
      @click="handleToggle"
    >
      <div class="relative z-10 flex items-center justify-between">
        <div class="flex items-center space-x-3">
          <!-- 公告图标 -->
          <div
            class="w-8 h-8 rounded-xl flex items-center justify-center flex-shrink-0"
            :style="{
              background: hasUnread
                ? 'linear-gradient(135deg, #F59E0B 0%, #D97706 100%)'
                : 'linear-gradient(135deg, #D4AD82 0%, #C4A882 100%)',
              boxShadow: hasUnread
                ? '4px 4px 8px rgba(245, 158, 11, 0.2), -2px -2px 4px rgba(255, 255, 255, 0.3), inset 1px 1px 2px rgba(255, 255, 255, 0.4)'
                : '4px 4px 8px rgba(136, 79, 34, 0.1), -2px -2px 4px rgba(255, 255, 255, 0.5)',
            }"
          >
            <svg
              class="w-4 h-4 text-white"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M11 5.882V19.24a1.76 1.76 0 01-3.417.592l-2.147-6.15M18 13a3 3 0 100-6M5.436 13.683A4.001 4.001 0 017 6h1.832c4.1 0 7.625-1.234 9.168-3v14c-1.543-1.766-5.067-3-9.168-3H7a3.988 3.988 0 01-1.564-.317z"
              />
            </svg>
          </div>
          <span class="text-sm font-bold text-clay-foreground">
            {{ hasUnread ? `${unreadCount} 条新公告` : '系统公告' }}
          </span>
        </div>
        <!-- 展开箭头 -->
        <div
          class="w-6 h-6 rounded-full flex items-center justify-center flex-shrink-0"
          style="
            background: rgba(255, 255, 255, 0.6);
            box-shadow:
              3px 3px 6px rgba(136, 79, 34, 0.06),
              -2px -2px 4px rgba(255, 255, 255, 0.8);
          "
        >
          <svg
            class="w-3 h-3 text-clay-muted"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2.5"
              d="M19 9l-7 7-7-7"
            />
          </svg>
        </div>
      </div>
    </div>

    <!-- 展开状态：完整公告列表 -->
    <div v-else class="clay-card p-6 sm:p-8">
      <div class="relative z-10">
        <!-- 标题栏 -->
        <div class="flex items-center justify-between mb-5">
          <div class="flex items-center space-x-3">
            <div
              class="w-10 h-10 rounded-2xl flex items-center justify-center flex-shrink-0"
              :style="{
                background: hasUnread
                  ? 'linear-gradient(135deg, #F59E0B 0%, #D97706 100%)'
                  : 'linear-gradient(135deg, #D4AD82 0%, #C4A882 100%)',
                boxShadow: hasUnread
                  ? '6px 6px 12px rgba(245, 158, 11, 0.2), -3px -3px 6px rgba(255, 255, 255, 0.3), inset 2px 2px 4px rgba(255, 255, 255, 0.4)'
                  : '6px 6px 12px rgba(136, 79, 34, 0.1), -3px -3px 6px rgba(255, 255, 255, 0.5)',
              }"
            >
              <svg
                class="w-5 h-5 text-white"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M11 5.882V19.24a1.76 1.76 0 01-3.417.592l-2.147-6.15M18 13a3 3 0 100-6M5.436 13.683A4.001 4.001 0 017 6h1.832c4.1 0 7.625-1.234 9.168-3v14c-1.543-1.766-5.067-3-9.168-3H7a3.988 3.988 0 01-1.564-.317z"
                />
              </svg>
            </div>
            <h3
              class="text-base font-bold text-clay-foreground"
              style="font-family: 'Nunito', sans-serif"
            >
              系统公告
              <span
                v-if="hasUnread"
                class="ml-2 inline-flex items-center justify-center px-2 py-0.5 text-xs font-bold text-white rounded-full"
                style="
                  background: linear-gradient(135deg, #F59E0B, #D97706);
                  box-shadow: 2px 2px 4px rgba(245, 158, 11, 0.3);
                "
              >
                {{ unreadCount }} 条未读
              </span>
            </h3>
          </div>
          <!-- 收起按钮 -->
          <button
            class="w-8 h-8 rounded-full flex items-center justify-center flex-shrink-0 transition-all duration-200 hover:scale-110"
            style="
              background: rgba(255, 255, 255, 0.6);
              box-shadow:
                4px 4px 8px rgba(136, 79, 34, 0.06),
                -3px -3px 6px rgba(255, 255, 255, 0.8);
            "
            title="收起公告"
            @click="handleToggle"
          >
            <svg
              class="w-4 h-4 text-clay-muted"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2.5"
                d="M5 15l7-7 7 7"
              />
            </svg>
          </button>
        </div>

        <!-- 公告列表 -->
        <div class="space-y-3">
          <div
            v-for="item in allAnnouncements"
            :key="item.id"
            class="rounded-2xl p-4 transition-all duration-300"
            :style="{
              background: item.important
                ? 'linear-gradient(135deg, rgba(245, 158, 11, 0.08) 0%, rgba(245, 158, 11, 0.15) 100%)'
                : 'linear-gradient(135deg, rgba(136, 79, 34, 0.04) 0%, rgba(136, 79, 34, 0.08) 100%)',
              boxShadow: [
                '6px 6px 12px rgba(136, 79, 34, 0.05)',
                '-4px -4px 8px rgba(255, 255, 255, 0.8)',
                'inset 2px 2px 4px rgba(255, 255, 255, 0.5)',
                'inset -2px -2px 4px rgba(136, 79, 34, 0.02)',
              ].join(', '),
              borderLeft: item.important
                ? '3px solid #F59E0B'
                : !isRead(item.id)
                  ? '3px solid #10B981'
                  : '3px solid transparent',
            }"
          >
            <div class="flex items-start justify-between gap-2">
              <h4
                class="text-sm font-bold"
                :class="
                  isRead(item.id)
                    ? 'text-clay-muted'
                    : 'text-clay-foreground'
                "
              >
                {{ item.title }}
              </h4>
              <span
                class="text-xs text-clay-muted flex-shrink-0 font-medium mt-0.5"
              >
                {{ item.date }}
              </span>
            </div>
            <p
              class="text-sm mt-2 leading-relaxed font-medium"
              :class="
                isRead(item.id) ? 'text-clay-muted/70' : 'text-clay-muted'
              "
            >
              {{ item.content }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
