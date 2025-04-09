import type {baseResponse, paramsType, listDataType} from "@/api/index";
import {useAxios} from "@/api/index";

export interface loginEmailType {
    user_name: string
    password: string
}

export function loginEmailApi(request: loginEmailType): Promise<baseResponse<string>> {
    return useAxios.post("/api/email_login", request)
}

export function logoutApi(): Promise<baseResponse<string>> {
    return useAxios.post("/api/logout")
}

export function loginQQPathApi(): Promise<baseResponse<string>> {
    return useAxios.get("/api/qq_login_path")
}

export function qqLoginApi(code: string): Promise<baseResponse<string>> {
    return useAxios.post("/api/login", null, {params: {code}})
}

export interface userInfoType {
    id: number
    created_at: string
    nick_name: string
    user_name: string
    avatar: string
    email: string
    tel: string
    addr: string
    token: string
    ip: string
    role: string
    role_id: number
    sign_status: "QQ" | "gitee" | "邮箱"
    integral: number // 积分
    sign: string // slogan
    link: string // 跳转链接
}

export function userInfoApi(): Promise<baseResponse<userInfoType>> {
    return useAxios.get("/api/user_info")
}


export function userListApi(params: paramsType): Promise<baseResponse<listDataType<userInfoType>>> {
    return useAxios.get("/api/users", {params})
}

export interface userCreateRequest {
    nick_name: string
    password: string
    role: number
    user_name: string
}

export function userCreateApi(data: userCreateRequest): Promise<baseResponse<string>> {
    return useAxios.post("/api/users", data)
}


export interface userUpdateRequest {
    nick_name: string
    role: number
    user_id: number
}

export function userUpdateApi(data: userUpdateRequest): Promise<baseResponse<string>> {
    return useAxios.put("/api/user_role", data)
}

export interface userInfoUpdateType {
    link: string
    nick_name: string
    sign: string
    avatar: string
}

export function userInfoUpdateApi(data: userInfoUpdateType): Promise<baseResponse<string>> {
    return useAxios.put("/api/user_info", data)
}

export interface userUpdatePasswordType {
    old_pwd: string
    pwd: string

    re_pwd: string
}

export function userUpdatePasswordApi(data: userUpdatePasswordType): Promise<baseResponse<string>> {
    return useAxios.put("/api/user_password", data)
}

export interface userBindEmailType {
    code?: string
    email: string
    password?: string
}

export function userBindEmailApi(data: userBindEmailType): Promise<baseResponse<string>> {
    return useAxios.post("/api/user_bind_email", data)
}