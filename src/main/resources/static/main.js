// Author: d1y<chenhonzhou@gmail.com>

const eleName = "app"

/**
 * 节流
 * 参考: https://juejin.im/post/6844903848230780941
 */
const debounce = (fn, wait)=> {
    let timeout = null
    return function() {
        if(timeout !== null) clearTimeout(timeout)
        timeout = setTimeout(fn, wait);
    }
}

/** 节流时间 */
const debounceTime = 4e3

/**
 * 获取路由
 */
const getRouter = ()=> {
    try {
        const path = window.location.pathname
        const cache = path.split("/")
        return cache[2]
    } catch (error) {
        throw new Error(error)
    }
}

const App = new Vue({
    el: document.querySelector(`#${ eleName }`),
    data: {
        content: data,
    },
    computed: {
        formData() {
            const data = App.content
            const router = getRouter()
            // https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams
            const params = new URLSearchParams();
            params.append(`router`, router);
            params.append(`content`, data);
            return params
        }
    },
    methods: {
        handlePushData: debounce(function() {
            const data = App.formData
            axios.post(`/api/router`, data)
        }, debounceTime)
    }
})