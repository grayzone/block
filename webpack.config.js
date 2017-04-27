const path = require('path');
var node_modules_dir = path.join(__dirname, 'node_modules');

var config = {
  entry: {
    app: path.resolve(__dirname, 'js/main.jsx')
  },
  output: {
    filename: 'block.js',
    path: path.resolve(__dirname, 'static/dist')
  },
  resolve: {
    extensions: [
      '.js', '.jsx'
    ],
    alias: {}
  },
  module: {
    rules: [
      {
        test: /\.jsx?$/,
        exclude: /node_modules/,
        use: [
          {
            loader: "babel-loader"
          }
        ]
      }, {
        test: /\.css$/,
        use: ['style-loader', 'css-loader?modules', 'postcss-loader']
      }, {
        test: /\.less$/,
        use: [
          'style-loader', {
            loader: 'css-loader',
            options: {
              importLoaders: 1
            }
          },
          'less-loader'
        ]
      }
    ]
  }
};

module.exports = config;