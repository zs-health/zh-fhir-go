import { define } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default define({
  title: "ZARISH HEALTH FHIR",
  description: "FHIR R5 Implementation for Bangladesh Healthcare",
  
  head: [
    ['link', { rel: 'icon', href: '/favicon.svg' }]
  ],

  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Getting Started', link: '/guide/introduction' },
      { text: 'API', link: '/api/overview' },
      { text: 'FHIR Resources', link: '/fhir/overview' },
      { text: 'Terminology', link: '/terminology/overview' },
    ],

    sidebar: {
      '/guide/': [
        {
          text: 'Getting Started',
          items: [
            { text: 'Introduction', link: '/guide/introduction' },
            { text: 'Installation', link: '/guide/installation' },
            { text: 'Quick Start', link: '/guide/quickstart' },
          ]
        }
      ],
      '/api/': [
        {
          text: 'API Reference',
          items: [
            { text: 'Overview', link: '/api/overview' },
            { text: 'FHIR Server', link: '/api/server' },
            { text: 'REST Endpoints', link: '/api/endpoints' },
          ]
        }
      ],
      '/fhir/': [
        {
          text: 'FHIR Resources',
          items: [
            { text: 'Overview', link: '/fhir/overview' },
            { text: 'Patient', link: '/fhir/patient' },
            { text: 'Bangladesh Profiles', link: '/fhir/profiles' },
          ]
        }
      ],
      '/terminology/': [
        {
          text: 'Terminology Service',
          items: [
            { text: 'Overview', link: '/terminology/overview' },
            { text: 'ICD-11', link: '/terminology/icd11' },
            { text: 'Bangladesh Divisions', link: '/terminology/bangladesh' },
          ]
        }
      ],
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/zs-health/zh-fhir-go' }
    ],

    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © 2024-2026 ZARISH HEALTH'
    },

    outline: {
      level: [2, 3]
    }
  }
})
