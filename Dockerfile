FROM oven/bun

WORKDIR /app

COPY package.json .
COPY bun.lockb .

RUN bun install

COPY src src
COPY tsconfig.json .
COPY .env .
COPY drizzle.config.ts .
COPY app.d.ts .

ENV NODE_ENV development
CMD ["bun", "--hot", "run", "src/index.ts"]

EXPOSE 3000
