import { GetEventDetailRequest } from "../../dto/event/getEventDetail.dto";
import { ServiceResponse } from "../../types/common";

const getEventDetailService = async ({
    id,
    jwt
}: GetEventDetailRequest): Promise<ServiceResponse> => {
    const url = `${process.env.TICKET_SERVICE_BASE_URL}/api/v1/event/${encodeURIComponent(id)}`;

    const response = await fetch(url, {
        method: 'GET',
        headers: {
            Authorization: `Bearer ${jwt}`
        },
        credentials: 'include'
    });
    
    const responseData: any = await response.json();

    if (responseData.code !== 200) {
        return {
            code: responseData.code,
            message: responseData.message,
            data: null,
        }
    }

    return {
        code: 200,
        message: "Successfully Fetch Event Detail",
        data: responseData.data.event
    }
}

export default getEventDetailService;