import { ServiceResponse } from "../../types/common";
import { GetProfileServicePayload } from "../../types/profile";

const getProfileService = async ({
	userId,
}: GetProfileServicePayload): Promise<ServiceResponse> => {
	return {
		code: 200,
		message: "Success",
		data: userId,
	};
};
export default getProfileService;
