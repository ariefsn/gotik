<script setup lang="ts">
import type { VideoItem } from '@/graphql/generated'
import { formatNumber } from '@/helper'
import { computed, ref } from 'vue'

const props = defineProps({
  video: {
    type: Object as () => VideoItem,
    required: true
  }
})

const item = computed(() => props.video.item)

const icons = ref([
  {
    icon: 'fc-like',
    value: formatNumber(item.value?.stats?.diggCount ?? 0)
  },
  {
    icon: 'fc-film-reel',
    value: formatNumber(item.value?.stats?.playCount ?? 0)
  },
  {
    icon: 'fc-comments',
    value: formatNumber(item.value?.stats?.commentCount ?? 0)
  },
  {
    icon: 'fc-share',
    value: formatNumber(item.value?.stats?.shareCount ?? 0)
  }
])
</script>

<template>
  <div
    class="shadow-md rounded-md hover:cursor-pointer hover:shadow-lg relative h-[296px]"
  >
    <img
      :src="video.item?.video?.cover || video.item?.video?.originCover || video.item?.video?.dynamicCover || ''"
      class="w-full h-64 object-cover rounded-t-md"
    />
    <div class="m-2 truncate">
      {{ video.item?.desc }}
    </div>

    <div class="absolute top top-0 right-0 bg-black w-[108px] h-[100px] rounded-tr-md rounded-bl-md opacity-50" />
    <div class="absolute top-0 right-0 text-right rounded-tr-md text-white text-xs">
      <div class="p-2 rounded-tr-md">
        <div v-for="ic in icons" :key="ic.icon">
          <span class="mr-2">{{ ic.value }}</span>
          <VIcon :name="ic.icon" />
        </div>
      </div>
    </div>
  </div>
</template>
