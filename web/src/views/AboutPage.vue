<!-- TODO: ページの内容を足す (戦績、成果物etc) -->
<template>
  <main-view-frame>
    <content-wrapper>
      <section>
        <h1>東工大マリオカートサークルについて</h1>
        <p>
          チーム「TT」は、東工大で活動するマリオカートサークルです。<br />
          初心者・上級者や学年問わずメンバー同士仲良く日々活動に取り組んでいます。
        </p>
      </section>

      <section>
        <h2>活動内容</h2>
        <ul>
          <li>交流戦 (6vs6)</li>
          <li>模擬戦 (基本的に模擬と呼びます)</li>
          <li>MK8DX Lounge (ガチ勢のレート戦です)</li>
          <li>大会 (MKC, 大学対抗戦)</li>
          <li>内戦・イベント (年に数回企画しています)</li>
          <li>オフ会 (飲み会など)</li>
        </ul>
        <p>
          その他、マリオカートに限らず別のゲームや競技プログラミングを楽しんだり、Discord
          Bot等の開発をしているメンバーもいます。
          自分の趣味をサークルのメンバーと気軽に共有して楽しめるのも本サークルの良いところです。
        </p>
      </section>
    </content-wrapper>
  </main-view-frame>
</template>

<script lang="ts" setup>
import ContentWrapper from '/@/components/UI/ContentWrapper.vue'
import MainViewFrame from '/@/components/UI/MainViewFrame.vue'
</script>

<style lang="scss" module>
h1,
h2 {
  padding: 0.5rem 0;
  margin: 0.5rem 0;
  border-bottom: solid 1px white;
}

h1 {
  @include size-h1;
}

h2 {
  @include size-h2;
}

section {
  padding: 0.5rem 0;
}

p {
  line-height: 1.5rem;
  @include size-body1;
}
</style>
