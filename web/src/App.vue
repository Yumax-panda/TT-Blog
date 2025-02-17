<template>
  <div :class="$style.app">
    <router-view v-slot="{ Component }">
      <component :is="Component" />
    </router-view>
  </div>
</template>

<style lang="scss" module>
.app {
  background-image: url('/@/assets/background.webp');
  background-repeat: repeat;
}
</style>
