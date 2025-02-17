<template>
  <div :class="$style.container">
    <latest-news />
    <post-list />
  </div>
</template>

<script lang="ts" setup>
import LatestNews from './LatestNews/LatestNews.vue'
import PostList from './PostList/PostList.vue'
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
</style>
