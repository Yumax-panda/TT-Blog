<template>
  <router-link
    :class="$style['panel-row']"
    :to="constructPostDetailPath(postId)"
  >
    <img
      :src="thumbnailUrl"
      :alt="`${title} feature image`"
      :class="$style.thumbnail"
    />
    <div :class="$style['news-text-container']">
      <small>{{ getFullDayString(datetime) }}</small>
      <span>{{ title }}</span>
    </div>
  </router-link>
</template>

<script setup lang="ts">
import { getFullDayString } from '/@/lib/basic/date'
import { constructPostDetailPath } from '/@/router'

defineProps<{
  postId: string
  thumbnailUrl: string
  title: string
  datetime: Readonly<Date>
}>()
</script>

<style lang="scss" module>
.panel-row {
  display: flex;
  flex-direction: row;
  text-decoration: none;
  border-top: solid 1px hsl(0, 0%, 87%);
  padding: 0.5em 1em;
  transition: background-color 0.25s;
  @include color-common-text-white-primary;
}

.panel-row:hover {
  background-color: rgba(255, 255, 255, 0.15);
}

.panel-row:last-child {
  border-bottom: solid 1px #dddddd;
}

.panel-row > div > span {
  font-weight: bold;
}

.thumbnail {
  width: 80px;
  height: 80px;
}

.news-text-container {
  display: flex;
  flex-direction: column;
  margin-left: 20px;
  justify-content: center;
}
</style>
