// theme.js または別のファイルにこの設定を保存
import { extendTheme } from '@chakra-ui/react';

const theme = extendTheme({
  fonts: {
    body: 'Inter, GenJyuuGothicL, "Hiragino Kaku Gothic ProN", "Hiragino Sans", Meiryo, sans-serif',
    heading: 'Inter, GenJyuuGothicL, "Hiragino Kaku Gothic ProN", "Hiragino Sans", Meiryo, sans-serif',
  },
});

export default theme;