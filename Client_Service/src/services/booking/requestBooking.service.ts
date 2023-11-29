import { db } from "../../configs/drizzle";
import { sql } from 'drizzle-orm'
import { RequestBookingServicePayload } from "../../dto/booking/requestBooking.dto";
import { ServiceResponse } from "../../types/common";

const requestBookingService = async ({
    ticketId,
    userId,
    jwt,
}: RequestBookingServicePayload): Promise<ServiceResponse> => {
    const checkQuery = sql`
        SELECT "status"
        FROM "booking_history"
        WHERE "ticket_id" = ${ticketId}
        AND "user_id" = ${userId}
        AND "status" != 'FAILED'
        LIMIT 1
    `;

    let checkResult;
    try {
        checkResult = await db.execute(checkQuery);
    } catch (e: any) {
        console.log(e);

        return {
            code: 500,
            message: e.message,
        }
    }

    if (checkResult.rows.length > 0) {
        console.log(checkResult.rows)

        return {
            code: 400,
            message: `You have already booked this ticket with status (${checkResult.rows[0]['status']})`,
        }
    }

    const query = sql`
        INSERT INTO "booking_history" ("ticket_id", "user_id", "status")
        VALUES (${ticketId}, ${userId}, 'WAITING FOR PAYMENT')
        RETURNING "id"
    `;

    let result;
    try {
        result = await db.execute(query);
    } catch (e: any) {
        console.log(e);

        return {
            code: 500,
            message: e.message,
        }
    }

    const bookingId = result.rows[0]['id'];
    const ticket_url = `${process.env.TICKET_SERVICE_BASE_URL}/api/v1/tickets/${encodeURIComponent(ticketId)}/book`;

    let response;
    try {
        response = await fetch(ticket_url, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + jwt,
            },
            credentials: 'include',
            body: JSON.stringify({
                bookingId: bookingId,
            }),
        });
    } catch (e: any) {
        console.log(e);

        return {
            code: 500,
            message: e.message,
        }
    }

    const responseData = await response.json();
    if (!(responseData.data.status && responseData.data.status !== 'FAILED')) {
        const updateQuery = sql`
            UPDATE "booking_history"
            SET "status" = 'FAILED'
            WHERE "id" = ${bookingId}
        `;

        try {
            await db.execute(updateQuery);
        } catch (e: any) {
            console.log(e);

            return {
                code: 500,
                message: e.message,
            }
        }

        return {
            code: 500,
            message: responseData.message,
            data: responseData.data,
        }
    }

    return {
        code: 200,
        message: "Request Booking Success",
        data: responseData.data,
    }
}

export default requestBookingService;
