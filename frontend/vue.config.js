module.exports = {
    css: {
        loaderOptions: {
            css: {
                test: /\.styl$/,
                loader: 'css-loader!stylus-loader?paths=node_modules/bootstrap-stylus/stylus/',
            },
            postcss: {
                // options here will be passed to postcss-loader
            },
        },
    },
};
