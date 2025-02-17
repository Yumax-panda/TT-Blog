<template>
  <div :class="$style.container">
    <div :class="$style['message-wrapper']">
      <div :class="$style.welcome">
        Welcome to <span :class="$style.tt">TT !</span>
      </div>
      <div :class="$style.message">
        チーム「TT」は, 東工大で活動するマリオカートサークルです.<br />
        初心者, 上級者, 学年問わず仲良く日々活動しています.
      </div>
      <router-link :to="'/about'" :class="$style.link">
        もっと詳しく!!
      </router-link>
    </div>
    <img :src="inkling" alt="inkling" :class="$style.inkling" />
  </div>
</template>

<script lang="ts" setup>
import inkling from '/@/assets/inkling.webp'
</script>

<style lang="scss" module scoped>
$inkling-ratio: calc(348 / 492);

.container {
  @include color-common-text-white-primary;
  background: linear-gradient(to right, #2c363f, #738da5);
  padding: 12px 24px;
  display: flex;
  gap: 10px;
  justify-content: space-around;
  align-items: center;
  flex-direction: column;

  @include mq(sm) {
    flex-direction: row;
  }
}

.message-wrapper {
  gap: 24px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.welcome {
  font-weight: bold;
  @include size-h2;
}

.tt {
  @include size-h1;
}

.message {
  @include size-body1;
}

.link {
  @include color-common-bg-white-primary;
  color: #2c363f;
  font-size: 16px;
  padding: 4px;
  border-radius: 2px;
  font-weight: bold;
  flex: none;
}

.inkling {
  flex-shrink: 1;
  flex-grow: 0;
  width: 90%;
  height: calc(90vw * $inkling-ratio);

  @include mq(sm) {
    width: clamp(250px, 30vw, 500px);
    height: calc(clamp(250px, 30vw, 500px) * $inkling-ratio);
  }
}
</style>
