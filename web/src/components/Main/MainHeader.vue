<template>
  <header>
    <div :class="$style.container">
      <router-link :to="'/'" :class="$style.inner">
        <img :src="egg" alt="Yoshi's egg" />
        <span>東工大マリオカートサークル</span>
      </router-link>
    </div>
  </header>
</template>

<script lang="ts" setup>
import egg from '/@/assets/yoshi_egg.webp'
</script>

<style lang="scss" module scoped>
$header-background-color: #2f6665;

header {
  top: 0;
  position: sticky;
}

.container {
  font-weight: bold;
  width: 100%;
  display: flex;
  background-color: $header-background-color;
  height: 44px;

  @include color-common-text-white-primary;
  @include mq(md) {
    height: 80px;
    font-size: 24px;
  }
}

.container:hover {
  cursor: pointer;
}

.inner {
  margin: auto;
  margin-left: 12px;
  display: flex;
  gap: 5px;
  height: 24px;
  line-height: 24px;
  @include mq(md) {
    gap: 10px;
    height: 40px;
    line-height: 40px;
  }
}
</style>
