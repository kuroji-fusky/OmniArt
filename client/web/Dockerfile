FROM node:22.14.0 AS base

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /client

FROM base AS prod-deps
COPY package*.json .
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --prod --frozen-lockfile

FROM base AS build
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm run build

FROM base
COPY --from=prod-deps ./client/node_modules /node_modules
COPY --from=build ./client/dist /dist
EXPOSE 3000

CMD [ "node", "build" ]