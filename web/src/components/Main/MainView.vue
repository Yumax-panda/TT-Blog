<template>
  <main-header />
  <div :class="$style.container">
    <welcome-section />
    <main-section />
  </div>
</template>

<script lang="ts" setup>
import MainSection from './MainSection/MainSection.vue'
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
