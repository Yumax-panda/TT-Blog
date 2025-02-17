<template>
  <div :class="$style.container">
    <latest-news />
    <section>
      <div :class="$style.title">投稿記事</div>
      <post-list :posts="posts" />
    </section>
  </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import LatestNews from './LatestNews/LatestNews.vue'
import PostList from '/@/components/UI/PostList/PostList.vue'

// TODO: implement
const posts = reactive([
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
.container {
  display: flex;
  flex-direction: column;
  gap: 38px;
  padding-top: 38px;
  padding-bottom: 56px;

  @include mq(lg) {
    gap: 80px;
    padding-top: 90px;
    padding-bottom: 128px;
  }

  & {
    @include container;
  }
}

.title {
  text-align: center;
  color: white;
  @include size-h1;
  @include color-common-bg-green;
}
</style>
