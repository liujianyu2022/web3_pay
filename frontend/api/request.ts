import axios from "axios"

const request = axios.create({
    baseURL: "/",                   // 使用 `/` 即可，Next.js 会自动转发
    timeout: 100 * 1000
})

request.interceptors.request.use(
    config => {
        return config
    },
    error => {
        return Promise.reject(error)
    }
)

request.interceptors.response.use(
    response => {
        return response
    },
    error => {
        return Promise.reject(error)
    }
)


export default request
