import Elysia from "elysia";

const passwordSignInController = new Elysia().post(
	"/password/signin",
	async () => {},
);
export default passwordSignInController;
