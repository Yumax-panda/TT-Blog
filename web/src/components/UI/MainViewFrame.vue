<template>
  <main-header />
  <div :class="$style.container">
    <slot></slot>
  </div>
  <main-footer />
</template>

<script lang="ts" setup>
import MainFooter from './MainFooter.vue'
import MainHeader from './MainHeader.vue'
</script>

<style lang="scss" module>
.container {
  min-height: 100vh;
}
</style>
