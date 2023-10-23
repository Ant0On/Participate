const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    compress: false,
    host: "0.0.0.0",
    proxy: "http://backend:3000"
  }
})
