import { useEditor } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'

type Props = Omit<NonNullable<Parameters<typeof useEditor>[0]>, 'extensions'>

const useContentEditor = (props: Props = {}) => {
  const editor = useEditor({
    extensions: [StarterKit],
    ...props
  })

  return { editor }
}

export default useContentEditor
