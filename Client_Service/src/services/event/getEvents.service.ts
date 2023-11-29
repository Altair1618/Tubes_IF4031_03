import { GetEventsRequest } from "../../dto/event/getEvents.dto"
import { ServiceResponse } from "../../types/common"

const getEventsService = async ({
    query,
    page,
    jwt
}: GetEventsRequest): Promise<ServiceResponse> => {
    if (!query) query = ''
    const url = `${process.env.TICKET_SERVICE_BASE_URL}/api/v1/event?query=${encodeURIComponent(query)}&page=${page}`;

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
        message: "Successfully Fetch Events",
        data: responseData.data.events
    }
}

export default getEventsService;
