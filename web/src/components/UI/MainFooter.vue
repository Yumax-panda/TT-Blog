<template>
  <footer>
    <div :class="$style.container">
      <div :class="$style.sitemap">
        <router-link :to="'/'" :class="$style.name"
          >東工大マリオカートサークル</router-link
        >
        <router-link :to="'/about'">サークルについて</router-link>
        <router-link :to="'/pages'">記事一覧</router-link>
        <router-link :to="'/contact'">お問い合わせ</router-link>
      </div>
      <div :class="$style.copy">Copyright Yumax-panda All rights reserved.</div>
    </div>
  </footer>
</template>

<script lang="ts" setup></script>

<style lang="scss" module>
footer {
  border-top: solid 1px white;
}

.container {
  text-decoration: none;
  vertical-align: middle;
  display: flex;
  flex-direction: column;
  gap: 12px;
  @include color-common-text-white-primary;
  @include container;
}

.sitemap {
  justify-content: center;
  display: flex;
  flex-direction: column;
  margin-left: 1rem;
  gap: 0.5rem;

  @include mq(sm) {
    margin-left: 0;
    align-items: center;
    flex-direction: row;
  }
}

.name {
  font-weight: bold;
  @include size-h2;
}

.copy {
  text-align: center;
}
</style>
