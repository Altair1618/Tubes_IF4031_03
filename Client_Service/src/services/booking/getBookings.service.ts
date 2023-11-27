import { db } from "../../configs/drizzle";
import { ServiceResponse } from "../../types/common";
import { GetBookingsServicePayload } from "../../dto/booking/getBookings.dto";
import { TicketPriceAndEventName, BookingGroup } from "../../types/booking";
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

	const MAX_GROUP_PER_PAGE = 5

	const query = sql`
		WITH distinct_groups AS (
			SELECT DISTINCT "group_id"
			FROM "booking_history"
			WHERE "user_id" = ${userId}
			ORDER BY "group_id" ASC
			LIMIT ${MAX_GROUP_PER_PAGE}
		)
		
		SELECT "ticket_id", "status", "group_id", "report", "created_at"
		FROM "booking_history"
		WHERE "user_id" = ${userId} AND "group_id" IN (SELECT "group_id" FROM distinct_groups)
		ORDER BY "group_id" ASC;
	`;

	const bookingsData = await db.execute(query)

	const tiketIds = bookingsData.rows.map(elmt => elmt['ticket_id'] as string)
	
	const params = new URLSearchParams();

	// Append each value of the array as a separate parameter with the same key
	tiketIds.forEach(tiketId => {
		params.append('booking_ids[]', tiketId);
	});

	const response = await fetch(`${process.env.TICKET_SERVICE_BASE_URL}/tikets/ids?${params.toString()}`, {
		method: 'GET',
		headers: {
			Authorization: `Bearer ${jwt}`
		},
		credentials: 'include'
	});

	const responseData = await response.json() as TicketServiceResponse<TicketPriceAndEventName>;

	const tiketPricesAndEventNames = responseData.data
	const bookingGroups: BookingGroup[] = []
	const bookingDict: GroupDictionary = {}

	bookingsData.rows.forEach((elmt) => {

		const groupId = elmt['group_id'] as number
		const ticketId = elmt['ticket_id'] as string
		const status = elmt['status'] as string

		if (!bookingDict[groupId])
		{
			bookingDict[groupId] = true

			const date = elmt['date'] as string
			const eventName = tiketPricesAndEventNames[ticketId].eventName
			const totalPrice = tiketPricesAndEventNames[ticketId].price
			const overallStatus = status
			const paymentUrl = elmt['report'] as string | null

			const bookingGroup: BookingGroup = {
				groupId,
				date,
				eventName,
				totalPrice,
				overallStatus,
				paymentUrl
			}

			bookingGroups.push(bookingGroup)
		}

		else
		{
			const lastGroup = bookingGroups[bookingGroups.length - 1]
			lastGroup.totalPrice += tiketPricesAndEventNames[ticketId].price
			// UPDATE STATUS
		}
	})

	return {
		code: 200,
		message: "Success",
		data: bookingsData
	}
};

export default getBookingsService;
