<template>
  <header :class="$style.container">
    <div :class="$style.inner">
      <img :src="egg" />
      <span>東工大マリオカートサークル</span>
    </div>
  </header>
</template>

<script lang="ts" setup>
import egg from '/@/assets/yoshi_egg.webp'
</script>

<style lang="scss" module scoped>
$header-background-color: #2f6665;

.container {
  font-weight: bold;
  width: 100%;
  top: 0px;
  position: sticky;
  display: flex;
  background-color: $header-background-color;
  height: 44px;

  @include color-common-text-white-primary;
  @include mq(md) {
    height: 80px;
    font-size: 24px;
  }
}

.inner {
  margin: auto;
  margin-left: 12px;
  display: flex;
  height: 24px;
  line-height: 24px;
  @include mq(md) {
    height: 40px;
    line-height: 40px;
  }
}
</style>
