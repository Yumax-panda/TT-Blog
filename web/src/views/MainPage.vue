<template>
  <main-view-frame>
    <welcome-section />
    <main-section />
  </main-view-frame>
</template>

<script lang="ts" setup>
import MainViewFrame from '/@/components/Common/MainViewFrame.vue'
import MainSection from '/@/components/Main/MainSection/MainSection.vue'
import WelcomeSection from '/@/components/Main/WelcomeSection/WelcomeSection.vue'
</script>

<style lang="scss" module>
.container {
  min-height: 100vh;
}
</style>
