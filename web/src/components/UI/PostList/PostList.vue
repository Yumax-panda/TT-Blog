<template>
  <div :class="$style['post-list']">
    <template v-for="post in posts" :key="post.postId">
      <post-panel v-bind="post" />
    </template>
  </div>
</template>

<script lang="ts" setup>
import PostPanel from './PostPanel.vue'

type Row = {
  title: string
  postId: string
  datetime: Readonly<Date>
  thumbnailUrl: string
}

defineProps<{ posts: Row[] }>()
</script>

<style lang="scss" module>
.post-list {
  display: flex;
  flex-wrap: wrap;
  gap: 1px;
  @include color-common-bg-gray;
}
</style>
