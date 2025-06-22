import { ref, type Ref } from 'vue'

export let error: Ref<Error | null> = ref(null)

export function fatalError(err: Error) {
    error.value = err
}

export function clearError() {
    error.value = null
}
