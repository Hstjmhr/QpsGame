import type {baseResponse, paramsType} from "@/api/index";
import type {optionType} from "@/types";
import {useAxios} from "@/api/index";

export function roleIdListApi(params?:paramsType):Promise<baseResponse<optionType[]>>{
    return useAxios.get("/api/role_ids", {params})
}