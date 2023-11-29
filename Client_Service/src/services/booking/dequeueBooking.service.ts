import { sql } from "drizzle-orm";
import { DequeueBookingServicePayload } from "../../dto/booking/dequeueBooking.dto";
import { ServiceResponse } from "../../types/common";
import { db } from "../../configs/drizzle";

const dequeueBookingService = async ({
    ticketId,
}: DequeueBookingServicePayload): Promise<ServiceResponse> => {
    const query = sql`
        SELECT "id"
        FROM "booking_history"
        WHERE "ticket_id" = ${ticketId}
        AND "status" = 'IN QUEUE'
        ORDER BY "created_at" ASC
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

    if (result.rows.length === 0) {
        const url = `${process.env.TICKET_SERVICE_BASE_URL}/api/v1/tickets/${encodeURIComponent(ticketId)}`;
        const ticketResponse = await fetch(url, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                status: "OPEN",
            }),
        });

        const ticketResponseData: any = await ticketResponse.json();
        if (ticketResponse.status !== 200) {
            return {
                code: 500,
                message: ticketResponseData.message,
            }
        }

        return {
            code: 400,
            message: `There is no booking with ticket ID (${ticketId})`,
        }
    } else {
        const bookingHistoryId = result.rows[0].id;
        const query = sql`
            UPDATE "booking_history"
            SET "status" = 'WAITING FOR PAYMENT'
            WHERE "id" = ${bookingHistoryId}
        `;

        try {
            await db.execute(query);
        } catch (e: any) {
            console.log(e);

            return {
                code: 500,
                message: e.message,
            }
        }
    }

    return {
        code: 200,
        message: "Success",
    }
}

export default dequeueBookingService;
