<script setup lang="ts">
defineEmits(['click'])
</script>

<template>
  <button @click="$emit('click')">
    <slot />
  </button>
</template>

<style scoped>
button {
  @apply px-4 py-1 border border-gray-400 rounded hover:bg-black hover:text-white disabled:bg-gray-500 disabled:text-white disabled:cursor-not-allowed;
}
</style>
