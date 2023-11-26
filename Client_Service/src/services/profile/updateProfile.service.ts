import { db } from "../../configs/drizzle";
import {
	UpdateProfileServicePayload,
	updateProfileSchema,
} from "../../dto/profile/updateProfile.dto";
import { user } from "../../models/user";
import { ServiceResponse } from "../../types/common";

const updateProfileService = async ({
	name,
}: UpdateProfileServicePayload): Promise<ServiceResponse> => {
	const validateResult = updateProfileSchema.safeParse({ name });
	if (!validateResult.success) {
		return {
			code: 400,
			message: validateResult.error.issues[0].message,
		};
	}

	try {
		await db.update(user).set({ name: name });
	} catch (e) {
		console.log(e);
		return {
			code: 500,
			message: "Something went wrong, please try again later",
		};
	}

	return {
		code: 200,
		message: "Profile successfully updated",
	};
};
export default updateProfileService;
