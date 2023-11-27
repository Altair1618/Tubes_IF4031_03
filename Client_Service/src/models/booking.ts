import { pgTable, timestamp, varchar, uuid } from "drizzle-orm/pg-core";
import { user } from "./user"; 

export const booking = pgTable("booking_history", {
	id: uuid("id").primaryKey().defaultRandom(),
	ticketId: uuid("ticket_id").unique().notNull(),
	status: varchar("status").notNull(),
    report: varchar("report"),
    userId: varchar("user_id", {length: 15}).notNull().references(() => user.id),
	createdAt: timestamp("created_at", { mode: "string" }).notNull().defaultNow(),
	updatedAt: timestamp("updated_at", { mode: "string" }).notNull().defaultNow(),
});
