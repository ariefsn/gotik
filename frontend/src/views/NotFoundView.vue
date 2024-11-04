<script setup lang="ts">
import H1 from '@/components/atoms/H1.vue'
</script>

<template>
  <div class="w-screen h-screen flex items-center justify-center">
    <div class="prose text-center">
      <H1>404</H1>
      <p>Oops, the page you visited does not exist.</p>
      <router-link to="/">Go Home</router-link>
    </div>
  </div>
</template>
