<template>
  <main-header />
  <div :class="$style.container">
    <welcome-section />
    <main-section />
  </div>
  <main-footer />
</template>

<script lang="ts" setup>
import MainSection from './MainSection/MainSection.vue'
import MainFooter from './MainFooter.vue'
import MainHeader from './MainHeader.vue'
import WelcomeSection from './WelcomeSection.vue'
</script>

<style lang="scss" module>
.container {
  min-height: 100vh;
}
</style>
