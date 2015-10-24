module.exports = {
    entry: [
        './client/app/app.js'
    ],
    devtool: 'sourcemap',
    output: {
        filename: 'bundle.min.js'
    },
    module: {
        loaders: [{
            test: /\.js$/,
            exclude: [/app\/lib/, /node_modules/],
            loaders: ['babel']
        }, {
            test: /\.html$/,
            loader: 'raw'
        }, {
            test: /\.styl$/,
            loader: 'style!css!stylus'
        }, {
            test: /\.css$/,
            loader: 'style!css'
        }]
    }
};
