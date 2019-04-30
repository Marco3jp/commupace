const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const VueLoaderPlugin = require('vue-loader/lib/plugin');
const CopyWebpackPlugin = require('copy-webpack-plugin');

const ASSET_PATH = '/';
const dist = __dirname + "/dist";
const path = require('path');

module.exports = (isDev) => {
    return {
        entry: {
            'main': './src/scripts/ts/main.ts',
        },
        mode: isDev ? "development" : "production",
        output: {
            path: dist,
            filename: '[name]_[hash].js',
            publicPath: ASSET_PATH
        },
        resolve: {
            extensions: ['*', '.js', '.vue', '.json', '.tsx', '.ts'],
            alias: {
                '@components': path.resolve(__dirname, 'src/components'),
                '@view': path.resolve(__dirname, 'src/view'),
                '@scripts': path.resolve(__dirname, 'src/scripts'),
                '@store': path.resolve(__dirname, 'src/store'),
                '@images': path.resolve(__dirname, 'src/images'),
                '@scss': path.resolve(__dirname, 'src/scss'),
            }
        },
        plugins: [
            new VueLoaderPlugin(),
            new CopyWebpackPlugin([
                {
                    from: path.resolve(__dirname, 'src/images'),
                    to: path.resolve(dist, 'static/images'),
                },
            ]),
            new HtmlWebpackPlugin({
                inject: "body",
                chunks: ['main'],
                template: "./index.html",
                filename: "./index.html"
            }),
            new webpack.DefinePlugin({
                API_PREFIX: JSON.stringify(API_PREFIX),
                APP_VERSION: JSON.stringify(APP_VERSION),
                STUDENT_AUTH_API_EP: JSON.stringify(`/${API_PREFIX}/${API_VERSION}/auth`),
                STUDENT_API_EP: JSON.stringify(`/${API_PREFIX}/${API_VERSION}/service`),
            })
        ],
    }
};