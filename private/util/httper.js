import axios from 'axios'
import qs from 'qs'

let httper = {
    get(url, data) {
        return new Promise((resolve, reject) => {
            axios.get(url, {
                params: data
            }).then((response) => {
                if (response) {
                    resolve(response);
                }
            }).catch((error) => {
                reject(error);
            })
        })
    },

    post(url, data) {
        return new Promise((resolve, reject) => {
            axios.post(url, qs.stringify(data), {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                    'Accept': 'application/json'
                }
            }).then((response) => {
                if (response) {
                    resolve(response);
                }
            }).catch((error) => {
                reject(error);
            })
        })
    }
}

export default httper;