<!-- TODO: ページネーションを実装する -->
<template>
  <main-view-frame>
    <content-wrapper>
      <h1>投稿記事</h1>
      <post-list :posts="posts" />
    </content-wrapper>
  </main-view-frame>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import PostList from '/@/components/UI/PostList/PostList.vue'
import ContentWrapper from '/@/components/UI/ContentWrapper.vue'
import MainViewFrame from '/@/components/UI/MainViewFrame.vue'

// TODO: implement
const posts = reactive([
  {
    postId: '1',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '2',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '3',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '4',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '5',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '6',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '7',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '8',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '9',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '10',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '11',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  },
  {
    postId: '12',
    title: 'test',
    thumbnailUrl: 'https://avatars.githubusercontent.com/u/93650251?v=4',
    datetime: new Date()
  }
])
</script>
