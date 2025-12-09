import { handleFetchResponse, SERVER_BASE_URL } from "@/lib/helper";
import { Dashboard } from "@/lib/schemas";
import { Error } from "@/lib/types";



export const getDasboardData = async () : Promise<Dashboard> => {
    const BASE_URL = `${SERVER_BASE_URL}/dashboard`

    try {
        const response  = await fetch(BASE_URL, {
            method: "GET"
        })
        return await handleFetchResponse(response)
    } catch(e) {
        throw e as Error
    }
}