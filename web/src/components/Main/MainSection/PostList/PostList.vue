<template>
  <section>
    <div :class="$style.title">投稿記事</div>
    <div :class="$style['post-list']">
      <template v-for="post in posts" :key="post.postId">
        <post-panel v-bind="post" />
      </template>
    </div>
  </section>
</template>

<script lang="ts" setup>
import PostPanel from './PostPanel.vue'
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
.title {
  text-align: center;
  color: white;
  @include size-h1;
  @include color-common-bg-green;
}

.post-list {
  display: flex;
  flex-wrap: wrap;
  gap: 1px;
  @include color-common-bg-gray;
}
</style>
