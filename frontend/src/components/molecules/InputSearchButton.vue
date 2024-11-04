<script setup lang="ts">
import Button from '@/components/atoms/Button.vue'
import InputText from '@/components/atoms/InputText.vue'
import { computed } from 'vue'

const props = defineProps({
  modelValue: {
    type: String,
    required: true
  },
  placeholder: {
    type: String,
    default: null
  },
  buttonTitle: {
    type: String,
    default: 'Search'
  },
  buttonDisabled: {
    type: Boolean,
    default: false
  }
})

const emits = defineEmits(['update:modelValue', 'search', 'submit'])

const value = computed({
  get: () => props.modelValue,
  set: (v) => {
    emits('update:modelValue', v)
  }
})
</script>

<template>
  <div class="flex items-center">
    <InputText class="w-full" v-model="value" :placeholder="placeholder" @submit="$emit('submit', value)" />
    <Button class="ml-2" :disabled="buttonDisabled" @click="$emit('search', value)">{{ buttonTitle }}</Button>
  </div>
</template>
