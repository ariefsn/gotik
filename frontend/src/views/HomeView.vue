<script setup lang="ts">
import H2 from '@/components/atoms/H2.vue'
import InputSearchButton from '@/components/molecules/InputSearchButton.vue'
import type { VideoItem } from '@/graphql/generated'
import { tiktokSearch, tiktokSearchSubscription } from '@/graphql/tiktok'
import { defineAsyncComponent, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import type { Subscription } from 'wonka'

const router = useRouter()

const keyword = ref('')
const tiktokSub = ref<Subscription | null>(null)
const videos = ref<VideoItem[]>([])
const isLoading = ref(false)

onMounted(() => {
  const q = router.currentRoute.value.query.q
  if (q) {
    keyword.value = q as string
    onSearch()
  }
})

const onSubscribe = () => {
  tiktokSub.value?.unsubscribe()

  tiktokSub.value = tiktokSearchSubscription(keyword.value).subscribe(res => {
    const allIds = videos.value.map(v => v.common?.doc_id_str)
    const newVideos = res?.data?.subSearchTiktok || []
    const tmp: VideoItem[] = []
    newVideos.forEach(v => {
      if (!allIds.includes(v.common?.doc_id_str)) {
        tmp.push(v)
      }
    })
    videos.value = [...videos.value, ...tmp]
    if (tmp.length > 0) {
      isLoading.value = false
    }
  })
}

const onSearch = async () => {
  videos.value = []

  if (!keyword.value) {
    isLoading.value = false
    return
  }

  onSubscribe()

  router.push({
    query: {
      q: keyword.value
    }
  })

  isLoading.value = true
  const res = await tiktokSearch(keyword.value)
  videos.value = [...(res.data?.searchTiktok || [])]

  isLoading.value = videos.value.length === 0
}

const LazyCardVideo = defineAsyncComponent(() => import('@/components/molecules/CardVideo.vue'))

</script>

<template>
  <div class="p-4 relative min-h-screen flex">
    <div class="fixed z-20 w-full top-0 left-0 bg-white">
      <div class="flex w-full justify-center">
        <div class="container p-4">
          <InputSearchButton
            v-model="keyword"
            placeholder="Search here..."
            :button-disabled="!keyword"
            @search="onSearch"
            @submit="onSearch"
          />
        </div>
      </div>
    </div>

    <div v-if="videos.length === 0 && !isLoading" class="flex w-full mt-28 justify-center items-center">
      <H2>
        Your video will appear here.
      </H2>
    </div>

    <div v-else-if="isLoading" class="flex w-full mt-28 justify-center items-center">
      <H2>
        Please wait...
      </H2>
    </div>

    <div v-else class="grid grid-cols-12 gap-4 mt-12 pt-3">
      <RouterLink
        v-for="(video, i) in videos"
        :key="i"
        class="col-span-12 sm:col-span-6 md:col-span-4 lg:col-span-3 hover:rotate-3"
        :to="'/video/' + video.item?.id"
        target="_blank"
      >
        <LazyCardVideo :video="video" v-motion-pop-visible-once />
      </RouterLink>
    </div>
  </div>
</template>
