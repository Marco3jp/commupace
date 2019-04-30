const merge = require('webpack-merge');
const baseConfig = require('./webpack.config.base.js');

module.exports = merge(baseConfig(true), {
    module: {
        rules: [
            {
                test: /\.ts$/,
                use: [
                    {
                        loader: 'ts-loader',
                        options: {
                            appendTsSuffixTo: [/\.vue$/]
                        }
                    }
                ]
            },
            {
                test: /\.scss/,
                use: [
                    {
                        loader: 'style-loader',
                        options: {
                            hmr: true,
                            singleton: true,
                        }
                    },
                    {
                        loader: 'css-loader',
                        options: {
                            url: false,
                            sourceMap: true,
                            importLoaders: 2
                        },
                    },
                    {
                        loader: 'sass-loader',
                        options: {
                            sourceMap: true,
                            /*
                            data: `
                                @import './src/scss/admin/_values.scss';
                                @import './src/scss/_animate.scss';
                                `
                             */
                        }
                    }
                ],
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            },
            {
                test: /\.(html)$/,
                use: {
                    loader: 'html-loader',
                    options: {
                        attrs: [':data-src']
                    }
                }
            },
        ]
    },
    devtool: 'inline-source-map',
    devServer: {
        contentBase: 'dist',
        port: 4500,
        host: '0.0.0.0',
        disableHostCheck: true,
        historyApiFallback: true,
        inline: true,
    },
    resolve: {
        alias: {
            vue: 'vue/dist/vue.js',
        }
    },
    plugins: [
        new webpack.NamedModulesPlugin(),
        new webpack.HotModuleReplacementPlugin(),
    ]
});

if (module.hot) {
    module.hot.accept();
}
