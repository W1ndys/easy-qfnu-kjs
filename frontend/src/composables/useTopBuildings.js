import { onMounted, ref } from 'vue'
import { getTopBuildings } from '@/api'

/**
 * 获取搜索排行前 N 的教学楼，用于快捷预选按钮
 */
export function useTopBuildings() {
  const topBuildings = ref([])
  const loaded = ref(false)

  async function fetchTopBuildings() {
    try {
      const data = await getTopBuildings()
      topBuildings.value = (data.buildings || []).map((item) => item.name)
      loaded.value = true
    } catch {
      topBuildings.value = []
      loaded.value = true
    }
  }

  onMounted(fetchTopBuildings)

  return { topBuildings, loaded }
}
