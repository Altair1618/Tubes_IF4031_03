import { pgTable, timestamp, varchar, uuid, pgEnum } from "drizzle-orm/pg-core";
import { user } from "./user"; 

export enum BookingStatus {
	SUCCESS = "SUCCESS",
	IN_QUEUE = "IN QUEUE", 
	WAITING_FOR_PAYMENT = "WAITING FOR PAYMENT", 
	PURCHASING = "PURCHASING", 
	FAILED = "FAILED"
}

export const bookingStatus = pgEnum(
	"booking_status", 
	Object.values(BookingStatus) as [string]
)

export const booking = pgTable("booking_history", {
	id: uuid("id").primaryKey().defaultRandom(),
	ticketId: uuid("ticket_id").notNull(),
	status: bookingStatus("status").notNull(),
    report: varchar("report"),
	paymentUrl: varchar("payment_url"),
    userId: varchar("user_id", {length: 15}).notNull().references(() => user.id),
	createdAt: timestamp("created_at", { mode: "string" }).notNull().defaultNow(),
	updatedAt: timestamp("updated_at", { mode: "string" }).notNull().defaultNow(),
});
