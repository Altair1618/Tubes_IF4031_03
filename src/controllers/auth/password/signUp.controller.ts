import Elysia from "elysia";

const passwordSignUpController = new Elysia().post(
	"/password/signup",
	async () => {},
);

export default passwordSignUpController;
