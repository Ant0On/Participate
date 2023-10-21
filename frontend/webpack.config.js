const path = require('path');
const { VueLoaderPlugin } = require('vue-loader')

module.exports = {
    mode: "development",
    entry: ["babel-polyfill", path.resolve("src", "js", "index.js")],
    output: {
        filename: "bundle.js",
        path: path.join(__dirname, "static/js/"),
        publicPath: "/js"
    },
    optimization: {
        minimize: false,
    },
    module: {
        rules: [
            {
                test: /\.vue$/,
                loader: "vue-loader"
            },
            {
                test: /\.js$/,
                loader: "babel-loader"
            },
            {
                test: /\.s[ac]ss$/i,
                use: [
                    // Creates `style` nodes from JS strings
                    'vue-style-loader',
                    // Translates CSS into CommonJS
                    'css-loader',
                    // Compiles Sass to CSS
                    'sass-loader',
                ],
            }
        ]
    },
    resolve: {
        extensions: [".js", ".json", ".jsx", ".vue"],
        alias: {
            'vue': '@vue/runtime-dom'
        }
    },
    devServer: {
        static: './',
        compress: false,
        host: "0.0.0.0",
        proxy: {
            "/api": "http://localhost:8080"
        }
    },
    plugins: [new VueLoaderPlugin()]
};