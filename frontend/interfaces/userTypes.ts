
export interface RegisterRequestType {
    email: string
    nickname: string
    password: string
    re_password: string
}

export interface LoginRequestType {
    email: string
    password: string
}