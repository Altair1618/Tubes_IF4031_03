import { GetEventTicketsRequest } from "../../dto/event/getEventTickets.dto";
import { ServiceResponse } from "../../types/common";

const getEventTicketsService = async ({
    id,
    jwt,
}: GetEventTicketsRequest): Promise<ServiceResponse> => {
    const url = `${process.env.TICKET_SERVICE_BASE_URL}/api/v1/event/${encodeURIComponent(id)}/tickets`;

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
        message: "Successfully Fetch Event Tickets",
        data: responseData.data
    }
}

export default getEventTicketsService;