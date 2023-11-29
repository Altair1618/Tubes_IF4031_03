import { GetEventsRequest } from "../../dto/event/getEvents.dto"
import { ServiceResponse } from "../../types/common"

const getEventsService = async ({
    query,
    page,
    jwt
}: GetEventsRequest): Promise<ServiceResponse> => {
    if (!query) query = ''
    const url = `${process.env.TICKET_SERVICE_BASE_URL}/api/v1/event?query=${encodeURIComponent(query)}&page=${page}`;

    let response;
    try {
        response = await fetch(url, {
            method: 'GET',
            headers: {
                Authorization: `Bearer ${jwt}`
            },
            credentials: 'include'
        });
    } catch (error) {
        console.log(error)

        return {
            code: 500,
            message: "Internal Server Error",
            data: null,
        }
    }
    
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
