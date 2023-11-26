import { eq } from "drizzle-orm";
import { db } from "../../configs/drizzle";
import { ServiceResponse } from "../../types/common";
import { GetProfileServicePayload } from "../../types/profile";
import { user } from "../../models/user";

const getProfileService = async ({
	userId,
}: GetProfileServicePayload): Promise<ServiceResponse> => {
	const userData = await db.query.user.findFirst({
		where: eq(user.id, userId),
	});

	if (!userData) {
		return {
			code: 404,
			message: "Profile not found",
		};
	}

	return {
		code: 200,
		message: "Success",
		data: {
			id: userData.id,
			email: userData.email,
			name: userData.name,
			picture: userData.picture,
		},
	};
};
export default getProfileService;
