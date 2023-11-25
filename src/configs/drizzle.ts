import { Pool } from "pg";
import { drizzle } from "drizzle-orm/node-postgres";
import { user } from "../models/user";

export const pool = new Pool({
	user: process.env.POSTGRES_USER ?? "",
	password: process.env.POSTGRES_PASSWORD ?? "",
	host: process.env.POSTGRES_HOST ?? "client_service_db",
	port: process.env.POSTGRES_PORT
		? parseInt(process.env.POSTGRES_PORT) ?? 5432
		: 5432,
	database: process.env.POSTGRES_DB ?? "tessera",
});

export const db = drizzle(pool, { schema: { user } });
