import type { TGqlInput } from '@/entities';
import { client } from '..';
import type { GetTiktokQuery, SearchTiktokQuery, SubSearchTiktokSubscription } from '../generated';
import { TIKTOK_GET_BY_ID, TIKTOK_SEARCH, TIKTOK_SUB_SEARCH } from './operations.gql';

export const tiktokSearch = (
  input: string
) => {
  return client.query<SearchTiktokQuery, TGqlInput<string>>(
    TIKTOK_SEARCH,
    { input }
  );
};

export const tiktokById = (
  input: string
) => {
  return client.query<GetTiktokQuery, TGqlInput<string>>(
    TIKTOK_GET_BY_ID,
    { input }
  );
};

export const tiktokSearchSubscription = (
  input: string
) => {
  return client.subscription<SubSearchTiktokSubscription, TGqlInput<string>>(
    TIKTOK_SUB_SEARCH,
    { input }
  );
};
