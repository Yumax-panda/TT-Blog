<template>
  <div :class="$style.wrapper">
    <div :class="$style.content">
      <slot></slot>
    </div>
  </div>
</template>

<style lang="scss" module>
.wrapper {
  @include content-wrapper;
}

.content {
  margin: 5rem auto;
  width: 80%;
}
</style>
