import { bigint, pgTable, varchar } from "drizzle-orm/pg-core";
import { user } from "./user";

export const session = pgTable("user_session", {
	id: varchar("id", { length: 128 }).primaryKey(),
	userId: varchar("user_id", { length: 15 })
		.notNull()
		.references(() => user.id),
	activeExpires: bigint("active_expires", { mode: "number" }).notNull(),
	idleExpires: bigint("idle_expires", { mode: "number" }).notNull(),
});
