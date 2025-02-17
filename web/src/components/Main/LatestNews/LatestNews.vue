<template>
  <section>
    <div :class="$style.title">最新情報</div>
    <div :class="$style.content">
      <headline-list />
      <twitter-embed />
    </div>
  </section>
</template>

<script setup lang="ts">
import HeadlineList from './HeadlineList.vue'
import TwitterEmbed from './TwitterEmbed.vue'
</script>

<style lang="scss" module>
.title {
  color: white;
  text-align: center;

  @include size-h1;
  @include color-common-bg-green;
}

.content {
  display: flex;
  flex-wrap: wrap;
  padding: 1rem 0;
  @include color-common-bg-gray;
}
</style>
