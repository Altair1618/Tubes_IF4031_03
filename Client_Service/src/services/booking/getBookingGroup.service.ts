import { db } from "../../configs/drizzle";
import { and, eq } from "drizzle-orm";
import { ServiceResponse } from "../../types/common";
import { booking } from "../../models/booking";
import { GetBookingGroupServicePayload } from "../../dto/booking/getBookingGroup.dto";
import { TicketPriceAndEvent } from "../../types/booking";
import { sql } from 'drizzle-orm' 
import { TicketServiceResponse } from "../../types/common";

const getBookingGroupService = async ({
	userId,
	groupId,
	jwt,

}: GetBookingGroupServicePayload): Promise<ServiceResponse> => {

	const bookingsData = await db.query.booking.findMany({
		where: and(
			eq(booking.userId, userId),
			eq(booking.groupId, groupId)
		)
	});

	const tiketIds = bookingsData.map((elmt) => elmt.id)
	
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

	const responseData = await response.json() as TicketServiceResponse<TicketPriceAndEvent>;

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
			const totalPage = elmt['total_page'] as number

			const bookingGroup: BookingGroup = {
				groupId,
				date,
				eventName,
				totalPrice,
				overallStatus,
				paymentUrl,
				totalPage
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
