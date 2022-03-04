"use strict";
const path = require("path");
const pkg = require("./package.json");
const name = pkg.name;
const port = 9527;

// 访问域名
// const serviceUrl_1 = 'http://127.0.0.1:8081';  // 打包
const serviceUrl_1 = "http://192.168.10.20:8081"; // 开发

// const serviceUrl_2 = 'http://127.0.0.1:8081';  // 打包
const serviceUrl_2 = "http://192.168.12.131:8889"; // 开发



function resolve(dir) {
  return path.join(__dirname, dir);
}

module.exports = {
  publicPath: "/",
  outputDir: "web-app",
  assetsDir: "static",
  lintOnSave: process.env.NODE_ENV === "development" ? "error" : false,
  productionSourceMap: false,
  devServer: {
    port: port,
    open: false,
    overlay: {
      warnings: false,
      errors: true
    },
    proxy: {
        [process.env.VUE_APP_LAN]: {
            target: serviceUrl_1 + "/app/",
            changeOrigin: true,
            pathRewrite: {
              ["^" + process.env.VUE_APP_LAN]: ""
            }
        },
        [process.env.VUE_APP_AUTH]: {
            target: serviceUrl_2 + "/sys_auth/",
            changeOrigin: true,
            pathRewrite: {
            ["^" + process.env.VUE_APP_AUTH]: ""
            }
        }
    },
    after(app) {
      require("@babel/register");
      const bodyParser = require("body-parser");

      app.use(bodyParser.json());
      app.use(
        bodyParser.urlencoded({
          extended: true
        })
      );
    }
  },
  configureWebpack: {
    name: name,
    resolve: {
      alias: {
        "@": resolve("src")
      }
    }
  },
  chainWebpack(config) {
    config.plugins.delete("preload");
    config.plugins.delete("prefetch");

    config.module
      .rule("vue")
      .use("vue-loader")
      .loader("vue-loader")
      .tap(options => {
        options.compilerOptions.preserveWhitespace = true;
        return options;
      })
      .end();

    config.when(process.env.NODE_ENV === "development", config =>
      config.devtool("cheap-source-map")
    );

    config.when(process.env.NODE_ENV !== "development", config => {
      config
        .plugin("ScriptExtHtmlWebpackPlugin")
        .after("html")
        .use("script-ext-html-webpack-plugin", [
          {
            inline: /runtime\..*\.js$/
          }
        ])
        .end();
      config.optimization.splitChunks({
        chunks: "all",
        cacheGroups: {
          libs: {
            name: "chunk-libs",
            test: /[\\/]node_modules[\\/]/,
            priority: 10,
            chunks: "initial"
          },
          elementUI: {
            name: "chunk-elementUI", // split elementUI into a single package
            priority: 20, // the weight needs to be larger than libs and app or it will be packaged into libs or app
            test: /[\\/]node_modules[\\/]element-ui[\\/]/
          },
          commons: {
            name: "chunk-commons",
            test: resolve("src/components"), // can customize your rules
            minChunks: 3, //  minimum common number
            priority: 5,
            reuseExistingChunk: true
          }
        }
      });
      config.optimization.runtimeChunk("single");
    });
  }
};
