import http from "k6/http"
import { check } from "k6"

export const options = {
    stages: [
        { duration: "2m", target: 500 },
        { duration: "3m", target: 500 },
        { duration: "1m", target: 0 }
    ]
}


export default function () {
    TestFrontend()
}

function TestFrontend() {
    const url = __ENV.FRONTEND_URL
    const response = http.get(url)
    check(response, {
        "is status 200": r => r.status == 200,
    })
}

function TestAuthService() {
    const url = __ENV.AUTH_SVC_URL
    const response = http.get(url)
    check(response, {
        "Is status 200", r => r.sta
    })
}

function TestUserService() {
    
}

export function teardown(data) {
    console.log(data)
    console.log("Done")
}