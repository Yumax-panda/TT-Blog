<template>
  <div :class="$style.container">
    <latest-news />
  </div>
</template>

<script lang="ts" setup>
import LatestNews from './LatestNews/LatestNews.vue'
</script>

<style lang="scss" module>
.container {
  display: flex;
  flex-direction: column;
  gap: 38px;
  padding: 38px 24px 56px;
  @include mq(lg) {
    gap: 80px;
    padding: 90px 128px;
  }
}
</style>
