<template>
  <div>
    <div :class="$style['title-wrapper']">
      <span :class="$style.title">最新情報</span>
    </div>
    <div :class="$style.content">
      <headline-list />
      <twitter-embed />
    </div>
  </div>
</template>

<script setup lang="ts">
import HeadlineList from './HeadlineList.vue'
import TwitterEmbed from './TwitterEmbed.vue'
</script>

<style lang="scss" module>
.title-wrapper {
  color: white;
  display: flex;

  @include size-h1;
  @include color-common-bg-green;
}

.title {
  margin: auto;
}

.content {
  display: flex;
  flex-wrap: wrap;
  padding: 1rem 0;
  @include color-common-bg-gray;
}
</style>
