import { eq } from "drizzle-orm";
import { db } from "../../configs/drizzle";
import { ServiceResponse } from "../../types/common";
import { GetBookingGroupServicePayload } from "../../dto/booking/getBookingGroup.dto";
import { booking } from "../../models/booking";

const getBookingGroupService = async ({
	userId,
}: GetBookingGroupServicePayload): Promise<ServiceResponse> => {
	const bookingsData = await db.query.booking.findMany({
		where: eq(user.id, userId),
	});

	return {
		code: 200,
		message: "Success",
		data: bookingsData.map((elmt) => {
            
        })
	};
};
export default getBookingGroupService;
