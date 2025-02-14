<template>
  <div :class="$style.panel">
    <template v-for="post in posts" :key="post.postId">
      <latest-news-panel-row v-bind="post" />
    </template>
  </div>
</template>

<script lang="ts" setup>
import LatestNewsPanelRow from './LatestNewsPanelRow.vue'
import { ref } from 'vue'

type Row = {
  title: string
  postId: string
  datetime: Readonly<Date>
  thumbnailUrl: string
}

// TODO: implement
const posts = ref<Row[]>([
  {
    postId: '1',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '2',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '3',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  }
])
</script>

<style lang="scss" module>
.panel {
  display: flex;
  margin: auto;
  flex-direction: column;
  width: 100%;
  padding-left: 0.5rem;

  @include mq(lg) {
    width: 50%;
  }
}
</style>
