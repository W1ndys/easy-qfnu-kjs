import { reactive } from 'vue'

export function useAlertDialog() {
  const alertState = reactive({
    open: false,
    title: '提示',
    message: '',
    buttonText: '我知道了',
  })

  function showAlert(message, options = {}) {
    alertState.title = options.title || '提示'
    alertState.message = message
    alertState.buttonText = options.buttonText || '我知道了'
    alertState.open = true
  }

  function closeAlert() {
    alertState.open = false
  }

  return {
    alertState,
    showAlert,
    closeAlert,
  }
}
