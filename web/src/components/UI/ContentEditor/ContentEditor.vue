<template>
  <editor-content :editor="editor" />
</template>

<script lang="ts" setup>
import { EditorContent } from '@tiptap/vue-3'
import useContentEditor from './composables/useContentEditor'

const { editor } = useContentEditor()
</script>
