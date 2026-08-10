FROM node:22-alpine AS builder
WORKDIR /app
# 代理（容器内通过 host.docker.internal 访问宿主机代理）
ARG HTTP_PROXY=http://host.docker.internal:7897
ARG HTTPS_PROXY=http://host.docker.internal:7897
ENV HTTP_PROXY=$HTTP_PROXY
ENV HTTPS_PROXY=$HTTPS_PROXY
COPY package.json package-lock.json ./
RUN npm config set proxy "$HTTP_PROXY" && \
    npm config set https-proxy "$HTTPS_PROXY" && \
    npm ci
COPY tsconfig.json next.config.js tailwind.config.js postcss.config.js ./
COPY src ./src
COPY public ./public
RUN npm run build

FROM node:22-alpine
WORKDIR /app
ENV NODE_ENV=production
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./
COPY --from=builder /app/public ./public
EXPOSE 3001
CMD ["npm", "start"]
