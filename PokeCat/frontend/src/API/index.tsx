import axios, { AxiosError } from "axios"

const API_URL = "http://localhost:8080/api"

async function test(setState: (data:any)=>void) {
    await axios.get(`${API_URL}/test`).then(r => {setState(r.data)})
}

async function sendAction(action: string, setState: (data: any) => void) {
    await axios.post(`${API_URL}/player/action`, {key: action}).then((r) => {
        console.log(r.data)
        setState(r.data.key)
    }).catch((error: AxiosError) => {
        console.log(error.response)
    })
    
}

async function testNoParams() {
    await axios.get(`${API_URL}/test`).then(r => console.log(r.data)).catch((error: AxiosError) => {
        console.log(error.response)
    })
}

async function getWorld(setState: (data: any) => void) {
    await axios.get(`${API_URL}/world`).then(r => {
        console.log(r.data.world)
        setState(r.data.world)
    }).catch((error: AxiosError) => {
        console.log(error.response)
    })
}

export {test, testNoParams, sendAction, getWorld}