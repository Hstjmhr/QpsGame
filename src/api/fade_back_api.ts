import type {baseResponse, listDataType, paramsType} from "@/api/index";
import {useAxios} from "@/api/index";

export interface fadeBackCreateType {
    email: string
    content: string
}

export interface fadeBackType {
    id?: number
    content: string
    created_at: string
    email?: string
}

export function feedbackCreateApi(data: fadeBackCreateType):Promise<baseResponse<string>>{
    return useAxios.post("/api/feedback", data)
}

export function feedbackListApi(params?: paramsType):Promise<baseResponse<listDataType<fadeBackType>>>{
    return useAxios.get("/api/feedback", {params})
}