import type { Config } from "drizzle-kit";
import * as dotenv from "dotenv";
dotenv.config();

export default {
	schema: "./src/models/*",
	out: "./drizzle",
	driver: "pg",
	dbCredentials: {
		user: process.env.POSTGRES_USER ?? "",
		password: process.env.POSTGRES_PASSWORD ?? "",
		host: "localhost",
		port: process.env.POSTGRES_PORT
			? parseInt(process.env.POSTGRES_PORT) ?? 5432
			: 5432,
		database: process.env.POSTGRES_DB ?? "tessera",
	},
} satisfies Config;
