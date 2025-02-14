<template>
  <div>
    <div :class="$style['title-wrapper']">
      <span :class="$style.title">最新情報</span>
    </div>
    <div :class="$style.content">
      <latest-news-list />
      <twitter-embed />
    </div>
  </div>
</template>

<script setup lang="ts">
import LatestNewsList from './LatestNewsList.vue'
import TwitterEmbed from './TwitterEmbed.vue'
</script>

<style lang="scss" module>
.title-wrapper {
  color: white;
  background-color: rgba($color: #00803e, $alpha: 0.4);
  display: flex;

  @include size-h1;
}

.title {
  margin: auto;
}

.content {
  background-color: rgba($color: #738da5, $alpha: 0.4);
  display: flex;
  flex-wrap: wrap;
  padding: 1rem 0;
}
</style>
