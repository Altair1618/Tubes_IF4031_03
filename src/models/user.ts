import { pgTable, timestamp, varchar } from "drizzle-orm/pg-core";

export const user = pgTable("auth_user", {
	id: varchar("id", { length: 15 }).primaryKey(),
	email: varchar("email").unique().notNull(),
	name: varchar("name").notNull(),
	picture: varchar("picture").notNull(),
	createdAt: timestamp("created_at", { mode: "string" }).notNull().defaultNow(),
	updatedAt: timestamp("updated_at", { mode: "string" }).notNull().defaultNow(),
});
