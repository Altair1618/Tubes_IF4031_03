import { bigint, pgTable, varchar } from "drizzle-orm/pg-core";

export const user = pgTable("auth_user", {
	id: varchar("id", { length: 15 }).primaryKey(),
	email: varchar("email").unique().notNull(),
});
