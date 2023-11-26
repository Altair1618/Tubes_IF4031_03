import { httpResponse } from "../utils/httpResponse";

const authMiddleware = (success: boolean, message?: string) => {
	if (!success) {
		return httpResponse({ code: 401, message: message as string });
	}
};
export default authMiddleware;
