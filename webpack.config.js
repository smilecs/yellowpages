const webpack = require('webpack'); //to access built-in plugins
const path = require('path');

let commonsPlugin = new webpack.optimize.CommonsChunkPlugin('common-chunks');

module.exports = {
  entry: {
    admin: "./ui/src/admin.js"
  },
  output: {
    path: path.resolve(__dirname, "./ui/assets/bin"),
    filename: "[name]-bundle.js"
  },
  plugins: [ commonsPlugin ],
  module: {
      loaders: [{
          test: /\.js$/,
          exclude: /node_modules/,
          loader: 'babel-loader'
      }]
  }
}
