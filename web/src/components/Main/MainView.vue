<template>
  <div :class="$style.container">
    <main-header />
    <welcome-section />
  </div>
</template>

<script lang="ts" setup>
import MainHeader from './MainHeader.vue'
import WelcomeSection from './WelcomeSection.vue'
</script>

<style lang="scss" module>
.container {
  background-image: url('/@/assets/background.webp');
  background-repeat: repeat;
  min-height: 100vh;
}
</style>
