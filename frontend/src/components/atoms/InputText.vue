<script setup lang="ts">
import { computed, defineProps } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    required: true
  },
  placeholder: {
    type: String,
    default: null
  },
})

const emits = defineEmits(['update:modelValue', 'submit'])

const value = computed({
  get: () => props.modelValue,
  set: (v) => {
    emits('update:modelValue', v)
  }
})

</script>

<template>
  <input v-model="value" :placeholder="placeholder" @keyup.enter="$emit('submit', value)" />
</template>

<style scoped>
input {
  @apply px-2 py-1 border border-gray-400 rounded hover:border-black;
}
</style>
