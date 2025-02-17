<template>
  <router-link :to="constructPostDetailPagePath(postId)" :class="$style.post">
    <img
      :src="thumbnailUrl"
      :alt="`${title} feature image`"
      :class="$style.thumbnail"
    />
    <div :class="$style.meta">
      <small>{{ getFullDayString(datetime) }}</small>
      <br />
      <strong :class="$style.title">{{ title }}</strong>
    </div>
  </router-link>
</template>

<script lang="ts" setup>
import { constructPostDetailPagePath } from '/@/router'
import { getFullDayString } from '/@/lib/basic/date'

defineProps<{
  postId: string
  title: string
  thumbnailUrl: string
  datetime: Readonly<Date>
}>()
</script>

<style lang="scss" module>
.post {
  background-position: 50%;
  background-size: cover;
  flex: 1 0 auto;
  position: relative;
  height: 50vh;
  max-height: 20em;
  text-decoration: none;
  width: 100%;

  @include mq(sm) {
    width: calc(33.33333% - 3px);
  }
}

.thumbnail {
  height: 100%;
  width: 100%;
  object-fit: cover;
}

.meta {
  position: absolute;
  bottom: 0;
  width: 100%;
  padding: 1em;
  background-color: rgba($color: rgb(51, 51, 51), $alpha: 0.8);

  @include color-common-text-white-primary;
}

.title {
  @include size-h3;
}
</style>
