import { Client, cacheExchange, fetchExchange, subscriptionExchange } from '@urql/vue';
import { SubscriptionClient } from 'subscriptions-transport-ws';

const url = (): string => import.meta.env.VITE_API_URL
const urlWs = (): string => import.meta.env.VITE_API_URL_WS

export const client = new Client({
  url: url(),
  exchanges: [
    cacheExchange,
    fetchExchange,
    subscriptionExchange({
      forwardSubscription(operation) {
        const wsClient = new SubscriptionClient(urlWs(), {
          reconnect: true,
        })
        return wsClient.request(operation)
      }
    })
  ],
})
