import { AxiosResponse } from "axios"
import { LoginRequestType, RegisterRequestType } from "../interfaces/userTypes"
import request from "./request"
import { ResponseType } from "../interfaces/commonTypes"

export const register = (data: RegisterRequestType): Promise<AxiosResponse<ResponseType>> => {
    return request({
        url: "/api/register",
        method: "post",
        data
    })
}

export const login = (data: LoginRequestType): Promise<AxiosResponse<ResponseType>> => {
    return request({
        url: "/api/login",
        method: "post",
        data
    })
}