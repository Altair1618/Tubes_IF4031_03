import { db } from "../../configs/drizzle";
import { ServiceResponse } from "../../types/common";
import { GetBookingsServicePayload } from "../../dto/booking/getBookings.dto";
import { BookingDataWithTotalPage, TicketPriceAndEvent} from "../../types/booking";
import { sql } from 'drizzle-orm' 
import { TicketServiceResponse } from "../../types/common";

interface GroupDictionary {
	[key: number]: boolean
}

const getBookingsService = async ({
	userId,
	jwt,
	page
}: GetBookingsServicePayload): Promise<ServiceResponse> => {

	const MAX_BOOKING_PER_PAGE = 10
	const query = sql`
		WITH total_page as (
			SELECT CEIL(CAST(COUNT("id") AS FLOAT) / ${MAX_BOOKING_PER_PAGE}) AS count
			FROM "booking_history"
		)

		SELECT "id", "ticket_id", "status", "report", "created_at", total_page.count as "total_page"
		FROM "booking_history", "total_page"
		WHERE "user_id" = ${userId}
		OFFSET ${MAX_BOOKING_PER_PAGE * (parseInt(page) - 1)}
		LIMIT ${MAX_BOOKING_PER_PAGE}
	`;

	const bookingsData = await db.execute(query)

	const ticketIds = bookingsData.rows.map(elmt => elmt['ticket_id'] as string)
	const idsParam = ticketIds.join(',');
	console.log(`${process.env.TICKET_SERVICE_BASE_URL}/api/v1/tickets?ids=${idsParam}`)
	const response = await fetch(`${process.env.TICKET_SERVICE_BASE_URL}/api/v1/tickets?ids=${idsParam}`, {
		method: 'GET',
		headers: {
			Authorization: `Bearer ${jwt}`
		},
		credentials: 'include'
	});

	const responseData = await response.json() as TicketServiceResponse<TicketPriceAndEvent>;

	const tiketPricesAndEventNames = responseData.data;
	const bookings: BookingDataWithTotalPage[]  = []

	bookingsData.rows.forEach((elmt) => {
		const id = elmt['id'] as string;
		const ticketId = elmt['ticket_id'] as string;
		const status = elmt['status'] as string;
		const report = elmt['report'] as string | null;
		const createdAt = elmt['created_at'] as string;
		const totalPage = elmt['total_page'] as number;

		bookings.push({
			id,
			ticketId,
			status,
			report,
			createdAt,
			totalPage,
			...(tiketPricesAndEventNames[ticketId] ?? {})
		})
	})

	return {
		code: 200,
		message: "Success",
		data: bookings
	}
};

export default getBookingsService;
