import axios, { AxiosResponse } from "axios"

const API_URL = "http://localhost:8080/api"

async function test(setState: (data:any)=>void) {
    await axios.get(`${API_URL}/test`).then(r => {setState(r.data)})
}

export {test}