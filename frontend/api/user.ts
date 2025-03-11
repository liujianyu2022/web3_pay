import { RegisterRequestType } from "../interfaces/userTypes"
import request from "./request"

export const register = (data: RegisterRequestType) => {
    return request({
        url: "/api/register",
        method: "post",
        data
    })
}