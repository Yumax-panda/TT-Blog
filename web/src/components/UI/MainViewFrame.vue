<template>
  <div :class="$style.wrapper">
    <main-header />
    <div :class="$style.container">
      <slot></slot>
    </div>
    <main-footer />
  </div>
</template>

<script lang="ts" setup>
import MainFooter from './MainFooter.vue'
import MainHeader from './MainHeader.vue'
</script>

<style lang="scss" module>
.wrapper {
  background-image: url('/@/assets/background.webp');
  background-repeat: repeat;
}
.container {
  min-height: 100vh;
}
</style>
