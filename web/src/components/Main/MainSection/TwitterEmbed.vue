<template>
  <div :class="$style.embed">
    <blockquote
      ref="timelineRef"
      class="twitter-tweet"
      data-lang="ja"
      data-theme="dark"
    >
      <p lang="ja" dir="ltr">
        マリオカートに興味がある東工大生募集してます。興味ある方はぜひ！
      </p>
      &mdash; 東工大マリカサークル (@TokyoTech_MK)
      <a
        href="https://twitter.com/TokyoTech_MK/status/1387773813312704531?ref_src=twsrc%5Etfw"
        >2021年4月29日</a
      >
    </blockquote>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

const timelineRef = ref<HTMLElement>()

onMounted(() => {
  if (!timelineRef.value) return

  const script = document.createElement('script')
  script.src = 'https://platform.twitter.com/widgets.js'
  script.async = true
  timelineRef.value.appendChild(script)
})
</script>

<style lang="scss" module>
.embed {
  width: calc(100% - 1rem);
  margin: auto;
  margin-top: 0;

  @include mq(lg) {
    width: calc(50% - 1rem);
  }
}
</style>
