<template>
  <main-header />
  <div :class="$style.container"></div>
</template>

<script lang="ts" setup>
import MainHeader from './MainHeader.vue'
</script>

<style lang="scss" module>
.container {
  background-image: url('/@/assets/background.webp');
  background-repeat: repeat;
  min-height: 100vh;
}
</style>
