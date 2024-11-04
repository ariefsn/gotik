import './assets/main.css'

import { createPinia } from 'pinia'
import { createApp } from 'vue'

import Vue3VideoPlayer from '@cloudgeek/vue3-video-player'
import '@cloudgeek/vue3-video-player/dist/vue3-video-player.css'
import { VideoPlayer } from '@videojs-player/vue'
import { MotionPlugin } from '@vueuse/motion'
import { OhVueIcon, addIcons } from 'oh-vue-icons'
import { FcComments, FcFilmReel, FcLike, FcShare } from 'oh-vue-icons/icons'
import 'video.js/dist/video-js.css'
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
app.use(Vue3VideoPlayer, {
  lang: 'en'
})
app.use(MotionPlugin)
app.component('VIcon', OhVueIcon)
app.component('VideoPlayer', VideoPlayer)
addIcons(FcLike, FcShare, FcFilmReel, FcComments)

app.mount('#app')
