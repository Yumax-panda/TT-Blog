<template>
  <div :class="$style.wrapper">
    <div :class="$style.container">
      <slot></slot>
    </div>
  </div>
</template>

<style lang="scss" module>
.wrapper {
  padding: 5rem 0;
}

.container {
  @include colored-container;
  padding: 1rem;
}
</style>
