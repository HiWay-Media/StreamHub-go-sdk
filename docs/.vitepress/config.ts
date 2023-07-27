import { defineConfig } from 'vitepress'
//
// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'streamhub-go-sdk',
  description: 'StreamHub Go SDK is a powerful library developed by HiWay Media that enables seamless integration with the StreamHub platform. StreamHub provides scalable and reliable video streaming solutions, and this SDK allows Go developers to interact with StreamHub\'s APIs effortlessly.',
  base: '/streamhub-go-sdk/',
  themeConfig: {
    outline: [2, 3],
    socialLinks: [
      { icon: 'github', link: 'https://github.com/HiWay-Media/streamhub-go-sdk' },
    ],
    sidebar: [
      { text: 'Introduction', link: '/' },
      { text: 'Getting started', link: '/getting-started' },
    ],
  },
})