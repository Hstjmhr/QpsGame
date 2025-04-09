import type {baseResponse, listDataType, paramsType} from "@/api/index";
import {useAxios} from "@/api/index";

export interface promotionType {
    created_at: string
    href: string
    id: number
    images: string
    is_show: boolean
    title: string
}

export function promotionListApi(params?: paramsType): Promise<baseResponse<listDataType<promotionType>>> {
    return useAxios.get("/api/adverts", {
        params: params, headers: {
            "Gvb_referer": location.pathname,
        }
    })
}

export interface promotionCreateType {
    href: string
    images: string
    is_show: boolean
    title: string
    id?: number
}

export const defaultPromotionForm = {
    href: "",
    images: "",
    is_show: false,
    title: "",
}

export function promotionCreateApi(data: promotionCreateType): Promise<baseResponse<string>> {
    return useAxios.post("/api/adverts", data)
}

export function promotionUpdateApi(id: number, data: promotionCreateType): Promise<baseResponse<string>> {
    return useAxios.put("/api/adverts/" + id.toString(), data)
}