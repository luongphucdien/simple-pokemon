import axios, { AxiosError } from "axios"
import { UserData } from "../Components/Login"

const API_URL = "http://localhost:8080/api"

async function test(setState: (data:any)=>void) {
    await axios.get(`${API_URL}/test`).then(r => {setState(r.data)})
}

async function sendAction(action: string, username: string) {
    await axios.post(`${API_URL}/player/action`, {action: action, username: username}).then((r) => {
        console.log(r.data["world-data"])
        // setWorldGrid(r.data["world-data"]["WorldGrid"])
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
        console.log(r.data["world_data"])
        setState(r.data.world)
    }).catch((error: AxiosError) => {
        console.log(error.response)
    })
}

async function addPlayer(userData: UserData) {
    await axios.post(`${API_URL}/player`, userData).then(r => {
        console.log(r.data)
    }).catch((err: AxiosError) => {
        console.log(err.response)
    })
}

async function removePlayer(username: string) {
    await axios.post(`${API_URL}/player/offline`, {"username": username}).then(r => {
        console.log(r.data["world-data"])
    }).catch((err: AxiosError) => {
        console.error(err.response)
    })
}

export {test, testNoParams, sendAction, getWorld, addPlayer, removePlayer}