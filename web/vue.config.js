const path = require("path");
module.exports = {
    chainWebpack: (config) => {
        config.resolve.alias.set("@", path.resolve(__dirname, "dashboard/src"));
    },
    publicPath: "/benchmark-kafka-go-clients/",
};
