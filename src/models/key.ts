import { pgTable, varchar } from "drizzle-orm/pg-core";
import { user } from "./user";

export const key = pgTable("user_key", {
	id: varchar("id", { length: 255 }).primaryKey(),
	userId: varchar("user_id", { length: 15 })
		.notNull()
		.references(() => user.id),
	hashedPassword: varchar("hashed_password", {
		length: 255,
	}),
});
