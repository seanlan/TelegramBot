module.exports = {
  title: 'Telegram Bot Manager',
  s3config: {
    cdn: 'https://cdn.taadu.com/',
    endpoint: 'https://0fa13440a3a7294b98ac1ac586e286dc.r2.cloudflarestorage.com',
    access_key: '0baade1a8d3b5abefc194f263761ec24',
    secret_key: 'e8ce67f94a123280eb05585d33773365d246b6e0c63c88ca0db017c78a85e871',
    version: 'v4',
    bucket: 'www'
  },

  /**
   * @type {boolean} true | false
   * @description Whether show the settings right-panel
   */
  showSettings: false,

  /**
   * @type {boolean} true | false
   * @description Whether need tagsView
   */
  tagsView: false,

  /**
   * @type {boolean} true | false
   * @description Whether fix the header
   */
  fixedHeader: false,

  /**
   * @type {boolean} true | false
   * @description Whether show the logo in sidebar
   */
  sidebarLogo: false,

  /**
   * @type {string | array} 'production' | ['production', 'development']
   * @description Need show err logs component.
   * The default is only used in the production env
   * If you want to also use it in dev, you can pass ['production', 'development']
   */
  errorLog: 'production'
}
