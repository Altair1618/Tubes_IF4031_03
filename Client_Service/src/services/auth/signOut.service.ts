import { auth } from "../../configs/lucia";
import { ServiceResponse } from "../../types/common";

const signOutService = async (sessionId: string): Promise<ServiceResponse> => {
	await auth.invalidateSession(sessionId);
	return {
		code: 200,
		message: "Sign Out success",
	};
};
export default signOutService;
