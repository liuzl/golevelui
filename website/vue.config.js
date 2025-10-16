// vue.config.js
module.exports = {
  outputDir: '../static', // 构建输出目录
  publicPath: "/golevelui/static/",
  assetsDir: 'assets',
  productionSourceMap: false,
  devServer: {
    proxy: {
      '': {
        target: 'http://127.0.0.1:4333',
      },
    }
  },
  configureWebpack: {
    performance: {
      hints: 'warning',
      maxAssetSize: 1024 * 1024, // 1 MiB
      maxEntrypointSize: 1024 * 1024, // 1 MiB
    },
  },
};
