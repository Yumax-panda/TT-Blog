<template>
  <post-template>
    <editor-content :editor="editor" />
  </post-template>
</template>

<script lang="ts" setup>
import { EditorContent } from '@tiptap/vue-3'
import PostTemplate from '/@/components/UI/Template/PostTemplate.vue'
import useContentEditor from './composables/useContentEditor'

const { editor } = useContentEditor()
</script>

<style lang="scss">
.tiptap {
  min-height: 100vh;
}
</style>
