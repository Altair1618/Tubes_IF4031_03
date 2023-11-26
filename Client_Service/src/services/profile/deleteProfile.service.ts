import { auth } from "../../configs/lucia";
import {
	DeleteProfileServicePayload,
	deleteProfileSchema,
} from "../../dto/profile/deleteProfile.dto";
import { ServiceResponse } from "../../types/common";

const deleteProfileService = async ({
	userId,
}: DeleteProfileServicePayload): Promise<ServiceResponse> => {
	const validateResult = deleteProfileSchema.safeParse({ userId });
	if (!validateResult.success) {
		return {
			code: 400,
			message: validateResult.error.issues[0].message,
		};
	}

	await auth.deleteUser(userId);

	return {
		code: 200,
		message: "Account deleted successfully",
	};
};
export default deleteProfileService;
