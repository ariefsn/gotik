<script setup lang="ts">
import type { VideoItem } from '@/graphql/generated'
import { tiktokById } from '@/graphql/tiktok'
import { formatNumber } from '@/helper'
import type videojs from 'video.js'
import { computed, onMounted, ref, shallowRef } from 'vue'
import { useRoute, useRouter } from 'vue-router'

type VideoJsPlayer = ReturnType<typeof videojs>

const route = useRoute()
const router = useRouter()
const id = route.params.id
const url = computed(() => String(import.meta.env.VITE_API_URL ?? '').replace('/query', '/video') + '?url=' + data.value?.item?.video?.playAddrPublic)
const isLoading = ref(false)
const data = ref<VideoItem | null>(null)
const player = shallowRef<VideoJsPlayer>()

const icons = computed(() => [
  {
    icon: 'fc-like',
    value: formatNumber(data.value?.item?.stats?.diggCount ?? 0)
  },
  {
    icon: 'fc-film-reel',
    value: formatNumber(data.value?.item?.stats?.playCount ?? 0)
  },
  {
    icon: 'fc-comments',
    value: formatNumber(data.value?.item?.stats?.commentCount ?? 0)
  },
  {
    icon: 'fc-share',
    value: formatNumber(data.value?.item?.stats?.shareCount ?? 0)
  },
])

onMounted(async () => {
  isLoading.value = true
  try {
    const res = await tiktokById(id as string)
    const d = res.data?.getTiktok
    if (!d) {
      router.replace('/404')
      return
    }
    data.value = d
  } finally {
    isLoading.value = false
  }
})

const onPlayerMounted = (payload: any) => {
  player.value = payload.player
}

const onPlayerEvent = (log: any) => {
  console.log(log)
}
</script>

<template>
  <div v-if="data" class="w-full h-screen flex justify-center items-center">
    <div class="grid grid-cols-12 gap-4">
      <div class="col-span-12 md:col-span-2 lg:col-span-3"></div>
      <div class="col-span-12 md:col-span-8 lg:col-span-6">
        <VideoPlayer
          :src="url + '.mp4'"
          :poster="data.item?.video?.cover ?? data.item?.video?.originCover ?? data.item?.video?.dynamicCover ?? ''"
          playsinline
          controls
          class="w-full h-[600px]"
          @mounted="onPlayerMounted"
          @timeupdate="onPlayerEvent(player?.currentTime())"
        />

        <div class="border">
          <div class="p-3 border-b-2 border-gray-300">
            <div class="flex justify-start align-top">
              <div>
                <img
                  :src="data.item?.author?.avatarThumb ?? data.item?.author?.avatarMedium ?? data.item?.author?.avatarLarger ?? ''"
                  alt=""
                  srcset=""
                  class="w-24 rounded-full"
                >
              </div>
              <div class="ml-4">
                <div>{{ data.item?.author?.nickname }} (@{{ data.item?.author?.uniqueId }})</div>
                <div>Followers: {{ formatNumber(data.item?.authorStats?.followerCount ?? 0) }}</div>
                <div>Videos: {{ formatNumber(data.item?.authorStats?.videoCount ?? 0) }}</div>
                <!-- <div class="flex">
                  <v-icon name="fc-like" />
                  <span class="ml-1 text-sm">
                    {{ formatNumber(data.item?.authorStats?.diggCount ?? 0) }}
                  </span>
                </div> -->
                <div>{{ data.item?.author?.signature }}</div>
              </div>
            </div>
          </div>

          <div class="p-3">
            <div v-for="item in icons" :key="item.icon">
              <VIcon :name="item.icon" />
              <span class="ml-2">{{ item.value }}</span>
            </div>
            <div>{{ data.item?.desc }}</div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
