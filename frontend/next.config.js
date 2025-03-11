module.exports = {
    async rewrites() {
        return [
            {
                source: '/api/:path*',                              // 匹配 /api 开头的所有请求
                destination: 'http://localhost:9000/api/:path*'     // 转发到后端服务
            }
        ]
    },

    webpack(config, options) {
        const { dev, isServer } = options;

        return config;
    },
};